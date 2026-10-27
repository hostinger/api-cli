package hosting

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/hostinger/api-cli/site"
)

func TestParseRemote(t *testing.T) {
	cases := map[string][3]string{
		"https://github.com/acme/site.git":       {"github.com", "acme", "site"},
		"git@github.com:acme/site.git":           {"github.com", "acme", "site"},
		"ssh://git@gitlab.com/group/sub/app.git": {"gitlab.com", "group/sub", "app"},
		"https://GitLab.com/acme/site":           {"gitlab.com", "acme", "site"},
		"github.com:acme/site.git":               {"github.com", "acme", "site"},
		"gh-work:acme/site.git":                  {"gh-work", "acme", "site"},
	}
	for raw, want := range cases {
		host, owner, name, ok := parseRemote(raw)
		if !ok || host != want[0] || owner != want[1] || name != want[2] {
			t.Errorf("%s -> %q %q %q ok=%v", raw, host, owner, name, ok)
		}
	}
	if _, _, _, ok := parseRemote("not a url"); ok {
		t.Error("garbage parsed")
	}
}

// gitProject makes a committed, pushed checkout of acme/site with a static page.
func gitProject(t *testing.T) (string, *git.Repository) {
	t.Helper()
	dir := t.TempDir()
	linkHere(t, dir)
	repo, err := git.PlainInit(dir, false)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := repo.CreateRemote(&config.RemoteConfig{Name: "origin", URLs: []string{"git@github.com:acme/site.git"}}); err != nil {
		t.Fatal(err)
	}
	write(t, dir, "index.html", "<h1>hi</h1>")
	commitAll(t, repo, "init")
	push(t, repo)
	return dir, repo
}

func commitAll(t *testing.T, repo *git.Repository, msg string) plumbing.Hash {
	t.Helper()
	wt, err := repo.Worktree()
	if err != nil {
		t.Fatal(err)
	}
	if err := wt.AddGlob("."); err != nil {
		t.Fatal(err)
	}
	hash, err := wt.Commit(msg, &git.CommitOptions{Author: &object.Signature{Name: "t", Email: "t@example.com", When: time.Now()}})
	if err != nil {
		t.Fatal(err)
	}
	return hash
}

// push pretends origin has the current branch at HEAD.
func push(t *testing.T, repo *git.Repository) {
	t.Helper()
	head, err := repo.Head()
	if err != nil {
		t.Fatal(err)
	}
	ref := plumbing.NewHashReference(plumbing.NewRemoteReferenceName("origin", head.Name().Short()), head.Hash())
	if err := repo.Storer.SetReference(ref); err != nil {
		t.Fatal(err)
	}
}

func TestDetectRepo(t *testing.T) {
	if r, err := DetectRepo(t.TempDir()); err != nil || r != nil {
		t.Errorf("plain dir: %+v %v", r, err)
	}
	dir, _ := gitProject(t)
	r, err := DetectRepo(dir)
	if err != nil || r == nil || r.FullName() != "acme/site" || r.Branch != "master" {
		t.Fatalf("checkout: %+v %v", r, err)
	}
	// What the working tree holds does not matter: a Git deploy builds the pushed branch.
	write(t, dir, "new.html", "")
	if r, err = DetectRepo(dir); err != nil || r == nil {
		t.Errorf("with uncommitted work: %+v %v", r, err)
	}
}

func gitFake(t *testing.T) *fakeAPI {
	api := newFake(t)
	api.installations = []Installation{{UUID: "inst-1", Provider: "github", AccountLogin: "acme", Status: "active"}}
	api.repos = []Repository{{Name: "site", FullName: "acme/site", Owner: "acme", DefaultBranch: "master"}}
	return api
}

func TestGitStaticDeployNonInteractive(t *testing.T) {
	fastPolling(t)
	api := gitFake(t)
	dir, _ := gitProject(t)
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Source: SourceGit})
	if err != nil {
		t.Fatal(err)
	}
	if res.Source != SourceGit || res.Repository != "acme/site" || res.Branch != "master" || res.Archive != "" {
		t.Errorf("result = %+v", res)
	}
	calls := strings.Join(api.calls, "\n")
	if !strings.Contains(calls, "autodeploy u1 example.com acme/site@master") || strings.Contains(calls, "creds") {
		t.Errorf("calls:\n%s", calls)
	}
	if api.autoDeploy == nil || !api.autoDeploy.IsEnabled || api.autoDeploy.InstallationUUID != "inst-1" {
		t.Errorf("auto deploy = %+v", api.autoDeploy)
	}
	p, _ := site.Load(dir)
	if p == nil || p.Source != "git" {
		t.Errorf("project = %+v", p)
	}
}

