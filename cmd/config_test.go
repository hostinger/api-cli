package cmd

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func setToken(m map[string]any) { m["api_token"] = "secret" }

func assertPrivate(t *testing.T, path string) {
	t.Helper()
	info, err := os.Lstat(path)
	if err != nil {
		t.Fatal(err)
	}
	if runtime.GOOS != "windows" && info.Mode().Perm() != 0o600 {
		t.Fatalf("mode = %v, want 0600", info.Mode().Perm())
	}
}

func TestWriteConfigNeverWritesTokenIntoReadableInode(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".hostinger.yaml")
	if err := os.WriteFile(path, []byte("format: json\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0o644); err != nil { // umask may have narrowed it
		t.Fatal(err)
	}
	reader, err := os.Open(path) // another user opened it while it was 0644
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()

	if err := writeConfig(path, setToken); err != nil {
		t.Fatal(err)
	}

	held, _ := io.ReadAll(reader)
	if strings.Contains(string(held), "secret") {
		t.Fatalf("token reached the previously opened file: %q", held)
	}
	got, _ := os.ReadFile(path)
	if !strings.Contains(string(got), "api_token: secret") || !strings.Contains(string(got), "format: json") {
		t.Fatalf("config = %q", got)
	}
	assertPrivate(t, path)
}

func TestWriteConfigCreatesPrivateFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sub", ".hostinger.yaml")
	if err := writeConfig(path, setToken); err != nil {
		t.Fatal(err)
	}
	assertPrivate(t, path)
	entries, _ := os.ReadDir(filepath.Dir(path))
	if len(entries) != 1 {
		t.Fatalf("temp file left behind: %v", entries)
	}
}

func TestWriteConfigFollowsSymlink(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("symlinks need privileges on Windows")
	}
	dir := t.TempDir()
	target := filepath.Join(dir, "dotfiles.yaml")
	link := filepath.Join(dir, ".hostinger.yaml")
	if err := os.WriteFile(target, []byte("format: json\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(target, link); err != nil {
		t.Fatal(err)
	}

	if err := writeConfig(link, setToken); err != nil {
		t.Fatal(err)
	}

	if info, _ := os.Lstat(link); info.Mode()&os.ModeSymlink == 0 {
		t.Fatal("symlink was replaced by a regular file")
	}
	got, _ := os.ReadFile(target)
	if !strings.Contains(string(got), "api_token: secret") {
		t.Fatalf("target = %q", got)
	}
	assertPrivate(t, target)
}

func TestWriteConfigCreatesDanglingSymlinkTarget(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("symlinks need privileges on Windows")
	}
	dir := t.TempDir()
	link := filepath.Join(dir, ".hostinger.yaml")
	if err := os.Symlink("dotfiles.yaml", link); err != nil { // relative, target missing
		t.Fatal(err)
	}

	if err := writeConfig(link, setToken); err != nil {
		t.Fatal(err)
	}

	if info, _ := os.Lstat(link); info.Mode()&os.ModeSymlink == 0 {
		t.Fatal("symlink was replaced by a regular file")
	}
	assertPrivate(t, filepath.Join(dir, "dotfiles.yaml"))
}

func TestWriteConfigRejectsBadYAML(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".hostinger.yaml")
	if err := os.WriteFile(path, []byte("{"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := writeConfig(path, setToken); err == nil {
		t.Fatal("want error for malformed YAML")
	}
}

func TestWriteConfigFailsInUnwritableDir(t *testing.T) {
	if runtime.GOOS == "windows" || os.Geteuid() == 0 {
		t.Skip("needs POSIX permissions and a non-root user")
	}
	dir := t.TempDir()
	if err := os.Chmod(dir, 0o500); err != nil {
		t.Fatal(err)
	}
	defer os.Chmod(dir, 0o700)
	if err := writeConfig(filepath.Join(dir, ".hostinger.yaml"), setToken); err == nil {
		t.Fatal("want error when the directory is not writable")
	}
}
