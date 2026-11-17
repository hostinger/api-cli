package preview

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/utils"
	"github.com/hostinger/hcdn-connector/pkg/connector"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// execute runs Cmd with args against a fake RunQuick and returns stdout,
// stderr, the exit code the CLI would use and the options RunQuick got.
func execute(t *testing.T, fake func(context.Context, connector.QuickOptions) (bool, error), d *detector, args ...string) (string, string, int, *connector.QuickOptions) {
	t.Helper()
	resetCmdFlags(t)
	var got *connector.QuickOptions
	oldRun, oldDetector, oldInteractive := runQuick, newDetector, interactive
	t.Cleanup(func() { runQuick, newDetector, interactive = oldRun, oldDetector, oldInteractive })
	isInteractive := testInteractive
	if !sourcesStubbed {
		stubSources(t, "test-token", &fakeOAuth{}, nil)
	}
	interactive = func(*cobra.Command) bool { return isInteractive }
	runQuick = func(ctx context.Context, o connector.QuickOptions) (bool, error) {
		got = &o
		if fake == nil {
			return false, nil
		}
		return fake(ctx, o)
	}
	newDetector = func(*cobra.Command, bool) detector {
		if d == nil {
			t.Error("detection ran")
			return detector{dial: newFakeDialer().dial, ports: probePorts, binary: "hostinger"}
		}
		return *d
	}
	var out, errOut bytes.Buffer
	Cmd.SetOut(&out)
	Cmd.SetErr(&errOut)
	Cmd.SetArgs(args)
	t.Cleanup(func() { Cmd.SetOut(nil); Cmd.SetErr(nil); Cmd.SetArgs(nil) })
	cmd, err := Cmd.ExecuteC()
	if cmd == nil {
		cmd = Cmd
	}
	code := utils.HandleError(&errOut, cmd, err)
	return out.String(), errOut.String(), code, got
}

// testInteractive is what interactive reports to execute; see asPerson.
var testInteractive bool

// asPerson makes the next execute run as if a person were at a terminal.
func asPerson(t *testing.T) {
	t.Helper()
	testInteractive = true
	t.Cleanup(func() { testInteractive = false })
}

func resetCmdFlags(t *testing.T) {
	t.Helper()
	Cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if sv, ok := f.Value.(pflag.SliceValue); ok {
			_ = sv.Replace(nil)
		} else {
			_ = f.Value.Set(f.DefValue)
		}
		f.Changed = false
	})
}

func TestPreviewArguments(t *testing.T) {
	cases := []struct {
		name string
		args []string
		code int
		msg  string
	}{
		{"mistyped subcommand", []string{"stop"}, utils.ExitUsage, `invalid port "stop": want a number from 1 to 65535`},
		{"port zero", []string{"0"}, utils.ExitUsage, `invalid port "0"`},
		{"port too large", []string{"70000"}, utils.ExitUsage, `invalid port "70000"`},
		{"two ports", []string{"3000", "5173"}, utils.ExitUsage, "accepts at most 1 arg(s)"},
		{"port and --url", []string{"3000", "--url", "http://localhost:5173"}, utils.ExitUsage, "pass a port or --url, not both"},
		{"blocked port", []string{"6379"}, utils.ExitUsage, "port 6379 is not shared"},
		{"blocked --url", []string{"--url", "http://127.0.0.1:3306"}, utils.ExitUsage, "port 3306 is not shared"},
		{"debugger via 0", []string{"--url", "http://0:9229"}, utils.ExitUsage,
			"Error: port 9229 is not shared: it is the usual port of the Node inspector, which must not be reachable from the internet\n"},
		{"debugger via 127.1", []string{"--url", "http://127.1:9229"}, utils.ExitUsage, "the usual port of the Node inspector"},
		{"DevTools port", []string{"9222"}, utils.ExitUsage, "the usual port of Chrome DevTools"},
		{"debugger in JSON mode", []string{"--url", "http://[::ffff:127.0.0.1]:6060", "--json"}, utils.ExitUsage, "the usual port of Go pprof"},
		{"--url scheme", []string{"--url", "ftp://localhost:21"}, utils.ExitUsage, "want an http or https origin"},
		{"--url without host", []string{"--url", "http://:8080"}, utils.ExitUsage, "want an http or https origin"},
		{"--url with path", []string{"--url", "http://localhost:8080/app"}, utils.ExitUsage, "want an origin only"},
		{"--url with credentials", []string{"--url", "http://u:p@localhost:8080"}, utils.ExitUsage, "want an origin only"},
		{"--url port", []string{"--url", "http://localhost:0"}, utils.ExitUsage, `invalid port "0"`},
		{"--ports not numbers", []string{"3000", "--ports", "a,b"}, utils.ExitUsage, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, stderr, code, got := execute(t, nil, nil, c.args...)
			if code != c.code {
				t.Errorf("exit %d, want %d (stderr %q)", code, c.code, stderr)
			}
			if !strings.Contains(stderr, c.msg) {
				t.Errorf("stderr %q, want it to contain %q", stderr, c.msg)
			}
			if got != nil {
				t.Error("the tunnel started")
			}
		})
	}
}

