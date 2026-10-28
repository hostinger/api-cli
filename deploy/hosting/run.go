package hosting

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hostinger/api-cli/site"
)

// NodeOverrides are build-setting flags; zero values keep what was detected.
type NodeOverrides = site.NodeSettings

// Options is one deploy; most fields mirror the up command's flags.
type Options struct {
	// Path is a directory, or an archive (then Type is required).
	Path       string
	Domain     string
	Create     bool
	OrderID    int
	Datacenter string
	Type       Type
	Yes        bool
	Settings   bool
	Source     Source
	DryRun     bool
	NoWait     bool
	Node       NodeOverrides
	// Headless is the kind of headless website to create (see HeadlessKinds).
	Headless string
}

// Result is what a deploy produced. Fields Outcome already reports
// are left out of its JSON.
type Result struct {
	Domain     string `json:"domain"`
	Username   string `json:"username"`
	OrderID    int    `json:"orderId,omitempty"`
	Type       Type   `json:"type"`
	URL        string `json:"-"`
	Created    bool   `json:"created"`
	Headless   bool   `json:"headless,omitempty"`
	Source     Source `json:"source"`
	Repository string `json:"repository,omitempty"`
	Branch     string `json:"branch,omitempty"`
	Archive    string `json:"archive,omitempty"`
	BuildUUID  string `json:"buildUuid,omitempty"`
	BuildState string `json:"buildState,omitempty"`
	// Settings are what the build ran with. In a dry run, only what is known
	// before upload; SettingsSource says where the rest will come from.
	Settings       *NodeSettings `json:"settings,omitempty"`
	SettingsSource string        `json:"settingsSource,omitempty"`
	SkippedFiles   []string      `json:"skippedFiles,omitempty"`
	Symlinks       []string      `json:"symlinks,omitempty"` // left out: not followed
	IgnoredFiles   int           `json:"ignoredFiles,omitempty"`
	ProjectFile    string        `json:"-"`
	NewLink        bool          `json:"-"`
	DryRun         bool          `json:"-"`
	Files          int           `json:"files,omitempty"`
	// OverwritesContent is set by a dry run when the first deploy to this
	// website would replace content: the real run asks, or needs --yes.
	OverwritesContent bool `json:"overwritesContent,omitempty"`
	// SuggestedHeadless is the headless website a dry run would offer to create
	// in a terminal (pass it as --headless); "" is a regular one.
	SuggestedHeadless string `json:"suggestedHeadless,omitempty"`
	// Reachable is false when https://Domain has no DNS yet.
	Reachable     bool   `json:"-"`
	ReachableNote string `json:"-"`
	// Next says how to follow a build left running (--no-wait).
	Next string `json:"-"`
}

// Tests lower these.
var (
	SetupPollInterval = 3 * time.Second
	SetupTimeout      = 15 * time.Minute
	BuildPollInterval = 5 * time.Second
	BuildPollMax      = 20 * time.Second
	BuildTimeout      = 20 * time.Minute
	Now               = time.Now
)

// Run deploys opts.Path: find or create the website, detect, pack, upload,
// deploy or build, and link the directory to the website.
func Run(ctx context.Context, a API, ui UI, opts Options) (*Result, error) {
	d := &deployment{ctx: ctx, a: a, ui: ui, opts: opts}
	if err := d.locate(); err != nil {
		return nil, err
	}
	detected, err := d.detect()
	if err != nil {
		return nil, err
	}
	if err := d.list(); err != nil {
		return nil, err
	}
	if err := d.target(detected); err != nil {
		return nil, d.keepCreated(err)
	}
	if err := d.prepare(); err != nil {
		return nil, d.keepCreated(err)
	}
	if opts.DryRun {
		return d.dryRun(), nil
	}
	// Link before uploading, so a failed deploy reruns against the same website.
	d.link(d.remembered, d.rememberedSource)
	if d.source == SourceGit {
		return d.deployFromGit()
	}
	return d.deployUpload()
}

