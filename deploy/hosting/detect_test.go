package hosting

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func write(t *testing.T, dir, name, content string) {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestDetect(t *testing.T) {
	cases := []struct {
		name  string
		files map[string]string
		want  Type
	}{
		{"plain html", map[string]string{"index.html": "<h1>hi</h1>"}, Static},
		{"php", map[string]string{"index.php": "<?php echo 1;"}, Static},
		{"package.json without scripts", map[string]string{"index.html": "", "package.json": `{"devDependencies":{"prettier":"1"}}`}, Static},
		{"build script", map[string]string{"package.json": `{"scripts":{"build":"vite build"}}`}, NodeJS},
		{"start script", map[string]string{"package.json": `{"scripts":{"start":"node server.js"}}`}, NodeJS},
		{"server dependency", map[string]string{"package.json": `{"dependencies":{"express":"4"}}`}, NodeJS},
		{"wordpress plugin beats package.json", map[string]string{"my-plugin.php": "<?php\n/*\nPlugin Name: Mine\n*/", "package.json": `{"scripts":{"build":"x"}}`}, WordPressPlugin},
		{"wordpress theme", map[string]string{"style.css": "/*\nTheme Name: Mine\n*/", "index.php": ""}, WordPressTheme},
		{"wordpress site", map[string]string{"wp-config.php": "<?php", "wp-content/index.php": ""}, WordPressSite},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			for name, content := range tc.files {
				write(t, dir, name, content)
			}
			got, _, err := Detect(dir)
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Errorf("Detect = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestDetectAmbiguousAndBroken(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, "style.css", "/* Theme Name: x */")
	write(t, dir, "wp-config.php", "")
	if _, _, err := Detect(dir); err == nil {
		t.Error("two WordPress signals should be ambiguous")
	}

	// A package.json that cannot be read says nothing: static, with a note.
	dir = t.TempDir()
	write(t, dir, "package.json", "{not json")
	typ, note, err := Detect(dir)
	if err != nil || typ != Static || !strings.Contains(note, "package.json could not be read") || !strings.Contains(note, "--type nodejs") {
		t.Errorf("broken package.json: %s %q %v", typ, note, err)
	}
}