// Git was only on offer: when the installations cannot be listed, upload.
func TestGitLookupFailureFallsBackToUpload(t *testing.T) {
	fastPolling(t)
	api := gitFake(t)
	api.installsErr = errors.New("429 too many requests")
	dir, _ := gitProject(t)
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{true, true}}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil || res.Source != SourceUpload {
		t.Fatalf("source = %v, err = %v", res, err)
	}
	if !strings.Contains(strings.Join(ui.infos, "\n"), "uploading this folder") {
		t.Errorf("infos = %v", ui.infos)
	}
	// Asked for Git: the failure is the answer.
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Source: SourceGit, Yes: true}); err == nil || !strings.Contains(err.Error(), "429") {
		t.Errorf("--source git: %v", err)
	}
}

// Auto-deploy clones nothing onto a Node.js website, so a static Git deploy there is refused.
func TestStaticGitDeployOntoNodeAppIsRefused(t *testing.T) {
	fastPolling(t)
	api := gitFake(t)
	api.websites[0].WebsiteType = "nodejs"
	dir, _ := gitProject(t)
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Source: SourceGit, Type: Static, Yes: true})
	if err == nil || !strings.Contains(err.Error(), "deploys nothing") || strings.Contains(strings.Join(api.calls, "\n"), "autodeploy") {
		t.Errorf("err = %v, calls = %v", err, api.calls)
	}
}

// A linked worktree reads origin from the main repository.
func TestDetectRepoInALinkedWorktree(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed")
	}
	dir, _ := gitProject(t)
	wt := filepath.Join(t.TempDir(), "wt")
	if out, err := exec.Command("git", "-C", dir, "worktree", "add", "-b", "feature", wt).CombinedOutput(); err != nil {
		t.Fatalf("git worktree add: %v\n%s", err, out)
	}
	r, err := DetectRepo(wt)
	if err != nil || r == nil || r.FullName() != "acme/site" || r.Branch != "feature" {
		t.Errorf("repo = %+v, err = %v", r, err)
	}
}

func TestGitNodeDeployUsesStoredSettings(t *testing.T) {
	fastPolling(t)
	api := gitFake(t)
	api.stored = &NodeSettings{AppType: "express", NodeVersion: 20, EntryFile: "server.js"}
	dir, repo := gitProject(t)
	write(t, dir, "package.json", `{"scripts":{"start":"node server.js"}}`)
	write(t, dir, "server.js", "")
	commitAll(t, repo, "node")
	push(t, repo)

	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Source: SourceGit, Node: NodeOverrides{NodeVersion: 22}})
	if err != nil {
		t.Fatal(err)
	}
	if res.Type != NodeJS || res.BuildState != "completed" || api.gitStarted == nil || api.gitStarted.Branch != "master" {
		t.Errorf("result = %+v started = %+v", res, api.gitStarted)
	}
	if api.started.NodeVersion != 22 || api.started.AppType != "express" {
		t.Errorf("settings = %+v", api.started)
	}
	calls := strings.Join(api.calls, "\n")
	if !strings.Contains(calls, "update-settings") || strings.Contains(calls, "creds") {
		t.Errorf("calls:\n%s", calls)
	}
}