func TestPreviewOptions(t *testing.T) {
	_, stderr, code, got := execute(t, nil, nil, "5173", "--host-header", "localhost",
		"--bootstrap", "a.example.com,b.example.com", "--ports", "7844,443", "--server-ca-file", "/tmp/ca.pem")
	if code != 0 || got == nil {
		t.Fatalf("exit %d, options %v", code, got)
	}
	if got.URL != "http://localhost:5173" || got.HostHeader != "localhost" || got.ServerCAFile != "/tmp/ca.pem" {
		t.Errorf("options %+v", got)
	}
	wantErr := "Warning: --host-header localhost bypasses the dev server's host check, and its debug consoles become public\n" +
		"Preview ended\n"
	if stderr != wantErr {
		t.Errorf("stderr %q, want %q", stderr, wantErr)
	}
	if !reflect.DeepEqual(got.Bootstrap, []string{"a.example.com", "b.example.com"}) || !reflect.DeepEqual(got.Ports, []int{7844, 443}) {
		t.Errorf("bootstrap %v ports %v", got.Bootstrap, got.Ports)
	}
	if !strings.HasPrefix(got.Version, "hostinger-cli/") {
		t.Errorf("Version = %q", got.Version)
	}
	if got.Logger == nil || got.Owner == nil || got.OnReady == nil || got.OnReachable == nil || got.OnPause == nil {
		t.Errorf("missing logger, owner callback or event hooks: %+v", got)
	}

	// The defaults leave the connector's own: the public bootstrap and port 7844.
	_, _, _, got = execute(t, nil, nil, "--url", "https://LOCALHOST:8443/")
	if got.URL != "https://localhost:8443" || len(got.Bootstrap) != 0 || len(got.Ports) != 0 || got.HostHeader != "" {
		t.Errorf("defaults: %+v", got)
	}
}

func TestPreviewNoOwner(t *testing.T) {
	_, _, _, got := execute(t, nil, nil, "3000", "--no-owner")
	if _, err := got.Owner(context.Background(), "t-1", time.Now().Add(time.Hour)); !errors.Is(err, connector.ErrOwnerOptedOut) {
		t.Fatalf("owner callback err = %v, want the connector's ErrOwnerOptedOut", err)
	}
}

func TestPreviewHiddenFlags(t *testing.T) {
	for _, name := range []string{"bootstrap", "server-ca-file", "ports", "no-owner"} {
		if f := Cmd.Flags().Lookup(name); f == nil || !f.Hidden {
			t.Errorf("--%s should exist and be hidden", name)
		}
	}
	for _, name := range []string{"url", "host-header", "json"} {
		if f := Cmd.Flags().Lookup(name); f == nil || f.Hidden {
			t.Errorf("--%s should be visible", name)
		}
	}
	if Cmd.Flags().Lookup("ttl") != nil {
		t.Error("--ttl is not supported yet")
	}
}

func TestPreviewDetects(t *testing.T) {
	d := testDetector(newFakeDialer("[::1]:5173"), false, nil)
	_, stderr, code, got := execute(t, nil, &d)
	if code != 0 || got == nil || got.URL != "http://[::1]:5173" {
		t.Fatalf("exit %d, options %+v", code, got)
	}
	if stderr != "Found a server at http://[::1]:5173\nPreview ended\n" {
		t.Errorf("stderr %q", stderr)
	}
	if _, stderr, _, _ = execute(t, nil, &d, "--json"); stderr != "" {
		t.Errorf("--json stderr %q", stderr)
	}

	d = testDetector(newFakeDialer(), false, nil)
	_, stderr, code, got = execute(t, nil, &d)
	if code != utils.ExitError || got != nil || !strings.Contains(stderr, `no server found on ports 5173, 3000`) {
		t.Fatalf("exit %d, stderr %q", code, stderr)
	}
}

