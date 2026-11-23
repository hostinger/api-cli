package preview

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/hostinger/api-cli/deploy/hosting"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/hcdn-connector/pkg/connector"
)

// States of a remote deploy. hPanel hears each as the website event
// "preview.deploy." + state; --json prints it as a "deploy" event.
const (
	stateAccepted  = "accepted"
	stateSucceeded = "succeeded"
	stateFailed    = "failed"
)

// Refusals of a verified command, besides hosting's failure reasons. The
// edge's answer never reaches hPanel, so every one is reported as failed.
const (
	reasonTargetChanged = "target_changed"
	reasonBusy          = "busy"
)

// eventTimeout bounds one status event call.
const eventTimeout = 15 * time.Second

// Seams for tests.
var (
	// runDeploy is `hostinger up`'s deploy.
	runDeploy = func(ctx context.Context, ui hosting.UI, opts hosting.Options) (*hosting.Result, error) {
		return hosting.Run(ctx, hosting.NewAPI(site.Client()), ui, opts)
	}
	// sendEvent posts a website event through public-api. A failed call is
	// not retried.
	sendEvent = func(ctx context.Context, username, domain, name string, params map[string]string) error {
		path := "/api/hosting/v1/accounts/" + url.PathEscape(username) + "/websites/" + url.PathEscape(domain) + "/events/trigger"
		body := map[string]any{"event": map[string]any{"name": name, "params": params}}
		return site.Client().JSON(ctx, http.MethodPost, path, nil, body, nil)
	}
)

// remote runs the deploys hPanel asks a Preview for. Copies of one command
// arrive at once over the Preview's sessions on several PoPs; the first one
// runs and the others are answered as done.
type remote struct {
	project  *project // as found at the start, for hPanel's listing
	verifier verifier
	p        *printer
	// base outlives the Preview's stop signal: a deploy that the Preview's
	// end cancels still reports it.
	base context.Context

	mu       sync.Mutex
	hostname string
	seen     map[string]bool    // ids (jti) of the commands handled: one per click
	cancel   context.CancelFunc // the running deploy's; nil when none runs
	// cancelled: Ctrl+C cancelled the running deploy, so the next one stops
	// the Preview.
	cancelled bool
	wg        sync.WaitGroup // the running deploy and status events being sent
}

// newRemote returns the remote deploys of a Preview started with o, or nil
// when it takes none: --no-remote-deploy, --url (an origin, not this
// folder), --no-owner (no keys to check a command with), or a working
// directory that is no project hPanel can deploy.
func newRemote(base context.Context, o options, p *printer) *remote {
	if o.noRemote || o.url != "" || o.noOwner {
		return nil
	}
	proj := currentProject()
	if proj == nil {
		return nil
	}
	r := &remote{project: proj, p: p, base: base, seen: map[string]bool{}}
	r.verifier.now = p.now
	return r
}

// attach makes the Preview take commands: it adds the command feature, the
// project's metadata for hPanel, and the owner ticket hooks that verify
// commands.
func (r *remote) attach(qo *connector.QuickOptions, owner OwnerOptions) {
	owner.OnTicket = r.verifier.setKeys
	qo.Owner = r.verifier.owner(NewOwnerTicketFunc(owner))
	qo.OnCommand = r.handle
	qo.AppMeta = r.project.appMeta()
	ready := qo.OnReady
	qo.OnReady = func(t connector.QuickTunnel) {
		r.mu.Lock()
		r.hostname = t.Hostname
		r.mu.Unlock()
		if ready != nil {
			ready(t)
		}
	}
}

// handle is the connector's OnCommand. It runs on a session's request
// goroutine while the edge waits, at most 10 s, so it answers at once and
// the deploy runs on its own, outside ctx.
func (r *remote) handle(ctx context.Context, c connector.Command) error {
	cmd, err := r.verifier.verify(ctx, c.Body)
	if err != nil {
		return err
	}
	return r.start(cmd)
}

// start runs a verified command once: a copy of it is answered as done.
func (r *remote) start(cmd *command) error {
	cl := cmd.claims
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.seen[cl.Jti] {
		return nil
	}
	r.seen[cl.Jti] = true

	j := &job{r: r, cmd: cmd}
	if r.cancel != nil {
		return j.decline(http.StatusConflict, reasonBusy, "a deploy is already running")
	}
	// Read again: the folder may have been linked, or changed, since the start.
	p := currentProject()
	switch {
	case p == nil:
		return j.decline(http.StatusConflict, hosting.ReasonUnsupportedType, "this folder is no longer a project hPanel can deploy")
	case p.linkedElsewhere(cl.Username, cl.Domain):
		return j.decline(http.StatusConflict, reasonTargetChanged, "this folder is linked to "+p.link.Domain)
	}
	ctx, cancel := context.WithCancel(r.base)
	r.cancel = cancel
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		j.run(ctx, p)
		cancel()
		r.mu.Lock()
		r.cancel, r.cancelled = nil, false
		r.mu.Unlock()
	}()
	return nil
}

