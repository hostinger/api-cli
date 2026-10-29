package output

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/pterm/pterm"
)

// captureStdout returns what fn writes to os.Stdout and pterm's output (table
// and tree rendering write there directly).
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdout
	os.Stdout = w
	pterm.SetDefaultOutput(w)
	done := make(chan string)
	go func() {
		b, _ := io.ReadAll(r)
		done <- string(b)
	}()
	fn()
	os.Stdout = old
	pterm.SetDefaultOutput(old)
	w.Close()
	return <-done
}

const hostileBody = `{"data":[{"name":"Alice\u001b]52;c;aGk=\u0007","note\u001b[2K":"line1\nline2\u001b]0;pwned\u001b\\","nested":{"k\u0007ey":"v\u009b2K","list":["x\u001b[2K"]}}],
"meta":{"total":1,"title\u001b]0;t\u0007":"Café 🚀"}}`

func assertSafe(t *testing.T, out string) {
	t.Helper()
	for _, bad := range []string{"\x1b]", "\x1b[2K", "\x07", "\u009b"} {
		if strings.Contains(out, bad) {
			t.Errorf("output contains %q:\n%q", bad, out)
		}
	}
	for _, want := range []string{"Alice", "line1", "line2", "Café 🚀", "\uFFFD"} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestTableStripsControlCharacters(t *testing.T) {
	assertSafe(t, captureStdout(t, func() { printTable([]byte(hostileBody)) }))
}

func TestTableStripsControlCharactersInArray(t *testing.T) {
	out := captureStdout(t, func() { printTable([]byte(`[{"Alice\u001b[2K":"Bob\u001b]52;c;aGk=\u0007"}]`)) })
	if strings.ContainsAny(out, "\x07") || strings.Contains(out, "\x1b]") || strings.Contains(out, "\x1b[2K") {
		t.Errorf("output contains control characters: %q", out)
	}
	if !strings.Contains(out, "Bob") {
		t.Errorf("output missing Bob: %s", out)
	}
}

func TestTreeStripsControlCharacters(t *testing.T) {
	assertSafe(t, captureStdout(t, func() { printTreeEnhanced([]byte(hostileBody)) }))
}
