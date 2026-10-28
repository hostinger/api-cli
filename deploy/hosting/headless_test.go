package hosting

import (
	"context"
	"strings"
	"testing"

	"github.com/hostinger/api-cli/site"
)

// headlessFake has one headless website, cms.example.com, where the server
// detects a Next.js app.
func headlessFake(t *testing.T) *fakeAPI {
	api := newFake(t)
	api.websites = []site.Website{{Domain: "cms.example.com", Username: "u1", OrderID: 11, WebsiteType: "headless"}}
	out, build := ".next", "build"
	api.settings = NodeSettings{AppType: "next", NodeVersion: 20, RootDirectory: ".", OutputDirectory: &out, BuildScript: &build, PackageManager: "npm"}
	return api
}

func nextProject(t *testing.T) string {
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "package.json", `{"scripts":{"build":"next build","start":"next start"},"dependencies":{"next":"15","react":"19"}}`)
	write(t, dir, "app/page.tsx", "")
	return dir
}

func TestHeadlessSiteBuildsTheFrontendAsNodeJS(t *testing.T) {
	fastPolling(t)
	api := headlessFake(t)
	dir := nextProject(t)
	ui := &fakeUI{t: t}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "cms.example.com", Yes: true})
	if err != nil {
		t.Fatal(err)
	}
	if res.Type != NodeJS || !res.Headless || res.BuildState != "completed" {
		t.Errorf("result = %+v", res)
	}
	s := api.started
	if s.AppType != "next" || deref(s.OutputDirectory) != ".next" || deref(s.BuildScript) != "build" || s.PackageManager != "npm" || s.NodeVersion != 20 {
		t.Errorf("started settings = %+v (output %q, build %q)", s, deref(s.OutputDirectory), deref(s.BuildScript))
	}
	infos := strings.Join(ui.infos, "\n")
	for _, want := range []string{"as nodejs (headless)", "backend (such as WordPress in /cms) is left alone"} {
		if !strings.Contains(infos, want) {
			t.Errorf("infos lack %q:\n%s", want, infos)
		}
	}
	if p, _ := site.Load(dir); p == nil || !p.Headless || p.Type != "nodejs" {
		t.Errorf("link = %+v", p)
	}
}

func TestHeadlessSiteRefusesStaticDeploys(t *testing.T) {
	fastPolling(t)
	api := headlessFake(t)
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: nextProject(t), Domain: "cms.example.com", Type: Static, Yes: true})
	if err == nil || !strings.Contains(err.Error(), "takes no static deploy") {
		t.Errorf("--type static: err = %v", err)
	}
	_, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Domain: "cms.example.com", Yes: true})
	if err == nil || !strings.Contains(err.Error(), "has no package.json") {
		t.Errorf("plain HTML: err = %v", err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "creds") {
		t.Error("uploaded after a refusal")
	}
}

func TestLinkRemembersAHeadlessSiteThatNowReportsNodeJS(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.websites = []site.Website{{Domain: "cms.example.com", Username: "u1", OrderID: 11, WebsiteType: "nodejs"}}
	dir := staticProject(t)
	if err := site.Save(dir, &site.Project{Domain: "cms.example.com", Username: "u1", Type: "static", Headless: true}); err != nil {
		t.Fatal(err)
	}
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir})
	if err == nil || !strings.Contains(err.Error(), "has no package.json") {
		t.Errorf("err = %v", err)
	}
}

func TestSitePickerPrefersAHeadlessSiteForANodeApp(t *testing.T) {
	fastPolling(t)
	api := headlessFake(t)
	api.websites = append([]site.Website{{Domain: "pages.example.com", Username: "u1", OrderID: 11, WebsiteType: "other"}}, api.websites...)
	ui := &fakeUI{t: t, interactive: true, selects: []int{-1 /* default */}, confirms: []bool{true /* replace */, true /* settings */}}
	res, err := Run(context.Background(), api, ui, Options{Path: nextProject(t)})
	if err != nil {
		t.Fatal(err)
	}
	if res.Domain != "cms.example.com" {
		t.Errorf("default picked %s", res.Domain)
	}
	if !strings.Contains(strings.Join(ui.options[0], "\n"), "cms.example.com (headless)") {
		t.Errorf("options = %v", ui.options[0])
	}
}