func TestGitNodeDeployDetectsSettingsFromRepository(t *testing.T) {
	fastPolling(t)
	api := gitFake(t)
	api.fromRepo = NodeSettings{AppType: "next", NodeVersion: 22, RootDirectory: ".", BuildScript: strPtr("build"), OutputDirectory: strPtr(".next"), PackageManager: "npm"}
	dir, repo := gitProject(t)
	write(t, dir, "package.json", `{"scripts":{"build":"next build","start":"next start"},"dependencies":{"next":"^16"}}`)
	commitAll(t, repo, "next app")
	push(t, repo)

	// No stored settings and no terminal: the API's detection from the branch is enough.
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Source: SourceGit})
	if err != nil {
		t.Fatal(err)
	}
	s := api.started
	if res.Type != NodeJS || s.AppType != "next" || s.NodeVersion != 22 || deref(s.BuildScript) != "build" || deref(s.OutputDirectory) != ".next" || s.PackageManager != "npm" {
		t.Errorf("started with %+v", s)
	}
	if calls := strings.Join(api.calls, "\n"); !strings.Contains(calls, "settings-from-repo acme/site@master inst-1") || strings.Contains(calls, "creds") {
		t.Errorf("calls:\n%s", calls)
	}
	wipe := func() { _ = os.RemoveAll(filepath.Join(dir, site.Dir)) } // settings remembered by a successful build
	wipe()
	// Flags still win over detection.
	api = gitFake(t)
	api.fromRepo = NodeSettings{AppType: "next", NodeVersion: 22, RootDirectory: ".", BuildScript: strPtr("build"), OutputDirectory: strPtr(".next")}
	_, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Source: SourceGit, Node: NodeOverrides{NodeVersion: 20, OutputDirectory: "out"}})
	if err != nil || api.started.NodeVersion != 20 || deref(api.started.OutputDirectory) != "out" || api.started.AppType != "next" {
		t.Errorf("err=%v started with %+v", err, api.started)
	}
	wipe()
	// Detection failing is reported; without a terminal the framework must then come from a flag.
	api = gitFake(t)
	api.fromRepoErr = errors.New("[Hosting:9999] branch has no package.json (HTTP 422)")
	_, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Source: SourceGit})
	if err == nil || !strings.Contains(err.Error(), "--app-type") {
		t.Errorf("err = %v", err)
	}
	wipe()
	api = gitFake(t)
	api.fromRepoErr = errors.New("[Hosting:9999] branch has no package.json (HTTP 422)")
	_, err = Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Source: SourceGit, Node: NodeOverrides{AppType: "next", BuildScript: "build"}})
	if err != nil || api.started.AppType != "next" {
		t.Errorf("err=%v started with %+v", err, api.started)
	}
}

func strPtr(s string) *string { return &s }

func TestGitDeployWarnsAboutUnpushedWork(t *testing.T) {
	fastPolling(t)
	api := gitFake(t)
	dir, repo := gitProject(t)
	// Clean and pushed: nothing to say.
	ui := &fakeUI{t: t}
	if _, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com", Source: SourceGit}); err != nil {
		t.Fatal(err)
	}
	if infos := strings.Join(ui.infos, "\n"); strings.Contains(infos, "push to deploy") {
		t.Errorf("warned on a clean checkout: %v", ui.infos)
	}
	// Uncommitted and unpushed work: a warning, and the deploy still runs.
	write(t, dir, "draft.html", "")
	write(t, dir, "committed.html", "")
	commitAll(t, repo, "more") // draft.html is committed too; make a new draft
	write(t, dir, "draft2.html", "")
	ui = &fakeUI{t: t}
	if _, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com", Source: SourceGit}); err != nil {
		t.Fatal(err)
	}
	infos := strings.Join(ui.infos, "\n")
	if !strings.Contains(infos, "commits not pushed to origin/master and uncommitted changes") || !strings.Contains(infos, "push to deploy your latest code") {
		t.Errorf("infos:\n%s", infos)
	}
	if len(ui.asked) != 0 {
		t.Errorf("asked %v; the warning must not prompt", ui.asked)
	}
}

func TestSourceAskedOnceThenRemembered(t *testing.T) {
	fastPolling(t)
	api := gitFake(t)
	dir, _ := gitProject(t)
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{true /* type */, true /* replace */}, selects: []int{0 /* git */}}
	if _, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"}); err != nil {
		t.Fatal(err)
	}
	if len(ui.asked) != 3 || !strings.Contains(ui.asked[2], "How do you want to deploy") {
		t.Errorf("asked = %v", ui.asked)
	}
	again := &fakeUI{t: t, interactive: true}
	res, err := Run(context.Background(), api, again, Options{Path: dir})
	if err != nil {
		t.Fatal(err)
	}
	if res.Source != SourceGit || len(again.asked) != 0 {
		t.Errorf("redeploy: source=%s asked=%v", res.Source, again.asked)
	}
	// --source upload overrides for one run and is then remembered as the new answer.
	res, err = Run(context.Background(), api, &fakeUI{t: t, interactive: true}, Options{Path: dir, Source: SourceUpload})
	if err != nil || res.Source != SourceUpload || res.Archive == "" {
		t.Errorf("upload override: %+v %v", res, err)
	}
}

