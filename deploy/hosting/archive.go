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
	"slices"
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
	Ignored int
	// Restored are the files .gitignore lists that a WordPress import packs anyway.
	Restored     []string
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
	return listFiles(ctx, dir, nil)
}

// listWordPress is ListFiles for a WordPress import, which replaces the whole
// site: what the site needs (wordPressNeeds) is packed even when .gitignore
// lists it, as WordPress repos often ignore wp-config.php, uploads, plugins or
// core. In plugins, themes and mu-plugins everything but logs and caches is
// packed: their vendor/ and dist/ are code the site runs. Elsewhere only a rule
// hiding one of those roots, or all of a root's contents (uploads/*), is
// overridden (rulesBelowRoots): the rest of .gitignore applies, *.log and
// uploads/wc-logs/ included. IgnoreFile applies in full.
func listWordPress(ctx context.Context, dir string) (*Listing, error) {
	return listFiles(ctx, dir, wordPressNeeds)
}

// wordPressNeeds is the WordPress site itself: core, wp-config.php and the parts
// of wp-content that hold content and code (not cache, upgrade or debug.log).
func wordPressNeeds(rel string) bool {
	parts := strings.Split(rel, "/")
	top := strings.ToLower(parts[0])
	if len(parts) == 1 {
		return top == "wp-content" || top == "wp-admin" || top == "wp-includes" || top == ".htaccess" ||
			top == "index.php" || top == "xmlrpc.php" || (strings.HasPrefix(top, "wp-") && strings.HasSuffix(top, ".php"))
	}
	switch top {
	case "wp-admin", "wp-includes":
		return true
	case "wp-content":
		sub := strings.ToLower(parts[1])
		if len(parts) == 2 && strings.HasSuffix(sub, ".php") {
			return true // index.php and drop-ins such as db.php
		}
		switch sub {
		case "uploads", "plugins", "themes", "mu-plugins", "languages":
			return true
		}
	}
	return false
}