// A whole run in JSON: ready, a pause, then the edge takes the Preview down.
func TestPreviewRunJSON(t *testing.T) {
	fake := func(_ context.Context, o connector.QuickOptions) (bool, error) {
		o.OnReady(connector.QuickTunnel{URL: testURL, ExpiresAt: testExpires, Reachable: true, Scope: connector.ScopePop,
			Owner: connector.OwnerLink{Mode: connector.OwnerModeLog, Absent: connector.OwnerAbsentNoToken}})
		o.OnPause(connector.QuickPause{Reason: connector.PauseCapExceeded, RetryAt: time.Now().Add(-time.Second)})
		return false, &connector.QuickError{Kind: connector.ErrTakenDown, Err: &connector.ExitError{Code: 78, Reason: "revoked"}}
	}
	stdout, stderr, code, _ := execute(t, fake, nil, "3000", "--json")
	want := `{"event":"ready","url":"https://quick-7k2m9x.example.com","expires_at":"2026-10-30T11:59:20Z","reachable":true,"owner":{"linked":false,"reason":"no_token"}}` + "\n" +
		`{"event":"paused","reason":"cap_exceeded","retry_after":0}` + "\n" +
		`{"event":"ended","reason":"taken_down","message":"this Preview was taken down"}` + "\n"
	if stdout != want {
		t.Errorf("stdout:\n%s\nwant:\n%s", stdout, want)
	}
	if code != ExitTakenDown || stderr != "" {
		t.Errorf("exit %d stderr %q; want %d and nothing", code, stderr, ExitTakenDown)
	}
}

// Text mode: Ctrl+C ends with exit 0 and "Preview ended" on stderr.
func TestPreviewRunTextStopped(t *testing.T) {
	fake := func(_ context.Context, o connector.QuickOptions) (bool, error) {
		o.OnReady(connector.QuickTunnel{URL: testURL, ExpiresAt: time.Now().Add(time.Hour), Reachable: true,
			Owner: connector.OwnerLink{Mode: connector.OwnerModeLog, Status: "ok", ClientID: "1"}})
		return false, nil
	}
	stdout, stderr, code, _ := execute(t, fake, nil, "3000")
	if code != 0 || stderr != "Preview ended\n" {
		t.Errorf("exit %d stderr %q", code, stderr)
	}
	for _, want := range []string{"Preview ready: " + testURL, "Forwarding to  http://localhost:3000", "Linked to your Hostinger account"} {
		if !strings.Contains(stdout, want) {
			t.Errorf("stdout lacks %q:\n%s", want, stdout)
		}
	}
}

func TestPreviewRunTextSignInRequired(t *testing.T) {
	fake := func(context.Context, connector.QuickOptions) (bool, error) {
		return false, &connector.QuickError{Kind: connector.ErrOwnerRequired, Err: connector.ErrOwnerRequired,
			Owner: connector.OwnerLink{Mode: connector.OwnerModeRequire, Absent: connector.OwnerAbsentNoToken}}
	}
	stdout, stderr, code, _ := execute(t, fake, nil, "3000")
	want := "Error: Sign in to start a Preview: run \"preview auth login\" or set HOSTINGER_API_TOKEN\n"
	if code != ExitSignInRequired || stderr != want || stdout != "" {
		t.Errorf("exit %d stdout %q stderr %q; want %d, %q", code, stdout, stderr, ExitSignInRequired, want)
	}
}

// A person at a terminal who is not signed in signs in through the browser
// before the Preview starts, so the owner ticket has a token.
func TestPreviewSignsInFirst(t *testing.T) {
	oauth := fakeOAuth{err: fmt.Errorf("%w: refresh token rejected (status 400)", api.ErrNotSignedIn)}
	stubSources(t, "", &oauth, nil)
	asPerson(t)
	var loginsAtStart int
	fake := func(context.Context, connector.QuickOptions) (bool, error) {
		loginsAtStart = oauth.logins
		return false, nil
	}
	_, stderr, code, got := execute(t, fake, nil, "3000")
	if code != 0 || got == nil || loginsAtStart != 1 {
		t.Fatalf("exit %d, started %v, logins before the Preview %d; stderr %q", code, got != nil, loginsAtStart, stderr)
	}
	if !strings.HasPrefix(stderr, "Sign in to your Hostinger account to share your site. Opening your browser...\nSigned in.\n") {
		t.Errorf("stderr %q", stderr)
	}
}

