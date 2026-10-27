package hosting

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/hostinger/api-cli/site"
)

// fakeAPI records calls and plays back canned answers.
type fakeAPI struct {
	orders   []Order
	websites []site.Website
	setups   []Setup
	settings NodeSettings
	states   []string // build states returned by successive BuildDetails calls
	calls    []string
	upload   *httptest.Server
	started  NodeSettings

	fromRepo    NodeSettings // what BuildSettingsFromRepository returns
	fromRepoErr error

	installations   []Installation
	repos           []Repository
	onInstallations func() // runs before each installations list, to change the fake between calls
	installsErr     error
	stored          *NodeSettings
	autoDeploy      *GitSettings
	gitStarted      *Repo
	owns            bool
	failSetup       bool
	detailsErrs     int // BuildDetails calls to fail before answering
	setupErrs       int // ListSetups calls to fail, after the first, before answering
	content         bool
	contentErr      error
	datacenters     []Datacenter
}

func (f *fakeAPI) log(format string, args ...any) {
	f.calls = append(f.calls, fmt.Sprintf(format, args...))
}

func (f *fakeAPI) ListOrders(context.Context) ([]Order, error) { f.log("orders"); return f.orders, nil }
func (f *fakeAPI) ListWebsites(_ context.Context, domain string) ([]site.Website, error) {
	f.log("websites %s", domain)
	var out []site.Website
	for _, w := range f.websites {
		if domain == "" || strings.Contains(w.Domain, domain) {
			out = append(out, w)
		}
	}
	return out, nil
}
func (f *fakeAPI) ListDatacenters(_ context.Context, orderID int) ([]Datacenter, error) {
	f.log("datacenters %d", orderID)
	return f.datacenters, nil
}
func (f *fakeAPI) StartSetup(_ context.Context, orderID int, domain, datacenter string) (Setup, error) {
	f.log("setup %d %q dc=%s", orderID, domain, datacenter)
	if domain == "" {
		domain = "fresh-site.hostingersite.com"
	}
	status := "running"
	if f.failSetup {
		status = "failed"
	}
	f.setups = []Setup{{Domain: domain, Username: "u123", Status: status}}
	f.websites = append(f.websites, site.Website{Domain: domain, Username: "u123", OrderID: orderID})
	return f.setups[0], nil
}
func (f *fakeAPI) ListSetups(_ context.Context, domain string) ([]Setup, error) {
	f.log("setups %s", domain)
	if f.setupErrs > 0 && strings.Count(strings.Join(f.calls, "\n"), "setups ") > 1 { // the first check passes
		f.setupErrs--
		return nil, errors.New("setups unavailable (HTTP 502)")
	}
	out := f.setups
	if len(out) > 0 && out[0].Status == "running" {
		f.setups = []Setup{{Domain: out[0].Domain, Username: out[0].Username, Status: "completed"}}
	}
	return out, nil
}
func (f *fakeAPI) UploadCredentials(_ context.Context, username, domain string) (Credentials, error) {
	f.log("creds %s %s", username, domain)
	return Credentials{URL: f.upload.URL, AuthKey: "auth", RestAuthKey: "rest"}, nil
}
func (f *fakeAPI) DeployArchive(_ context.Context, username, domain, archive string) error {
	f.log("deploy %s %s %s", username, domain, archive)
	return nil
}
func (f *fakeAPI) BuildSettingsFromArchive(_ context.Context, username, domain, archive string) (NodeSettings, error) {
	f.log("settings %s", archive)
	return f.settings, nil
}
func (f *fakeAPI) BuildSettingsFromRepository(_ context.Context, username, domain string, r *Repo) (NodeSettings, error) {
	f.log("settings-from-repo %s@%s %s", r.FullName(), r.Branch, r.installUID)
	return f.fromRepo, f.fromRepoErr
}
func (f *fakeAPI) StartBuild(_ context.Context, username, domain string, s NodeSettings, archive string) (Build, error) {
	f.log("build %s %s %s", username, domain, archive)
	f.started = s
	return Build{UUID: "b-1", State: "pending"}, nil
}
func (f *fakeAPI) BuildDetails(context.Context, string, string, string) (Build, error) {
	f.log("details")
	if f.detailsErrs > 0 {
		f.detailsErrs--
		return Build{}, errors.New("status unavailable (HTTP 502)")
	}
	state := f.states[0]
	if len(f.states) > 1 {
		f.states = f.states[1:]
	}
	return Build{UUID: "b-1", State: state}, nil
}
func (f *fakeAPI) BuildLogs(context.Context, string, string, string) (string, error) {
	return "npm ERR! missing script", nil
}
func (f *fakeAPI) BuildAnalysis(context.Context, string, string, string) (string, string, error) {
	return "No build script", "Add one", nil
}
func (f *fakeAPI) HasContent(_ context.Context, username, domain string) (bool, error) {
	f.log("content %s", domain)
	return f.content, f.contentErr
}
func (f *fakeAPI) OwnsDomain(_ context.Context, domain string) (bool, error) {
	f.log("owns %s", domain)
	return f.owns, nil
}
func (f *fakeAPI) ListGitInstallations(context.Context) ([]Installation, error) {
	f.log("installations")
	if f.onInstallations != nil {
		f.onInstallations()
	}
	return f.installations, f.installsErr
}
func (f *fakeAPI) ListInstallationRepositories(_ context.Context, uuid string) ([]Repository, error) {
	f.log("repositories %s", uuid)
	return f.repos, nil
}
func (f *fakeAPI) UpdateGitAutoDeploy(_ context.Context, username, domain string, s GitSettings) error {
	f.log("autodeploy %s %s %s/%s@%s", username, domain, s.Owner, s.Repository, s.Branch)
	f.autoDeploy = &s
	return nil
}
func (f *fakeAPI) BuildSettings(context.Context, string, string) (*NodeSettings, error) {
	f.log("stored-settings")
	return f.stored, nil
}
func (f *fakeAPI) UpdateBuildSettings(_ context.Context, username, domain string, s NodeSettings) error {
	f.log("update-settings")
	f.started = s
	return nil
}
func (f *fakeAPI) StartGitBuild(_ context.Context, username, domain string, s NodeSettings, r *Repo) (Build, error) {
	f.log("git-build %s %s %s@%s", username, domain, r.FullName(), r.Branch)
	f.started = s
	f.gitStarted = r
	return Build{UUID: "b-1", State: "pending"}, nil
}

// fakeUI answers prompts from a script and fails on any unexpected one.
type fakeUI struct {
	t           *testing.T
	interactive bool
	selects     []int
	inputs      []string
	confirms    []bool
	asked       []string
	options     [][]string // the options of each Select, in order
	infos       []string
}

