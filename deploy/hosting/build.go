package hosting

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
)

// BuildError is a failed Node.js build with the server's logs and analysis.
type BuildError struct {
	UUID     string
	Logs     string
	Analysis string
	Solution string
}

func (e *BuildError) Error() string {
	msg := "build " + e.UUID + " failed"
	if e.Analysis != "" {
		msg += ": " + e.Analysis
	}
	if e.Solution != "" {
		msg += "\nSuggested fix: " + e.Solution
	}
	return msg
}

// followHint tells how to follow a running build; a linked directory needs no args.
func followHint(domain, uuid string, linked bool) string {
	if linked {
		return "check the build with `hostinger hosting nodejs build` and read its log with `hostinger hosting nodejs build-logs` (this directory remembers it)"
	}
	return "check the build with `hostinger hosting nodejs build " + domain + " " + uuid + "` and read its log with `hostinger hosting nodejs build-logs " + domain + " " + uuid + "`"
}

// checkNodeOverrides rejects build-setting flags the build API would refuse.
func checkNodeOverrides(n NodeOverrides) error {
	if n.NodeVersion != 0 && !slices.Contains(NodeVersions, n.NodeVersion) {
		versions := make([]string, len(NodeVersions))
		for i, v := range NodeVersions {
			versions[i] = strconv.Itoa(v)
		}
		return fmt.Errorf("--node-version %d is not offered (one of: %s)", n.NodeVersion, strings.Join(versions, ", "))
	}
	if n.AppType != "" && !slices.Contains(AppTypes, n.AppType) {
		return fmt.Errorf("--app-type %s is not known (one of: %s)", n.AppType, strings.Join(AppTypes, ", "))
	}
	if n.PackageManager != "" && !slices.Contains(PackageManagers, n.PackageManager) {
		return fmt.Errorf("--package-manager %s is not known (one of: %s)", n.PackageManager, strings.Join(PackageManagers, ", "))
	}
	return nil
}

// runBuild builds the uploaded archive with detected settings, then remembered
// ones, then flags; a first deploy or --settings reviews them.
func runBuild(ctx context.Context, a API, ui UI, w *site.Website, detected NodeSettings, start func(NodeSettings) (Build, error), opts Options, remembered *site.NodeSettings, packed *Listing) (Build, *NodeSettings, error) {
	settings := detected
	var err error
	if remembered != nil {
		settings = applyOverrides(settings, *remembered)
	}
	settings = applyOverrides(settings, opts.Node)
	if settings.NodeVersion == 0 {
		settings.NodeVersion = 22
	}
	if settings.RootDirectory == "" {
		settings.RootDirectory = "."
	}

	validate := func(NodeSettings) []string { return nil }
	if packed != nil {
		validate = func(s NodeSettings) []string { return validateSettings(s, packed) }
	}
	review := ui.Interactive() && (opts.Settings || remembered == nil)
	if !review {
		if problems := validate(settings); len(problems) > 0 {
			if !ui.Interactive() {
				return Build{}, nil, errors.New("the build settings do not match the project: " + strings.Join(problems, "; ") + ". Pass --settings in a terminal, or the matching flags")
			}
			ui.Info("The remembered build settings no longer match the project: " + strings.Join(problems, "; "))
			review = true
		} else if remembered != nil && !opts.Settings {
			ui.Info("Using the remembered build settings (" + summarize(settings) + "); pass --settings to review them")
		}
	}
	if review {
		settings, err = reviewSettings(ui, settings, validate)
		if err != nil {
			return Build{}, nil, err
		}
	}
	if settings.AppType == "" {
		return Build{}, nil, errors.New("the framework could not be detected: pass --app-type (for example express, next or vite)")
	}

	for attempt := 1; ; attempt++ {
		build, err := startAndWait(ctx, a, ui, w, start, settings, opts.NoWait)
		if err == nil && opts.NoWait {
			return build, nil, nil // not finished: settings are not remembered yet
		}
		var berr *BuildError
		if !errors.As(err, &berr) || !ui.Interactive() || attempt >= 3 {
			if err != nil {
				return build, nil, err
			}
			return build, &settings, nil
		}
		ui.Info(berr.Error())
		again, cerr := ui.Confirm("Review the build settings and build again?", true)
		if cerr != nil || !again {
			return build, nil, err
		}
		settings, err = reviewSettings(ui, settings, validate)
		if err != nil {
			return build, nil, err
		}
	}
}