// Without an account, scripts and --json end with exit 77 before any
// Preview and never open a browser.
func TestPreviewNotSignedInWithoutAPerson(t *testing.T) {
	const msg = `Sign in to start a Preview: run "preview auth login" or set HOSTINGER_API_TOKEN`
	cases := []struct {
		name   string
		person bool
		args   []string
		stdout string
		stderr string
	}{
		{"not a terminal", false, []string{"3000"}, "", "Error: " + msg + "\n"},
		{"json", true, []string{"3000", "--json"},
			`{"event":"ended","reason":"sign_in_required","message":"` + strings.ReplaceAll(msg, `"`, `\"`) + `"}` + "\n", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			oauth := fakeOAuth{err: api.ErrNotSignedIn}
			stubSources(t, "", &oauth, nil)
			if tc.person {
				asPerson(t)
			}
			stdout, stderr, code, got := execute(t, nil, nil, tc.args...)
			if code != ExitSignInRequired || got != nil || oauth.logins != 0 || stdout != tc.stdout || stderr != tc.stderr {
				t.Errorf("exit %d, started %v, logins %d, stdout %q, stderr %q; want %d, %q, %q",
					code, got != nil, oauth.logins, stdout, stderr, ExitSignInRequired, tc.stdout, tc.stderr)
			}
		})
	}
}

// The hidden --no-owner flag skips the account check, for testing the service.
func TestPreviewNoOwnerSkipsSignIn(t *testing.T) {
	oauth := fakeOAuth{err: api.ErrNotSignedIn}
	stubSources(t, "", &oauth, nil)
	asPerson(t)
	_, _, code, got := execute(t, nil, nil, "3000", "--no-owner")
	if code != 0 || got == nil || oauth.called || oauth.logins != 0 {
		t.Errorf("exit %d, started %v, token checked %v, logins %d", code, got != nil, oauth.called, oauth.logins)
	}
}

// A sign-in that does not finish, as over SSH, ends with exit 77 before any
// Preview and names the API token as the way around the browser.
func TestPreviewSignInFails(t *testing.T) {
	oauth := fakeOAuth{err: api.ErrNotSignedIn, loginErr: errors.New("authorization denied")}
	stubSources(t, "", &oauth, nil)
	asPerson(t)
	_, stderr, code, got := execute(t, nil, nil, "3000")
	want := "Sign in to your Hostinger account to share your site. Opening your browser...\n" +
		"Error: Sign-in did not finish (authorization denied). Run the same command to try again, or set HOSTINGER_API_TOKEN to use an API token\n"
	if code != ExitSignInRequired || got != nil || stderr != want {
		t.Errorf("exit %d, started %v, stderr %q; want %d and %q", code, got != nil, stderr, ExitSignInRequired, want)
	}
}

// Ctrl+C during the sign-in stops the run with exit 0, as it would a
// Preview, and names the API token as the way around the browser.
func TestPreviewSignInCanceled(t *testing.T) {
	oauth := fakeOAuth{err: api.ErrNotSignedIn, loginErr: context.Canceled}
	stubSources(t, "", &oauth, nil)
	asPerson(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	Cmd.SetContext(ctx)
	t.Cleanup(func() { Cmd.SetContext(context.Background()) })
	_, stderr, code, got := execute(t, nil, nil, "3000")
	const want = "Sign-in canceled. To skip the browser, set HOSTINGER_API_TOKEN\n"
	if code != 0 || got != nil || oauth.logins != 1 || !strings.HasSuffix(stderr, want) {
		t.Errorf("exit %d, started %v, logins %d, stderr %q; want 0, not started, 1, %q", code, got != nil, oauth.logins, stderr, want)
	}
}

// A sign-in that waits out its time says so instead of "context deadline
// exceeded"; an auth-server request that timed out keeps its own message.
func TestNotSignedInTimeout(t *testing.T) {
	cases := []struct {
		name string
		err  error
		why  string
	}{
		{"browser wait", context.DeadlineExceeded, "timed out after 5 minutes"},
		{"auth server", fmt.Errorf("token exchange: %w", context.DeadlineExceeded), "token exchange: context deadline exceeded"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := notSignedIn(context.Background(), &signInError{tc.err}, "hostinger", "HOSTINGER_API_TOKEN")
			want := "Sign-in did not finish (" + tc.why + "). Run the same command to try again, or set HOSTINGER_API_TOKEN to use an API token"
			if e.code != ExitSignInRequired || e.message != want {
				t.Errorf("ending %+v, want exit %d and %q", e, ExitSignInRequired, want)
			}
		})
	}
}

