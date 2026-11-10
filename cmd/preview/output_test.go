package preview

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hostinger/api-cli/utils"
	"github.com/hostinger/hcdn-connector/pkg/connector"
)

var (
	testNow     = time.Date(2026, 10, 29, 12, 0, 0, 0, time.UTC)
	testExpires = testNow.Add(23*time.Hour + 59*time.Minute + 20*time.Second)
	testURL     = "https://quick-7k2m9x.example.com"
)

func testPrinter(json bool) (*printer, *bytes.Buffer, *bytes.Buffer) {
	var out, errOut bytes.Buffer
	return &printer{stdout: &out, stderr: &errOut, json: json, target: "http://localhost:5173",
		now: func() time.Time { return testNow }, loc: time.UTC}, &out, &errOut
}

func tunnel(link connector.OwnerLink) connector.QuickTunnel {
	return connector.QuickTunnel{TunnelID: "t-1", Hostname: strings.TrimPrefix(testURL, "https://"),
		URL: testURL, ExpiresAt: testExpires, Owner: link, Reachable: true, Scope: connector.ScopePop, Pop: "bnk"}
}

// connecting is a tunnel announced before it answers.
func connecting(link connector.OwnerLink) connector.QuickTunnel {
	t := tunnel(link)
	t.Reachable, t.Scope, t.Pop = false, "", ""
	return t
}

// The owner outcomes: every reason the line can carry, and the edge not asking.
var ownerCases = []struct {
	name     string
	link     connector.OwnerLink
	line     string // the text owner line, "" for none
	jsonPart string // the JSON owner field, "" when omitted
}{
	{"not asked", connector.OwnerLink{Absent: connector.OwnerAbsentNotRequested},
		"", ""},
	{"linked", connector.OwnerLink{Mode: connector.OwnerModeLog, Ticket: "a.b.c", Status: "ok", ClientID: "123456"},
		"  Linked to your Hostinger account\n", `,"owner":{"linked":true,"client_id":"123456"}`},
	{"no token", connector.OwnerLink{Mode: connector.OwnerModeLog, Absent: connector.OwnerAbsentNoToken},
		"  Not linked to an account (no token)\n", `,"owner":{"linked":false,"reason":"no_token"}`},
	{"opted out", connector.OwnerLink{Mode: connector.OwnerModeLog, Absent: connector.OwnerAbsentOptedOut},
		"  Not linked to an account (--no-owner)\n", `,"owner":{"linked":false,"reason":"opted_out"}`},
	{"fetch failed", connector.OwnerLink{Mode: connector.OwnerModeLog, Absent: connector.OwnerAbsentFetchFailed},
		"  Not linked to an account (ticket request failed)\n", `,"owner":{"linked":false,"reason":"fetch_failed"}`},
	{"other environment", connector.OwnerLink{Mode: connector.OwnerModeLog, Ticket: "a.b.c", Status: "invalid",
		Reason: "unknown_kid", OtherEnvironment: true},
		"  Not linked to an account (ticket from another environment)\n", `,"owner":{"linked":false,"reason":"other_environment"}`},
	{"edge label", connector.OwnerLink{Mode: connector.OwnerModeLog, Ticket: "a.b.c", Status: "invalid", Reason: "expired"},
		"  Not linked to an account (expired)\n", `,"owner":{"linked":false,"reason":"expired"}`},
	{"status only", connector.OwnerLink{Mode: connector.OwnerModeLog, Ticket: "a.b.c", Status: "missing"},
		"  Not linked to an account (missing)\n", `,"owner":{"linked":false,"reason":"missing"}`},
	{"not verified", connector.OwnerLink{Mode: connector.OwnerModeLog, Ticket: "a.b.c"},
		"  Not linked to an account (ticket not verified)\n", `,"owner":{"linked":false,"reason":"not_verified"}`},
}

func TestReadyText(t *testing.T) {
	for _, c := range ownerCases {
		t.Run(c.name, func(t *testing.T) {
			p, out, errOut := testPrinter(false)
			p.ready(tunnel(c.link))
			want := "\nPreview ready: https://quick-7k2m9x.example.com\n\n" +
				"  Forwarding to  http://localhost:5173\n" +
				"  Expires        2026-10-30 11:59 UTC (in 23h59m)\n" +
				c.line +
				"\nAnyone with the URL can reach this server, including debug consoles and dev tools.\n" +
				"Press Ctrl+C to stop.\n"
			if out.String() != want {
				t.Errorf("stdout:\n%s\nwant:\n%s", out, want)
			}
			if errOut.Len() > 0 {
				t.Errorf("stderr: %q", errOut)
			}
		})
	}
}

