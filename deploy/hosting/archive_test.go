package hosting

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestListFilesSkipsJunkAndSecrets(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"index.html", "css/app.css", "node_modules/x/index.js", ".git/HEAD", ".hostinger/project.json", ".DS_Store", ".env", ".env.production", ".env.example", "src/.env.local", ".envrc", "sub/.git", ".netrc", ".git-credentials", "keys/id_ed25519", "certs/server.pem", "tls.key", ".svn/entries", ".hg/store"} {
		write(t, dir, name, "x")
	}
	if err := os.Symlink(filepath.Join(dir, "index.html"), filepath.Join(dir, "link.html")); err != nil {
		t.Skip("symlinks not supported:", err)
	}
	l, err := ListFiles(context.Background(), dir)
	if err != nil {
		t.Fatal(err)
	}
	wantFiles := []string{".env.example", "css/app.css", "index.html"}
	wantSkipped := []string{".env", ".env.production", ".envrc", ".git-credentials", ".netrc", "certs/server.pem", "keys/id_ed25519", "src/.env.local", "tls.key"}
	if !reflect.DeepEqual(l.Files, wantFiles) {
		t.Errorf("files = %v, want %v", l.Files, wantFiles)
	}
	if !reflect.DeepEqual(l.Skipped, wantSkipped) {
		t.Errorf("skipped = %v, want %v", l.Skipped, wantSkipped)
	}
	if l.Ignored != 0 || l.IgnoreSource != "" {
		t.Errorf("no ignore file: ignored=%d source=%q", l.Ignored, l.IgnoreSource)
	}
}

func TestListFilesRespectsGitignore(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"index.html", "src/app.js", "dist/bundle.js", "dist/assets/a.css", "logs/x.log", "keep.log", "docs/notes.md", "docs/draft.tmp"} {
		write(t, dir, name, "x")
	}
	write(t, dir, ".gitignore", "dist/\n*.log\n!keep.log\n")
	write(t, dir, "docs/.gitignore", "*.tmp\n")
	l, err := ListFiles(context.Background(), dir)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{".gitignore", "docs/.gitignore", "docs/notes.md", "index.html", "keep.log", "src/app.js"}
	if !reflect.DeepEqual(l.Files, want) {
		t.Errorf("files = %v, want %v", l.Files, want)
	}
	if l.Ignored != 4 || l.IgnoreSource != ".gitignore" {
		t.Errorf("ignored=%d source=%q", l.Ignored, l.IgnoreSource)
	}

	// The named directory itself is never hidden by its parent's rules.
	l, err = ListFiles(context.Background(), filepath.Join(dir, "dist"))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(l.Files, []string{"assets/a.css", "bundle.js"}) {
		t.Errorf("dist files = %v", l.Files)
	}
}

func TestHostingerignoreWinsOverGitignore(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"index.html", "dist/bundle.js", "scratch/n.txt"} {
		write(t, dir, name, "x")
	}
	write(t, dir, ".gitignore", "dist/\n")
	write(t, dir, IgnoreFile, "# local clutter\nscratch/\n")
	l, err := ListFiles(context.Background(), dir)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{".gitignore", ".hostingerignore", "dist/bundle.js", "index.html"}
	if !reflect.DeepEqual(l.Files, want) || l.IgnoreSource != IgnoreFile || l.Ignored != 1 {
		t.Errorf("files=%v ignored=%d source=%q", l.Files, l.Ignored, l.IgnoreSource)
	}
}

func TestZipRoundTrip(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "index.html", "<h1>hi</h1>")
	write(t, dir, "css/app.css", "body{}")
	dest := filepath.Join(t.TempDir(), "site.zip")
	if err := Zip(context.Background(), dir, []string{"css/app.css", "index.html"}, dest); err != nil {
		t.Fatal(err)
	}
	r, err := zip.OpenReader(dest)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	var names []string
	for _, f := range r.File {
		names = append(names, f.Name)
	}
	if !reflect.DeepEqual(names, []string{"css/app.css", "index.html"}) {
		t.Errorf("zip entries = %v", names)
	}
}

func TestArchiveName(t *testing.T) {
	now := time.Date(2026, 10, 20, 9, 30, 0, 0, time.UTC)
	old := randomSuffix
	randomSuffix = func() string { return "r4nd" }
	t.Cleanup(func() { randomSuffix = old })
	if got := ArchiveName("/tmp/My Site!", now); got != "My-Site_20261020T093000Z_r4nd.zip" {
		t.Errorf("ArchiveName = %q", got)
	}
	if got := ArchiveName("/tmp/!!!", now); got != "site_20261020T093000Z_r4nd.zip" {
		t.Errorf("ArchiveName = %q", got)
	}
}