func (u *fakeUI) Interactive() bool { return u.interactive }
func (u *fakeUI) Select(title string, options []string, def int) (int, error) {
	u.asked = append(u.asked, title)
	u.options = append(u.options, options)
	if len(u.selects) == 0 {
		u.t.Fatalf("unexpected select %q %v", title, options)
	}
	idx := u.selects[0]
	u.selects = u.selects[1:]
	if idx < 0 {
		return def, nil
	}
	return idx, nil
}
func (u *fakeUI) Input(title, def string) (string, error) {
	u.asked = append(u.asked, title)
	if len(u.inputs) == 0 {
		u.t.Fatalf("unexpected input %q", title)
	}
	v := u.inputs[0]
	u.inputs = u.inputs[1:]
	return v, nil
}
func (u *fakeUI) Confirm(title string, def bool) (bool, error) {
	u.asked = append(u.asked, title)
	if len(u.confirms) == 0 {
		u.t.Fatalf("unexpected confirm %q", title)
	}
	v := u.confirms[0]
	u.confirms = u.confirms[1:]
	return v, nil
}
func (u *fakeUI) Info(msg string)   { u.infos = append(u.infos, msg) }
func (u *fakeUI) Start(string) Step { return noStep{} }

type noStep struct{}

func (noStep) Update(string) {}
func (noStep) Done(string)   {}
func (noStep) Fail(string)   {}

func fastPolling(t *testing.T) {
	t.Helper()
	SetupPollInterval, BuildPollInterval, BuildPollMax = time.Millisecond, time.Millisecond, time.Millisecond
	delays := RetryDelays
	RetryDelays = []time.Duration{time.Millisecond, time.Millisecond}
	t.Cleanup(func() {
		SetupPollInterval, BuildPollInterval, BuildPollMax = 3*time.Second, 5*time.Second, 20*time.Second
		RetryDelays = delays
	})
}

// linkHere makes dir the current directory for Run, so links land in it.
func linkHere(t *testing.T, dir string) {
	t.Helper()
	old := Getwd
	Getwd = func() (string, error) { return dir, nil }
	t.Cleanup(func() { Getwd = old })
}

func staticProject(t *testing.T) string {
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "index.html", "<h1>hi</h1>")
	write(t, dir, ".env", "SECRET=1")
	return dir
}

func newFake(t *testing.T) *fakeAPI {
	srv := httptest.NewServer(&tusServer{})
	t.Cleanup(srv.Close)
	prev := LookupHost
	LookupHost = func(_ context.Context, host string) ([]string, error) {
		if strings.HasSuffix(host, ".nodns") {
			return nil, &net.DNSError{Err: "no such host", Name: host, IsNotFound: true}
		}
		if strings.HasSuffix(host, ".offline") {
			return nil, &net.DNSError{Err: "i/o timeout", Name: host, IsTimeout: true}
		}
		return []string{"203.0.113.10"}, nil
	}
	t.Cleanup(func() { LookupHost = prev })
	return &fakeAPI{
		orders:   []Order{{ID: 11, Status: "active"}},
		websites: []site.Website{{Domain: "example.com", Username: "u1", OrderID: 11}},
		upload:   srv,
		states:   []string{"running", "completed"},
	}
}

func TestRunStaticToExistingWebsiteNonInteractive(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := staticProject(t)
	ui := &fakeUI{t: t}

	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Username != "u1" || res.Type != Static || res.URL != "https://example.com" || res.Created {
		t.Errorf("result = %+v", res)
	}
	if len(res.SkippedFiles) != 1 || res.SkippedFiles[0] != ".env" {
		t.Errorf("skipped = %v", res.SkippedFiles)
	}
	joined := strings.Join(api.calls, "\n")
	if !strings.Contains(joined, "creds u1 example.com") || !strings.Contains(joined, "deploy u1 example.com "+res.Archive) {
		t.Errorf("calls:\n%s", joined)
	}
	p, err := site.Load(dir)
	if err != nil || p == nil || p.Domain != "example.com" || p.Username != "u1" || p.Type != "static" {
		t.Errorf("project file = %+v, %v", p, err)
	}
}

func TestRunRedeploysFromProjectFileWithoutAsking(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := staticProject(t)
	if err := site.Save(dir, &site.Project{Domain: "example.com", Username: "u1", OrderID: 11, Type: "static"}); err != nil {
		t.Fatal(err)
	}
	ui := &fakeUI{t: t, interactive: true} // any prompt would fail the test
	res, err := Run(context.Background(), api, ui, Options{Path: dir})
	if err != nil {
		t.Fatal(err)
	}
	if res.Domain != "example.com" || res.Type != Static {
		t.Errorf("result = %+v", res)
	}
}

func TestRunDomainOverrideIsOneOff(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.websites = append(api.websites, site.Website{Domain: "staging.net", Username: "u1", OrderID: 11})
	dir := staticProject(t)
	if err := site.Save(dir, &site.Project{Domain: "example.com", Username: "u1", OrderID: 11, Type: "static"}); err != nil {
		t.Fatal(err)
	}
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "staging.net"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Domain != "staging.net" || res.ProjectFile != "" {
		t.Errorf("result = %+v", res)
	}
	p, _ := site.Load(dir)
	if p == nil || p.Domain != "example.com" {
		t.Errorf("link moved to %+v", p)
	}

	// --create from a linked directory moves the link to the new site.
	res, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Create: true, OrderID: 11, Type: Static})
	if err != nil {
		t.Fatal(err)
	}
	p, _ = site.Load(dir)
	if res.ProjectFile == "" || p == nil || p.Domain != res.Domain {
		t.Errorf("link after create = %+v, result %+v", p, res)
	}
}

func TestRunNonInteractiveNeedsADomain(t *testing.T) {
	api := newFake(t)
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t)})
	if err == nil || !strings.Contains(err.Error(), "--domain") {
		t.Errorf("err = %v", err)
	}
	_, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Domain: "missing.com"})
	if err == nil || !strings.Contains(err.Error(), "--create") {
		t.Errorf("err = %v", err)
	}
}

func TestRunInteractiveCreatesWebsiteWithFreeSubdomain(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.orders = append(api.orders, Order{ID: 22, Status: "active"}, Order{ID: 33, Status: "suspended"})
	dir := staticProject(t)
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{true /* detected type */},
		selects: []int{0 /* create new */, 1 /* second plan */},
		inputs:  []string{"" /* free subdomain */},
	}
	res, err := Run(context.Background(), api, ui, Options{Path: dir})
	if err != nil {
		t.Fatal(err)
	}
	if !res.Created || res.Domain != "fresh-site.hostingersite.com" || res.OrderID != 22 || res.Username != "u123" {
		t.Errorf("result = %+v", res)
	}
	joined := strings.Join(api.calls, "\n")
	if !strings.Contains(joined, `setup 22 ""`) {
		t.Errorf("calls:\n%s", joined)
	}
	// The setup was running once, then completed: Run waited for it.
	if strings.Count(joined, "setups fresh-site") < 2 {
		t.Errorf("expected setup polling, calls:\n%s", joined)
	}
	if len(ui.asked) != 4 {
		t.Errorf("asked = %v", ui.asked)
	}
}