func TestHeadlessFlagCreatesAHeadlessWebsite(t *testing.T) {
	fastPolling(t)
	api := headlessFake(t)
	dir := nextProject(t)
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Headless: "wordpress", Domain: "shop.example.com", OrderID: 11})
	if err != nil {
		t.Fatal(err)
	}
	if !res.Created || !res.Headless || res.Domain != "shop.example.com" || api.started.AppType != "next" {
		t.Errorf("result = %+v, started %+v", res, api.started)
	}
	if !strings.Contains(strings.Join(api.calls, "\n"), "setup-type headless_wordpress") {
		t.Errorf("calls:\n%s", strings.Join(api.calls, "\n"))
	}
	if p, _ := site.Load(dir); p == nil || !p.Headless || p.Domain != "shop.example.com" {
		t.Errorf("link = %+v", p)
	}
}

func TestHeadlessFlagCreatesOnlyWhatItCan(t *testing.T) {
	fastPolling(t)
	api := headlessFake(t)
	api.websites = append(api.websites, site.Website{Domain: "example.com", Username: "u1", OrderID: 11, WebsiteType: "other"})

	// A static folder fails before a website is created.
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Headless: "empty", Create: true})
	if err == nil || !strings.Contains(err.Error(), "has no package.json") {
		t.Errorf("static folder: err = %v", err)
	}
	// An existing website that is not headless is not turned into one.
	_, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: nextProject(t), Headless: "empty", Domain: "example.com", Yes: true})
	if err == nil || !strings.Contains(err.Error(), "is not a headless website") {
		t.Errorf("regular website: err = %v", err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "setup ") {
		t.Errorf("created a website:\n%s", strings.Join(api.calls, "\n"))
	}
	// An existing headless website takes the deploy: running again after a failure.
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: nextProject(t), Headless: "wordpress", Domain: "cms.example.com", Yes: true})
	if err != nil || res.Created || !res.Headless {
		t.Errorf("headless website: %+v, %v", res, err)
	}
}

func TestInteractiveCreateAsksANodeAppWhatKindOfWebsite(t *testing.T) {
	fastPolling(t)
	api := headlessFake(t)
	api.websites = nil
	ui := &fakeUI{t: t, interactive: true,
		confirms: []bool{false /* not a regular website */, true /* settings */},
		selects:  []int{indexOf(headlessChoices, "Headless for a PocketBase backend")},
		inputs:   []string{"" /* free subdomain */},
	}
	res, err := Run(context.Background(), api, ui, Options{Path: nextProject(t)})
	if err != nil {
		t.Fatal(err)
	}
	if !res.Created || !res.Headless {
		t.Errorf("result = %+v", res)
	}
	if !strings.Contains(strings.Join(api.calls, "\n"), "setup-type headless_pocketbase") {
		t.Errorf("calls:\n%s", strings.Join(api.calls, "\n"))
	}
}

func TestDryRunOfAHeadlessCreateCreatesNothing(t *testing.T) {
	api := headlessFake(t)
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: nextProject(t), Create: true, Headless: "empty", OrderID: 11, DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	if !res.DryRun || !res.Created || !res.Headless || res.Type != NodeJS || res.Settings != nil || !strings.Contains(res.SettingsSource, "at deploy") {
		t.Errorf("plan = %+v", res)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "setup ") {
		t.Errorf("dry run created a website:\n%s", strings.Join(api.calls, "\n"))
	}
}

func TestParseHeadless(t *testing.T) {
	if k, err := ParseHeadless("wordpress"); err != nil || k != "wordpress" {
		t.Errorf("wordpress = %q, %v", k, err)
	}
	if _, err := ParseHeadless("headless_wordpress"); err == nil || !strings.Contains(err.Error(), "one of: empty, wordpress, ecommerce, pocketbase") {
		t.Errorf("err = %v", err)
	}
}