// deployment is one Run in progress.
type deployment struct {
	ctx  context.Context
	a    API
	ui   UI
	opts Options

	project *site.Project // the link found from the current directory, if any
	linkDir string
	src     string // the directory or archive to deploy
	isDir   bool

	website  *site.Website
	created  bool
	headless bool // from the website, or the link: a built headless site can report nodejs
	typ      Type

	source           Source
	repo             *Repo
	packed           *Listing // for Git, only listed to check build settings against
	remembered       *site.NodeSettings
	rememberedSource Source
	lastBuild        *site.BuildRef // the build this run started, remembered for the build commands
	result           *Result
}

// locate finds the link and the path to deploy.
func (d *deployment) locate() error {
	linkDir, err := Getwd()
	if err != nil {
		return err
	}
	linkDir, err = filepath.Abs(linkDir)
	if err != nil {
		return err
	}
	project, projectDir, err := site.Find(linkDir)
	if err != nil {
		return fmt.Errorf("the project link could not be read: %v (fix the file, or run `hostinger unlink`)", err)
	}
	if project != nil {
		linkDir = projectDir
	}
	d.project, d.linkDir = project, linkDir
	path, fromLink := d.opts.Path, ""
	if (path == "" || path == ".") && project != nil {
		// A linked project deploys the folder it was linked with, from anywhere inside it.
		// The link can come with a cloned repository: never pack outside the project.
		if escapes(project.Path) || filepath.IsAbs(project.Path) {
			return fmt.Errorf("the linked folder %q in %s must be a relative path inside the project", project.Path, site.Path(projectDir))
		}
		path = filepath.Join(projectDir, project.Path)
		// A symlinked folder can point anywhere (public -> ~/.ssh): compare real paths.
		if !insideReal(projectDir, path) {
			return fmt.Errorf("the linked folder %q in %s leads outside the project through a symlink", project.Path, site.Path(projectDir))
		}
		if project.Path != "" && project.Path != "." {
			fromLink = " (the folder " + site.Path(projectDir) + " links; run `hostinger link` again if it moved)"
		}
		if cwd, err := Getwd(); err == nil && filepath.Clean(cwd) != filepath.Clean(projectDir) {
			d.ui.Info("Deploying the linked folder " + path)
		}
	}
	if path == "" {
		path = "."
	}
	if d.src, err = filepath.Abs(path); err != nil {
		return err
	}
	if err := checkNodeOverrides(d.opts.Node); err != nil {
		return err
	}
	info, err := os.Stat(d.src)
	if err != nil {
		return fmt.Errorf("path not found: %s%s", d.src, fromLink)
	}
	d.isDir = info.IsDir()
	if !d.isDir {
		if !IsArchive(d.src) {
			return fmt.Errorf("%s is neither a directory nor an archive (zip, tar, tar.gz, tgz, 7z, gz)%s", d.src, fromLink)
		}
		if d.opts.Type == "" {
			return errors.New("an archive cannot be inspected: pass --type static or --type nodejs")
		}
	}
	if d.opts.Headless != "" {
		// Before a headless website is created for a folder that cannot go there.
		return checkHeadless(d.src, d.isDir, d.opts.Type, "a headless website")
	}
	return nil
}

// detect runs before the website is chosen: an undeployable folder must fail
// before a website is created. The result is a hint for the website picker.
func (d *deployment) detect() (Type, error) {
	if !d.isDir || d.opts.Type != "" {
		return "", nil
	}
	detected, note, err := Detect(d.src)
	// The remembered type is checked later, in chooseType, but only applies
	// when the link picks the website.
	if d.project != nil && d.project.Type != "" && d.opts.Domain == "" && !d.opts.Create {
		return detected, nil
	}
	if err != nil {
		return "", err
	}
	if err := undeployable(d.src, detected); err != nil {
		return "", err
	}
	if note != "" {
		d.ui.Info(note)
	}
	return detected, nil
}

