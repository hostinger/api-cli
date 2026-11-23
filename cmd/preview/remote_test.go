package preview

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/deploy/hosting"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/hcdn-connector/pkg/connector"
)

const testClientID = "123456"

var commandNow = time.Date(2026, 11, 18, 12, 0, 0, 0, time.UTC)

// signer makes command tokens as wh-api does.
type signer struct {
	kid string
	key ed25519.PrivateKey
}

func newSigner(t *testing.T, kid string) signer {
	t.Helper()
	_, key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	return signer{kid: kid, key: key}
}

func (s signer) public() CommandKey {
	return CommandKey{Kid: s.kid, PublicKey: base64.RawURLEncoding.EncodeToString(s.key.Public().(ed25519.PublicKey))}
}

// token signs claims over the defaults; a nil value drops a claim.
func (s signer) token(t *testing.T, typ string, over map[string]any) []byte {
	t.Helper()
	claims := map[string]any{
		"v": 1, "iss": "wh-api", "aud": commandAudience, "sub": testClientID, "tid": testTunnelID,
		"cmd": commandDeploy, "username": "u1", "domain": "example.com", "overwrite": false,
		"jti": "11111111-2222-4333-8444-555555555555", "iat": commandNow.Unix(), "exp": commandNow.Add(time.Minute).Unix(),
	}
	for k, v := range over {
		if v == nil {
			delete(claims, k)
		} else {
			claims[k] = v
		}
	}
	enc := func(v any) string {
		raw, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		return base64.RawURLEncoding.EncodeToString(raw)
	}
	signed := enc(map[string]string{"alg": "EdDSA", "kid": s.kid, "typ": typ}) + "." + enc(claims)
	return []byte(signed + "." + base64.RawURLEncoding.EncodeToString(ed25519.Sign(s.key, []byte(signed))))
}

type sentEvent struct {
	username, domain, name string
	params                 map[string]string
}

// harness is a remote with stubbed deploys and status events.
type harness struct {
	*remote
	out     *bytes.Buffer
	mu      sync.Mutex
	events  []sentEvent
	deploys []hosting.Options
}

func newHarness(t *testing.T, s signer, deploy func(context.Context, hosting.Options) (*hosting.Result, error)) *harness {
	t.Helper()
	h := &harness{out: &bytes.Buffer{}}
	p := &printer{log: h.out, now: func() time.Time { return commandNow }}
	h.remote = &remote{project: currentProject(), p: p, base: context.Background(), seen: map[string]bool{}}
	h.verifier.now = p.now
	h.verifier.setKeys(testClientID, []CommandKey{s.public()})
	h.verifier.tunnelID = testTunnelID
	h.hostname = "calm-river.hostingerpreview.com"

	oldDeploy, oldSend := runDeploy, sendEvent
	t.Cleanup(func() { runDeploy, sendEvent = oldDeploy, oldSend })
	runDeploy = func(ctx context.Context, _ hosting.UI, o hosting.Options) (*hosting.Result, error) {
		h.mu.Lock()
		h.deploys = append(h.deploys, o)
		h.mu.Unlock()
		if deploy == nil {
			return &hosting.Result{URL: "https://example.com", BuildUUID: "build-1"}, nil
		}
		return deploy(ctx, o)
	}
	sendEvent = func(_ context.Context, username, domain, name string, params map[string]string) error {
		h.mu.Lock()
		defer h.mu.Unlock()
		h.events = append(h.events, sentEvent{username, domain, name, params})
		return nil
	}
	// Before the seams are restored: cleanups run last first.
	t.Cleanup(h.wg.Wait)
	return h
}

// names waits for the deploy and the events, then lists the event names
// with the reason of a failure.
func (h *harness) names() []string {
	h.wg.Wait()
	h.mu.Lock()
	defer h.mu.Unlock()
	var out []string
	for _, e := range h.events {
		n := e.name
		if r := e.params["reason"]; r != "" {
			n += ":" + r
		}
		out = append(out, n)
	}
	return out
}

// projectStubbed: the test chose the working directory; execute leaves it.
var projectStubbed bool

// inProject makes dir the working directory of remote deploys.
func inProject(t *testing.T, dir string) {
	t.Helper()
	old := hosting.Getwd
	hosting.Getwd = func() (string, error) { return dir, nil }
	projectStubbed = true
	t.Cleanup(func() { hosting.Getwd, projectStubbed = old, false })
}

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func nodeFolder(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	writeFile(t, dir, "package.json", `{"scripts":{"start":"node server.js"}}`)
	return dir
}