func TestRunCreateFlagNonInteractive(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Create: true, OrderID: 11, Domain: "new.example.com", Type: Static})
	if err != nil {
		t.Fatal(err)
	}
	if !res.Created || res.Domain != "new.example.com" {
		t.Errorf("result = %+v", res)
	}
}

func TestRunNodeBuildWithOverrides(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "express", NodeVersion: 20, RootDirectory: "."}
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "package.json", `{"scripts":{"start":"node index.js","build":"echo"}}`)
	write(t, dir, "index.js", "")

	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Node: NodeOverrides{NodeVersion: 22, BuildScript: "build"}})
	if err != nil {
		t.Fatal(err)
	}
	if res.Type != NodeJS || res.BuildUUID != "b-1" || res.BuildState != "completed" {
		t.Errorf("result = %+v", res)
	}
	if api.started.NodeVersion != 22 || api.started.AppType != "express" || api.started.BuildScript == nil || *api.started.BuildScript != "build" || api.started.OutputDirectory != nil {
		t.Errorf("started settings = %+v", api.started)
	}
}

func TestRunNodeBuildFailureCarriesLogs(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "next"}
	api.states = []string{"failed"}
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "package.json", `{"scripts":{"build":"next build"}}`)

	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com"})
	var berr *BuildError
	if !errors.As(err, &berr) {
		t.Fatalf("err = %v", err)
	}
	if berr.Logs == "" || berr.Solution != "Add one" || !strings.Contains(err.Error(), "No build script") {
		t.Errorf("build error = %+v", berr)
	}
	if res == nil || res.BuildUUID != "b-1" || res.BuildState != "failed" {
		t.Errorf("result = %+v", res)
	}
}

func TestRunRefusesWordPressAndArchivesWithoutType(t *testing.T) {
	api := newFake(t)
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "wp-config.php", "")
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com"}); err == nil || !strings.Contains(err.Error(), "wordpress") {
		t.Errorf("err = %v", err)
	}
	archive := filepath.Join(t.TempDir(), "site.zip")
	if err := os.WriteFile(archive, []byte("PK"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: archive, Domain: "example.com"}); err == nil || !strings.Contains(err.Error(), "--type") {
		t.Errorf("err = %v", err)
	}
}

func TestRunArchiveAsIs(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	archive := filepath.Join(t.TempDir(), "site.zip")
	if err := os.WriteFile(archive, []byte("PK"), 0o644); err != nil {
		t.Fatal(err)
	}
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: archive, Domain: "example.com", Type: Static})
	if err != nil {
		t.Fatal(err)
	}
	// Uploaded under a name of its own, so it never replaces a site.zip already there.
	if !strings.HasPrefix(res.Archive, "site_") || !strings.HasSuffix(res.Archive, ".zip") || res.ProjectFile != "" {
		t.Errorf("result = %+v", res)
	}
}

// A cloned repo's link cannot reach outside the project through a symlink.
func TestLinkedSymlinkOutsideTheProjectIsRefused(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := staticProject(t)
	outside := t.TempDir()
	write(t, outside, "id_secret", "x")
	if err := os.Symlink(outside, filepath.Join(dir, "public")); err != nil {
		t.Skip("symlinks not supported:", err)
	}
	if err := site.Save(dir, &site.Project{Kind: "website", Domain: "example.com", Username: "u1", Path: "public", Type: "static"}); err != nil {
		t.Fatal(err)
	}
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{})
	if err == nil || !strings.Contains(err.Error(), "outside the project through a symlink") || strings.Contains(strings.Join(api.calls, "\n"), "creds") {
		t.Errorf("err = %v, calls = %v", err, api.calls)
	}
}

// A folder outside the current directory is deployed but not linked: the next
// run could not find it from here.
func TestFolderOutsideTheLinkDirIsNotLinked(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	cwd := t.TempDir()
	linkHere(t, cwd)
	other := t.TempDir()
	write(t, other, "index.html", "x")
	ui := &fakeUI{t: t}
	res, err := Run(context.Background(), api, ui, Options{Path: other, Domain: "example.com", Type: Static})
	if err != nil {
		t.Fatal(err)
	}
	if res.ProjectFile != "" || !strings.Contains(strings.Join(ui.infos, "\n"), "Not linked") {
		t.Errorf("project file = %q, infos = %v", res.ProjectFile, ui.infos)
	}
	if p, _ := site.Load(cwd); p != nil {
		t.Errorf("link written: %+v", p)
	}
}

func TestRunNoPlan(t *testing.T) {
	api := newFake(t)
	api.orders = nil
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Create: true, Type: Static})
	if err == nil || !strings.Contains(err.Error(), "no Web or Cloud hosting plan") {
		t.Errorf("err = %v", err)
	}
}

func nodeProject(t *testing.T) string {
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "package.json", `{"scripts":{"start":"node server.js","build":"echo build"}}`)
	write(t, dir, "server.js", "")
	write(t, dir, "index.js", "")
	return dir
}

func TestFirstNodeDeployReviewsSettingsAndRemembersThem(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "express", NodeVersion: 18, EntryFile: "server.js"}
	dir := nodeProject(t)
	ui := &fakeUI{t: t, interactive: true,
		selects:  []int{indexOf(AppTypes, "express"), 2 /* Node 22 */, 0 /* npm */},
		confirms: []bool{true /* type */, true /* replace */, false /* change them */},
		inputs:   []string{"." /* root */, "build" /* build script */, "-" /* output dir */, "index.js" /* entry */},
	}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if api.started.NodeVersion != 22 || api.started.AppType != "express" || deref(api.started.BuildScript) != "build" || api.started.OutputDirectory != nil || api.started.EntryFile != "index.js" {
		t.Errorf("build ran with %+v", api.started)
	}
	if res.Settings == nil || res.Settings.NodeVersion != 22 {
		t.Errorf("result settings = %+v", res.Settings)
	}
	p, _ := site.Load(dir)
	if p == nil || p.Node == nil || p.Node.NodeVersion != 22 || p.Node.BuildScript != "build" || p.Node.OutputDirectory != none || p.Node.EntryFile != "index.js" {
		t.Errorf("remembered = %+v", p.Node)
	}
	if p == nil || p.Build == nil || p.Build.UUID != "b-1" || p.Build.State != "completed" || p.Build.At.IsZero() {
		t.Errorf("remembered build = %+v", p.Build)
	}

	// Second deploy: remembered settings win over detection, nothing is asked.
	api.settings = NodeSettings{AppType: "express", NodeVersion: 18, EntryFile: "server.js"}
	quiet := &fakeUI{t: t, interactive: true}
	if _, err := Run(context.Background(), api, quiet, Options{Path: dir}); err != nil {
		t.Fatal(err)
	}
	if api.started.NodeVersion != 22 || api.started.EntryFile != "index.js" {
		t.Errorf("redeploy ignored remembered settings: %+v", api.started)
	}
	if !strings.Contains(strings.Join(quiet.infos, "\n"), "--settings") {
		t.Errorf("redeploy should say how to review settings, infos = %v", quiet.infos)
	}

	// --settings reviews again; accepting keeps them.
	again := &fakeUI{t: t, interactive: true, confirms: []bool{true}}
	if _, err := Run(context.Background(), api, again, Options{Path: dir, Settings: true}); err != nil {
		t.Fatal(err)
	}
	if len(again.asked) != 1 {
		t.Errorf("asked = %v", again.asked)
	}
}