// target picks the website and the type, and asks before replacing content.
func (d *deployment) target(detected Type) error {
	opts, project := d.opts, d.project
	suggested := ""
	if d.isDir {
		suggested = SuggestHeadless(d.src)
	}
	website, created, err := resolveWebsite(d.ctx, d.a, d.ui, opts, project, detected, suggested)
	if err != nil {
		return err
	}
	d.website, d.created = website, created
	linked := project != nil && strings.EqualFold(project.Domain, website.Domain)
	d.headless = website.IsHeadless() || (linked && project.Headless)
	if opts.Headless != "" && !d.headless {
		return fmt.Errorf("%s already exists and is not a headless website: pass --create for a new headless website, or leave --headless out", website.Domain)
	}
	d.result = &Result{Domain: website.Domain, Username: website.Username, OrderID: website.OrderID, URL: "https://" + website.Domain, Created: created, Headless: d.headless, Reachable: true}
	if !opts.DryRun {
		if err := waitForSetup(d.ctx, d.a, d.ui, website.Domain, created); err != nil {
			return err
		}
	}

	fromLink := opts.Domain == "" && !opts.Create && project != nil && strings.EqualFold(project.Domain, website.Domain)
	d.typ = opts.Type
	switch {
	case d.headless:
		// Static files would land in the document root, next to the backend.
		if err := checkHeadless(d.src, d.isDir, d.typ, website.Domain); err != nil {
			return err
		}
		d.typ = NodeJS
	case d.typ == "":
		if d.typ, err = chooseType(d.src, d.ui, opts, project, website.Domain); err != nil {
			return err
		}
	}
	if created && opts.DryRun && opts.Headless == "" && d.typ == NodeJS {
		d.result.SuggestedHeadless = suggested
	}
	line := "Deploying to " + website.Domain + " as " + string(d.typ)
	var notes []string
	if d.headless {
		notes = append(notes, "headless")
	}
	switch {
	case fromLink:
		notes = append(notes, "linked")
	case created:
		notes = append(notes, "new")
	}
	if len(notes) > 0 {
		line += " (" + strings.Join(notes, ", ") + ")"
	}
	d.ui.Info(line)
	if d.headless && !created {
		d.ui.Info("This deploys the website's frontend; its backend (such as WordPress in /cms) is left alone")
	}
	if project != nil && !strings.EqualFold(project.Domain, website.Domain) && !shouldLink(project, website, opts) {
		d.ui.Info("One-off deploy: this directory stays linked to " + project.Domain)
	}
	switch {
	case website.WebsiteType == "nodejs" && d.typ == Static:
		// Checked even when linked: the app keeps running whatever the link says.
		return confirmStaticOverApp(d.ui, website, opts)
	case !fromLink && !created && opts.DryRun:
		// The plan says what the real run would ask, or need --yes for without a terminal.
		d.result.OverwritesContent, _ = hasContent(d.ctx, d.a, website, d.headless)
	case !fromLink && !created:
		return confirmOverwrite(d.ctx, d.a, d.ui, website, opts, d.headless)
	}
	return nil
}

// prepare starts the result, chooses upload or Git, and lists the files.
func (d *deployment) prepare() error {
	w := d.website
	d.result.Type = d.typ
	if !d.opts.DryRun {
		d.result.Reachable, d.result.ReachableNote = reachability(d.ctx, w.Domain)
	}
	if d.project != nil && strings.EqualFold(d.project.Domain, w.Domain) {
		d.remembered, d.rememberedSource = d.project.Node, Source(d.project.Source)
	}
	var err error
	if d.source, d.repo, err = chooseSource(d.ctx, d.a, d.ui, d.opts, d.src, d.isDir, d.rememberedSource); err != nil {
		return err
	}
	d.result.Source = d.source
	if d.source == SourceGit && d.typ == Static && w.WebsiteType == "nodejs" {
		// Auto-deploy saves settings but clones nothing onto a Node.js website.
		return errors.New(w.Domain + " runs a Node.js app, where a static Git deploy deploys nothing: deploy with --source upload, or --type nodejs to build the app from Git")
	}
	if !d.isDir {
		return nil
	}
	if d.typ == Static {
		d.packed.leaveOutStatic()
	}
	d.result.Files, d.result.SkippedFiles, d.result.IgnoredFiles = len(d.packed.Files), d.packed.Skipped, d.packed.Ignored
	d.result.Symlinks = d.packed.Symlinks
	if d.source != SourceGit {
		return d.packed.nonEmpty()
	}
	return nil
}

