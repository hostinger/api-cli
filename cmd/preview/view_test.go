package preview

import (
	"context"
	"io"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/hostinger/hcdn-connector/pkg/connector"
	"github.com/hostinger/hcdn-connector/pkg/quickview"
	"github.com/spf13/viper"
)

var escapes = regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z]`)

// withView makes the next execute run for a person and draw the live view on
// its stdout, as at a terminal, and reports how many views were made.
func withView(t *testing.T) *int {
	t.Helper()
	asPerson(t)
	// These tests may run inside an agent's shell.
	t.Setenv("AI_AGENT", "")
	t.Setenv("CLAUDECODE", "")
	made := 0
	old := newView
	t.Cleanup(func() { newView = old })
	newView = func(w io.Writer, o quickview.Options) *quickview.View {
		made++
		return quickview.New(w, o)
	}
	return &made
}

func TestPreviewViewStopped(t *testing.T) {
	withView(t)
	fake := func(_ context.Context, o connector.QuickOptions) (bool, error) {
		o.Logger.Warn("edge is slow")
		o.OnReady(connector.QuickTunnel{URL: testURL, ExpiresAt: time.Now().Add(time.Hour), Reachable: true,
			Owner: connector.OwnerLink{Mode: connector.OwnerModeLog, Status: "ok", ClientID: "1"}})
		return false, nil
	}
	stdout, stderr, code, _ := execute(t, fake, nil, "3000")
	screen := escapes.ReplaceAllString(stdout, "")
	if code != 0 {
		t.Fatalf("exit %d, stderr %q", code, stderr)
	}
	for _, want := range []string{"Sharing http://localhost:3000", testURL, "Linked to your Hostinger account",
		"Preview ended. The link no longer works."} {
		if !strings.Contains(screen, want) {
			t.Errorf("view lacks %q:\n%s", want, screen)
		}
	}
	if strings.Contains(screen, "Preview ready:") {
		t.Errorf("plain lines next to the view:\n%s", screen)
	}
	// The log line goes above the view; the end is in the view, not on stderr.
	if !strings.Contains(stderr, "edge is slow") || strings.Contains(stderr, "Preview ended") {
		t.Errorf("stderr %q", stderr)
	}
}

// The port and host header warnings print through the view's log writer,
// below the logo, once each.
func TestPreviewViewWarnings(t *testing.T) {
	withView(t)
	_, stderr, code, _ := execute(t, nil, nil, "3000", "--host-header", "localhost")
	want := "Warning: --host-header localhost bypasses the dev server's host check, and its debug consoles become public\n"
	if code != 0 || stderr != want {
		t.Errorf("exit %d stderr %q, want %q", code, stderr, want)
	}
}

// A link announced while connecting turns live, then the Preview expires.
func TestPreviewViewExpired(t *testing.T) {
	withView(t)
	fake := func(_ context.Context, o connector.QuickOptions) (bool, error) {
		o.OnReady(connector.QuickTunnel{URL: testURL, ExpiresAt: time.Now().Add(time.Hour)})
		o.OnReachable(connector.QuickTunnel{URL: testURL, ExpiresAt: time.Now().Add(time.Hour), Reachable: true})
		return true, nil
	}
	stdout, stderr, code, _ := execute(t, fake, nil, "3000")
	screen := escapes.ReplaceAllString(stdout, "")
	if code != 0 || stderr != "" {
		t.Errorf("exit %d stderr %q", code, stderr)
	}
	for _, want := range []string{"Starting up.", "Your link is live", "Preview expired. The link no longer works."} {
		if !strings.Contains(screen, want) {
			t.Errorf("view lacks %q:\n%s", want, screen)
		}
	}
}

// A failure ends the view, then prints its reason as without the view. A
// revoked Preview with an owner gets the owner-stop message.
func TestPreviewViewFailed(t *testing.T) {
	withView(t)
	fake := func(_ context.Context, o connector.QuickOptions) (bool, error) {
		o.OnReady(connector.QuickTunnel{URL: testURL, ExpiresAt: time.Now().Add(time.Hour)})
		o.OnPause(connector.QuickPause{Reason: connector.PauseCapExceeded, RetryAt: time.Now().Add(time.Hour)})
		return false, &connector.QuickError{Kind: connector.ErrTakenDown, Err: connector.ErrTakenDown,
			Owner: connector.OwnerLink{Ticket: "ticket"}}
	}
	stdout, stderr, code, _ := execute(t, fake, nil, "3000")
	screen := escapes.ReplaceAllString(stdout, "")
	if code != ExitTakenDown || stderr != "Error: "+ownerStopped+"\n" {
		t.Errorf("exit %d stderr %q", code, stderr)
	}
	for _, want := range []string{"Too many Previews from your network. Retrying at", "Preview stopped"} {
		if !strings.Contains(screen, want) {
			t.Errorf("view lacks %q:\n%s", want, screen)
		}
	}
}

// A detected server is the view's first step instead of a stderr line.
func TestPreviewViewDetected(t *testing.T) {
	withView(t)
	d := testDetector(newFakeDialer("[::1]:5173"), false, nil)
	stdout, stderr, code, _ := execute(t, nil, &d)
	if code != 0 || stderr != "" {
		t.Errorf("exit %d stderr %q", code, stderr)
	}
	if screen := escapes.ReplaceAllString(stdout, ""); !strings.Contains(screen, "Found your server at http://[::1]:5173") {
		t.Errorf("view:\n%s", screen)
	}
}

// JSON keeps its own output, and so do an AI agent's shell and a run without
// a person (CI, --no-input): no view.
func TestPreviewViewSkipped(t *testing.T) {
	made := withView(t)
	if _, _, code, _ := execute(t, nil, nil, "3000", "--json"); code != 0 {
		t.Errorf("--json exit %d", code)
	}
	t.Setenv("CLAUDECODE", "1")
	if _, stderr, code, _ := execute(t, nil, nil, "3000"); code != 0 || stderr != "Preview ended\n" {
		t.Errorf("agent shell exit %d stderr %q", code, stderr)
	}
	t.Setenv("CLAUDECODE", "")
	testInteractive = false
	if _, stderr, code, _ := execute(t, nil, nil, "3000"); code != 0 || stderr != "Preview ended\n" {
		t.Errorf("non-interactive exit %d stderr %q", code, stderr)
	}
	if *made != 0 {
		t.Errorf("made %d views", *made)
	}
}

// --verbose keeps the view; the connector's debug log prints above it.
func TestPreviewViewVerbose(t *testing.T) {
	made := withView(t)
	viper.Set("verbose", true)
	t.Cleanup(func() { viper.Set("verbose", false) })
	fake := func(_ context.Context, o connector.QuickOptions) (bool, error) {
		o.Logger.Debug("dialing edge")
		return false, nil
	}
	_, stderr, code, _ := execute(t, fake, nil, "3000")
	if code != 0 || *made != 1 || !strings.Contains(stderr, "dialing edge") {
		t.Errorf("exit %d, %d views, stderr %q", code, *made, stderr)
	}
}
