package site

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestSaveFindLoad(t *testing.T) {
	root := t.TempDir()
	nested := filepath.Join(root, "a", "b")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	if p, dir, err := Find(nested); err != nil || p != nil || dir != "" {
		t.Fatalf("Find before Save = %v %q %v", p, dir, err)
	}
	if err := Save(root, &Project{OrderID: 7, Domain: "example.com", Username: "u1", Type: "static"}); err != nil {
		t.Fatal(err)
	}
	p, dir, err := Find(nested)
	if err != nil {
		t.Fatal(err)
	}
	if dir != root || p == nil || p.Domain != "example.com" || p.Username != "u1" || p.OrderID != 7 {
		t.Errorf("Find = %+v in %q", p, dir)
	}
	info, _ := os.Stat(Path(root))
	if info.Mode().Perm() != 0o600 {
		t.Errorf("project file mode = %o", info.Mode().Perm())
	}
	// Not a git checkout: no .gitignore is created.
	if _, err := os.Stat(filepath.Join(root, ".gitignore")); err == nil {
		t.Error(".gitignore created outside git")
	}
}

func TestSaveAddsGitignoreEntry(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, ".gitignore"), []byte("node_modules"), 0o644); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		if err := Save(root, &Project{Domain: "example.com"}); err != nil {
			t.Fatal(err)
		}
	}
	raw, _ := os.ReadFile(filepath.Join(root, ".gitignore"))
	if got := string(raw); got != "node_modules\n.hostinger/\n" {
		t.Errorf(".gitignore = %q", got)
	}

	// A subfolder of the checkout gets its own .gitignore, created if missing.
	sub := filepath.Join(root, "apps", "web")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := Save(sub, &Project{Domain: "example.com"}); err != nil {
		t.Fatal(err)
	}
	raw, _ = os.ReadFile(filepath.Join(sub, ".gitignore"))
	if string(raw) != ".hostinger/\n" {
		t.Errorf("subfolder .gitignore = %q", raw)
	}
}