func TestReadyJSON(t *testing.T) {
	for _, c := range ownerCases {
		t.Run(c.name, func(t *testing.T) {
			p, out, _ := testPrinter(true)
			p.ready(tunnel(c.link))
			want := `{"event":"ready","url":"https://quick-7k2m9x.example.com","expires_at":"2026-10-30T11:59:20Z","reachable":true` + c.jsonPart + "}\n"
			if out.String() != want {
				t.Errorf("stdout %s want %s", out, want)
			}
		})
	}
}

// Announced while still connecting, then reachable: once, whatever the callers do.
func TestReadyConnectingThenReachableText(t *testing.T) {
	p, out, errOut := testPrinter(false)
	p.ready(connecting(connector.OwnerLink{Mode: connector.OwnerModeLog, Absent: connector.OwnerAbsentNoToken}))
	p.reachable(tunnel(connector.OwnerLink{}))
	p.reachable(tunnel(connector.OwnerLink{}))
	want := "\nPreview URL: https://quick-7k2m9x.example.com (still connecting; it may answer 530 for a few seconds)\n\n" +
		"  Forwarding to  http://localhost:5173\n" +
		"  Expires        2026-10-30 11:59 UTC (in 23h59m)\n" +
		"  Not linked to an account (no token)\n" +
		"\nAnyone with the URL can reach this server, including debug consoles and dev tools.\n" +
		"Press Ctrl+C to stop.\n" +
		"Preview ready.\n"
	if out.String() != want {
		t.Errorf("stdout:\n%s\nwant:\n%s", out, want)
	}
	if errOut.Len() > 0 {
		t.Errorf("stderr: %q", errOut)
	}
}

func TestReadyConnectingThenReachableJSON(t *testing.T) {
	p, out, _ := testPrinter(true)
	p.ready(connecting(connector.OwnerLink{Absent: connector.OwnerAbsentNotRequested}))
	p.reachable(tunnel(connector.OwnerLink{}))
	p.reachable(tunnel(connector.OwnerLink{}))
	want := `{"event":"ready","url":"https://quick-7k2m9x.example.com","expires_at":"2026-10-30T11:59:20Z","reachable":false}` + "\n" +
		`{"event":"reachable"}` + "\n"
	if out.String() != want {
		t.Errorf("stdout:\n%s\nwant:\n%s", out, want)
	}
}

func TestReadyConnectingStyles(t *testing.T) {
	p, out, _ := testPrinter(false)
	p.bold = true
	p.ready(connecting(connector.OwnerLink{}))
	if !strings.Contains(out.String(), "Preview URL: \x1b[1m"+testURL+"\x1b[0m (still connecting") {
		t.Errorf("URL not emphasised:\n%q", out)
	}
}

func TestReadyTextSanitisesAndStyles(t *testing.T) {
	p, out, _ := testPrinter(false)
	p.bold = true
	tun := tunnel(connector.OwnerLink{Mode: connector.OwnerModeLog, Ticket: "a.b.c", Status: "invalid", Reason: "x\x1b]0;t\x07"})
	p.ready(tun)
	if !strings.Contains(out.String(), "Preview ready: \x1b[1m"+testURL+"\x1b[0m\n") {
		t.Errorf("URL not emphasised:\n%q", out)
	}
	if strings.Contains(out.String(), "\x1b]") || strings.Contains(out.String(), "\x07") {
		t.Errorf("control characters from the edge reached the terminal: %q", out)
	}
}

func TestPaused(t *testing.T) {
	retry := testNow.Add(90*time.Second + 300*time.Millisecond)
	p, out, errOut := testPrinter(false)
	p.paused(connector.QuickPause{Reason: connector.PauseCapExceeded, RetryAt: retry})
	p.paused(connector.QuickPause{Reason: "disabled", RetryAt: retry})
	wantErr := "too many Previews from your network; retrying at 12:01\nPreview paused (disabled); retrying at 12:01\n"
	if errOut.String() != wantErr || out.Len() > 0 {
		t.Errorf("stderr %q stdout %q; want stderr %q", errOut, out, wantErr)
	}

	p, out, _ = testPrinter(true)
	p.paused(connector.QuickPause{Reason: connector.PauseCapExceeded, RetryAt: retry})
	p.paused(connector.QuickPause{Reason: connector.PauseCapExceeded, RetryAt: testNow.Add(-time.Second)})
	want := `{"event":"paused","reason":"cap_exceeded","retry_after":91}` + "\n" +
		`{"event":"paused","reason":"cap_exceeded","retry_after":0}` + "\n"
	if out.String() != want {
		t.Errorf("stdout %s want %s", out, want)
	}
}

// quickErr builds the error RunQuick returns for kind.
func quickErr(kind, cause error) error {
	return &connector.QuickError{Kind: kind, Err: cause}
}