func startAndWait(ctx context.Context, a API, ui UI, w *site.Website, start func(NodeSettings) (Build, error), settings NodeSettings, noWait bool) (Build, error) {
	step := ui.Start("Starting build")
	build, err := start(settings)
	if err != nil {
		step.Fail("Could not start build")
		return Build{}, err
	}
	if noWait {
		step.Done("Build " + build.UUID + " started")
		return build, nil
	}
	step.Update("Building (" + build.UUID + ")")
	interval := BuildPollInterval
	deadline := Now().Add(BuildTimeout)
	failures := 0
	for build.State != "completed" && build.State != "failed" {
		select {
		case <-ctx.Done():
			step.Fail("Interrupted; the build keeps running on the server")
			return build, ctx.Err()
		case <-time.After(interval):
		}
		if interval < BuildPollMax {
			interval = min(interval*3/2, BuildPollMax)
		}
		next, err := a.BuildDetails(ctx, w.Username, w.Domain, build.UUID)
		if err != nil {
			// One failed poll does not end a wait that can last minutes.
			if failures++; failures < 3 {
				continue
			}
			step.Fail("Could not read build status")
			return build, fmt.Errorf("%w; the build may still be running: follow it with `hostinger hosting nodejs build %s %s`", err, w.Domain, build.UUID)
		}
		failures = 0
		build = next
		if build.State != "completed" && build.State != "failed" && Now().After(deadline) {
			step.Fail("Timed out")
			return build, fmt.Errorf("the build %s is still running; follow it with `hostinger hosting nodejs build %s %s`", build.UUID, w.Domain, build.UUID)
		}
	}
	if build.State == "failed" {
		step.Fail("Build failed")
		berr := &BuildError{UUID: build.UUID}
		berr.Logs, _ = a.BuildLogs(ctx, w.Username, w.Domain, build.UUID)
		analysis, solution, _ := a.BuildAnalysis(ctx, w.Username, w.Domain, build.UUID)
		berr.Analysis, berr.Solution = PlainText(analysis), PlainText(solution)
		return build, berr
	}
	step.Done("Build completed")
	return build, nil
}

// validateSettings catches bad settings against the packed files before a build.
func validateSettings(s NodeSettings, packed *Listing) []string {
	var problems []string
	root := strings.Trim(strings.TrimPrefix(filepath.ToSlash(s.RootDirectory), "./"), "/")
	if root == "." {
		root = ""
	}
	inRoot := func(rel string) string {
		if root == "" {
			return rel
		}
		return root + "/" + rel
	}
	if escapes(root) || strings.HasPrefix(s.RootDirectory, "/") {
		problems = append(problems, "root directory must be a relative path inside the project")
	} else if !packed.Has(inRoot("package.json")) {
		if root == "" {
			problems = append(problems, "package.json is not in the project root")
		} else {
			problems = append(problems, "root directory "+s.RootDirectory+" has no package.json")
		}
	}
	// A build usually produces the entry file, so check it only without one.
	if s.EntryFile != "" && deref(s.BuildScript) == "" && !packed.Has(inRoot(strings.TrimPrefix(filepath.ToSlash(s.EntryFile), "./"))) {
		problems = append(problems, "entry file "+s.EntryFile+" does not exist")
	}
	if script := deref(s.BuildScript); script != "" {
		if scripts := packageScripts(packed, inRoot("package.json")); scripts != nil {
			if _, ok := scripts[script]; !ok {
				problems = append(problems, "build script "+script+" is not in package.json scripts (have: "+strings.Join(sortedKeys(scripts), ", ")+")")
			}
		}
	}
	if out := deref(s.OutputDirectory); out != "" && (strings.HasPrefix(out, "/") || escapes(out)) {
		problems = append(problems, "output directory must be a relative path")
	}
	return problems
}

// escapes reports a ".." path segment; "my..app" is a fine name.
func escapes(p string) bool {
	for _, seg := range strings.Split(filepath.ToSlash(p), "/") {
		if seg == ".." {
			return true
		}
	}
	return false
}

