// Package site resolves the website a command acts on and remembers it per
// directory, like `vercel link`.
package site

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/index"
)

const (
	Dir = ".hostinger"
	// File is shared with the deploy-to-hosting skill, so Save keeps its fields.
	File = "site.json"
	// legacyFile is what early builds of this CLI wrote.
	legacyFile = "project.json"
)

type Project struct {
	// Kind is "website"; empty in older files means the same.
	Kind string `json:"kind,omitempty"`
	// Path is the deployed folder relative to the link; empty is the directory itself.
	Path     string `json:"path,omitempty"`
	OrderID  int    `json:"orderId,omitempty"`
	Domain   string `json:"domain"`
	Username string `json:"username,omitempty"`
	Type     string `json:"type,omitempty"`
	// Headless is set for a headless website, which only takes Node.js deploys.
	Headless bool `json:"headless,omitempty"`
	// Source is "upload" or "git".
	Source string `json:"source,omitempty"`
	// Node is only written after a successful build.
	Node *NodeSettings `json:"node,omitempty"`
	// Build is the last build up started, so the build commands default to it.
	Build *BuildRef `json:"build,omitempty"`
}

type BuildRef struct {
	UUID  string    `json:"uuid"`
	State string    `json:"state,omitempty"`
	At    time.Time `json:"at"`
}

// NodeSettings are remembered build settings; empty strings mean "not set".
type NodeSettings struct {
	NodeVersion     int    `json:"nodeVersion,omitempty"`
	AppType         string `json:"appType,omitempty"`
	RootDirectory   string `json:"rootDirectory,omitempty"`
	OutputDirectory string `json:"outputDirectory,omitempty"`
	BuildScript     string `json:"buildScript,omitempty"`
	EntryFile       string `json:"entryFile,omitempty"`
	PackageManager  string `json:"packageManager,omitempty"`
}

func Path(dir string) string {
	return filepath.Join(dir, Dir, File)
}

// Load returns nil, nil when dir has no link.
func Load(dir string) (*Project, error) {
	path := Path(dir)
	raw, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		path = filepath.Join(dir, Dir, legacyFile)
		raw, err = os.ReadFile(path)
	}
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var p Project
	if err := json.Unmarshal(raw, &p); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if p.Domain == "" {
		return nil, nil
	}
	// A cloned repository can ship the link, and the domain reaches the
	// terminal and prompts.
	if p.Domain, err = canonicalDomain(p.Domain); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if p.Type == "php" {
		p.Type = "static" // the skill's name for files served as they are
	}
	return &p, nil
}

// canonicalDomain drops a trailing dot and keeps the rest as given, so the
// link matches the domain as the API spells it (IDNs included), then refuses
// what is not a hostname.
func canonicalDomain(d string) (string, error) {
	a := strings.TrimSuffix(d, ".")
	if !validDomain(a) {
		return "", fmt.Errorf("%q is not a valid domain", d)
	}
	return a, nil
}

// validDomain accepts hostnames whose labels are 1-63 runes of letters,
// digits, combining marks and inner hyphens, in Unicode or punycode.
func validDomain(d string) bool {
	if len(d) > 253 {
		return false
	}
	for _, label := range strings.Split(d, ".") {
		if n := utf8.RuneCountInString(label); n == 0 || n > 63 || label[0] == '-' || label[len(label)-1] == '-' {
			return false
		}
		for _, c := range label {
			if !(unicode.IsLetter(c) || unicode.IsDigit(c) || unicode.IsMark(c) || c == '-') {
				return false
			}
		}
	}
	return true
}

// Find looks for a link in start and its parents, returning it and its directory.
func Find(start string) (*Project, string, error) {
	dirs, err := LinkDirs(start)
	if err != nil {
		return nil, "", err
	}
	for _, dir := range dirs {
		p, err := Load(dir)
		if err != nil {
			return nil, "", err
		}
		if p != nil {
			return p, dir, nil
		}
	}
	return nil, "", nil
}

// CommittedLink returns the link file in dir that Load reads when Git tracks
// it, else "". A cloned repository can ship one naming someone else's website,
// so a committed link is not consent to replace it. A repository that cannot
// be read counts as tracking it.
func CommittedLink(dir string) string {
	file := Path(dir)
	if _, err := os.Lstat(file); errors.Is(err, os.ErrNotExist) {
		file = filepath.Join(dir, Dir, legacyFile)
	}
	repo, err := git.PlainOpenWithOptions(dir, &git.PlainOpenOptions{DetectDotGit: true, EnableDotGitCommonDir: true})
	if errors.Is(err, git.ErrRepositoryNotExists) {
		return ""
	}
	if err != nil {
		return file
	}
	var idx *index.Index
	wt, err := repo.Worktree()
	if err == nil {
		idx, err = repo.Storer.Index()
	}
	if err != nil {
		return file
	}
	// Resolved, so a symlinked .hostinger cannot hide a tracked file. An
	// unresolvable root is "", which Rel refuses.
	root, _ := filepath.EvalSymlinks(wt.Filesystem.Root())
	real, err := filepath.EvalSymlinks(file)
	if err != nil {
		return file
	}
	rel, err := filepath.Rel(root, real)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return file
	}
	// Case-folded: macOS and Windows read a committed .HOSTINGER/SITE.json as
	// the link too.
	rel = filepath.ToSlash(rel)
	for _, e := range idx.Entries {
		if strings.EqualFold(e.Name, rel) {
			return file
		}
	}
	return ""
}