func TestHeadlessBackendIsNotContentADeployReplaces(t *testing.T) {
	fastPolling(t)
	api := headlessFake(t)
	api.content = true // the backend in /cms fills the document root

	// No frontend built yet: nothing to replace, so no --yes is needed.
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: nextProject(t), Domain: "cms.example.com"}); err != nil {
		t.Fatalf("first frontend: %v", err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "content ") {
		t.Errorf("checked the document root:\n%s", strings.Join(api.calls, "\n"))
	}
	plan, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: nextProject(t), Domain: "cms.example.com", DryRun: true})
	if err != nil || plan.OverwritesContent {
		t.Errorf("dry run = %+v, %v", plan, err)
	}

	// A built frontend is replaced, and the messages say the backend stays.
	out := ".next"
	api.stored = &NodeSettings{AppType: "next", NodeVersion: 22, OutputDirectory: &out}
	_, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: nextProject(t), Domain: "cms.example.com"})
	if err == nil || !strings.Contains(err.Error(), "already has a frontend") || !strings.Contains(err.Error(), "backend, such as WordPress in /cms, is kept") || !strings.Contains(err.Error(), "--yes") {
		t.Errorf("built frontend without --yes: err = %v", err)
	}
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{false}}
	if _, err := Run(context.Background(), api, ui, Options{Path: nextProject(t), Domain: "cms.example.com"}); err == nil || !strings.Contains(err.Error(), "cancelled") {
		t.Errorf("declined: err = %v", err)
	}
	if len(ui.asked) != 1 || !strings.Contains(ui.asked[0], "already has a frontend, which this deploy replaces (its backend") {
		t.Errorf("asked = %v", ui.asked)
	}
}

// frontendFor is a Next.js project depending on dep, in dependencies or devDependencies.
func frontendFor(t *testing.T, field, dep string) string {
	dir := nextProject(t)
	write(t, dir, "package.json", `{"scripts":{"build":"next build"},"dependencies":{"next":"15"},"`+field+`":{"`+dep+`":"1"}}`)
	return dir
}

func TestSuggestHeadless(t *testing.T) {
	// A backend client wins over the frontend framework next to it.
	for _, tc := range []struct{ field, dep, want string }{
		{"dependencies", "pocketbase", "pocketbase"},
		{"dependencies", "@faustwp/core", "wordpress"},
		{"devDependencies", "@wordpress/api-fetch", "wordpress"},
		{"dependencies", "graphql-request", "empty"},
	} {
		if got := SuggestHeadless(frontendFor(t, tc.field, tc.dep)); got != tc.want {
			t.Errorf("next + %s %s = %q, want %q", tc.field, tc.dep, got, tc.want)
		}
	}
	for pkg, want := range map[string]string{
		`{"devDependencies":{"vite":"6"}}`:                     "empty",
		`{"dependencies":{"@sveltejs/kit":"2"}}`:               "empty",
		`{"dependencies":{"express":"4"}}`:                     "",
		`{"scripts":{"start":"node server.js"}}`:               "",
		`{"dependencies":{"express":"4","pocketbase":"0.21"}}`: "pocketbase",
	} {
		dir := t.TempDir()
		write(t, dir, "package.json", pkg)
		if got := SuggestHeadless(dir); got != want {
			t.Errorf("%s = %q, want %q", pkg, got, want)
		}
	}
	if got := SuggestHeadless(t.TempDir()); got != "" {
		t.Errorf("no package.json = %q", got)
	}
}

// apiProject is a Node.js API server: no frontend, so no headless suggestion.
func apiProject(t *testing.T) string {
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "package.json", `{"scripts":{"start":"node server.js","build":"tsc"},"dependencies":{"express":"4"}}`)
	write(t, dir, "server.js", "")
	return dir
}

// createFromTerminal creates a website for dir in a terminal, answering the
// kind question with confirm and, when declined, the picker's default.
func createFromTerminal(t *testing.T, dir string, confirm bool) (*fakeAPI, *fakeUI) {
	t.Helper()
	fastPolling(t)
	api := headlessFake(t)
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{confirm, true /* settings */}, inputs: []string{"" /* free subdomain */}}
	if !confirm {
		ui.selects = []int{-1 /* default */}
	}
	if _, err := Run(context.Background(), api, ui, Options{Path: dir, Create: true, OrderID: 11, Type: NodeJS}); err != nil {
		t.Fatal(err)
	}
	return api, ui
}

func setupType(api *fakeAPI) string {
	for _, c := range api.calls {
		if k, ok := strings.CutPrefix(c, "setup-type "); ok {
			return k
		}
	}
	return ""
}

