package hosting

import (
	"archive/zip"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

// archiveExtensions are what the deploy endpoints accept as-is.
var archiveExtensions = []string{".zip", ".tar", ".tar.gz", ".tgz", ".7z", ".gz", ".gzip"}

func IsArchive(path string) bool {
	lower := strings.ToLower(path)
	for _, ext := range archiveExtensions {
		if strings.HasSuffix(lower, ext) {
			return true
		}
	}
	return false
}

// isSecretFile matches env files (examples such as .env.example are kept),
// credentials and private keys: a static deploy would serve them publicly,
// so no ignore file can bring them back.
func isSecretFile(name string) bool {
	switch name {
	case ".env", ".envrc", ".netrc", ".git-credentials", ".pgpass", "auth.json", ".my.cnf", ".pypirc", ".vault-token": // .envrc: direnv
		return true
	}
	// Shell and REPL histories, and Terraform state (it stores secrets in plain text).
	if strings.HasSuffix(name, "_history") || strings.HasSuffix(name, ".tfstate") || strings.HasSuffix(name, ".tfstate.backup") {
		return true
	}
	for _, key := range []string{"id_rsa", "id_dsa", "id_ecdsa", "id_ed25519"} {
		if strings.HasPrefix(name, key) {
			return true
		}
	}
	if ext := filepath.Ext(name); ext == ".pem" || ext == ".key" {
		return true
	}
	if !strings.HasPrefix(name, ".env.") {
		return false
	}
	switch name {
	case ".env.example", ".env.sample", ".env.template", ".env.dist":
		return false
	}
	return true
}

// IgnoreFile replaces the .gitignore files when present (same syntax).
const IgnoreFile = ".hostingerignore"

type Listing struct {
	// Files are sorted, forward-slash and relative to Dir.
	Files []string
	// Skipped are the secret files left out (see isSecretFile).
	Skipped []string
	// Symlinks are not followed, so what they point to is left out.
	Symlinks []string
	// Ignored files were left out by IgnoreSource's rules.
	Ignored      int
	IgnoreSource string
	Dir          string
}

func (l *Listing) Has(rel string) bool {
	i := sort.SearchStrings(l.Files, rel)
	return i < len(l.Files) && l.Files[i] == rel
}

// ListFiles returns the files to ship: always without node_modules, .git,
// .hostinger, .env files and symlinks, then filtered by IgnoreFile or the
// .gitignore files. A parent's .gitignore cannot hide the named directory.
// A cancelled ctx stops the walk.
func ListFiles(ctx context.Context, dir string) (*Listing, error) {
	// WalkDir does not descend into a symlinked root (./current -> release).
	if real, err := filepath.EvalSymlinks(dir); err == nil {
		dir = real
	}
	matcher, source, err := ignoreRules(dir)
	if err != nil {
		return nil, err
	}
	out := &Listing{IgnoreSource: source}
	walk := func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		name := d.Name()
		rel = filepath.ToSlash(rel)
		if d.IsDir() {
			if name == "node_modules" || name == ".git" || name == ".svn" || name == ".hg" || name == ".hostinger" {
				return filepath.SkipDir
			}
			if matcher != nil && matcher.Match(strings.Split(rel, "/"), true) {
				n, cerr := countFiles(ctx, path)
				if cerr != nil {
					return cerr
				}
				out.Ignored += n
				return filepath.SkipDir
			}
			return nil
		}
		if d.Type()&fs.ModeSymlink != 0 {
			if matcher != nil && matcher.Match(strings.Split(rel, "/"), false) {
				out.Ignored++
			} else {
				out.Symlinks = append(out.Symlinks, rel)
			}
			return nil
		}
		// In a worktree or submodule .git is a file.
		if !d.Type().IsRegular() || name == ".DS_Store" || name == ".git" {
			return nil
		}
		if isSecretFile(name) {
			out.Skipped = append(out.Skipped, rel)
			return nil
		}
		if matcher != nil && matcher.Match(strings.Split(rel, "/"), false) {
			out.Ignored++
			return nil
		}
		out.Files = append(out.Files, rel)
		return nil
	}
	if err := filepath.WalkDir(dir, walk); err != nil {
		return nil, err
	}
	sort.Strings(out.Files)
	sort.Strings(out.Skipped)
	sort.Strings(out.Symlinks)
	return out, nil
}

// leaveOutStatic moves .npmrc files to Skipped: they often hold registry tokens,
// which a static site would serve. A Node.js build needs them for npm install.
func (l *Listing) leaveOutStatic() {
	kept := l.Files[:0]
	for _, rel := range l.Files {
		if path.Base(rel) == ".npmrc" {
			l.Skipped = append(l.Skipped, rel)
		} else {
			kept = append(kept, rel)
		}
	}
	l.Files = kept
	sort.Strings(l.Skipped)
}