func linkedFolder(t *testing.T, domain, typ string) string {
	t.Helper()
	dir := t.TempDir()
	writeFile(t, dir, "index.html", "<h1>hi</h1>")
	if err := site.Save(dir, &site.Project{Domain: domain, Username: "u1", Type: typ}); err != nil {
		t.Fatal(err)
	}
	return dir
}

func code(err error) string {
	var ce *connector.CommandError
	if !errors.As(err, &ce) {
		return fmt.Sprint(err)
	}
	return fmt.Sprintf("%d %s", ce.Status, ce.Code)
}

func TestCommandDeploysAFirstNodeProjectOnce(t *testing.T) {
	inProject(t, nodeFolder(t))
	s := newSigner(t, "k1")
	h := newHarness(t, s, nil)
	body := s.token(t, commandType, map[string]any{"overwrite": true})

	// Copies arrive at once over several sessions: one deploy, every copy 202.
	var wg sync.WaitGroup
	errs := make([]error, 5)
	for i := range errs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs[i] = h.handle(context.Background(), connector.Command{Body: body})
		}()
	}
	wg.Wait()
	for i, err := range errs {
		if err != nil {
			t.Errorf("copy %d: %v", i, err)
		}
	}
	if got := h.names(); strings.Join(got, ",") != "preview.deploy.accepted,preview.deploy.succeeded" {
		t.Errorf("events = %v", got)
	}
	if len(h.deploys) != 1 {
		t.Fatalf("deploys = %d", len(h.deploys))
	}
	o := h.deploys[0]
	if o.Domain != "example.com" || !o.Remote || !o.Yes || o.Path != "." || o.Message != "Deployed from hPanel Preview · calm-river" {
		t.Errorf("options = %+v", o)
	}
	e := h.events[1]
	want := map[string]string{"command": string(body), "command_id": "11111111-2222-4333-8444-555555555555", "tunnel_id": testTunnelID, "build_uuid": "build-1"}
	if e.username != "u1" || e.domain != "example.com" || fmt.Sprint(e.params) != fmt.Sprint(want) {
		t.Errorf("succeeded = %+v", e)
	}
	if _, ok := h.events[0].params["build_uuid"]; ok {
		t.Errorf("accepted carries a build: %v", h.events[0].params)
	}
	if !strings.Contains(h.out.String(), "Deployed to https://example.com") {
		t.Errorf("output:\n%s", h.out)
	}
}

func TestCommandDeploysALinkedProjectToItsLink(t *testing.T) {
	inProject(t, linkedFolder(t, "example.com", "static"))
	s := newSigner(t, "k1")
	h := newHarness(t, s, func(context.Context, hosting.Options) (*hosting.Result, error) {
		return &hosting.Result{URL: "https://example.com"}, nil
	})
	if err := h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, map[string]any{"domain": "EXAMPLE.com"})}); err != nil {
		t.Fatal(err)
	}
	h.names()
	if o := h.deploys[0]; o.Domain != "" || o.Yes {
		t.Errorf("a linked project deploys from its link: %+v", o)
	}
	if p := h.events[1].params; p["build_uuid"] != "" {
		t.Errorf("static deploy has a build: %v", p)
	}
}

func TestCommandRefusedBeforeItIsTrusted(t *testing.T) {
	inProject(t, nodeFolder(t))
	s := newSigner(t, "k1")
	other := newSigner(t, "k1")
	stranger := newSigner(t, "k9")
	for name, tc := range map[string]struct {
		body []byte
		want string
	}{
		"garbage":       {[]byte("not.a"), "400 bad_command"},
		"forged":        {other.token(t, commandType, nil), "403 bad_signature"},
		"unknown kid":   {stranger.token(t, commandType, nil), "403 unknown_kid"},
		"owner ticket":  {s.token(t, "hcdn-owner+jwt", map[string]any{"aud": "hcdn-tunnel"}), "403 wrong_audience"},
		"audience":      {s.token(t, commandType, map[string]any{"aud": "hcdn-tunnel"}), "403 wrong_audience"},
		"command":       {s.token(t, commandType, map[string]any{"cmd": "stop"}), "400 unsupported_command"},
		"tunnel":        {s.token(t, commandType, map[string]any{"tid": "0192f3a4-5b6c-8d7e-9f80-000000000000"}), "403 wrong_tunnel"},
		"owner":         {s.token(t, commandType, map[string]any{"sub": "999"}), "403 wrong_owner"},
		"expired":       {s.token(t, commandType, map[string]any{"exp": commandNow.Add(-61 * time.Second).Unix()}), "403 expired"},
		"no domain":     {s.token(t, commandType, map[string]any{"domain": nil}), "400 bad_command"},
		"no command id": {s.token(t, commandType, map[string]any{"jti": nil}), "400 bad_command"},
	} {
		t.Run(name, func(t *testing.T) {
			h := newHarness(t, s, nil)
			if got := code(h.handle(context.Background(), connector.Command{Body: tc.body})); got != tc.want {
				t.Errorf("answer = %s, want %s", got, tc.want)
			}
			if got := h.names(); len(got) > 0 || len(h.deploys) > 0 {
				t.Errorf("events %v, deploys %v", got, h.deploys)
			}
		})
	}
}