func TestLoadIgnoresEmptyFile(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, Dir), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(Path(root), []byte(`{}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if p, err := Load(root); err != nil || p != nil {
		t.Errorf("Load = %+v %v", p, err)
	}
	if err := os.WriteFile(Path(root), []byte(`{broken`), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(root); err == nil || !strings.Contains(err.Error(), "invalid") {
		t.Errorf("broken file err = %v", err)
	}
}

func TestLoadRejectsInvalidDomain(t *testing.T) {
	label := strings.Repeat("a", 63)
	long := strings.Join([]string{label, label, label, label}, ".") // 255
	cases := []struct {
		domain string
		ok     bool
		want   string
	}{
		{"example.com", true, "example.com"},
		{"Shop-1.Example.COM", true, "Shop-1.Example.COM"},
		{"xn--bcher-kva.example", true, "xn--bcher-kva.example"},
		{"bücher.example", true, "bücher.example"},
		{"e\u0301cole.fr", true, "e\u0301cole.fr"},
		{strings.Repeat("ü", 63) + ".lt", true, strings.Repeat("ü", 63) + ".lt"},
		{"example.com.", true, "example.com"},
		{"localhost", true, "localhost"},
		{label + ".com", true, label + ".com"},
		{"a\u001b[2Jb.com", false, ""},
		{"evil.com\u0007", false, ""},
		{"a b.com", false, ""},
		{"-a.com", false, ""},
		{"a-.com", false, ""},
		{"a..com", false, ""},
		{".com", false, ""},
		{"a\u202eb.com", false, ""},
		{"a/b.com", false, ""},
		{"a:b.com", false, ""},
		{"bü-.lt", false, ""},
		{strings.Repeat("ü", 64) + ".lt", false, ""},
		{label + "a.com", false, ""},
		{long, false, ""},
	}
	for _, c := range cases {
		root := t.TempDir()
		if err := os.MkdirAll(filepath.Join(root, Dir), 0o755); err != nil {
			t.Fatal(err)
		}
		raw, _ := json.Marshal(Project{Domain: c.domain}) // escapes control characters, so they reach Load
		if err := os.WriteFile(Path(root), raw, 0o644); err != nil {
			t.Fatal(err)
		}
		p, err := Load(root)
		if c.ok && (err != nil || p == nil || p.Domain != c.want) {
			t.Errorf("%q: Load = %+v, %v", c.domain, p, err)
		}
		if !c.ok && (err == nil || p != nil || !strings.Contains(err.Error(), Path(root)) || strings.ContainsRune(err.Error(), 0x1b)) {
			t.Errorf("%q: Load = %+v, %v; want an error naming the file", c.domain, p, err)
		}
	}

	// The legacy file is named when it is the one read.
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, Dir), 0o755); err != nil {
		t.Fatal(err)
	}
	legacy := filepath.Join(root, Dir, legacyFile)
	if err := os.WriteFile(legacy, []byte(`{"domain":"a b"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(root); err == nil || !strings.Contains(err.Error(), legacy) {
		t.Errorf("legacy err = %v", err)
	}
}

func TestSaveKeepsDomainAsGiven(t *testing.T) {
	for in, want := range map[string]string{"bücher.lt": "bücher.lt", "Example.com.": "Example.com", "xn--bcher-kva.lt": "xn--bcher-kva.lt"} {
		root := t.TempDir()
		p := &Project{Kind: "website", Domain: in}
		if err := Save(root, p); err != nil {
			t.Fatalf("%q: %v", in, err)
		}
		if p.Domain != in {
			t.Errorf("Save changed the caller's domain to %q", p.Domain)
		}
		got, err := Load(root)
		if err != nil || got == nil || got.Domain != want {
			t.Errorf("%q: Load = %+v, %v; want %q", in, got, err, want)
		}
	}
}

func TestSaveRefusesInvalidDomain(t *testing.T) {
	for _, d := range []string{"a\u001b[2Jb.com", "evil.com\u0007", "a b.com", "-a.com", "a..com", ".", strings.Repeat("a", 64) + ".com", strings.Repeat("a.", 127) + "aa"} {
		root := t.TempDir()
		err := Save(root, &Project{Kind: "website", Domain: d})
		if err == nil || !strings.Contains(err.Error(), "not a valid domain") || strings.ContainsRune(err.Error(), 0x1b) {
			t.Errorf("%q: Save err = %v", d, err)
		}
		if _, err := os.Stat(Path(root)); err == nil {
			t.Errorf("%q: link written", d)
		}
	}
}

func TestSaveEmptyDomainIsNoLink(t *testing.T) {
	root := t.TempDir()
	if err := Save(root, &Project{Kind: "website"}); err != nil {
		t.Fatal(err)
	}
	if p, err := Load(root); err != nil || p != nil {
		t.Errorf("Load = %+v, %v; want no link", p, err)
	}
}

func TestSaveMergesAgentFieldsAndReadsLegacy(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, Dir), 0o755); err != nil {
		t.Fatal(err)
	}
	// An agent following the deploy skill wrote site.json first, with its own fields.
	agent := `{"domain":"example.com","username":"u1","type":"php","urls":{"store_api":"https://x"}}`
	if err := os.WriteFile(Path(root), []byte(agent), 0o644); err != nil {
		t.Fatal(err)
	}
	p, err := Load(root)
	if err != nil || p == nil || p.Type != "static" {
		t.Fatalf("Load = %+v %v", p, err)
	}
	p.Kind, p.OrderID, p.Source = "website", 7, "upload"
	if err := Save(root, p); err != nil {
		t.Fatal(err)
	}
	raw, _ := os.ReadFile(Path(root))
	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got["urls"] == nil || got["orderId"] != float64(7) || got["domain"] != "example.com" {
		t.Errorf("merged = %s", raw)
	}
	// Linked to another website: the agent's fields described the old one.
	if err := Save(root, &Project{Kind: "website", Domain: "other.com"}); err != nil {
		t.Fatal(err)
	}
	raw, _ = os.ReadFile(Path(root))
	got = nil
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got["urls"] != nil || got["domain"] != "other.com" {
		t.Errorf("after re-link to another website = %s", raw)
	}

	// A link written by an early build as project.json is still read, and replaced on save.
	legacy := t.TempDir()
	if err := os.MkdirAll(filepath.Join(legacy, Dir), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(legacy, Dir, legacyFile), []byte(`{"domain":"old.com"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	p, err = Load(legacy)
	if err != nil || p == nil || p.Domain != "old.com" {
		t.Fatalf("legacy Load = %+v %v", p, err)
	}
	if err := Save(legacy, p); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(legacy, Dir, legacyFile)); err == nil {
		t.Error("legacy file kept after save")
	}
	if _, err := os.Stat(Path(legacy)); err != nil {
		t.Error("site.json not written")
	}
}

func TestInteractiveHonoursCIAndFlag(t *testing.T) {
	t.Setenv("HOSTINGER_NONINTERACTIVE", "")
	t.Setenv("CI", "true")
	if Interactive() {
		t.Error("CI=true should disable prompts")
	}
	t.Setenv("CI", "false")
	viper.Set("no_input", true)
	t.Cleanup(func() { viper.Set("no_input", false) })
	if Interactive() {
		t.Error("--no-input (no_input) should disable prompts")
	}
}

func TestLinkSearchStaysInTheProject(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	// A link made once in the home directory does not apply to projects below it.
	if err := Save(home, &Project{Domain: "home.example.com"}); err != nil {
		t.Fatal(err)
	}
	proj := filepath.Join(home, "code", "site")
	if err := os.MkdirAll(proj, 0o755); err != nil {
		t.Fatal(err)
	}
	if p, _, err := Find(proj); err != nil || p != nil {
		t.Errorf("home link used: %+v %v", p, err)
	}
	// Nor does a link above the repository root.
	if err := Save(filepath.Join(home, "code"), &Project{Domain: "parent.example.com"}); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(proj, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	sub := filepath.Join(proj, "src")
	if err := os.Mkdir(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if p, _, err := Find(sub); err != nil || p != nil {
		t.Errorf("link above the repository used: %+v %v", p, err)
	}
	// A link at the repository root applies to its subfolders.
	if err := Save(proj, &Project{Domain: "site.example.com"}); err != nil {
		t.Fatal(err)
	}
	if p, dir, err := Find(sub); err != nil || p == nil || p.Domain != "site.example.com" || dir != proj {
		t.Errorf("repository link: %+v %s %v", p, dir, err)
	}
}

func TestHeadlessLinkIsShown(t *testing.T) {
	if !(Website{WebsiteType: "headless"}).IsHeadless() || (Website{WebsiteType: "nodejs"}).IsHeadless() {
		t.Error("IsHeadless")
	}
	c := &Context{Domain: "cms.example.com", DomainSource: "link", Link: &Project{Domain: "cms.example.com", Type: "nodejs", Headless: true}}
	if lines := strings.Join(c.Lines(), "\n"); !strings.Contains(lines, "nodejs (headless website)") {
		t.Errorf("lines:\n%s", lines)
	}
}

func symlinkOrSkip(t *testing.T, target, link string) {
	t.Helper()
	if err := os.Symlink(target, link); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}
}

func TestSaveRefusesSymlinkedFiles(t *testing.T) {
	outside := t.TempDir()
	victim := filepath.Join(outside, "bashrc")
	if err := os.WriteFile(victim, []byte("export PATH=/usr/bin\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	// A cloned checkout ships .hostinger/site.json as a link out of it.
	repo := t.TempDir()
	if err := os.MkdirAll(filepath.Join(repo, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(repo, Dir), 0o755); err != nil {
		t.Fatal(err)
	}
	symlinkOrSkip(t, victim, Path(repo))
	if err := Save(repo, &Project{Domain: "example.com"}); err == nil || !strings.Contains(err.Error(), "is a symlink") {
		t.Errorf("symlinked site.json err = %v", err)
	}

	// Or .gitignore.
	repo = t.TempDir()
	if err := os.MkdirAll(filepath.Join(repo, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	symlinkOrSkip(t, victim, filepath.Join(repo, ".gitignore"))
	if err := Save(repo, &Project{Domain: "example.com"}); err == nil || !strings.Contains(err.Error(), "is a symlink") {
		t.Errorf("symlinked .gitignore err = %v", err)
	}

	if raw, _ := os.ReadFile(victim); string(raw) != "export PATH=/usr/bin\n" {
		t.Errorf("outside file changed: %q", raw)
	}
}

func TestSaveSymlinkedLinkDir(t *testing.T) {
	// Pointing outside the checkout: nothing is written or removed there.
	outside := t.TempDir()
	legacy := filepath.Join(outside, legacyFile)
	if err := os.WriteFile(legacy, []byte(`{"domain":"keep.com"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	repo := t.TempDir()
	if err := os.Mkdir(filepath.Join(repo, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	symlinkOrSkip(t, outside, filepath.Join(repo, Dir))
	if err := Save(repo, &Project{Domain: "example.com"}); err == nil || !strings.Contains(err.Error(), "links outside") {
		t.Errorf("outside link dir err = %v", err)
	}
	if _, err := os.Stat(filepath.Join(outside, File)); err == nil {
		t.Error("site.json written outside the checkout")
	}
	if _, err := os.Stat(legacy); err != nil {
		t.Error("legacy file outside the checkout removed")
	}

	// Outside dir, but no checkout either: the project is the directory itself.
	plain := t.TempDir()
	symlinkOrSkip(t, outside, filepath.Join(plain, Dir))
	if err := Save(plain, &Project{Domain: "example.com"}); err == nil || !strings.Contains(err.Error(), "links outside") {
		t.Errorf("outside link dir without git err = %v", err)
	}

	// A dangling link.
	dangling := t.TempDir()
	symlinkOrSkip(t, filepath.Join(dangling, "missing"), filepath.Join(dangling, Dir))
	if err := Save(dangling, &Project{Domain: "example.com"}); err == nil || !strings.Contains(err.Error(), "refusing") {
		t.Errorf("dangling link dir err = %v", err)
	}

	// A folder shared inside the checkout still works, from a subfolder too.
	shared := filepath.Join(repo, "shared")
	app := filepath.Join(repo, "apps", "web")
	for _, d := range []string{shared, app} {
		if err := os.MkdirAll(d, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	symlinkOrSkip(t, filepath.Join("..", "..", "shared"), filepath.Join(app, Dir))
	if err := Save(app, &Project{Domain: "example.com"}); err != nil {
		t.Fatal(err)
	}
	if p, err := Load(app); err != nil || p == nil || p.Domain != "example.com" {
		t.Errorf("Load through shared dir = %+v %v", p, err)
	}
	if _, err := os.Stat(filepath.Join(shared, File)); err != nil {
		t.Error("site.json not written into the shared folder")
	}
}

func TestWriteFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "f")
	if err := writeFile(path, []byte("a"), 0o600); err != nil {
		t.Fatal(err)
	}
	if info, _ := os.Stat(path); info.Mode().Perm() != 0o600 {
		t.Errorf("new file mode = %o", info.Mode().Perm())
	}
	// An existing file keeps its mode and is replaced, with no temp file left.
	if err := os.Chmod(path, 0o640); err != nil {
		t.Fatal(err)
	}
	if err := writeFile(path, []byte("b"), 0o600); err != nil {
		t.Fatal(err)
	}
	if info, _ := os.Stat(path); info.Mode().Perm() != 0o640 && runtime.GOOS != "windows" {
		t.Errorf("existing file mode = %o", info.Mode().Perm())
	}
	if raw, _ := os.ReadFile(path); string(raw) != "b" {
		t.Errorf("content = %q", raw)
	}
	if entries, _ := os.ReadDir(dir); len(entries) != 1 {
		t.Errorf("left behind: %v", entries)
	}

	if err := writeFile(dir, []byte("x"), 0o600); err == nil || !strings.Contains(err.Error(), "not a regular file") {
		t.Errorf("directory err = %v", err)
	}
	if err := writeFile(filepath.Join(path, "below"), []byte("x"), 0o600); err == nil {
		t.Error("path below a file: no error")
	}
	if err := writeFile(filepath.Join(dir, "missing", "f"), []byte("x"), 0o600); err == nil {
		t.Error("missing folder: no error")
	}
}

func TestRemove(t *testing.T) {
	write := func(dir string, names ...string) {
		t.Helper()
		for _, name := range names {
			if err := os.WriteFile(filepath.Join(dir, name), []byte(`{"domain":"example.com"}`), 0o644); err != nil {
				t.Fatal(err)
			}
		}
	}

	// A plain folder: the link files go, then the folder once empty.
	plain := t.TempDir()
	if err := os.Mkdir(filepath.Join(plain, Dir), 0o755); err != nil {
		t.Fatal(err)
	}
	write(filepath.Join(plain, Dir), File, legacyFile)
	if err := Remove(plain); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Lstat(filepath.Join(plain, Dir)); err == nil {
		t.Error("empty link folder kept")
	}
	// Nothing linked is not an error.
	if err := Remove(plain); err != nil {
		t.Errorf("missing link err = %v", err)
	}

	// Other tools' files keep the folder.
	other := t.TempDir()
	if err := os.Mkdir(filepath.Join(other, Dir), 0o755); err != nil {
		t.Fatal(err)
	}
	write(filepath.Join(other, Dir), File, "workspace.json")
	if err := Remove(other); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(other, Dir, "workspace.json")); err != nil {
		t.Error("other tool's file removed")
	}

	// A folder that cannot be read fails.
	bad := t.TempDir()
	if err := os.MkdirAll(filepath.Join(bad, Dir, File, "x"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := Remove(bad); err == nil {
		t.Error("non-empty site.json directory removed without error")
	}
}

func TestRemoveSymlinks(t *testing.T) {
	// A symlinked folder outside the checkout: refused, nothing removed there.
	outside := t.TempDir()
	for _, name := range []string{File, legacyFile} {
		if err := os.WriteFile(filepath.Join(outside, name), []byte("keep"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	repo := t.TempDir()
	if err := os.Mkdir(filepath.Join(repo, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	symlinkOrSkip(t, outside, filepath.Join(repo, Dir))
	if err := Remove(repo); err == nil || !strings.Contains(err.Error(), "links outside") {
		t.Errorf("outside link dir err = %v", err)
	}
	for _, name := range []string{File, legacyFile} {
		if _, err := os.Stat(filepath.Join(outside, name)); err != nil {
			t.Errorf("%s outside the checkout removed", name)
		}
	}

	// A folder shared inside the checkout: files go, the link stays.
	shared := filepath.Join(repo, "shared")
	app := filepath.Join(repo, "apps", "web")
	for _, d := range []string{shared, app} {
		if err := os.MkdirAll(d, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	for _, name := range []string{File, legacyFile} {
		if err := os.WriteFile(filepath.Join(shared, name), []byte(`{"domain":"example.com"}`), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	symlinkOrSkip(t, filepath.Join("..", "..", "shared"), filepath.Join(app, Dir))
	if err := Remove(app); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{File, legacyFile} {
		if _, err := os.Stat(filepath.Join(shared, name)); err == nil {
			t.Errorf("%s left in the shared folder", name)
		}
	}
	if info, err := os.Lstat(filepath.Join(app, Dir)); err != nil || info.Mode()&os.ModeSymlink == 0 {
		t.Errorf("shared link folder not kept: %v", err)
	}

	// A symlinked site.json is removed as a link; its target stays.
	victim := filepath.Join(outside, File)
	linked := t.TempDir()
	if err := os.Mkdir(filepath.Join(linked, Dir), 0o755); err != nil {
		t.Fatal(err)
	}
	symlinkOrSkip(t, victim, Path(linked))
	if err := Remove(linked); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Lstat(Path(linked)); err == nil {
		t.Error("symlinked site.json kept")
	}
	if raw, _ := os.ReadFile(victim); string(raw) != "keep" {
		t.Errorf("link target changed: %q", raw)
	}
}