// cancelDeploy cancels the running deploy, and reports whether it did: the
// first Ctrl+C during a deploy ends the deploy, not the Preview.
func (r *remote) cancelDeploy() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.cancel == nil || r.cancelled {
		return false
	}
	r.cancel()
	r.cancelled = true
	r.p.say("Cancelling the deploy; the Preview keeps running. Press Ctrl+C again to stop the Preview.")
	return true
}

// close cancels a running deploy when the Preview ends, and waits for its
// status event and any other still being sent.
func (r *remote) close() {
	r.mu.Lock()
	if r.cancel != nil {
		r.cancel()
	}
	r.mu.Unlock()
	r.wg.Wait()
}

// previewName names the Preview in the build message: its host name's first
// label, as hPanel lists it, else the start of its tunnel id.
func (r *remote) previewName(tunnelID string) string {
	r.mu.Lock()
	defer r.mu.Unlock()
	name, _, _ := strings.Cut(r.hostname, ".")
	if name == "" {
		name, _, _ = strings.Cut(tunnelID, "-")
	}
	return name
}

// job is one command: it runs the deploy and reports how it went.
type job struct {
	r   *remote
	cmd *command
}

// report prints state and tells hPanel. reason goes with stateFailed, and
// buildUUID with stateSucceeded or stateFailed when a build started: a failed
// build's logs, or one that keeps running after Ctrl+C.
func (j *job) report(state, reason, buildUUID string) {
	cl := j.cmd.claims
	j.r.p.deploy(state, cl.Domain, reason)
	params := map[string]string{"command": j.cmd.raw, "command_id": cl.Jti, "tunnel_id": cl.Tid}
	if reason != "" {
		params["reason"] = reason
	}
	if buildUUID != "" {
		params["build_uuid"] = buildUUID
	}
	ctx, cancel := context.WithTimeout(context.WithoutCancel(j.r.base), eventTimeout)
	defer cancel()
	if err := sendEvent(ctx, cl.Username, cl.Domain, "preview.deploy."+state, params); err != nil {
		j.r.p.say("Could not tell hPanel how the deploy went: " + err.Error())
	}
}

// decline refuses the command and reports it. The caller holds r.mu.
func (j *job) decline(status int, reason, why string) error {
	j.r.p.say("Refused a deploy from hPanel to " + j.cmd.claims.Domain + ": " + why + ".")
	// In the background: it must not hold up the edge's answer.
	j.r.wg.Add(1)
	go func() {
		defer j.r.wg.Done()
		j.report(stateFailed, reason, "")
	}()
	return commandError(status, reason)
}

// run deploys p as `hostinger up` would.
func (j *job) run(ctx context.Context, p *project) {
	cl := j.cmd.claims
	j.r.p.say("Deploying " + p.dir + " to " + cl.Domain + ", as asked in hPanel. Press Ctrl+C to cancel the deploy.")
	j.report(stateAccepted, "", "")
	opts := p.deployOptions(cl.Domain, cl.Overwrite, "Deployed from hPanel Preview · "+j.r.previewName(cl.Tid))
	res, err := runDeploy(ctx, lineUI{j.r.p}, opts)
	var buildUUID string
	if res != nil {
		buildUUID = res.BuildUUID
	}
	switch {
	case errors.Is(err, context.Canceled):
		j.r.p.say("Deploy cancelled.")
	case err != nil:
		j.r.p.say("Deploy failed: " + err.Error())
	default:
		j.r.p.say("Deployed to " + res.URL)
		j.report(stateSucceeded, "", buildUUID)
		return
	}
	j.report(stateFailed, hosting.FailureReason(err), buildUUID)
}

// lineUI shows a remote deploy's steps as plain lines and asks nothing:
// hosting.Options.Remote refuses what it would ask.
type lineUI struct{ p *printer }

var errNoPrompt = errors.New("a deploy from hPanel asks nothing")

func (lineUI) Interactive() bool                         { return false }
func (lineUI) Select(string, []string, int) (int, error) { return 0, errNoPrompt }
func (lineUI) Input(string, string) (string, error)      { return "", errNoPrompt }
func (lineUI) Confirm(string, bool) (bool, error)        { return false, errNoPrompt }
func (u lineUI) Info(msg string)                         { u.p.say("  " + msg) }
func (u lineUI) Start(string) hosting.Step               { return lineStep{u} }

// lineStep prints how a step ended; progress updates would only fill the
// terminal.
type lineStep struct{ u lineUI }

func (lineStep) Update(string)     {}
func (s lineStep) Done(msg string) { s.u.Info(msg) }
func (s lineStep) Fail(msg string) { s.u.Info(msg) }
