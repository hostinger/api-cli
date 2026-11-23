package preview

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/hostinger/api-cli/deploy/hosting"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/hcdn-connector/pkg/connector"
)

func TestFindProject(t *testing.T) {
	node := nodeFolder(t)
	if p := findProject(node); p == nil || p.link != nil || p.typ != hosting.NodeJS || p.target() != "" {
		t.Errorf("unlinked Node.js project = %+v", p)
	}
	// Any other unlinked folder would deploy as static files: not offered.
	plain := t.TempDir()
	writeFile(t, plain, "notes.txt", "hi")
	if p := findProject(plain); p != nil {
		t.Errorf("unlinked plain folder = %+v", p)
	}

	linked := linkedFolder(t, "example.com", "static")
	sub := filepath.Join(linked, "src")
	if err := os.Mkdir(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if p := findProject(sub); p == nil || p.dir != linked || p.typ != hosting.Static || p.target() != "example.com" {
		t.Errorf("from inside a linked project = %+v", p)
	}
	// hPanel is told the type the deploy will use, not the one remembered.
	became := linkedFolder(t, "example.com", "static")
	writeFile(t, became, "package.json", `{"scripts":{"build":"vite build"}}`)
	if p := findProject(became); p == nil || p.typ != hosting.NodeJS {
		t.Errorf("linked static folder that became a Node.js app = %+v", p)
	}
	wp := linkedFolder(t, "example.com", "wordpress-site")
	writeFile(t, wp, "wp-config.php", "<?php")
	if p := findProject(wp); p != nil {
		t.Errorf("linked WordPress site = %+v", p)
	}

	headless := t.TempDir()
	if err := site.Save(headless, &site.Project{Domain: "example.com", Headless: true}); err != nil {
		t.Fatal(err)
	}
	if p := findProject(headless); p == nil || p.typ != hosting.NodeJS {
		t.Errorf("linked headless website = %+v", p)
	}
}

func TestAppMeta(t *testing.T) {
	p := &project{dir: "/home/me/shop", link: &site.Project{Domain: "example.com"}, typ: hosting.Static}
	if got, want := p.appMeta(), `{"target":"example.com","type":"static","folder":"shop"}`; got != want {
		t.Errorf("appMeta = %s, want %s", got, want)
	}
	p = &project{dir: "/home/me/my app", typ: hosting.NodeJS}
	if got, want := p.appMeta(), `{"type":"nodejs","folder":"my app"}`; got != want {
		t.Errorf("unlinked appMeta = %s, want %s", got, want)
	}
}

// The edge drops a hello whose metadata is not printable ASCII within the
// cap, and the connector refuses to start with one.
func TestAppMetaIsPrintableASCIIWithinTheCap(t *testing.T) {
	for _, folder := range []string{
		"Café 🚀\x7f\t",
		strings.Repeat("ž", 300),
		strings.Repeat("🚀", 200),
	} {
		got := (&project{dir: "/x/" + folder, link: &site.Project{Domain: strings.Repeat("a", 240) + ".com"}, typ: hosting.NodeJS}).appMeta()
		if len(got) > connector.MaxAppMetaBytes {
			t.Errorf("%d bytes", len(got))
		}
		for _, c := range []byte(got) {
			if c < 0x20 || c > 0x7e {
				t.Fatalf("byte %#x in %q", c, got)
			}
		}
		var m struct{ Folder string }
		if err := json.Unmarshal([]byte(got), &m); err != nil || !strings.HasPrefix(folder, m.Folder) || !shownAsIs(m.Folder) {
			t.Errorf("folder %q decodes to %q (%v)", folder, m.Folder, err)
		}
	}
}

// hPanel's listing drops the whole metadata, the target with it, for a
// folder name outside the rule below, so the CLI never sends one.
func TestFolderNameKeepsTheTarget(t *testing.T) {
	for folder, want := range map[string]string{
		"shop":                               "shop",
		"my app":                             "my app",
		"Café 🚀":                             "Café 🚀",
		" padded\t":                          "padded",
		"zero\u200bwidth\u202e":              "zerowidth",
		"no\u00a0break":                      "no break",
		"Tom's site":                         "",
		"R&D":                                "",
		"<b>":                                "",
		"..":                                 "",
		strings.Repeat("ž", 127) + " end":    strings.Repeat("ž", 127),
		strings.Repeat("🚀", 100):             strings.Repeat("🚀", 63),
		"." + strings.Repeat(" ", 300) + "x": "",
		"bad\xffname":                        "bad\ufffdname",
	} {
		got := folderName(folder, maxFolderBytes)
		if got != want {
			t.Errorf("folderName(%q) = %q, want %q", folder, got, want)
		}
		if got != "" && !shownAsIs(got) {
			t.Errorf("folderName(%q) = %q, which the listing drops", folder, got)
		}
	}
}

// shownAsIs is the listing's rule for a folder name.
func shownAsIs(folder string) bool {
	if len(folder) > 255 || !utf8.ValidString(folder) || folder == "." || folder == ".." ||
		strings.TrimSpace(folder) != folder {
		return false
	}
	return !strings.ContainsFunc(folder, func(r rune) bool {
		return strings.ContainsRune(`/\<>"'&`+"`", r) || !unicode.IsPrint(r)
	})
}