// list packs the file list before any website is created: an empty folder
// fails here unless it may deploy from Git, which ships what is pushed.
func (d *deployment) list() error {
	if !d.isDir {
		return nil
	}
	var err error
	if d.packed, err = ListFiles(d.ctx, d.src); err != nil {
		return err
	}
	d.packed.Dir = d.src
	git := d.opts.Source == SourceGit || (d.opts.Source == "" && d.project != nil && d.project.Source == string(SourceGit))
	if git {
		return nil
	}
	return d.packed.nonEmpty()
}

// keepCreated links a website this run created before failing, so the next
// run deploys to it instead of creating another.
func (d *deployment) keepCreated(err error) error {
	if !d.created || d.opts.DryRun || d.result == nil {
		return err
	}
	d.link(nil, d.opts.Source)
	again := "run again to deploy to it"
	switch {
	case d.result.ProjectFile == "": // nothing remembers it
		again = "run again with --domain " + d.website.Domain + " to deploy to it"
	case d.opts.Create:
		again = "run again without --create to deploy to it"
	}
	return fmt.Errorf("%w (%s was created; %s)", err, d.website.Domain, again)
}

// link writes the directory's link when this deploy should move or keep it.
func (d *deployment) link(node *site.NodeSettings, source Source) {
	if d.isDir && shouldLink(d.project, d.website, d.opts) {
		build := d.lastBuild
		if build == nil && d.project != nil && strings.EqualFold(d.project.Domain, d.website.Domain) {
			build = d.project.Build // this run started none: keep the last one
		}
		d.result.ProjectFile = saveLink(d.linkDir, d.src, d.website, d.typ, d.headless, node, source, build, d.ui)
		d.result.NewLink = d.result.ProjectFile != "" && (d.project == nil || !strings.EqualFold(d.project.Domain, d.website.Domain))
	}
}

// dryRun reports the plan, with what is known of the build settings.
func (d *deployment) dryRun() *Result {
	result := d.result
	result.DryRun = true
	switch {
	case d.source == SourceGit:
		result.Repository, result.Branch = d.repo.FullName(), d.repo.Branch
	case d.isDir:
		result.Archive = ArchiveName(d.src, Now())
	default:
		result.Archive = uploadName(d.src)
	}
	if d.typ != NodeJS {
		return result
	}
	s := NodeSettings{}
	if d.remembered != nil {
		s = applyOverrides(s, *d.remembered)
	}
	s = applyOverrides(s, d.opts.Node)
	detected := "detected from the uploaded files at deploy"
	if d.source == SourceGit {
		detected = "detected from the repository at deploy"
	}
	switch {
	case s == (NodeSettings{}):
		result.SettingsSource = detected // nothing is known yet: no half-empty settings
	case d.remembered != nil && d.opts.Node != (NodeOverrides{}):
		result.Settings, result.SettingsSource = &s, "remembered, with flags applied; the rest "+detected
	case d.remembered != nil:
		result.Settings, result.SettingsSource = &s, "remembered from the last deploy"
	default:
		result.Settings, result.SettingsSource = &s, "from flags; the rest "+detected
	}
	return result
}