func TestFailedBuildDoesNotRememberSettings(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "next", NodeVersion: 20}
	dir := nodeProject(t)
	if err := site.Save(dir, &site.Project{Kind: "website", Domain: "example.com", Username: "u1", OrderID: 11, Type: "nodejs", Node: &site.NodeSettings{NodeVersion: 18, AppType: "express"}}); err != nil {
		t.Fatal(err)
	}
	api.states = []string{"failed"}
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Node: NodeOverrides{NodeVersion: 24}})
	if err == nil {
		t.Fatal("expected the build to fail")
	}
	p, _ := site.Load(dir)
	if p == nil || p.Domain != "example.com" {
		t.Fatalf("link lost after a failed build: %+v", p)
	}
	if p.Node == nil || p.Node.NodeVersion != 18 || p.Node.AppType != "express" {
		t.Errorf("failed build changed remembered settings: %+v", p.Node)
	}
}

func TestNonInteractiveNodeDeployAsksNothing(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "express", NodeVersion: 18}
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: nodeProject(t), Domain: "example.com", Settings: true})
	if err != nil {
		t.Fatal(err)
	}
	if res.Settings == nil || res.Settings.NodeVersion != 18 {
		t.Errorf("settings = %+v", res.Settings)
	}
}

func TestValidateSettingsAgainstPackedFiles(t *testing.T) {
	dir := nodeProject(t)
	write(t, dir, "app/package.json", `{"scripts":{"compile":"tsc"}}`)
	write(t, dir, "app/main.js", "")
	packed, err := ListFiles(context.Background(), dir)
	if err != nil {
		t.Fatal(err)
	}
	packed.Dir = dir
	str := func(v string) *string { return &v }
	cases := []struct {
		name     string
		settings NodeSettings
		problems int
	}{
		{"detected defaults", NodeSettings{RootDirectory: ".", EntryFile: "server.js", BuildScript: str("build")}, 0},
		{"junk everywhere", NodeSettings{RootDirectory: ".sdfds", EntryFile: "server.jssd", OutputDirectory: str("../x")}, 3},
		{"entry file the build makes", NodeSettings{RootDirectory: ".", EntryFile: "dist/server.js", BuildScript: str("build"), OutputDirectory: str("dist")}, 0},
		{"missing entry file, no build", NodeSettings{RootDirectory: ".", EntryFile: "dist/server.js"}, 1},
		{"nested root", NodeSettings{RootDirectory: "app", EntryFile: "main.js", BuildScript: str("compile")}, 0},
		{"nested root wrong script", NodeSettings{RootDirectory: "app/", BuildScript: str("build")}, 1},
		{"escaping root", NodeSettings{RootDirectory: "../other"}, 1},
	}
	for _, tc := range cases {
		if got := validateSettings(tc.settings, packed); len(got) != tc.problems {
			t.Errorf("%s: problems = %v, want %d", tc.name, got, tc.problems)
		}
	}
}

func TestPlainText(t *testing.T) {
	in := "The build failed during the <b>Preparing build</b> stage with <code>ERROR: x</code>.<ul><li>Re-create <code>a.zip</code></li><li>Change Entry File to: <code>server.js</code></li></ul>"
	want := "The build failed during the Preparing build stage with `ERROR: x`.\n  - Re-create `a.zip`\n  - Change Entry File to: `server.js`"
	if got := PlainText(in); got != want {
		t.Errorf("PlainText =\n%s\nwant\n%s", got, want)
	}
	if PlainText("") != "" {
		t.Error("empty stays empty")
	}
}

func TestReviewReasksUntilSettingsFit(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "express", NodeVersion: 22, EntryFile: "server.js"}
	dir := nodeProject(t)
	ui := &fakeUI{t: t, interactive: true,
		confirms: []bool{true /* type */, true /* replace */, false /* change */, false /* change again after the problems */},
		// first pass: junk; second pass: fixed
		selects: []int{0, 2, 0, 0, 2, 0},
		inputs:  []string{".sdfds", "-sdf", "-", "server.jssd", ".", "build", "-", "server.js"},
	}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Settings.EntryFile != "server.js" || deref(res.Settings.BuildScript) != "build" || res.Settings.RootDirectory != "." {
		t.Errorf("settings = %+v", res.Settings)
	}
	if !strings.Contains(strings.Join(ui.infos, "\n"), "would not build") {
		t.Errorf("problems were not explained: %v", ui.infos)
	}
}

func TestFailedBuildOffersRetryWithSameArchive(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "express", NodeVersion: 22, EntryFile: "server.js"}
	api.states = []string{"failed", "completed"}
	dir := nodeProject(t)
	ui := &fakeUI{t: t, interactive: true,
		confirms: []bool{true /* type */, true /* replace */, true /* use detected */, true /* build again */, true /* use these */},
	}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if res.BuildState != "completed" || strings.Count(strings.Join(api.calls, "\n"), "build u1 example.com") != 2 || strings.Count(strings.Join(api.calls, "\n"), "creds") != 1 {
		t.Errorf("calls = %v", api.calls)
	}
	p, _ := site.Load(dir)
	if p == nil || p.Node == nil {
		t.Errorf("successful retry should remember settings: %+v", p)
	}
}

func TestRememberedSettingsThatNoLongerFitAreReviewed(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "express", NodeVersion: 22}
	dir := nodeProject(t)
	if err := site.Save(dir, &site.Project{Kind: "website", Domain: "example.com", Username: "u1", OrderID: 11, Type: "nodejs", Node: &site.NodeSettings{AppType: "express", NodeVersion: 22, EntryFile: "gone.js"}}); err != nil {
		t.Fatal(err)
	}
	// Non-interactive: fail fast, no upload.
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir})
	if err == nil || !strings.Contains(err.Error(), "gone.js") {
		t.Errorf("err = %v", err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "build u1") {
		t.Error("a build was started with settings that cannot work")
	}
	// Interactive: review is forced; fix the entry file.
	ui := &fakeUI{t: t, interactive: true,
		confirms: []bool{false},
		selects:  []int{0, 2, 0},
		inputs:   []string{".", "-", "-", "server.js"},
	}
	if _, err := Run(context.Background(), api, ui, Options{Path: dir}); err != nil {
		t.Fatal(err)
	}
	if api.started.EntryFile != "server.js" {
		t.Errorf("started with %+v", api.started)
	}
}