func TestCreateConfirmsTheSuggestedKindOfWebsite(t *testing.T) {
	// A PocketBase frontend: Enter creates the suggested headless website.
	api, ui := createFromTerminal(t, frontendFor(t, "dependencies", "pocketbase"), true)
	if got := setupType(api); got != "headless_pocketbase" {
		t.Errorf("setup type = %q", got)
	}
	if ui.asked[0] != "This looks like a frontend for PocketBase. Create a headless website for a PocketBase backend?" || len(ui.options) != 0 {
		t.Errorf("asked = %q, options = %v", ui.asked, ui.options)
	}

	// Declined: the full question, offering a regular website first.
	api, ui = createFromTerminal(t, frontendFor(t, "devDependencies", "@faustwp/core"), false)
	if got := setupType(api); got != "" {
		t.Errorf("declined suggestion created %q", got)
	}
	if ui.asked[0] != "This looks like a headless WordPress frontend. Create a headless website with WordPress in /cms?" || ui.asked[1] != "What kind of website?" {
		t.Errorf("asked = %q", ui.asked)
	}
}

func TestCreateSuggestsAnEmptyHeadlessWebsiteForAFrontend(t *testing.T) {
	api, ui := createFromTerminal(t, nextProject(t), true)
	if got := setupType(api); got != "headless_empty" || ui.asked[0] != "This looks like a web frontend. Create a headless website (frontend only, add a backend later)?" || len(ui.options) != 0 {
		t.Errorf("setup type %q, asked %q, options %v", got, ui.asked, ui.options)
	}
	// Declined: the full question, offering a regular website first.
	api, ui = createFromTerminal(t, nextProject(t), false)
	if got := setupType(api); got != "" || ui.asked[1] != "What kind of website?" {
		t.Errorf("setup type %q, asked %q", got, ui.asked)
	}
}

func TestCreateWithoutASuggestionConfirmsARegularWebsite(t *testing.T) {
	// Enter: a regular website, no picker.
	api, ui := createFromTerminal(t, apiProject(t), true)
	if got := setupType(api); got != "" || ui.asked[0] != "Create a regular website?" || len(ui.options) != 0 {
		t.Errorf("setup type %q, asked %q, options %v", got, ui.asked, ui.options)
	}
	// Declined: the full question, offering a headless website first.
	api, ui = createFromTerminal(t, apiProject(t), false)
	if got := setupType(api); got != "headless_empty" || ui.asked[1] != "What kind of website?" {
		t.Errorf("setup type %q, asked %q", got, ui.asked)
	}
}

func TestCreateWithoutATerminalOnlyHintsTheSuggestion(t *testing.T) {
	fastPolling(t)
	api := headlessFake(t)
	ui := &fakeUI{t: t}
	if _, err := Run(context.Background(), api, ui, Options{Path: frontendFor(t, "dependencies", "pocketbase"), Create: true, OrderID: 11}); err != nil {
		t.Fatal(err)
	}
	if got := setupType(api); got != "" {
		t.Errorf("guessed %q without a terminal", got)
	}
	if !strings.Contains(strings.Join(ui.infos, "\n"), "pass --headless pocketbase") {
		t.Errorf("infos = %v", ui.infos)
	}
	// --headless asks nothing.
	api = headlessFake(t)
	if _, err := Run(context.Background(), api, &fakeUI{t: t, interactive: true, confirms: []bool{true /* settings */}, inputs: []string{""}}, Options{Path: frontendFor(t, "dependencies", "pocketbase"), Create: true, OrderID: 11, Headless: "empty"}); err != nil {
		t.Fatal(err)
	}
	if got := setupType(api); got != "headless_empty" {
		t.Errorf("--headless created %q", got)
	}
}

func TestDryRunReportsTheSuggestedKind(t *testing.T) {
	api := headlessFake(t)
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: frontendFor(t, "dependencies", "@faustwp/core"), Create: true, OrderID: 11, DryRun: true})
	if err != nil || res.SuggestedHeadless != "wordpress" || res.Headless {
		t.Errorf("plan = %+v, %v", res, err)
	}
	res, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: frontendFor(t, "dependencies", "@faustwp/core"), Create: true, OrderID: 11, DryRun: true, Headless: "wordpress"})
	if err != nil || res.SuggestedHeadless != "" || !res.Headless {
		t.Errorf("plan with --headless = %+v, %v", res, err)
	}
}