func TestNoInstallationFallsBackToUpload(t *testing.T) {
	fastPolling(t)
	api := gitFake(t)
	api.repos = nil
	dir, _ := gitProject(t)
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{true /* type */, true /* replace */}, selects: []int{0 /* upload */}}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Source != SourceUpload || len(ui.asked) != 3 || !strings.Contains(ui.asked[2], "Git checkout of acme/site, but no Git account linked in your hPanel account") {
		t.Errorf("source=%s asked=%v", res.Source, ui.asked)
	}
	// Non-interactive never asks and never looks installations up.
	api.calls = nil
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Source: SourceUpload}); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "installations") {
		t.Error("installations were listed for an upload")
	}
	_ = os.Remove(filepath.Join(dir, site.Dir))
}

// stubBrowser returns the URLs the browser was asked to open.
func stubBrowser(t *testing.T) *[]string {
	t.Helper()
	var opened []string
	prev := OpenBrowser
	OpenBrowser = func(url string) error { opened = append(opened, url); return nil }
	t.Cleanup(func() { OpenBrowser = prev })
	return &opened
}

func TestNoInstallationOffersToConnectInHPanel(t *testing.T) {
	fastPolling(t)
	opened := stubBrowser(t)
	api := gitFake(t)
	api.repos = nil
	dir, _ := gitProject(t)

	// Connect, open hPanel, check once, then stop: upload.
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{true /* type */, true /* replace */, true /* open browser */, true /* check again */, false /* stop */}, selects: []int{1 /* connect */}}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	infos := strings.Join(ui.infos, "\n")
	if res.Source != SourceUpload || len(*opened) != 1 || (*opened)[0] != HPanelIntegrations {
		t.Errorf("source=%s opened=%v", res.Source, *opened)
	}
	if !strings.Contains(infos, "Still no Git account linked in your hPanel account") || !strings.Contains(infos, "Uploading the folder instead") {
		t.Errorf("infos:\n%s", infos)
	}
	// The steps are printed before the browser question.
	steps, still := strings.Index(infos, "Account integrations page"), strings.Index(infos, "Still no Git account")
	if steps < 0 || still < steps || !strings.Contains(infos[steps:still], "at least to acme/site") || !strings.Contains(infos[steps:still], "Opened hPanel") || !strings.Contains(ui.asked[3], "Open "+HPanelIntegrations) {
		t.Errorf("asked=%v infos:\n%s", ui.asked, infos)
	}
	if n := strings.Count(strings.Join(api.calls, "\n"), "installations"); n != 2 {
		t.Errorf("installations listed %d times, want 2 (before and after the check)", n)
	}
	_ = os.RemoveAll(filepath.Join(dir, site.Dir))

	// Connected during the check: the deploy goes through Git.
	api = gitFake(t)
	api.repos = nil
	api.onInstallations = func() {
		if strings.Count(strings.Join(api.calls, "\n"), "installations") == 2 {
			api.repos = []Repository{{Name: "site", FullName: "acme/site", Owner: "acme", DefaultBranch: "master"}}
		}
	}
	ui = &fakeUI{t: t, interactive: true, confirms: []bool{true /* type */, true /* replace */, true /* open browser */, true /* check again */}, selects: []int{1 /* connect */}}
	res, err = Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Source != SourceGit || res.Repository != "acme/site" || api.autoDeploy == nil {
		t.Errorf("result=%+v autoDeploy=%+v", res, api.autoDeploy)
	}
	_ = os.RemoveAll(filepath.Join(dir, site.Dir))

	// --source git in a terminal gets the same walk-through; declining the browser
	// opens nothing, and stopping is an error, not an upload.
	api = gitFake(t)
	api.repos = nil
	ui = &fakeUI{t: t, interactive: true, confirms: []bool{true /* type */, true /* replace */, false /* no browser */, false /* stop */}}
	_, err = Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com", Source: SourceGit})
	if err == nil || !strings.Contains(err.Error(), "no Git account linked in your hPanel account can read acme/site") {
		t.Errorf("err = %v", err)
	}
	if len(ui.asked) != 4 || !strings.Contains(ui.asked[3], "Check again") || len(*opened) != 2 || !strings.Contains(strings.Join(ui.infos, "\n"), "yourself") {
		t.Errorf("asked = %v opened = %v infos = %v", ui.asked, *opened, ui.infos)
	}
}