func TestTypeConfirmationThenOverride(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := staticProject(t)
	write(t, dir, "package.json", `{"scripts":{"start":"node server.js"}}`)
	// Detected as nodejs; the person says no and picks static.
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{false /* not that type */, true /* replace */}, selects: []int{0}}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Type != Static || !strings.Contains(ui.asked[0], "Node.js app") {
		t.Errorf("type=%s asked=%v", res.Type, ui.asked)
	}
}

func TestOrderLabels(t *testing.T) {
	o := Order{ID: 7}
	o.Plan.Name = "hostinger_unlimited_v6"
	if got := o.Label(); got != "hostinger_unlimited_v6 (order 7)" {
		t.Errorf("plan key: %q", got)
	}
	o.DisplayName = "Unlimited Web Hosting"
	if got := o.Label(); got != "Unlimited Web Hosting (order 7)" {
		t.Errorf("display name: %q", got)
	}
	if got := (Order{ID: 1}).Label(); got != "Hosting plan (order 1)" {
		t.Errorf("empty: %q", got)
	}
}

func TestLinkWrittenBeforeUploadAndRecordsPath(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	root := t.TempDir()
	linkHere(t, root)
	dist := filepath.Join(root, "dist")
	write(t, dist, "index.html", "<h1>built</h1>")

	// First deploy of ./dist from the project root: the link lands in root.
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dist, Domain: "example.com", Type: Static})
	if err != nil {
		t.Fatal(err)
	}
	if res.ProjectFile != site.Path(root) {
		t.Errorf("project file at %s", res.ProjectFile)
	}
	p, _ := site.Load(root)
	if p == nil || p.Path != "dist" {
		t.Fatalf("project = %+v", p)
	}
	if _, err := os.Stat(site.Path(dist)); err == nil {
		t.Error("link was also written inside dist")
	}

	// A bare `up` from the root deploys dist again.
	api.calls = nil
	res, err = Run(context.Background(), api, &fakeUI{t: t}, Options{})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(res.Archive, "dist_") {
		t.Errorf("redeploy packed %s, want the dist folder", res.Archive)
	}

	// The link is written before the upload: a failing upload leaves it behind.
	fresh := t.TempDir()
	linkHere(t, fresh)
	write(t, fresh, "index.html", "")
	api.upload.Close() // every upload now fails
	_, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: fresh, Domain: "example.com", Type: Static})
	if err == nil {
		t.Fatal("expected the upload to fail")
	}
	if p, _ := site.Load(fresh); p == nil || p.Domain != "example.com" {
		t.Errorf("link missing after failed upload: %+v", p)
	}
}

func TestCorruptLinkIsReported(t *testing.T) {
	api := newFake(t)
	dir := staticProject(t)
	if err := os.MkdirAll(filepath.Join(dir, site.Dir), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(site.Path(dir), []byte("{broken"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Type: Static})
	if err == nil || !strings.Contains(err.Error(), "could not be read") {
		t.Errorf("err = %v", err)
	}
}

func TestUnownedCustomDomainIsConfirmed(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := staticProject(t)
	ui := &fakeUI{t: t, interactive: true,
		selects:  []int{0 /* create new */},
		inputs:   []string{"not-mine.com"},
		confirms: []bool{false /* do not create */},
	}
	_, err := Run(context.Background(), api, ui, Options{Path: dir, Type: Static})
	if err == nil || !strings.Contains(err.Error(), "no website created") {
		t.Errorf("err = %v", err)
	}
	if !strings.Contains(strings.Join(ui.asked, "\n"), "does not own not-mine.com") {
		t.Errorf("asked = %v", ui.asked)
	}
	// An owned domain is not questioned.
	api.owns = true
	ui = &fakeUI{t: t, interactive: true, selects: []int{0}, inputs: []string{"mine.com"}}
	if _, err := Run(context.Background(), api, ui, Options{Path: dir, Type: Static}); err != nil {
		t.Fatal(err)
	}
}

func TestSitePickerGroupsWebsitesByPlan(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.orders = []Order{{ID: 11, DisplayName: "Single Web Hosting"}, {ID: 12, DisplayName: "Unlimited Web Hosting"}, {ID: 13, DisplayName: "Single Web Hosting"}}
	// Interleaved by the API; the picker sorts them by plan, keeping each plan's own order.
	api.websites = []site.Website{
		{Domain: "a.example.com", Username: "u2", OrderID: 12, WebsiteType: "nodejs"},
		{Domain: "b.example.com", Username: "u1", OrderID: 11},
		{Domain: "c.example.com", Username: "u2", OrderID: 12},
		{Domain: "d.example.com", Username: "u3", OrderID: 13},
		{Domain: "e.example.com", Username: "u9", OrderID: 99}, // an order the list does not know
	}
	ui := &fakeUI{t: t, interactive: true, selects: []int{3 /* c */}, confirms: []bool{true, true /* replace */}}
	res, err := Run(context.Background(), api, ui, Options{Path: staticProject(t)})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"Create a new website",
		"Order 99: e.example.com",
		"Single Web Hosting (order 11): b.example.com",
		"Single Web Hosting (order 13): d.example.com",
		"Unlimited Web Hosting: a.example.com (nodejs)",
		"Unlimited Web Hosting: c.example.com",
	}
	if len(ui.options) != 1 || strings.Join(ui.options[0], "|") != strings.Join(want, "|") {
		t.Errorf("options = %q", ui.options)
	}
	if res.Domain != "d.example.com" {
		t.Errorf("picked %s", res.Domain)
	}

	// One plan: labels stay flat and the orders are not listed.
	api = newFake(t)
	api.websites = []site.Website{{Domain: "x.example.com", Username: "u1", OrderID: 11}, {Domain: "y.example.com", Username: "u1", OrderID: 11}}
	ui = &fakeUI{t: t, interactive: true, selects: []int{1}, confirms: []bool{true, true /* replace */}}
	if _, err := Run(context.Background(), api, ui, Options{Path: staticProject(t)}); err != nil {
		t.Fatal(err)
	}
	if ui.options[0][1] != "x.example.com" || strings.Contains(strings.Join(api.calls, "\n"), "orders") {
		t.Errorf("options = %q calls = %q", ui.options, api.calls)
	}
}

func TestSitePickerPrefersMatchingType(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.websites = []site.Website{
		{Domain: "node.example.com", Username: "u1", OrderID: 11, WebsiteType: "nodejs"},
		{Domain: "pages.example.com", Username: "u1", OrderID: 11, WebsiteType: "other"},
	}
	dir := staticProject(t)
	ui := &fakeUI{t: t, interactive: true, selects: []int{-1 /* default */}, confirms: []bool{true, true /* replace */}}
	res, err := Run(context.Background(), api, ui, Options{Path: dir})
	if err != nil {
		t.Fatal(err)
	}
	if res.Domain != "pages.example.com" {
		t.Errorf("default picked %s", res.Domain)
	}
}

