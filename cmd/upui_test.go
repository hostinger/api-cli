package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/pterm/pterm"
)

const hostile = "a\x1b]52;c;aGk=\x07b"

// captureStderr returns what f writes to os.Stderr and pterm's default output.
func captureStderr(t *testing.T, f func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stderr
	os.Stderr = w
	pterm.SetDefaultOutput(w)
	defer func() {
		os.Stderr = old
		pterm.SetDefaultOutput(old)
	}()
	f()
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out)
}

func assertSafe(t *testing.T, name, out string, wantLines int) {
	t.Helper()
	if strings.Contains(out, "\x1b]52") || strings.ContainsRune(out, '\a') {
		t.Errorf("%s printed the escape sequence: %q", name, out)
	}
	if got := strings.Count(out, "a�]52;c;aGk=�b"); got != wantLines {
		t.Errorf("%s printed %d sanitised messages, want %d: %q", name, got, wantLines, out)
	}
}

func TestTermUIStripsControlCharacters(t *testing.T) {
	u := &termUI{}
	out := captureStderr(t, func() {
		u.Info(hostile)
		step := u.Start(hostile)
		step.Update(hostile)
		step.Done(hostile)
		u.Start(hostile).Fail(hostile)
	})
	assertSafe(t, "plain", out, 5)

	var buf bytes.Buffer
	sp, err := pterm.DefaultSpinner.WithWriter(&buf).Start("x")
	if err != nil {
		t.Fatal(err)
	}
	step := &spinnerStep{sp: sp}
	step.Update(hostile)
	step.Done(hostile)
	if err := sp.Stop(); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(buf.String(), "\x1b]52") || !strings.Contains(buf.String(), "�]52") {
		t.Errorf("spinner = %q", buf.String())
	}
	buf.Reset()
	sp, _ = pterm.DefaultSpinner.WithWriter(&buf).Start("x")
	(&spinnerStep{sp: sp}).Fail(hostile)
	if strings.Contains(buf.String(), "\x1b]52") || !strings.Contains(buf.String(), "�]52") {
		t.Errorf("spinner fail = %q", buf.String())
	}
}

func TestTermUIPromptsNeedATerminal(t *testing.T) {
	u := &termUI{}
	if _, err := u.Select(hostile, []string{hostile}, 0); err != errNotInteractive {
		t.Errorf("Select err = %v", err)
	}
	if _, err := u.Input(hostile, ""); err != errNotInteractive {
		t.Errorf("Input err = %v", err)
	}
	if _, err := u.Confirm(hostile, false); err != errNotInteractive {
		t.Errorf("Confirm err = %v", err)
	}
}

func TestSafeAll(t *testing.T) {
	got := safeAll([]string{hostile, "ok"})
	if len(got) != 2 || got[0] != "a�]52;c;aGk=�b" || got[1] != "ok" {
		t.Errorf("safeAll = %q", got)
	}
}