// Detection that finds only Jupyter's port shares it, with a warning.
func TestPreviewDetectsJupyterWithAWarning(t *testing.T) {
	d := testDetector(newFakeDialer("127.0.0.1:8888"), false, nil)
	_, stderr, code, got := execute(t, nil, &d)
	if code != 0 || got == nil || got.URL != "http://127.0.0.1:8888" {
		t.Fatalf("exit %d, options %+v", code, got)
	}
	want := "Found a server at http://127.0.0.1:8888\n" +
		"Warning: port 8888 is the usual port of Jupyter; anyone with the URL can run code on this computer unless it asks for a token or password\n" +
		"Preview ended\n"
	if stderr != want {
		t.Errorf("stderr:\n%s\nwant:\n%s", stderr, want)
	}
}

// The whole text output of a run, from the ready block to the end.
func TestPreviewRunTextGolden(t *testing.T) {
	expires := time.Now().Add(2*time.Hour + 10*time.Second)
	fake := func(_ context.Context, o connector.QuickOptions) (bool, error) {
		o.OnReady(connector.QuickTunnel{URL: testURL, ExpiresAt: expires, Reachable: true, Scope: connector.ScopeSession})
		return true, nil
	}
	stdout, stderr, code, _ := execute(t, fake, nil, "3000")
	want := "\nPreview ready: https://quick-7k2m9x.example.com\n\n" +
		"  Forwarding to  http://localhost:3000\n" +
		"  Expires        " + expires.In(time.Local).Format("2006-01-02 15:04 MST") + " (in 2h)\n" +
		"\nAnyone with the URL can reach this server, including debug consoles and dev tools.\n" +
		"Press Ctrl+C to stop.\n"
	if code != 0 || stdout != want || stderr != "Preview ended\n" {
		t.Errorf("exit %d\nstdout:\n%s\nwant:\n%s\nstderr %q", code, stdout, want, stderr)
	}
}

// Announced while still connecting: the URL with a note, then "Preview ready."
// once the service reaches it.
func TestPreviewRunTextConnectingThenReachable(t *testing.T) {
	expires := time.Now().Add(2*time.Hour + 10*time.Second)
	fake := func(_ context.Context, o connector.QuickOptions) (bool, error) {
		o.OnReady(connector.QuickTunnel{URL: testURL, ExpiresAt: expires})
		o.OnReachable(connector.QuickTunnel{URL: testURL, ExpiresAt: expires, Reachable: true, Scope: connector.ScopeSession})
		return false, nil
	}
	stdout, stderr, code, _ := execute(t, fake, nil, "3000")
	want := "\nPreview URL: https://quick-7k2m9x.example.com (still connecting; it may answer 530 for a few seconds)\n\n" +
		"  Forwarding to  http://localhost:3000\n" +
		"  Expires        " + expires.In(time.Local).Format("2006-01-02 15:04 MST") + " (in 2h)\n" +
		"\nAnyone with the URL can reach this server, including debug consoles and dev tools.\n" +
		"Press Ctrl+C to stop.\n" +
		"Preview ready.\n"
	if code != 0 || stdout != want || stderr != "Preview ended\n" {
		t.Errorf("exit %d\nstdout:\n%s\nwant:\n%s\nstderr %q", code, stdout, want, stderr)
	}
}

func TestPreviewRunJSONConnectingThenReachable(t *testing.T) {
	fake := func(_ context.Context, o connector.QuickOptions) (bool, error) {
		o.OnReady(connector.QuickTunnel{URL: testURL, ExpiresAt: testExpires})
		o.OnReachable(connector.QuickTunnel{URL: testURL, ExpiresAt: testExpires, Reachable: true, Scope: connector.ScopePop})
		return false, nil
	}
	stdout, stderr, code, _ := execute(t, fake, nil, "3000", "--json")
	want := `{"event":"ready","url":"https://quick-7k2m9x.example.com","expires_at":"2026-10-30T11:59:20Z","reachable":false}` + "\n" +
		`{"event":"reachable"}` + "\n" +
		`{"event":"ended","reason":"stopped"}` + "\n"
	if code != 0 || stdout != want || stderr != "" {
		t.Errorf("exit %d stderr %q\nstdout:\n%s\nwant:\n%s", code, stderr, stdout, want)
	}
}