// nonEmpty is an error when nothing is left to deploy.
func (l *Listing) nonEmpty() error {
	switch {
	case len(l.Files) > 0:
		return nil
	case l.Ignored > 0:
		return fmt.Errorf("%s has no files to deploy: %s left out all %d of them", l.Dir, l.IgnoreSource, l.Ignored)
	}
	return fmt.Errorf("%s has no files to deploy", l.Dir)
}

// ignoreRules loads IgnoreFile when present, else every .gitignore under dir.
func ignoreRules(dir string) (gitignore.Matcher, string, error) {
	raw, err := os.ReadFile(filepath.Join(dir, IgnoreFile))
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, "", err
	}
	if err == nil {
		var patterns []gitignore.Pattern
		for _, line := range strings.Split(string(raw), "\n") {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			patterns = append(patterns, gitignore.ParsePattern(line, nil))
		}
		return gitignore.NewMatcher(patterns), IgnoreFile, nil
	}
	patterns, err := readGitignores(dir, nil)
	if err != nil {
		return nil, "", err
	}
	if len(patterns) == 0 {
		return nil, "", nil
	}
	return gitignore.NewMatcher(patterns), ".gitignore", nil
}

// readGitignores is gitignore.ReadPatterns without walking into the folders
// packing always skips (node_modules can be huge) or ignored ones, and without
// failing on an unreadable subfolder: packing reports it if it is shipped.
func readGitignores(root string, domain []string) ([]gitignore.Pattern, error) {
	dir := filepath.Join(append([]string{root}, domain...)...)
	var ps []gitignore.Pattern
	raw, err := os.ReadFile(filepath.Join(dir, ".gitignore"))
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	for _, line := range strings.Split(string(raw), "\n") {
		if line = strings.TrimRight(line, "\r"); strings.TrimSpace(line) != "" && !strings.HasPrefix(line, "#") {
			ps = append(ps, gitignore.ParsePattern(line, domain))
		}
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		if len(domain) > 0 {
			return ps, nil
		}
		return nil, err
	}
	for _, e := range entries {
		name := e.Name()
		if !e.IsDir() || name == "node_modules" || name == ".git" || name == ".hostinger" {
			continue
		}
		sub := append(append([]string{}, domain...), name)
		if gitignore.NewMatcher(ps).Match(sub, true) {
			continue
		}
		subps, err := readGitignores(root, sub)
		if err != nil {
			return nil, err
		}
		ps = append(ps, subps...)
	}
	return ps, nil
}

func countFiles(ctx context.Context, dir string) (int, error) {
	n := 0
	err := filepath.WalkDir(dir, func(_ string, d os.DirEntry, err error) error {
		if err != nil {
			return filepath.SkipDir // only a count: an unreadable folder is not worth failing over
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		if !d.IsDir() && d.Type().IsRegular() {
			n++
		}
		return nil
	})
	return n, err
}

// Zip packs files from dir into dest; a cancelled ctx stops it between files.
func Zip(ctx context.Context, dir string, files []string, dest string) (err error) {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := out.Close(); err == nil {
			err = cerr
		}
	}()
	w := zip.NewWriter(out)
	for _, rel := range files {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := addToZip(w, dir, rel); err != nil {
			return err
		}
	}
	return w.Close()
}

func addToZip(w *zip.Writer, dir, rel string) error {
	src, err := os.Open(filepath.Join(dir, filepath.FromSlash(rel)))
	if err != nil {
		return err
	}
	defer src.Close()
	info, err := src.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = rel
	header.Method = zip.Deflate
	entry, err := w.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(entry, src)
	return err
}

var unsafeName = regexp.MustCompile(`[^A-Za-z0-9_-]+`)

// ArchiveName is the folder name made safe plus a timestamp, so deploys never collide.
func ArchiveName(dir string, now time.Time) string {
	base := unsafeName.ReplaceAllString(filepath.Base(dir), "-")
	base = strings.Trim(base, "-")
	if base == "" {
		base = "site"
	}
	return fmt.Sprintf("%s_%s_%s.zip", base, now.UTC().Format("20060102T150405Z"), randomSuffix())
}

// uploadName is the name an archive is uploaded as: never the user's own file
// name, which could replace a file of that name in the document root.
func uploadName(archive string) string {
	name := filepath.Base(archive)
	ext := filepath.Ext(name)
	if strings.HasSuffix(strings.ToLower(name), ".tar.gz") {
		ext = name[len(name)-len(".tar.gz"):]
	}
	return strings.TrimSuffix(name, ext) + "_" + randomSuffix() + ext
}

// randomSuffix keeps uploaded archives in the document root from being guessed
// from the folder name and time. Tests replace it.
var randomSuffix = func() string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
