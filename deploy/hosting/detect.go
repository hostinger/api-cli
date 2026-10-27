package hosting

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// serverDependencies mark a Node.js app rather than front-end tooling.
var serverDependencies = []string{"express", "fastify", "koa", "hono", "next", "nuxt", "@nestjs/core", "@remix-run/node", "@sveltejs/kit", "astro"}

// head returns the first 8 KB lower-cased: where WordPress reads plugin and theme headers.
func head(path string) string {
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()
	buf := make([]byte, 8192)
	n, _ := f.Read(buf)
	return strings.ToLower(string(buf[:n]))
}

// Detect classifies a directory by its top-level files. WordPress wins over
// package.json (plugins carry one for assets). note explains a fallback to static.
func Detect(dir string) (typ Type, note string, err error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", "", fmt.Errorf("reading %s: %w", dir, err)
	}
	hasFile := func(name string) bool {
		for _, e := range entries {
			if !e.IsDir() && e.Name() == name {
				return true
			}
		}
		return false
	}
	hasDir := func(name string) bool {
		for _, e := range entries {
			if e.IsDir() && e.Name() == name {
				return true
			}
		}
		return false
	}

	var wordpress []Type
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".php") && strings.Contains(head(filepath.Join(dir, e.Name())), "plugin name:") {
			wordpress = append(wordpress, WordPressPlugin)
			break
		}
	}
	if hasFile("style.css") && strings.Contains(head(filepath.Join(dir, "style.css")), "theme name:") {
		wordpress = append(wordpress, WordPressTheme)
	}
	if hasFile("wp-config.php") || hasDir("wp-content") {
		wordpress = append(wordpress, WordPressSite)
	}
	if len(wordpress) > 1 {
		return "", "", fmt.Errorf("ambiguous directory: it looks like %s; pass --type", joinTypes(wordpress))
	}
	if len(wordpress) == 1 {
		return wordpress[0], "", nil
	}

	if !hasFile("package.json") {
		return Static, "", nil
	}
	raw, err := os.ReadFile(filepath.Join(dir, "package.json"))
	if err != nil {
		return "", "", err
	}
	var pkg struct {
		Scripts      map[string]string `json:"scripts"`
		Dependencies map[string]string `json:"dependencies"`
	}
	if err := json.Unmarshal(raw, &pkg); err != nil {
		return Static, fmt.Sprintf("package.json could not be read (%v), so this is deployed as static files; pass --type nodejs if it is a Node.js app", err), nil
	}
	if pkg.Scripts["build"] != "" || pkg.Scripts["start"] != "" {
		return NodeJS, "", nil
	}
	for _, dep := range serverDependencies {
		if _, ok := pkg.Dependencies[dep]; ok {
			return NodeJS, "", nil
		}
	}
	return Static, "", nil
}

func joinTypes(ts []Type) string {
	parts := make([]string, len(ts))
	for i, t := range ts {
		parts[i] = string(t)
	}
	return strings.Join(parts, " and ")
}