// deployFromGit turns on auto-deploy for a static site, or builds the pushed
// branch of a Node.js app.
func (d *deployment) deployFromGit() (*Result, error) {
	ctx, a, ui, w, repo := d.ctx, d.a, d.ui, d.website, d.repo
	d.result.Repository, d.result.Branch = repo.FullName(), repo.Branch
	// A Git deploy builds what is pushed; warn when the checkout is ahead.
	if work := unpushedWork(d.src, repo.Branch); work != "" {
		ui.Info("This checkout has " + work + ". A Git deploy builds origin/" + repo.Branch + ", so push to deploy your latest code.")
	}
	if d.typ == Static {
		step := ui.Start("Linking Git " + repo.FullName() + "@" + repo.Branch)
		err := a.UpdateGitAutoDeploy(ctx, w.Username, w.Domain, GitSettings{InstallationUUID: repo.installUID, Owner: repo.Owner, Repository: repo.Name, Branch: repo.Branch, IsEnabled: true})
		if err != nil {
			step.Fail("Could not set up the Git deploy")
			return nil, err
		}
		step.Done("Deployed from Git; every push to " + repo.Branch + " redeploys")
		d.link(d.remembered, d.source)
		return d.result, nil
	}
	stored, err := a.BuildSettings(ctx, w.Username, w.Domain)
	if err != nil {
		return nil, err
	}
	// Stored settings, else the API's detection from the pushed branch.
	var detected NodeSettings
	if stored != nil {
		detected = *stored
	} else {
		step := ui.Start("Reading build settings from " + repo.FullName() + "@" + repo.Branch)
		detected, err = a.BuildSettingsFromRepository(ctx, w.Username, w.Domain, repo)
		if err != nil {
			step.Fail("Could not detect the build settings: " + err.Error())
			detected = NodeSettings{}
		} else {
			step.Done("Build settings read")
		}
	}
	start := func(s NodeSettings) (Build, error) {
		if err := a.UpdateBuildSettings(ctx, w.Username, w.Domain, s); err != nil {
			return Build{}, err
		}
		return a.StartGitBuild(ctx, w.Username, w.Domain, s, repo)
	}
	return d.finishBuild(runBuild(ctx, a, ui, w, detected, start, d.opts, d.remembered, d.packed))
}

// deployUpload packs the directory (an archive goes as is), uploads it, and
// deploys the files or builds the app.
func (d *deployment) deployUpload() (*Result, error) {
	ctx, a, ui, w, result := d.ctx, d.a, d.ui, d.website, d.result
	archive := d.src
	if d.isDir {
		step := ui.Start("Packing " + d.src)
		tmp, err := os.MkdirTemp("", "hostinger-up-")
		if err != nil {
			step.Fail("Packing failed")
			return nil, err
		}
		defer os.RemoveAll(tmp)
		archive = filepath.Join(tmp, ArchiveName(d.src, Now()))
		if err := Zip(d.ctx, d.src, d.packed.Files, archive); err != nil {
			step.Fail("Packing failed")
			return nil, err
		}
		step.Done(packedSummary(d.packed))
		if skipped := d.packed.Skipped; len(skipped) > 0 {
			hint := "A static site cannot read them; keep secrets out of the document root."
			if d.typ == NodeJS {
				hint = "Set them with `hostinger hosting nodejs replace-environment-variables`."
			}
			ui.Info("Left out (secrets): " + strings.Join(skipped, ", ") + ". " + hint)
		}
		if links := d.packed.Symlinks; len(links) > 0 {
			ui.Info("Left out (symlinks are not followed): " + strings.Join(links, ", ") + ". Copy what they point to into the folder to deploy it.")
		}
		result.Archive = filepath.Base(archive)
	} else {
		result.Archive = uploadName(archive)
	}

	step := ui.Start("Uploading " + result.Archive)
	creds, err := a.UploadCredentials(ctx, w.Username, w.Domain)
	if err != nil {
		step.Fail("Upload failed")
		return nil, err
	}
	err = Upload(ctx, creds, archive, result.Archive, func(sent, total int64) {
		step.Update(fmt.Sprintf("Uploading %s (%d%%)", result.Archive, sent*100/max(total, 1)))
	})
	if err != nil {
		step.Fail("Upload failed")
		return nil, err
	}
	step.Done("Uploaded " + result.Archive)

	if d.typ == NodeJS {
		step := ui.Start("Reading build settings from the archive")
		detected, err := a.BuildSettingsFromArchive(ctx, w.Username, w.Domain, result.Archive)
		if err != nil {
			step.Fail("Could not detect build settings")
			return nil, err
		}
		step.Done("Build settings read")
		start := func(s NodeSettings) (Build, error) {
			return a.StartBuild(ctx, w.Username, w.Domain, s, result.Archive)
		}
		return d.finishBuild(runBuild(ctx, a, ui, w, detected, start, d.opts, d.remembered, d.packed))
	}
	step = ui.Start("Deploying files")
	if err := a.DeployArchive(ctx, w.Username, w.Domain, result.Archive); err != nil {
		step.Fail("Deploy failed")
		return nil, err
	}
	step.Done("Files deployed")
	d.link(nil, d.source)
	return result, nil
}

