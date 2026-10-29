package site

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-git/go-git/v5"
)

// trackedRepo inits a repository at root and force-adds paths (relative to root).
func trackedRepo(t *testing.T, root string, paths ...string) {
	t.Helper()
	repo, err := git.PlainInit(root, false)
	if err != nil {
		t.Fatal(err)
	}
	wt, err := repo.Worktree()
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range paths {
		if _, err := wt.Add(p); err != nil {
			t.Fatal(err)
		}
	}
}

func writeLink(t *testing.T, dir, name string) string {
	t.Helper()
	path := filepath.Join(dir, Dir, name)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(`{"domain":"example.com"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestCommittedLink(t *testing.T) {
	// Not a repository.
	plain := t.TempDir()
	writeLink(t, plain, File)
	if got := CommittedLink(plain); got != "" {
		t.Errorf("outside Git = %q", got)
	}

	// Untracked, as `hostinger link` leaves it.
	untracked := t.TempDir()
	trackedRepo(t, untracked)
	writeLink(t, untracked, File)
	if got := CommittedLink(untracked); got != "" {
		t.Errorf("untracked = %q", got)
	}

	// Tracked.
	tracked := t.TempDir()
	path := writeLink(t, tracked, File)
	trackedRepo(t, tracked, filepath.Join(Dir, File))
	if got := CommittedLink(tracked); got != path {
		t.Errorf("tracked = %q, want %q", got, path)
	}

	// Tracked in a subfolder of the repository, found from there.
	root := t.TempDir()
	sub := filepath.Join(root, "app")
	path = writeLink(t, sub, File)
	trackedRepo(t, root, filepath.Join("app", Dir, File))
	if got := CommittedLink(sub); got != path {
		t.Errorf("tracked in subfolder = %q, want %q", got, path)
	}

	// The legacy file, read when site.json is missing.
	legacy := t.TempDir()
	path = writeLink(t, legacy, legacyFile)
	trackedRepo(t, legacy, filepath.Join(Dir, legacyFile))
	if got := CommittedLink(legacy); got != path {
		t.Errorf("tracked legacy = %q, want %q", got, path)
	}
}

func TestCommittedLinkUnreadableRepositoryCounts(t *testing.T) {
	// A corrupt index.
	broken := t.TempDir()
	trackedRepo(t, broken)
	path := writeLink(t, broken, File)
	if err := os.WriteFile(filepath.Join(broken, ".git", "index"), []byte("garbage"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := CommittedLink(broken); got != path {
		t.Errorf("corrupt index = %q", got)
	}

	// A configuration that cannot be read.
	bad := t.TempDir()
	trackedRepo(t, bad)
	path = writeLink(t, bad, File)
	if err := os.WriteFile(filepath.Join(bad, ".git", "config"), []byte("[core\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := CommittedLink(bad); got != path {
		t.Errorf("unreadable config = %q", got)
	}
}

func TestCommittedLinkThroughSymlinks(t *testing.T) {
	// A tracked folder behind a symlinked .hostinger is still tracked.
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "real"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "real", File), []byte(`{"domain":"example.com"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	symlinkOrSkip(t, "real", filepath.Join(root, Dir))
	trackedRepo(t, root, filepath.Join("real", File))
	if got := CommittedLink(root); got != Path(root) {
		t.Errorf("symlinked dir = %q", got)
	}

	// A link resolving outside the repository, or nowhere, counts.
	outside := t.TempDir()
	writeLink(t, outside, File)
	escape := t.TempDir()
	trackedRepo(t, escape)
	symlinkOrSkip(t, filepath.Join(outside, Dir), filepath.Join(escape, Dir))
	if got := CommittedLink(escape); got != Path(escape) {
		t.Errorf("escaping = %q", got)
	}
	dangling := t.TempDir()
	trackedRepo(t, dangling)
	if err := os.MkdirAll(filepath.Join(dangling, Dir), 0o755); err != nil {
		t.Fatal(err)
	}
	symlinkOrSkip(t, "missing.json", Path(dangling))
	if got := CommittedLink(dangling); got != Path(dangling) {
		t.Errorf("dangling = %q", got)
	}
}

func TestCommittedLinkMatchesAnyCase(t *testing.T) {
	// The index names the file in another case than the one on disk, as a
	// repository committing .HOSTINGER/SITE.json looks on macOS and Windows.
	for _, name := range []string{".hostinger/SITE.json", ".HOSTINGER/site.json", ".Hostinger/Site.Json"} {
		dir := t.TempDir()
		path := writeLink(t, dir, File)
		trackedRepo(t, dir, filepath.Join(Dir, File))
		repo, err := git.PlainOpen(dir)
		if err != nil {
			t.Fatal(err)
		}
		idx, err := repo.Storer.Index()
		if err != nil {
			t.Fatal(err)
		}
		idx.Entries[0].Name = name
		if err := repo.Storer.SetIndex(idx); err != nil {
			t.Fatal(err)
		}
		if got := CommittedLink(dir); got != path {
			t.Errorf("%s = %q, want %q", name, got, path)
		}
	}
}