// packageScripts is nil when the package.json cannot be read.
func packageScripts(packed *Listing, rel string) map[string]string {
	if packed.Dir == "" || !packed.Has(rel) {
		return nil
	}
	raw, err := os.ReadFile(filepath.Join(packed.Dir, filepath.FromSlash(rel)))
	if err != nil {
		return nil
	}
	var pkg struct {
		Scripts map[string]string `json:"scripts"`
	}
	if json.Unmarshal(raw, &pkg) != nil {
		return nil
	}
	if pkg.Scripts == nil {
		return map[string]string{}
	}
	return pkg.Scripts
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

var htmlTag = regexp.MustCompile(`<[^>]+>`)

// PlainText renders the build analysis HTML for a terminal, without control characters.
func PlainText(html string) string {
	if html == "" {
		return ""
	}
	r := strings.NewReplacer(
		"<li>", "\n  - ", "</li>", "",
		"<ul>", "", "</ul>", "", "<ol>", "", "</ol>", "",
		"<br>", "\n", "<br/>", "\n", "<br />", "\n", "<p>", "\n", "</p>", "\n",
		"<code>", "`", "</code>", "`",
	)
	text := htmlTag.ReplaceAllString(r.Replace(html), "")
	text = strings.NewReplacer("&lt;", "<", "&gt;", ">", "&amp;", "&", "&quot;", `"`, "&#39;", "'", "&nbsp;", " ").Replace(text)
	var lines []string
	for _, line := range strings.Split(text, "\n") {
		if line = strings.TrimRight(line, " "); strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	return utils.SafeText(strings.Join(lines, "\n"))
}

// NodeVersions, AppTypes and PackageManagers are what the build API accepts.
var (
	NodeVersions    = []int{18, 20, 22, 24}
	AppTypes        = []string{"express", "next", "nuxt", "nest", "vite", "react", "create-react-app", "vue", "angular", "svelte", "svelte-kit", "astro", "gatsby", "parcel", "fastify", "hono", "react-router", "nitro", "other"}
	PackageManagers = []string{"npm", "yarn", "pnpm"}
)

// none clears a setting in the review.
const none = "-"

// reviewSettings asks again while validate finds problems.
func reviewSettings(ui UI, s NodeSettings, validate func(NodeSettings) []string) (NodeSettings, error) {
	for {
		edited, err := reviewOnce(ui, s)
		if err != nil {
			return s, err
		}
		problems := validate(edited)
		if len(problems) == 0 {
			return edited, nil
		}
		ui.Info("These settings would not build: " + strings.Join(problems, "; "))
		s = edited
	}
}

func reviewOnce(ui UI, s NodeSettings) (NodeSettings, error) {
	ui.Info("Build settings:")
	ui.Info("  Framework:         " + orNone(s.AppType))
	ui.Info("  Node.js version:   " + strconv.Itoa(s.NodeVersion))
	ui.Info("  Root directory:    " + orNone(s.RootDirectory))
	ui.Info("  Build script:      " + orNone(deref(s.BuildScript)))
	ui.Info("  Output directory:  " + orNone(deref(s.OutputDirectory)))
	ui.Info("  Entry file:        " + orNone(s.EntryFile))
	ui.Info("  Package manager:   " + orNone(s.PackageManager))
	if s.AppType == "" {
		ui.Info("The framework could not be detected; choose it.")
	} else {
		ok, err := ui.Confirm("Use these settings?", true)
		if err != nil || ok {
			return s, err
		}
	}

	idx, err := ui.Select("Framework", AppTypes, indexOf(AppTypes, s.AppType))
	if err != nil {
		return s, err
	}
	s.AppType = AppTypes[idx]

	versions := make([]string, len(NodeVersions))
	def := 0
	for i, v := range NodeVersions {
		versions[i] = strconv.Itoa(v)
		if v == s.NodeVersion {
			def = i
		}
	}
	if idx, err = ui.Select("Node.js version", versions, def); err != nil {
		return s, err
	}
	s.NodeVersion = NodeVersions[idx]

	if s.RootDirectory, err = ui.Input("Root directory (where package.json is)", s.RootDirectory); err != nil {
		return s, err
	}
	build, err := ui.Input("Build script ("+none+" for none)", orNone(deref(s.BuildScript)))
	if err != nil {
		return s, err
	}
	s.BuildScript = optional(build)
	out, err := ui.Input("Output directory ("+none+" for none)", orNone(deref(s.OutputDirectory)))
	if err != nil {
		return s, err
	}
	s.OutputDirectory = optional(out)
	entry, err := ui.Input("Entry file ("+none+" for none)", orNone(s.EntryFile))
	if err != nil {
		return s, err
	}
	if s.EntryFile = strings.TrimSpace(entry); s.EntryFile == none {
		s.EntryFile = ""
	}
	if idx, err = ui.Select("Package manager", PackageManagers, max(indexOf(PackageManagers, s.PackageManager), 0)); err != nil {
		return s, err
	}
	s.PackageManager = PackageManagers[idx]
	return s, nil
}

func summarize(s NodeSettings) string {
	return fmt.Sprintf("%s on Node.js %d", orNone(s.AppType), s.NodeVersion)
}

func orNone(v string) string {
	if v == "" {
		return none
	}
	return v
}

func deref(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func optional(v string) *string {
	v = strings.TrimSpace(v)
	if v == "" || v == none {
		return nil
	}
	return &v
}

func indexOf(list []string, v string) int {
	for i, item := range list {
		if item == v {
			return i
		}
	}
	return -1
}

func toRemembered(s *NodeSettings) *site.NodeSettings {
	if s == nil {
		return nil
	}
	return &site.NodeSettings{
		NodeVersion:     s.NodeVersion,
		AppType:         s.AppType,
		RootDirectory:   s.RootDirectory,
		OutputDirectory: orNone(deref(s.OutputDirectory)), // none: keep it cleared next time
		BuildScript:     orNone(deref(s.BuildScript)),
		EntryFile:       orNone(s.EntryFile),
		PackageManager:  s.PackageManager,
	}
}

func applyOverrides(s NodeSettings, o NodeOverrides) NodeSettings {
	if o.NodeVersion != 0 {
		s.NodeVersion = o.NodeVersion
	}
	if o.AppType != "" {
		s.AppType = o.AppType
	}
	if o.RootDirectory != "" {
		s.RootDirectory = o.RootDirectory
	}
	if o.OutputDirectory != "" {
		s.OutputDirectory = optional(o.OutputDirectory)
	}
	if o.BuildScript != "" {
		s.BuildScript = optional(o.BuildScript)
	}
	switch o.EntryFile {
	case "":
	case none: // cleared in an earlier review: keep it cleared
		s.EntryFile = ""
	default:
		s.EntryFile = o.EntryFile
	}
	if o.PackageManager != "" {
		s.PackageManager = o.PackageManager
	}
	return s
}