func TestEmptyFrameworkIsAskedFor(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{NodeVersion: 22} // nothing detected
	dir := nodeProject(t)
	ui := &fakeUI{t: t, interactive: true,
		confirms: []bool{true /* type */, true /* replace */},
		selects:  []int{indexOf(AppTypes, "express"), 2, 0},
		inputs:   []string{".", "-", "-", "server.js"},
	}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Settings.AppType != "express" {
		t.Errorf("settings = %+v", res.Settings)
	}
}

func TestPackedSummary(t *testing.T) {
	l := &Listing{Files: []string{"a"}, Ignored: 2, IgnoreSource: ".gitignore"}
	if got := packedSummary(l); got != "Packed 1 files (2 left out by .gitignore)" {
		t.Errorf("summary = %q", got)
	}
}

func TestDryRunChangesNothing(t *testing.T) {
	api := newFake(t)
	api.websites = nil // the site does not exist yet: a dry run must not create it
	dir := nodeProject(t)
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Create: true, OrderID: 11, DryRun: true, Node: NodeOverrides{NodeVersion: 22}})
	if err != nil {
		t.Fatal(err)
	}
	if !res.DryRun || !res.Created || res.Domain != "<free subdomain>" || res.Type != NodeJS || res.Files != 3 || res.Settings == nil || res.Settings.NodeVersion != 22 || !strings.HasPrefix(res.Archive, filepath.Base(dir)) {
		t.Errorf("plan = %+v", res)
	}
	joined := strings.Join(api.calls, "\n")
	for _, forbidden := range []string{"setup ", "creds", "build", "deploy"} {
		if strings.Contains(joined, forbidden) {
			t.Errorf("dry run called %q:\n%s", forbidden, joined)
		}
	}
	if p, _ := site.Load(dir); p != nil {
		t.Errorf("dry run wrote a link: %+v", p)
	}
}

func TestNoWaitReturnsStartedBuild(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.settings = NodeSettings{AppType: "express", NodeVersion: 22, EntryFile: "server.js"}
	dir := nodeProject(t)
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", NoWait: true})
	if err != nil {
		t.Fatal(err)
	}
	if res.BuildUUID != "b-1" || res.BuildState != "pending" || !strings.Contains(res.Next, "nodejs build`") || !strings.Contains(res.Next, "remembers") {
		t.Errorf("result = %+v", res)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "details") {
		t.Error("polled despite --no-wait")
	}
	p, _ := site.Load(dir)
	if p == nil || p.Domain != "example.com" || p.Node != nil {
		t.Errorf("link after no-wait = %+v", p)
	}
}

func TestFirstDeployToSiteWithContentAsks(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.content = true
	dir := staticProject(t)

	// Interactive, not linked yet: ask, and "no" changes nothing.
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{true /* type */, false /* do not replace */}}
	_, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err == nil || !strings.Contains(err.Error(), "cancelled") {
		t.Errorf("err = %v", err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "creds") {
		t.Error("uploaded after a refused overwrite")
	}
	if p, _ := site.Load(dir); p != nil {
		t.Errorf("refused deploy still linked: %+v", p)
	}

	// "yes" deploys and links; the redeploy from the link does not ask again.
	ui = &fakeUI{t: t, interactive: true, confirms: []bool{true, true}}
	if _, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"}); err != nil {
		t.Fatal(err)
	}
	again := &fakeUI{t: t, interactive: true}
	if _, err := Run(context.Background(), api, again, Options{Path: dir}); err != nil {
		t.Fatal(err)
	}

	// Without a terminal replacing content needs --yes; with it, a warning.
	fresh := staticProject(t)
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: fresh, Domain: "example.com"}); err == nil || !strings.Contains(err.Error(), "pass --yes") {
		t.Errorf("err = %v", err)
	}
	quiet := &fakeUI{t: t}
	if _, err := Run(context.Background(), api, quiet, Options{Path: fresh, Domain: "example.com", Yes: true}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(strings.Join(quiet.infos, "\n"), "already has content") {
		t.Errorf("no warning: %v", quiet.infos)
	}
	api.content = false
	ui = &fakeUI{t: t, interactive: true, confirms: []bool{true /* type */, true /* empty site, still asked */}}
	if _, err := Run(context.Background(), api, ui, Options{Path: staticProject(t), Domain: "example.com"}); err != nil {
		t.Fatal(err)
	}
}

func TestFirstWebsiteOnPlanPicksDatacenter(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.datacenters = []Datacenter{{Title: "Europe (Netherlands)", Code: "serverius"}, {Title: "North America (USA, MA)", Code: "boston"}}

	// Order 11 already has example.com: no datacenter is sent.
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Create: true, OrderID: 11, Type: Static})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range api.calls {
		if strings.HasPrefix(c, "setup 11") && !strings.HasSuffix(c, "dc=") {
			t.Errorf("datacenter sent for a plan with a website: %s", c)
		}
	}

	// Order 22 is empty: without a terminal the best match is used; in a terminal it is asked.
	api.calls = nil
	api.orders = append(api.orders, Order{ID: 22, Status: "active"})
	_, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Create: true, OrderID: 22, Type: Static})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(strings.Join(api.calls, "\n"), `setup 22 "" dc=serverius`) {
		t.Errorf("best match not used: %v", api.calls)
	}
	api.calls = nil
	api.websites = api.websites[:1] // forget the site the previous run created on order 22
	ui := &fakeUI{t: t, interactive: true, selects: []int{1 /* boston */}, inputs: []string{""}}
	_, err = Run(context.Background(), api, ui, Options{Path: staticProject(t), Create: true, OrderID: 22, Type: Static})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(strings.Join(api.calls, "\n"), `dc=boston`) {
		t.Errorf("chosen datacenter not sent: %v", api.calls)
	}
	// --datacenter wins and asks nothing.
	api.calls = nil
	api.websites = api.websites[:1]
	_, err = Run(context.Background(), api, &fakeUI{t: t, interactive: true, inputs: []string{""}}, Options{Path: staticProject(t), Create: true, OrderID: 22, Type: Static, Datacenter: "ukf"})
	if err != nil || !strings.Contains(strings.Join(api.calls, "\n"), "dc=ukf") {
		t.Errorf("flag: %v %v", err, api.calls)
	}
}