// LinkDirs are where a link for start can live: start, then its parents up
// to the repository root. It never climbs into the home directory or above,
// so a link made there once does not apply to every project below it.
func LinkDirs(start string) ([]string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return nil, err
	}
	home, _ := os.UserHomeDir()
	dirs := []string{dir}
	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dirs, nil // the repository root
		}
		parent := filepath.Dir(dir)
		if parent == dir || (home != "" && (parent == home || strings.HasPrefix(home, parent+string(filepath.Separator)))) {
			return dirs, nil
		}
		dir = parent
		dirs = append(dirs, dir)
	}
}

// Save writes the link, keeping other tools' fields, and gitignores Dir
// since it names the hosting account.
func Save(dir string, p *Project) error {
	// Load refuses a domain it does not accept, so refuse it here too.
	if p.Domain != "" {
		d, err := canonicalDomain(p.Domain)
		if err != nil {
			return fmt.Errorf("not linking: %w", err)
		}
		q := *p
		q.Domain = d
		p = &q
	}
	if err := checkLinkDir(dir); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(dir, Dir), 0o755); err != nil {
		return err
	}
	merged := map[string]any{}
	if raw, err := os.ReadFile(Path(dir)); err == nil {
		_ = json.Unmarshal(raw, &merged) // unreadable: start over
	}
	// Other tools' fields (website_uid...) describe the website the link named:
	// keep them for that website only.
	if old, _ := merged["domain"].(string); old != "" && !strings.EqualFold(old, p.Domain) {
		merged = map[string]any{}
	}
	ours, err := json.Marshal(p)
	if err != nil {
		return err
	}
	var fields map[string]any
	if err := json.Unmarshal(ours, &fields); err != nil {
		return err
	}
	for _, k := range []string{"kind", "path", "orderId", "domain", "username", "type", "source", "node", "build"} {
		delete(merged, k)
	}
	for k, v := range fields {
		merged[k] = v
	}
	raw, err := json.MarshalIndent(merged, "", "  ")
	if err != nil {
		return err
	}
	if err := writeFile(Path(dir), append(raw, '\n'), 0o600); err != nil {
		return err
	}
	_ = os.Remove(filepath.Join(dir, Dir, legacyFile))
	return ensureGitignore(dir)
}

// Remove deletes the link files only: Dir can also hold other tools' files
// (the headless skill's workspaces under ~/.hostinger). Dir goes too once
// empty, unless it is a symlinked, shared folder.
func Remove(dir string) error {
	if err := checkLinkDir(dir); err != nil {
		return err
	}
	for _, name := range []string{File, legacyFile} {
		if err := os.Remove(filepath.Join(dir, Dir, name)); err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	if info, err := os.Lstat(filepath.Join(dir, Dir)); err == nil && info.Mode()&os.ModeSymlink == 0 {
		_ = os.Remove(filepath.Join(dir, Dir))
	}
	return nil
}

// ensureGitignore adds Dir to .gitignore inside a Git checkout, as Vercel does.
func ensureGitignore(dir string) error {
	if !insideGit(dir) {
		return nil
	}
	path := filepath.Join(dir, ".gitignore")
	raw, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	for _, line := range strings.Split(string(raw), "\n") {
		line = strings.TrimSpace(line)
		if line == Dir || line == Dir+"/" || line == "/"+Dir || line == "/"+Dir+"/" {
			return nil
		}
	}
	content := string(raw)
	if content != "" && !strings.HasSuffix(content, "\n") {
		content += "\n"
	}
	content += Dir + "/\n"
	return writeFile(path, []byte(content), 0o644)
}

// checkLinkDir refuses a symlinked Dir that resolves outside the project: the
// Git checkout dir is in (below the home directory, as LinkDirs climbs), else
// dir itself. A checkout can ship the symlink; inside, it is a shared folder.
func checkLinkDir(dir string) error {
	path := filepath.Join(dir, Dir)
	info, err := os.Lstat(path)
	if err != nil || info.Mode()&os.ModeSymlink == 0 {
		return nil // missing, or not a link: MkdirAll creates it or fails
	}
	root := dir
	if dirs, err := LinkDirs(dir); err == nil {
		top := dirs[len(dirs)-1]
		if _, err := os.Stat(filepath.Join(top, ".git")); err == nil {
			root = top
		}
	}
	realRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return err
	}
	real, err := filepath.EvalSymlinks(path)
	if err != nil {
		return fmt.Errorf("refusing to use %s: %w", path, err)
	}
	if rel, err := filepath.Rel(realRoot, real); err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return fmt.Errorf("refusing to use %s: it links outside %s", path, root)
	}
	return nil
}

// writeFile replaces path with data, refusing a symlink or other non-regular
// file. It renames a temp file over path, so a link swapped in after the check
// is replaced rather than followed. An existing file keeps its mode.
func writeFile(path string, data []byte, perm os.FileMode) error {
	info, err := os.Lstat(path)
	switch {
	case err == nil && info.Mode()&os.ModeSymlink != 0:
		return fmt.Errorf("refusing to write %s: it is a symlink", path)
	case err == nil && !info.Mode().IsRegular():
		return fmt.Errorf("refusing to write %s: it is not a regular file", path)
	case err == nil:
		perm = info.Mode().Perm()
	case !errors.Is(err, os.ErrNotExist):
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), "."+filepath.Base(path)+".*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name()) // a no-op once renamed
	_, err = tmp.Write(data)
	if err == nil {
		err = tmp.Chmod(perm)
	}
	if cerr := tmp.Close(); err == nil {
		err = cerr
	}
	if err == nil {
		err = os.Rename(tmp.Name(), path)
	}
	return err
}

func insideGit(dir string) bool {
	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return false
		}
		dir = parent
	}
}