func TestIsArchive(t *testing.T) {
	for path, want := range map[string]bool{"a.zip": true, "a.tar.gz": true, "a.TGZ": true, "a.txt": false, "dir": false} {
		if IsArchive(path) != want {
			t.Errorf("IsArchive(%s) = %v", path, !want)
		}
	}
}

func TestListFilesFollowsASymlinkedRoot(t *testing.T) {
	release := t.TempDir()
	write(t, release, "index.html", "x")
	current := filepath.Join(t.TempDir(), "current")
	if err := os.Symlink(release, current); err != nil {
		t.Skip("symlinks not supported:", err)
	}
	l, err := ListFiles(context.Background(), current)
	if err != nil || len(l.Files) != 1 || l.Files[0] != "index.html" {
		t.Errorf("files = %v, err = %v", l, err)
	}
}

func TestUnreadableIgnoreFileIsAnError(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "index.html", "x")
	if err := os.Mkdir(filepath.Join(dir, IgnoreFile), 0o755); err != nil { // reading a directory fails
		t.Fatal(err)
	}
	if _, err := ListFiles(context.Background(), dir); err == nil {
		t.Error("fell back to .gitignore instead of failing")
	}
}

// .gitignore files are not looked for where packing never goes, so an
// unreadable folder there does not fail the deploy.
func TestGitignoresSkipWhatPackingSkips(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "index.html", "x")
	write(t, dir, ".gitignore", "*.log\n")
	locked := filepath.Join(dir, "node_modules", "locked")
	if err := os.MkdirAll(locked, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(locked, 0); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(locked, 0o755) })
	l, err := ListFiles(context.Background(), dir)
	if err != nil || !reflect.DeepEqual(l.Files, []string{".gitignore", "index.html"}) {
		t.Errorf("files = %v, err = %v", l, err)
	}
}

func TestNodeVersionErrorListsOfferedVersions(t *testing.T) {
	err := checkNodeOverrides(NodeOverrides{NodeVersion: 16})
	want := fmt.Sprint(NodeVersions)
	if err == nil || !strings.Contains(err.Error(), strings.ReplaceAll(strings.Trim(want, "[]"), " ", ", ")) {
		t.Errorf("err = %v, want the versions %s", err, want)
	}
}

// Ctrl-C while packing stops at once, not when the walk or zip finishes.
func TestPackingStopsWhenCancelled(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "index.html", "x")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := ListFiles(ctx, dir); !errors.Is(err, context.Canceled) {
		t.Errorf("ListFiles: %v", err)
	}
	if err := Zip(ctx, dir, []string{"index.html"}, filepath.Join(t.TempDir(), "a.zip")); !errors.Is(err, context.Canceled) {
		t.Errorf("Zip: %v", err)
	}
}

// Symlinks are not followed but are listed, so nothing goes missing silently.
func TestSymlinksAreListed(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "real.html", "x")
	if err := os.Symlink(filepath.Join(dir, "real.html"), filepath.Join(dir, "index.html")); err != nil {
		t.Skip("symlinks not supported:", err)
	}
	if err := os.Symlink(t.TempDir(), filepath.Join(dir, "assets")); err != nil {
		t.Fatal(err)
	}
	l, err := ListFiles(context.Background(), dir)
	if err != nil || !reflect.DeepEqual(l.Files, []string{"real.html"}) || !reflect.DeepEqual(l.Symlinks, []string{"assets", "index.html"}) {
		t.Errorf("listing = %+v, err = %v", l, err)
	}
}

// .npmrc ships with a Node.js build (npm install reads it) but not to a static site.
func TestNpmrcLeftOutOfStaticDeploys(t *testing.T) {
	l := &Listing{Files: []string{".npmrc", "index.html", "sub/.npmrc"}}
	l.leaveOutStatic()
	if !reflect.DeepEqual(l.Files, []string{"index.html"}) || !reflect.DeepEqual(l.Skipped, []string{".npmrc", "sub/.npmrc"}) {
		t.Errorf("listing = %+v", l)
	}
}

func TestUploadNameIsUnique(t *testing.T) {
	old := randomSuffix
	randomSuffix = func() string { return "r4nd" }
	t.Cleanup(func() { randomSuffix = old })
	for in, want := range map[string]string{"/x/site.zip": "site_r4nd.zip", "/x/App.tar.gz": "App_r4nd.tar.gz", "/x/b.tgz": "b_r4nd.tgz"} {
		if got := uploadName(in); got != want {
			t.Errorf("uploadName(%s) = %s, want %s", in, got, want)
		}
	}
}