// finishBuild records a Node.js build and links the directory, remembering
// only settings that built.
func (d *deployment) finishBuild(build Build, used *NodeSettings, err error) (*Result, error) {
	if build.UUID != "" {
		d.result.BuildUUID, d.result.BuildState = build.UUID, build.State
		d.lastBuild = &site.BuildRef{UUID: build.UUID, State: build.State, At: time.Now()}
	}
	d.result.Settings = used
	switch {
	case err != nil:
		d.link(d.remembered, d.source)
		return d.result, err
	case used == nil: // --no-wait
		d.link(d.remembered, d.source)
		d.result.Next = followHint(d.website.Domain, build.UUID, d.result.ProjectFile != "")
		return d.result, nil
	}
	d.link(toRemembered(used), d.source)
	return d.result, nil
}

// saveLink writes the link and returns its path, or "" when it failed.
func saveLink(linkDir, src string, website *site.Website, typ Type, headless bool, node *site.NodeSettings, source Source, build *site.BuildRef, ui UI) string {
	rel := ""
	r, err := filepath.Rel(linkDir, src)
	if err != nil || escapes(r) || filepath.IsAbs(r) {
		// The next run could not deploy it from here: a link must stay inside the project.
		ui.Info("Not linked: " + src + " is outside " + linkDir + "; run `hostinger up` from that folder to link it")
		return ""
	}
	if r != "." {
		rel = filepath.ToSlash(r)
	}
	p := &site.Project{Kind: "website", Path: rel, OrderID: website.OrderID, Domain: website.Domain, Username: website.Username, Type: string(typ), Headless: headless, Node: node, Source: string(source), Build: build}
	if err := site.Save(linkDir, p); err != nil {
		ui.Info("Could not write " + site.Path(linkDir) + ": " + err.Error())
		return ""
	}
	return site.Path(linkDir)
}

func packedSummary(l *Listing) string {
	msg := fmt.Sprintf("Packed %d files", len(l.Files))
	var notes []string
	if l.Ignored > 0 {
		notes = append(notes, fmt.Sprintf("%d left out by %s", l.Ignored, l.IgnoreSource))
	}
	if len(notes) > 0 {
		msg += " (" + strings.Join(notes, ", ") + ")"
	}
	return msg
}

// insideReal reports whether path, with symlinks resolved, is root or below it.
// A path that does not exist yet is left to the caller's not-found error.
func insideReal(root, path string) bool {
	realRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return false
	}
	real, err := filepath.EvalSymlinks(path)
	if err != nil {
		return errors.Is(err, fs.ErrNotExist)
	}
	rel, err := filepath.Rel(realRoot, real)
	return err == nil && !escapes(rel) && !filepath.IsAbs(rel)
}