func TestDetachedCheckoutUploadsUnlessGitIsAskedFor(t *testing.T) {
	fastPolling(t)
	api := gitFake(t)
	dir, repo := gitProject(t)
	head, err := repo.Head()
	if err != nil {
		t.Fatal(err)
	}
	// What actions/checkout leaves behind: HEAD at a commit, not a branch.
	wt, err := repo.Worktree()
	if err != nil {
		t.Fatal(err)
	}
	if err := wt.Checkout(&git.CheckoutOptions{Hash: head.Hash()}); err != nil {
		t.Fatal(err)
	}
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Source != SourceUpload || res.Archive == "" {
		t.Errorf("result = %+v", res)
	}
	_ = os.RemoveAll(filepath.Join(dir, site.Dir))
	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Source: SourceGit}); err == nil || !strings.Contains(err.Error(), "detached HEAD") {
		t.Errorf("--source git on a detached checkout: %v", err)
	}
}

func TestDryRunNeitherAsksForSourceNorConnectsGit(t *testing.T) {
	opened := stubBrowser(t)
	api := gitFake(t)
	dir, _ := gitProject(t)
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{true /* type */}}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com", DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	if res.Source != SourceUpload || strings.Contains(strings.Join(api.calls, "\n"), "installations") {
		t.Errorf("source=%s calls=%v", res.Source, api.calls)
	}
	// --source git looks the installation up but never walks through connecting one.
	api.repos = nil
	if _, err := Run(context.Background(), api, &fakeUI{t: t, interactive: true, confirms: []bool{true}}, Options{Path: dir, Domain: "example.com", DryRun: true, Source: SourceGit}); err == nil {
		t.Error("dry run with no installation for the repository succeeded")
	}
	if len(*opened) != 0 {
		t.Errorf("dry run opened %v", *opened)
	}
}

func TestInstallationMustMatchTheProvider(t *testing.T) {
	api := gitFake(t)
	api.installations = []Installation{{UUID: "lab", Provider: "gitlab", Status: "active"}, {UUID: "hub", Provider: "github", Status: "active"}}
	// Both installations list acme/site; the checkout's remote is on GitHub.
	inst, err := findInstallation(context.Background(), api, &Repo{Host: "github.com", Owner: "acme", Name: "site"})
	if err != nil || inst == nil || inst.UUID != "hub" {
		t.Errorf("installation = %+v, %v", inst, err)
	}
	// An SSH alias (git@github-work:acme/site.git) and ssh.github.com still find GitHub first.
	for _, host := range []string{"github-work", "ssh.github.com"} {
		inst, err = findInstallation(context.Background(), api, &Repo{Host: host, Owner: "acme", Name: "site"})
		if err != nil || inst == nil || inst.UUID != "hub" {
			t.Errorf("%s: installation = %+v, %v", host, inst, err)
		}
	}
	// A host that names no provider falls back to any installation listing the repository.
	api.installations = []Installation{{UUID: "lab", Provider: "gitlab", Status: "active"}}
	inst, err = findInstallation(context.Background(), api, &Repo{Host: "work-git", Owner: "acme", Name: "site"})
	if err != nil || inst == nil || inst.UUID != "lab" {
		t.Errorf("unnamed host: installation = %+v, %v", inst, err)
	}
}

func TestCheckoutBehindOriginIsNotUnpushed(t *testing.T) {
	dir, repo := gitProject(t)
	first, _ := repo.Head()
	write(t, dir, "next.html", "")
	commitAll(t, repo, "next")
	push(t, repo) // origin has "next"
	wt, err := repo.Worktree()
	if err != nil {
		t.Fatal(err)
	}
	if err := wt.Reset(&git.ResetOptions{Commit: first.Hash(), Mode: git.HardReset}); err != nil {
		t.Fatal(err)
	}
	if work := unpushedWork(dir, "master"); strings.Contains(work, "not pushed") {
		t.Errorf("behind origin reported as %q", work)
	}
}