func TestCommandWithinTheLeewayRuns(t *testing.T) {
	inProject(t, nodeFolder(t))
	s := newSigner(t, "k1")
	h := newHarness(t, s, nil)
	if err := h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, map[string]any{"exp": commandNow.Add(-59 * time.Second).Unix()})}); err != nil {
		t.Fatal(err)
	}
}

func TestCommandWhileADeployRunsIsBusy(t *testing.T) {
	inProject(t, nodeFolder(t))
	s := newSigner(t, "k1")
	release := make(chan struct{})
	h := newHarness(t, s, func(context.Context, hosting.Options) (*hosting.Result, error) {
		<-release
		return &hosting.Result{URL: "https://example.com"}, nil
	})
	if err := h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, nil)}); err != nil {
		t.Fatal(err)
	}
	second := s.token(t, commandType, map[string]any{"jti": "22222222-2222-4333-8444-555555555555"})
	if got := code(h.handle(context.Background(), connector.Command{Body: second})); got != "409 busy" {
		t.Errorf("answer = %s", got)
	}
	// Its copies are answered as done, not reported busy again.
	if err := h.handle(context.Background(), connector.Command{Body: second}); err != nil {
		t.Errorf("copy: %v", err)
	}
	close(release)
	// The refusal is reported in the background: only accepted and succeeded
	// keep their order.
	got := h.names()
	if !slices.Equal(slices.DeleteFunc(slices.Clone(got), func(n string) bool { return n == "preview.deploy.failed:busy" }),
		[]string{"preview.deploy.accepted", "preview.deploy.succeeded"}) || len(got) != 3 {
		t.Errorf("events = %v", got)
	}
}

func TestCommandToAnotherWebsiteThanTheLink(t *testing.T) {
	inProject(t, linkedFolder(t, "example.com", "static"))
	s := newSigner(t, "k1")
	h := newHarness(t, s, nil)
	got := code(h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, map[string]any{"domain": "other.com"})}))
	if got != "409 target_changed" {
		t.Errorf("answer = %s", got)
	}
	if names := h.names(); strings.Join(names, ",") != "preview.deploy.failed"+":target_changed" || len(h.deploys) > 0 {
		t.Errorf("events %v, deploys %v", names, h.deploys)
	}
	if e := h.events[0]; e.domain != "other.com" {
		t.Errorf("the event goes to the command's website: %+v", e)
	}
}

func TestCommandForAFolderThatIsNoLongerAProject(t *testing.T) {
	dir := nodeFolder(t)
	inProject(t, dir)
	s := newSigner(t, "k1")
	h := newHarness(t, s, nil)
	if err := os.Remove(filepath.Join(dir, "package.json")); err != nil {
		t.Fatal(err)
	}
	if got := code(h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, nil)})); got != "409 unsupported_type" {
		t.Errorf("answer = %s", got)
	}
}

func TestDeployFailuresAreReported(t *testing.T) {
	for _, tc := range []struct {
		err  error
		want string
	}{
		{&hosting.RefusedError{Reason: hosting.ReasonWouldOverwrite, Message: "has content"}, "would_overwrite"},
		{&hosting.BuildError{UUID: "b"}, "build_failed"},
		{fmt.Errorf("packing: %w", context.Canceled), "cancelled"},
		{errors.New("HTTP 500"), "deploy_failed"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			inProject(t, nodeFolder(t))
			s := newSigner(t, "k1")
			h := newHarness(t, s, func(context.Context, hosting.Options) (*hosting.Result, error) { return nil, tc.err })
			if err := h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, nil)}); err != nil {
				t.Fatal(err)
			}
			if got := h.names(); strings.Join(got, ",") != "preview.deploy.accepted"+","+"preview.deploy.failed"+":"+tc.want {
				t.Errorf("events = %v", got)
			}
		})
	}
}