// Tests replace Getwd.
var Getwd = os.Getwd

// shouldLink says whether to link the directory to this website. --domain on
// an already-linked directory is a one-off; `hostinger link` changes the link.
func shouldLink(project *site.Project, website *site.Website, opts Options) bool {
	if project == nil || strings.EqualFold(project.Domain, website.Domain) {
		return true
	}
	return opts.Domain == "" || opts.Create
}

// rememberedOrDetected keeps the remembered type unless the folder now looks
// like the other kind: a Node.js app deployed as static publishes its source.
func rememberedOrDetected(src string, ui UI, opts Options, remembered Type) (Type, error) {
	detected, _, err := Detect(src)
	if err != nil || detected == remembered || (detected != Static && detected != NodeJS) {
		return remembered, nil
	}
	msg := "This folder looks like a " + typeLabel(detected) + ", but it was last deployed as " + string(remembered)
	switch {
	case opts.Yes:
		ui.Info(msg + "; deploying it as " + string(detected))
		return detected, nil
	case !ui.Interactive():
		return "", fmt.Errorf("%s: pass --type %s or --type %s", msg, detected, remembered)
	}
	ok, err := ui.Confirm(msg+". Deploy it as "+string(detected)+"?", true)
	if err != nil {
		return "", err
	}
	if ok {
		return detected, nil
	}
	return remembered, nil
}

// checkHeadless refuses what cannot go onto a headless website (target names
// it): it serves a Node.js frontend, so a folder needs a package.json.
func checkHeadless(src string, isDir bool, typ Type, target string) error {
	if typ == Static {
		return fmt.Errorf("%s serves a Node.js frontend, so it takes no static deploy (its document root holds the backend, such as WordPress in /cms): leave --type out, or pass --type nodejs", target)
	}
	if !isDir {
		return nil
	}
	if _, err := os.Stat(filepath.Join(src, "package.json")); err != nil {
		return fmt.Errorf("%s serves a Node.js frontend, and %s has no package.json: deploy a Node.js project (such as Next.js, Nuxt, Astro or Vite), or deploy these files to a website that is not headless", target, src)
	}
	return nil
}

func undeployable(src string, detected Type) error {
	switch detected {
	case WordPressSite, WordPressPlugin, WordPressTheme:
		return fmt.Errorf("%s looks like a %s, which this command cannot deploy yet; use the hostinger wordpress commands", src, detected)
	}
	return nil
}

func chooseType(src string, ui UI, opts Options, project *site.Project, domain string) (Type, error) {
	if project != nil && strings.EqualFold(project.Domain, domain) {
		for _, t := range DeployableTypes {
			if string(t) == project.Type {
				return rememberedOrDetected(src, ui, opts, t)
			}
		}
	}
	detected, _, err := Detect(src) // the note was shown before the website was chosen
	if err != nil {
		return "", err
	}
	if err := undeployable(src, detected); err != nil {
		return "", err
	}
	if opts.Yes || !ui.Interactive() {
		return detected, nil
	}
	ok, err := ui.Confirm("This looks like a "+typeLabel(detected)+". Deploy it as "+string(detected)+"?", true)
	if err != nil {
		return "", err
	}
	if ok {
		return detected, nil
	}
	labels := []string{"static (HTML, CSS, JavaScript, PHP served as files)", "nodejs (built and run by Node.js)"}
	def := 0
	if other(detected) == NodeJS {
		def = 1 // the person said no: offer the other one first
	}
	idx, err := ui.Select("Deploy as", labels, def)
	if err != nil {
		return "", err
	}
	return DeployableTypes[idx], nil
}

func other(t Type) Type {
	if t == NodeJS {
		return Static
	}
	return NodeJS
}

func typeLabel(t Type) string {
	if t == NodeJS {
		return "Node.js app"
	}
	return "static site"
}