func TestUndeployableDirectoryFailsBeforeAnythingIsCreated(t *testing.T) {
	fastPolling(t)
	for name, files := range map[string]map[string]string{
		"wordpress site": {"wp-config.php": "<?php", "index.php": "<?php"},
	} {
		t.Run(name, func(t *testing.T) {
			api := newFake(t)
			api.websites = nil
			dir := t.TempDir()
			linkHere(t, dir)
			for f, body := range files {
				write(t, dir, f, body)
			}
			ui := &fakeUI{t: t, interactive: true}
			_, err := Run(context.Background(), api, ui, Options{Path: dir})
			if err == nil {
				t.Fatal("expected an error")
			}
			if calls := strings.Join(api.calls, "\n"); strings.Contains(calls, "setup") || strings.Contains(calls, "websites") {
				t.Errorf("API was called before the directory was checked:\n%s", calls)
			}
			if len(ui.asked) != 0 {
				t.Errorf("asked %v before the directory was checked", ui.asked)
			}
		})
	}
}

func TestBrokenPackageJSONDeploysAsStaticWithANote(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "index.html", "<h1>hi</h1>")
	write(t, dir, "package.json", "")
	ui := &fakeUI{t: t}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Type != Static || !strings.Contains(strings.Join(ui.infos, "\n"), "package.json could not be read") {
		t.Errorf("type=%s infos=%v", res.Type, ui.infos)
	}
}

func TestUnresolvedDomainIsReportedNotLive(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.websites = append(api.websites, site.Website{Domain: "shop.nodns", Username: "u1", OrderID: 11})
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Domain: "shop.nodns"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Reachable || !strings.Contains(res.ReachableNote, "no DNS records") || !strings.Contains(res.ReachableNote, "Preview the website in hPanel") {
		t.Errorf("reachable=%v note=%q", res.Reachable, res.ReachableNote)
	}
	// A free subdomain is never looked up, and a resolving domain is live.
	res, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Domain: "example.com"})
	if err != nil || !res.Reachable || res.ReachableNote != "" {
		t.Errorf("example.com: reachable=%v note=%q err=%v", res.Reachable, res.ReachableNote, err)
	}
	if ok, note := reachability(context.Background(), "tan-flamingo-307107.hostingersite.com"); !ok || note != "" {
		t.Errorf("platform subdomain: %v %q", ok, note)
	}
	// A lookup that fails for other reasons (offline, timeout) says nothing about the domain.
	if ok, note := reachability(context.Background(), "shop.offline"); !ok || note != "" {
		t.Errorf("offline: %v %q", ok, note)
	}
}

func TestDryRunSettingsSayWhereTheyComeFrom(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := nodeProject(t)
	// Nothing known: no settings object, just where they will come from.
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	if res.Settings != nil || !strings.Contains(res.SettingsSource, "detected from the uploaded files at deploy") {
		t.Errorf("settings=%+v source=%q", res.Settings, res.SettingsSource)
	}
	// Flags: the known part, and a note that the rest is detected.
	res, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", DryRun: true, Node: NodeOverrides{AppType: "express"}})
	if err != nil || res.Settings == nil || res.Settings.AppType != "express" || !strings.HasPrefix(res.SettingsSource, "from flags") {
		t.Errorf("settings=%+v source=%q err=%v", res.Settings, res.SettingsSource, err)
	}
}

// The plan tells an agent what the real run would ask: replacing content needs --yes.
func TestDryRunSaysItOverwritesContent(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.content = true
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Domain: "example.com", DryRun: true})
	if err != nil || !res.OverwritesContent {
		t.Errorf("with content: overwrites=%v err=%v", res != nil && res.OverwritesContent, err)
	}
	api.content = false
	res, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Domain: "example.com", DryRun: true})
	if err != nil || res.OverwritesContent {
		t.Errorf("empty site: overwrites=%v err=%v", res != nil && res.OverwritesContent, err)
	}
}

func TestRememberedTypeYieldsToAFolderOfTheOtherKind(t *testing.T) {
	fastPolling(t)
	dir := nodeProject(t)
	remember := func() {
		if err := site.Save(dir, &site.Project{Kind: "website", Domain: "example.com", Username: "u1", OrderID: 11, Type: "static"}); err != nil {
			t.Fatal(err)
		}
	}
	// Without a terminal a Node.js app never ships as static files: it stops.
	remember()
	api := newFake(t)
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir}); err == nil || !strings.Contains(err.Error(), "--type nodejs") {
		t.Fatalf("err = %v", err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "creds") {
		t.Error("uploaded although the type was in doubt")
	}
	// --yes takes what the folder is.
	api.settings = NodeSettings{AppType: "express", NodeVersion: 22, EntryFile: "server.js"}
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Yes: true})
	if err != nil || res.Type != NodeJS {
		t.Fatalf("--yes: type=%v err=%v", res, err)
	}
	// In a terminal the question defaults to what the folder is.
	remember()
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{true /* deploy as nodejs */, true /* use these settings */}}
	if res, err = Run(context.Background(), api, ui, Options{Path: dir, Settings: false}); err != nil || res.Type != NodeJS {
		t.Fatalf("interactive: %+v %v asked=%v", res, err, ui.asked)
	}
}

func TestSubfolderDeploysTheLinkedFolder(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := staticProject(t)
	if err := site.Save(dir, &site.Project{Kind: "website", Domain: "example.com", Username: "u1", OrderID: 11, Type: "static"}); err != nil {
		t.Fatal(err)
	}
	sub := filepath.Join(dir, "public")
	write(t, sub, "page.html", "")
	linkHere(t, sub) // the command runs inside public/
	ui := &fakeUI{t: t}
	res, err := Run(context.Background(), api, ui, Options{})
	if err != nil {
		t.Fatal(err)
	}
	if res.Files != 2 || !strings.Contains(strings.Join(ui.infos, "\n"), "Deploying the linked folder "+dir) {
		t.Errorf("files=%d infos=%v", res.Files, ui.infos)
	}
	if p, _ := site.Load(dir); p == nil || p.Path != "" {
		t.Errorf("the link moved: %+v", p)
	}
	if _, err := os.Stat(filepath.Join(sub, site.Dir)); err == nil {
		t.Error("a second link was written in the subfolder")
	}
}

func TestStaticDeployOntoNodeSiteAsks(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.websites = []site.Website{{Domain: "app.example.com", Username: "u1", OrderID: 11, WebsiteType: "nodejs"}}
	dir := staticProject(t)
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "app.example.com"})
	if err == nil || !strings.Contains(err.Error(), "the app keeps running") || !strings.Contains(err.Error(), "--yes") {
		t.Fatalf("err = %v", err)
	}
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{true /* static */, false /* not over the app */}}
	if _, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "app.example.com"}); err == nil || !strings.Contains(err.Error(), "cancelled") {
		t.Errorf("err = %v", err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "creds") {
		t.Error("uploaded after a refusal")
	}
	quiet := &fakeUI{t: t}
	if _, err := Run(context.Background(), api, quiet, Options{Path: dir, Domain: "app.example.com", Yes: true}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(strings.Join(quiet.infos, "\n"), "keeps running") {
		t.Errorf("infos = %v", quiet.infos)
	}
}

func TestEmptyFolderAndBadFlagsFailBeforeAnything(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := t.TempDir()
	linkHere(t, dir)
	for _, dry := range []bool{true, false} {
		if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", DryRun: dry}); err == nil || !strings.Contains(err.Error(), "no files to deploy") {
			t.Errorf("dry=%v err=%v", dry, err)
		}
	}
	if p, _ := site.Load(dir); p != nil {
		t.Errorf("a rejected folder was linked: %+v", p)
	}
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Domain: "example.com", DryRun: true, Node: NodeOverrides{NodeVersion: 17}}); err == nil || !strings.Contains(err.Error(), "--node-version 17") {
		t.Errorf("node 17: %v", err)
	}
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Create: true, OrderID: 99, DryRun: true}); err == nil || !strings.Contains(err.Error(), "order 99 was not found") {
		t.Errorf("order 99: %v", err)
	}
}