// fakeHostingAPI answers what hosting.Run asks on the way to a Node.js build;
// any other call panics on the nil API it embeds.
type fakeHostingAPI struct {
	hosting.API
	uploadURL string
	ctrlC     func() // pressed while the build runs; else the build fails
}

func (f *fakeHostingAPI) ListWebsites(context.Context, string) ([]site.Website, error) {
	return []site.Website{{Domain: "example.com", Username: "u1"}}, nil
}
func (f *fakeHostingAPI) ListSetups(context.Context, string) ([]hosting.Setup, error) {
	return nil, nil
}
func (f *fakeHostingAPI) UploadCredentials(context.Context, string, string) (hosting.Credentials, error) {
	return hosting.Credentials{URL: f.uploadURL}, nil
}
func (f *fakeHostingAPI) BuildSettingsFromArchive(context.Context, string, string, string) (hosting.NodeSettings, error) {
	return hosting.NodeSettings{AppType: "express"}, nil
}
func (f *fakeHostingAPI) StartBuild(context.Context, string, string, hosting.NodeSettings, string, string) (hosting.Build, error) {
	if f.ctrlC != nil {
		f.ctrlC()
		return hosting.Build{UUID: "b-1", State: "running"}, nil
	}
	return hosting.Build{UUID: "b-1", State: "failed"}, nil
}
func (f *fakeHostingAPI) BuildLogs(context.Context, string, string, string) (string, error) {
	return "npm ERR! missing script: build", nil
}
func (f *fakeHostingAPI) BuildAnalysis(context.Context, string, string, string) (string, string, error) {
	return "", "", nil
}

// A remote deploy runs the real hosting.Run on a linked static folder that
// has become a Node.js app: lineUI prints its steps as plain lines, and a
// deploy that ends after its build started reports that build.
func TestRemoteDeployThroughHostingRun(t *testing.T) {
	upload := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
		if r.Method == http.MethodPost {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}))
	t.Cleanup(upload.Close)
	for _, tc := range []struct {
		name   string
		ctrlC  bool
		reason string
		lines  []string
	}{{
		// No "build again?" prompt: lineUI is not interactive.
		name: "build fails", reason: hosting.ReasonBuildFailed,
		lines: []string{
			"  This folder looks like a Node.js app, but it was last deployed as static; deploying it as nodejs",
			"  Uploaded ",
			"  Build failed",
			"Deploy failed: build b-1 failed",
		},
	}, {
		// hPanel follows the build, which goes on without the Preview.
		name: "ctrl-c during the build", ctrlC: true, reason: hosting.ReasonCancelled,
		lines: []string{"  Interrupted; the build keeps running on the server", "Deploy cancelled."},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			dir := linkedFolder(t, "example.com", "static")
			writeFile(t, dir, "package.json", `{"scripts":{"start":"node server.js"}}`)
			inProject(t, dir)
			s := newSigner(t, "k1")
			h := newHarness(t, s, nil)
			api := &fakeHostingAPI{uploadURL: upload.URL}
			if tc.ctrlC {
				api.ctrlC = func() { h.cancelDeploy() }
			}
			// The harness's stub takes no UI; its cleanup restores runDeploy.
			runDeploy = func(ctx context.Context, ui hosting.UI, o hosting.Options) (*hosting.Result, error) {
				return hosting.Run(ctx, api, ui, o)
			}
			if err := h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, nil)}); err != nil {
				t.Fatal(err)
			}
			if got := strings.Join(h.names(), ","); got != "preview.deploy.accepted,preview.deploy.failed:"+tc.reason {
				t.Errorf("events = %s", got)
			}
			if got := h.events[1].params["build_uuid"]; got != "b-1" {
				t.Errorf("build_uuid = %q", got)
			}
			out := h.out.String()
			for _, want := range tc.lines {
				if !strings.Contains(out, "\n"+want) {
					t.Errorf("no line %q in:\n%s", want, out)
				}
			}
		})
	}
}