func TestClassify(t *testing.T) {
	signIn := `Sign in to start a Preview: run "hostinger auth login" or set HOSTINGER_API_TOKEN`
	cases := []struct {
		name     string
		expired  bool
		err      error
		reason   string
		code     int
		message  string
		contains bool // message is a prefix only
	}{
		{"stopped", false, nil, "stopped", 0, "Preview ended", false},
		{"expired", true, nil, "expired", 0, "Preview ended", false},
		{"cancelled", false, context.Canceled, "stopped", 0, "Preview ended", false},
		{"taken down", false, quickErr(connector.ErrTakenDown, &connector.ExitError{Code: 78, Reason: "revoked"}),
			"taken_down", 78, "this Preview was taken down", false},
		{"sign-in required", false, &connector.QuickError{Kind: connector.ErrOwnerRequired,
			Err: fmt.Errorf("%w: no owner ticket (no_token)", connector.ErrOwnerRequired)},
			"sign_in_required", 77, signIn, false},
		{"sign-in required, goaway", false, &connector.QuickError{Kind: connector.ErrOwnerRequired,
			Err: &connector.ExitError{Code: 77, Reason: "owner_required"}},
			"sign_in_required", 77, signIn, false},
		{"sign-in required, other environment", false, &connector.QuickError{Kind: connector.ErrOwnerRequired,
			Err: connector.ErrOwnerRequired, Owner: connector.OwnerLink{Mode: "require", OtherEnvironment: true}},
			"sign_in_required", 77, signIn + "; the owner ticket comes from another environment", false},
		{"unreachable", false, quickErr(connector.ErrEdgeUnreachable, errors.New("dial tcp: i/o timeout")),
			"unreachable", 1, "outbound TCP 7844 is blocked on this network", false},
		{"refused", false, quickErr(connector.ErrQuickRefused, errors.New("remote error: tls: bad certificate")),
			"error", 1, "the Preview service refused the tunnel: remote error: tls: bad certificate", false},
		{"not offered", false, quickErr(connector.ErrQuickNotOffered, errors.New(`class "named"`)),
			"error", 1, `the Preview service does not offer Previews to this client: class "named"`, false},
		{"cert rejected", false, quickErr(connector.ErrCertRejected, &connector.ExitError{Code: 78, Reason: "unknown_serial"}),
			"error", 1, "the Preview service no longer accepts this tunnel; run the command again for a new URL", false},
		{"other goaway", false, &connector.ExitError{Code: 75, Reason: "protocol_error"},
			"error", 1, "Preview failed: exit 75 (protocol_error)", false},
		{"bad option", false, errors.New("connector: invalid URL"),
			"error", 1, "Preview failed: connector: invalid URL", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			e := classify(c.expired, c.err, "hostinger", "HOSTINGER_API_TOKEN")
			if e.reason != c.reason || e.code != c.code || e.message != c.message {
				t.Errorf("classify = %+v, want {%s %d %q}", e, c.reason, c.code, c.message)
			}
		})
	}
}

func TestEnded(t *testing.T) {
	cases := []struct {
		e        ending
		json     bool
		stdout   string
		stderr   string
		exitCode int
	}{
		{ending{"stopped", 0, "Preview ended"}, false, "", "Preview ended\n", 0},
		{ending{"expired", 0, "Preview ended"}, true, `{"event":"ended","reason":"expired"}` + "\n", "", 0},
		{ending{"taken_down", 78, "this Preview was taken down"}, false, "", "Error: this Preview was taken down\n", 78},
		{ending{"taken_down", 78, "this Preview was taken down"}, true, `{"event":"ended","reason":"taken_down","message":"this Preview was taken down"}` + "\n", "", 78},
		{ending{"sign_in_required", 77, "Sign in"}, true, `{"event":"ended","reason":"sign_in_required","message":"Sign in"}` + "\n", "", 77},
		{ending{"unreachable", 1, "blocked"}, false, "", "Error: blocked\n", 1},
		{ending{"error", 1, "boom"}, true, `{"event":"ended","reason":"error","message":"boom"}` + "\n", "", 1},
	}
	for _, c := range cases {
		p, out, errOut := testPrinter(c.json)
		err := p.ended(c.e)
		// The command returns err; the root's handler prints it and picks the code.
		code := utils.HandleError(errOut, Cmd, err)
		if out.String() != c.stdout || errOut.String() != c.stderr || code != c.exitCode {
			t.Errorf("%+v json=%v: stdout %q stderr %q exit %d; want %q %q %d",
				c.e, c.json, out, errOut, code, c.stdout, c.stderr, c.exitCode)
		}
	}
}

func TestRoughDuration(t *testing.T) {
	for d, want := range map[time.Duration]string{
		20 * time.Second:              "under a minute",
		45 * time.Minute:              "45m",
		time.Hour:                     "1h",
		24*time.Hour - 30*time.Second: "24h",
		23*time.Hour + 59*time.Minute: "23h59m",
		-time.Minute:                  "under a minute",
		2*time.Hour + 5*time.Minute:   "2h5m",
	} {
		if got := roughDuration(d); got != want {
			t.Errorf("roughDuration(%s) = %q, want %q", d, got, want)
		}
	}
}