func TestContentCheckFailureCountsAsContent(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.contentErr = errors.New("listing failed (HTTP 500)")
	// Without a terminal an unknown site is not replaced without --yes.
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Domain: "example.com"})
	if err == nil || !strings.Contains(err.Error(), "could not check whether example.com has content") || !strings.Contains(err.Error(), "pass --yes") {
		t.Errorf("err = %v", err)
	}
	// In a terminal the question defaults to no, as for a site with content.
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{true /* type */, false}}
	if _, err := Run(context.Background(), api, ui, Options{Path: staticProject(t), Domain: "example.com"}); err == nil {
		t.Error("deployed after refusing")
	}
	if q := ui.asked[len(ui.asked)-1]; !strings.Contains(q, "already has content") {
		t.Errorf("question = %q", q)
	}
}

func TestStaticRedeployKeepsTheRememberedBuild(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := staticProject(t)
	ref := &site.BuildRef{UUID: "9b2f4c1e-0d3a-4e5f-8a6b-7c8d9e0f1a2b", State: "completed", At: time.Now()}
	if err := site.Save(dir, &site.Project{Kind: "website", Domain: "example.com", Username: "u1", OrderID: 11, Type: "static", Build: ref}); err != nil {
		t.Fatal(err)
	}
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir}); err != nil {
		t.Fatal(err)
	}
	if p, _ := site.Load(dir); p == nil || p.Build == nil || p.Build.UUID != ref.UUID {
		t.Errorf("build forgotten: %+v", p)
	}
}

func TestEmptyFolderFailsBeforeAWebsiteIsCreated(t *testing.T) {
	api := newFake(t)
	dir := t.TempDir()
	linkHere(t, dir)
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Create: true, OrderID: 11, Type: Static}); err == nil || !strings.Contains(err.Error(), "no files to deploy") {
		t.Errorf("err = %v", err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "setup ") {
		t.Errorf("a website was created for an empty folder: %v", api.calls)
	}
}

func TestCreatedWebsiteStaysLinkedWhenTheRunFails(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.failSetup = true
	dir := staticProject(t)
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Create: true, OrderID: 11, Type: Static})
	if err == nil || !strings.Contains(err.Error(), "fresh-site.hostingersite.com was created; run again without --create") {
		t.Errorf("err = %v", err)
	}
	if p, _ := site.Load(dir); p == nil || p.Domain != "fresh-site.hostingersite.com" {
		t.Errorf("new website not linked: %+v", p)
	}
}

func TestBuildPollingToleratesAFewErrors(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.detailsErrs = 2
	opts := Options{Path: nodeProject(t), Domain: "example.com", Type: NodeJS, Node: NodeOverrides{AppType: "express", NodeVersion: 22}}
	if res, err := Run(context.Background(), api, &fakeUI{t: t}, opts); err != nil || res.BuildState != "completed" {
		t.Fatalf("res = %+v, err = %v", res, err)
	}
	api = newFake(t)
	api.detailsErrs = 10
	opts.Path = nodeProject(t)
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, opts); err == nil || !strings.Contains(err.Error(), "may still be running") {
		t.Errorf("err = %v", err)
	}
}

func TestClearedSettingsAreRemembered(t *testing.T) {
	build := "build"
	r := toRemembered(&NodeSettings{AppType: "vite", NodeVersion: 22})
	if r.BuildScript != none || r.OutputDirectory != none {
		t.Errorf("remembered = %+v", r)
	}
	s := applyOverrides(NodeSettings{BuildScript: &build}, *r)
	if s.BuildScript != nil || s.OutputDirectory != nil {
		t.Errorf("a cleared setting came back: %+v", s)
	}
	if escapes("my..app") || !escapes("app/../other") {
		t.Error("escapes should only match a .. segment")
	}
}

func TestSetupWaitToleratesAFewErrors(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.setupErrs = 2
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: staticProject(t), Create: true, OrderID: 11, Type: Static})
	if err != nil || !res.Created {
		t.Errorf("res = %+v, err = %v", res, err)
	}
}

func TestUndeployableFolderFailsEvenWithARememberedType(t *testing.T) {
	api := newFake(t)
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "wp-config.php", "<?php")
	if err := site.Save(dir, &site.Project{Kind: "website", Domain: "example.com", Username: "u1", Type: "static"}); err != nil {
		t.Fatal(err)
	}
	// --create means the link does not pick the website, so its type does not apply.
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Create: true, OrderID: 11})
	if err == nil || !strings.Contains(err.Error(), "cannot deploy yet") {
		t.Errorf("err = %v", err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "setup ") {
		t.Errorf("a website was created first: %v", api.calls)
	}
}

func TestLinkedPathMustStayInsideTheProject(t *testing.T) {
	api := newFake(t)
	dir := staticProject(t)
	// A cloned repository can ship its own site.json pointing anywhere.
	for _, bad := range []string{"../../.ssh", "/etc", "dist/../../secrets"} {
		if err := site.Save(dir, &site.Project{Kind: "website", Domain: "example.com", Username: "u1", Type: "static", Path: bad}); err != nil {
			t.Fatal(err)
		}
		_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{})
		if err == nil || !strings.Contains(err.Error(), "must be a relative path inside the project") {
			t.Errorf("%s: err = %v", bad, err)
		}
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "creds") {
		t.Errorf("uploaded from outside the project: %v", api.calls)
	}
}

// An entry file cleared in the review stays cleared, and a name ending in - is kept.
func TestClearedEntryFileIsRemembered(t *testing.T) {
	remembered := toRemembered(&NodeSettings{AppType: "express", NodeVersion: 22})
	if got := applyOverrides(NodeSettings{EntryFile: "server.js"}, *remembered); got.EntryFile != "" {
		t.Errorf("cleared entry came back as %q", got.EntryFile)
	}
	remembered = toRemembered(&NodeSettings{EntryFile: "server-"})
	if got := applyOverrides(NodeSettings{}, *remembered); got.EntryFile != "server-" {
		t.Errorf("entry = %q", got.EntryFile)
	}
}
