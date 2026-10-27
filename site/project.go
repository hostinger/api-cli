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
	raw, err := os.ReadFile(Path(dir))
	if errors.Is(err, os.ErrNotExist) {
		raw, err = os.ReadFile(filepath.Join(dir, Dir, legacyFile))
	}
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var p Project
	if err := json.Unmarshal(raw, &p); err != nil {
		return nil, fmt.Errorf("%s: %w", Path(dir), err)
	}
	if p.Domain == "" {
		return nil, nil
	}
	if p.Type == "php" {
		p.Type = "static" // the skill's name for files served as they are
	}
	return &p, nil
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
	if err := os.MkdirAll(filepath.Join(dir, Dir), 0o755); err != nil {
		return err
	}
	merged := map[string]any{}
	if raw, err := os.ReadFile(Path(dir)); err == nil {
		_ = json.Unmarshal(raw, &merged) // unreadable: start over
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
	if err := os.WriteFile(Path(dir), append(raw, '\n'), 0o600); err != nil {
		return err
	}
	_ = os.Remove(filepath.Join(dir, Dir, legacyFile))
	return ensureGitignore(dir)
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
	return os.WriteFile(path, []byte(content), 0o644)
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
