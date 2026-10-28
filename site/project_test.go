package site

import (
	"encoding/json"
	"os"
	"path/filepath"
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