// The build message names the Preview the connector reports ready, and the
// connector's own OnReady still runs.
func TestBuildMessageNamesTheReadyPreview(t *testing.T) {
	inProject(t, nodeFolder(t))
	s := newSigner(t, "k1")
	h := newHarness(t, s, nil)
	var ready string
	qo := connector.QuickOptions{OnReady: func(t connector.QuickTunnel) { ready = t.Hostname }}
	h.attach(&qo, OwnerOptions{})
	qo.OnReady(connector.QuickTunnel{Hostname: "bold-lake.hostingerpreview.com"})
	if err := h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, nil)}); err != nil {
		t.Fatal(err)
	}
	h.names()
	if ready != "bold-lake.hostingerpreview.com" || h.deploys[0].Message != "Deployed from hPanel Preview · bold-lake" {
		t.Errorf("OnReady got %q, message %q", ready, h.deploys[0].Message)
	}
}

func TestLineUIAsksNothing(t *testing.T) {
	u := lineUI{}
	_, selErr := u.Select("", nil, 0)
	_, inErr := u.Input("", "")
	_, confErr := u.Confirm("", true)
	if u.Interactive() || selErr != errNoPrompt || inErr != errNoPrompt || confErr != errNoPrompt {
		t.Error("lineUI asked")
	}
}

func TestCtrlCCancelsTheDeployNotThePreview(t *testing.T) {
	inProject(t, nodeFolder(t))
	s := newSigner(t, "k1")
	started := make(chan struct{})
	h := newHarness(t, s, func(ctx context.Context, _ hosting.Options) (*hosting.Result, error) {
		close(started)
		<-ctx.Done()
		return nil, ctx.Err()
	})
	if h.cancelDeploy() {
		t.Error("cancelled a deploy while none ran")
	}
	if err := h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, nil)}); err != nil {
		t.Fatal(err)
	}
	<-started
	if !h.cancelDeploy() {
		t.Fatal("no deploy to cancel")
	}
	// While the cancelled deploy still reports, the next Ctrl+C stops the Preview.
	if h.cancelDeploy() {
		t.Error("a second Ctrl+C cancelled the deploy again")
	}
	if got := h.names(); strings.Join(got, ",") != "preview.deploy.accepted"+","+"preview.deploy.failed"+":cancelled" {
		t.Errorf("events = %v", got)
	}
	if p := h.events[1].params; p["build_uuid"] != "" {
		t.Errorf("no build started, yet build_uuid = %q", p["build_uuid"])
	}
	if h.cancelDeploy() {
		t.Error("the next Ctrl+C still cancels a deploy")
	}
}

func TestPreviewEndCancelsTheDeployAndReportsIt(t *testing.T) {
	inProject(t, nodeFolder(t))
	s := newSigner(t, "k1")
	started := make(chan struct{})
	h := newHarness(t, s, func(ctx context.Context, _ hosting.Options) (*hosting.Result, error) {
		close(started)
		<-ctx.Done()
		return nil, ctx.Err()
	})
	if err := h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, nil)}); err != nil {
		t.Fatal(err)
	}
	<-started
	h.close()
	h.mu.Lock()
	n := len(h.events)
	h.mu.Unlock()
	if n != 2 || h.events[1].params["reason"] != hosting.ReasonCancelled {
		t.Errorf("events = %+v", h.events)
	}
}

func TestUnknownKidRefreshesTheKeysOnce(t *testing.T) {
	inProject(t, nodeFolder(t))
	old, rotated := newSigner(t, "k1"), newSigner(t, "k2")
	srv := newFakePublicAPI(t, http.StatusCreated, fmt.Sprintf(`{"ticket":"eyJ.a.b","expires_at":"2026-10-23T11:05:00Z","client_id":%q,"command_keys":[{"kid":"k2","public_key":%q}]}`,
		testClientID, rotated.public().PublicKey), 0)
	h := newHarness(t, old, nil)
	var qo connector.QuickOptions
	h.attach(&qo, OwnerOptions{APIURL: srv.URL, Token: staticSource("tok")})
	if _, err := qo.Owner(context.Background(), testTunnelID, testNotAfter); err != nil {
		t.Fatal(err)
	}
	h.verifier.setKeys(testClientID, []CommandKey{old.public()}) // as if the run started before the rotation
	calls := srv.calls.Load()

	if err := h.handle(context.Background(), connector.Command{Body: rotated.token(t, commandType, nil)}); err != nil {
		t.Fatalf("rotated key: %v", err)
	}
	if srv.calls.Load() != calls+1 {
		t.Errorf("owner ticket calls = %d, want one more than %d", srv.calls.Load(), calls)
	}
	// Within ten minutes another unknown kid is refused without a call.
	stranger := newSigner(t, "k3")
	if got := code(h.handle(context.Background(), connector.Command{Body: stranger.token(t, commandType, nil)})); got != "403 unknown_kid" {
		t.Errorf("answer = %s", got)
	}
	if srv.calls.Load() != calls+1 {
		t.Errorf("refreshed again within ten minutes")
	}
	h.names()
}