func listFiles(ctx context.Context, dir string, needs func(rel string) bool) (*Listing, error) {
	// WalkDir does not descend into a symlinked root (./current -> release).
	if real, err := filepath.EvalSymlinks(dir); err == nil {
		dir = real
	}
	patterns, source, err := ignoreRules(dir)
	if err != nil {
		return nil, err
	}
	var matcher gitignore.Matcher
	if len(patterns) > 0 {
		matcher = gitignore.NewMatcher(patterns)
	}
	if source != ".gitignore" {
		needs = nil // IgnoreFile is written for deploys: it is followed in full
	}
	out := &Listing{IgnoreSource: source}
	// hidden is whether the rules hide rel or, when ancestors is set, a
	// folder above it: a folder kept open for what the site needs still hides
	// the rest of its files.
	hidden := func(m gitignore.Matcher, parts []string, isDir, ancestors bool) bool {
		if m.Match(parts, isDir) {
			return true
		}
		for i := 1; ancestors && i < len(parts); i++ {
			if m.Match(parts[:i], true) {
				return true
			}
		}
		return false
	}
	// below caches rulesBelowRoots' matcher: it depends only on the roots,
	// the first two parts, and whether they are folders.
	below := map[string]gitignore.Matcher{}
	// ignored is whether rel is left out, and restored whether it is packed
	// although .gitignore hides it.
	ignored := func(rel string, isDir bool) (hide, restored bool) {
		parts := strings.Split(rel, "/")
		if matcher == nil || !hidden(matcher, parts, isDir, needs != nil) {
			return false, false
		}
		if needs == nil || !needs(rel) {
			return true, false
		}
		// Plugins and themes are code the site runs: their gitignored vendor/
		// and dist/ are packed, only logs and caches stay out.
		if shipsWithCode(rel) {
			if isLogOrCache(parts, isDir) {
				return true, false
			}
			return false, true
		}
		// Elsewhere only the rules hiding a root the site needs are overridden;
		// *.log or uploads/wc-logs/ still apply.
		key := fmt.Sprintf("%s|%t", strings.Join(parts[:min(2, len(parts))], "/"), isDir || len(parts) > 2)
		m, ok := below[key]
		if !ok {
			m = gitignore.NewMatcher(rulesBelowRoots(patterns, parts, isDir, needs))
			below[key] = m
		}
		if hidden(m, parts, isDir, true) {
			return true, false
		}
		return false, true
	}
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
			if hide, _ := ignored(rel, true); hide {
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
			if hide, _ := ignored(rel, false); hide {
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
		hide, restored := ignored(rel, false)
		if hide {
			out.Ignored++
			return nil
		}
		if restored {
			out.Restored = append(out.Restored, rel)
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

// rulesBelowRoots drops the rules that hide a root of parts the site needs
// (wp-config.php, wp-content, wp-content/uploads, …) or everything in one
// (wp-content/*, uploads/*), which a made-up name in it shows. What is left
// targets something below the roots: *.log, uploads/wc-logs/.
func rulesBelowRoots(patterns []gitignore.Pattern, parts []string, isDir bool, needs func(string) bool) []gitignore.Pattern {
	var roots [][]string
	for i := 1; i <= min(2, len(parts)); i++ {
		if needs(strings.Join(parts[:i], "/")) {
			roots = append(roots, parts[:i])
		}
	}
	return slices.DeleteFunc(slices.Clone(patterns), func(p gitignore.Pattern) bool {
		for _, r := range roots {
			dir := isDir || len(r) < len(parts)
			if p.Match(r, dir) == gitignore.Exclude || dir && hidesAllIn(p, r) {
				return true
			}
		}
		return false
	})
}

// isLogOrCache: a log file, or the cache folder of a plugin or theme
// (wp-content/plugins/<p>/cache/), which it rebuilds and which can hold
// customer data. A deeper cache folder is code: vendor/psr/cache, src/Cache.
func isLogOrCache(parts []string, isDir bool) bool {
	if len(parts) > 3 && (len(parts) > 4 || isDir) {
		if p := strings.ToLower(parts[3]); p == "cache" || p == ".cache" {
			return true
		}
	}
	return !isDir && strings.HasSuffix(strings.ToLower(parts[len(parts)-1]), ".log")
}

// hidesAllIn is whether p hides anything in the folder dir (uploads/*), which
// a made-up name in it shows; uploads/wc-logs/ or *.log do not.
func hidesAllIn(p gitignore.Pattern, dir []string) bool {
	probe := append(slices.Clone(dir), "hostinger-probe")
	return p.Match(probe, false) == gitignore.Exclude || p.Match(probe, true) == gitignore.Exclude
}

// leaveOutStatic moves .npmrc files to Skipped: they often hold registry tokens,
// which a static or imported WordPress site would serve. A Node.js build needs
// them for npm install.
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
func ignoreRules(dir string) ([]gitignore.Pattern, string, error) {
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
		return patterns, IgnoreFile, nil
	}
	patterns, err := readGitignores(dir, nil)
	if err != nil {
		return nil, "", err
	}
	if len(patterns) == 0 {
		return nil, "", nil
	}
	return patterns, ".gitignore", nil
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

// uploadName is the name an archive or dump is uploaded as: never the user's
// own file name, which could replace a file of that name in the document root.
// The name is cleaned and kept short: the API refuses "..", and paths over 255
// characters, only after both files are uploaded.
func uploadName(archive string) string {
	name := filepath.Base(archive)
	ext := filepath.Ext(name)
	if strings.HasSuffix(strings.ToLower(name), ".tar.gz") {
		ext = name[len(name)-len(".tar.gz"):]
	}
	stem := unsafeUploadName.ReplaceAllString(strings.TrimSuffix(name, ext), "-")
	stem = strings.Trim(dotRuns.ReplaceAllString(stem, "."), "-.")
	if len(stem) > 100 {
		stem = strings.TrimRight(stem[:100], "-.")
	}
	if stem == "" {
		stem = "upload"
	}
	return stem + "_" + randomSuffix() + ext
}

var (
	unsafeUploadName = regexp.MustCompile(`[^A-Za-z0-9._-]+`)
	dotRuns          = regexp.MustCompile(`\.{2,}`)
)

// randomSuffix keeps uploaded archives in the document root from being guessed
// from the folder name and time. Tests replace it.
var randomSuffix = func() string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