func TestOwnerTicketKeepsTheCommandKeys(t *testing.T) {
	s := newSigner(t, "k1")
	srv := newFakePublicAPI(t, http.StatusCreated, fmt.Sprintf(`{"ticket":"eyJ.a.b","expires_at":"2026-10-23T11:05:00Z","client_id":"42","command_keys":[{"kid":"k1","public_key":%q},{"kid":"bad","public_key":"AAAA"}]}`, s.public().PublicKey), 0)
	var v verifier
	fn := NewOwnerTicketFunc(OwnerOptions{APIURL: srv.URL, Token: staticSource("tok"), OnTicket: v.setKeys})
	if _, err := fn(context.Background(), testTunnelID, testNotAfter); err != nil {
		t.Fatal(err)
	}
	if key, id, _ := v.get("k1"); key == nil || id != "42" {
		t.Errorf("k1 = %v, client id %q", key, id)
	}
	if key, _, _ := v.get("bad"); key != nil {
		t.Error("kept a key that is not an Ed25519 public key")
	}
}

func TestStatusEventsReachPublicAPI(t *testing.T) {
	srv := newFakePublicAPI(t, http.StatusOK, `{"message":"ok"}`, 0)
	old := site.Client
	site.Client = func() *api.RawClient { return api.NewRawClient(srv.URL, nil) }
	t.Cleanup(func() { site.Client = old })
	err := sendEvent(context.Background(), "u1", "example.com", "preview.deploy.failed", map[string]string{"command_id": "c", "reason": "busy"})
	if err != nil {
		t.Fatal(err)
	}
	got := srv.req.Load()
	if got.method != http.MethodPost || got.path != "/api/hosting/v1/accounts/u1/websites/example.com/events/trigger" {
		t.Errorf("request = %s %s", got.method, got.path)
	}
	if fmt.Sprint(got.body) != "map[event:map[name:preview.deploy.failed params:map[command_id:c reason:busy]]]" {
		t.Errorf("body = %v", got.body)
	}
}

func TestDeployEventsInJSON(t *testing.T) {
	inProject(t, nodeFolder(t))
	s := newSigner(t, "k1")
	h := newHarness(t, s, nil)
	var stdout bytes.Buffer
	h.p.json, h.p.stdout = true, &stdout
	if err := h.handle(context.Background(), connector.Command{Body: s.token(t, commandType, nil)}); err != nil {
		t.Fatal(err)
	}
	h.names()
	want := `{"event":"deploy","state":"accepted","domain":"example.com"}` + "\n" +
		`{"event":"deploy","state":"succeeded","domain":"example.com"}` + "\n"
	if stdout.String() != want {
		t.Errorf("stdout:\n%s\nwant:\n%s", stdout.String(), want)
	}
}

// The edge resets a command's stream after 10 s: a hung key refresh must not
// hold the answer that long.
func TestKeyRefreshIsBoundedWellUnderTheEdgeTimeout(t *testing.T) {
	inProject(t, nodeFolder(t))
	srv := newFakePublicAPI(t, http.StatusCreated, `{"ticket":"eyJ.a.b"}`, 30*time.Second)
	h := newHarness(t, newSigner(t, "k1"), nil)
	var qo connector.QuickOptions
	h.attach(&qo, OwnerOptions{APIURL: srv.URL, Token: staticSource("tok")})
	h.verifier.ticket = qo.Owner // as if the connector had asked for an owner
	start := time.Now()
	got := code(h.handle(context.Background(), connector.Command{Body: newSigner(t, "k9").token(t, commandType, nil)}))
	if got != "403 unknown_kid" || time.Since(start) > keyRefreshTimeout+time.Second {
		t.Errorf("answer = %s after %s", got, time.Since(start))
	}
}
