package cmd

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/hostinger/api-cli/brand"
	"github.com/hostinger/api-cli/deploy/hosting"
)

// lockedBuffer is written by the spinner goroutine and read by the test.
type lockedBuffer struct {
	mu sync.Mutex
	b  bytes.Buffer
}

func (l *lockedBuffer) Write(p []byte) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.b.Write(p)
}

func (l *lockedBuffer) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.b.String()
}

var (
	sgr      = regexp.MustCompile(`\x1b\[[0-9;]*m`)
	cursorUp = regexp.MustCompile(`^\x1b\[(\d+)A`)
)

// screen replays what the view wrote, as a terminal would show it: cursor up
// with clear to the end of the screen, and colours dropped.
func screen(raw string) string {
	var rows []string
	cur := ""
	s := sgr.ReplaceAllString(raw, "")
	for len(s) > 0 {
		if m := cursorUp.FindStringSubmatch(s); m != nil {
			var n int
			fmt.Sscan(m[1], &n)
			rows = rows[:max(len(rows)-n, 0)]
			cur = ""
			s = s[len(m[0]):]
			continue
		}
		switch {
		case strings.HasPrefix(s, "\r\x1b[J"):
			cur = ""
			s = s[4:]
		case s[0] == '\n':
			rows = append(rows, cur)
			cur = ""
			s = s[1:]
		default:
			cur += s[:1]
			s = s[1:]
		}
	}
	if cur != "" {
		rows = append(rows, cur)
	}
	return strings.Join(rows, "\n")
}

func testBrandView(out *lockedBuffer, width int) (*brandView, *time.Time) {
	now := time.Date(2026, 11, 17, 12, 0, 0, 0, time.UTC)
	return &brandView{
		out: out, pal: brand.Palette{}, width: func() int { return width },
		now: func() time.Time { return now },
	}, &now
}

func TestBrandViewSteps(t *testing.T) {
	var out lockedBuffer
	v, now := testBrandView(&out, 100)
	s := v.start("Uploading " + hostile)
	if got := screen(out.String()); !strings.Contains(got, "█") {
		t.Fatalf("no logo with the first step: %q", got)
	}
	s.Update("Uploading\nmore")
	if got := screen(out.String()); !strings.HasSuffix(got, "Uploading more\n\n  Press  Ctrl+C  to cancel") {
		t.Errorf("live block = %q", got)
	}
	v.mu.Lock() // the spinner goroutine reads the clock under v.mu
	*now = now.Add(65 * time.Second)
	v.mu.Unlock()
	s.Done("Uploaded")
	s.Done("again") // a step ends once: nothing more is printed
	s.Fail("again")
	v.start("Building").Fail("Build failed")
	got := screen(out.String())
	if !strings.HasSuffix(got, "  ✓ Uploaded  1m 05s\n  ✕ Build failed\n\n  Press  Ctrl+C  to cancel") {
		t.Errorf("screen = %q", got)
	}
	if strings.Contains(out.String(), "\x1b]52") {
		t.Errorf("unsafe step text: %q", out.String())
	}
}

func TestBrandViewPrintsAboveLiveBlock(t *testing.T) {
	var out lockedBuffer
	v, _ := testBrandView(&out, 100)
	s := v.start("Packing")
	v.info("Left out: .env\n  second " + hostile)
	fmt.Fprint(v.logWriter(), "> GET https://developers.hostinger.com/x\n")
	got := screen(out.String())
	want := "\n  › Left out: .env\n      second a�]52;c;aGk=�b\n> GET https://developers.hostinger.com/x\n  ⠋ Packing\n\n  Press  Ctrl+C  to cancel"
	if !strings.HasSuffix(got, want) {
		t.Errorf("screen =\n%s\nwant\n%s", got, want)
	}
	s.Done("Packed")
}

func TestBrandViewHidesForPrompts(t *testing.T) {
	var out lockedBuffer
	v, _ := testBrandView(&out, 100)
	v.hide() // a question before any step still comes after the logo
	before := out.String()
	if !strings.Contains(before, "█") {
		t.Errorf("no logo before the first question: %q", before)
	}
	v.mu.Lock()
	v.redraw() // a spinner tick while the prompt shows
	v.mu.Unlock()
	if out.String() != before {
		t.Errorf("drew over a prompt: %q", strings.TrimPrefix(out.String(), before))
	}
	v.show()
	if got := screen(out.String()); !strings.HasSuffix(got, "Press  Ctrl+C  to cancel") {
		t.Errorf("live block not back: %q", got)
	}
}

func TestBrandViewSpacesPrompts(t *testing.T) {
	var out lockedBuffer
	v, _ := testBrandView(&out, 100)
	v.start("Looking up your websites").Done("6 websites")
	for _, q := range []string{"Which website?", "Create a regular website?"} {
		v.hide()
		fmt.Fprintln(&out, q) // what pterm prints
		v.show()
	}
	v.start("Creating website").Done("Created example.com")
	got := screen(out.String())
	want := "  ✓ 6 websites\n\nWhich website?\n\nCreate a regular website?\n\n  ✓ Created example.com\n\n  Press  Ctrl+C  to cancel"
	if !strings.HasSuffix(got, want) {
		t.Errorf("screen ends\n%s\nwant\n%s", got[max(len(got)-len(want)-20, 0):], want)
	}
}

func TestBrandViewFail(t *testing.T) {
	for _, c := range []struct {
		cancelled bool
		want      string
	}{
		{false, "  ✕ Building\n\n  ✕ Deploy failed\n"},
		{true, "  ○ Building\n\n  ○ Deploy cancelled\n"},
	} {
		var out lockedBuffer
		v, _ := testBrandView(&out, 100)
		v.start("Building")
		v.fail(c.cancelled)
		v.fail(c.cancelled) // once only
		if got := screen(out.String()); !strings.HasSuffix(got, "\n"+c.want) {
			t.Errorf("cancelled=%v: %q, want it to end %q", c.cancelled, got, c.want)
		}
	}
	// A deploy that fails before showing anything (a mistyped path) leaves
	// only its error: no logo, no status line.
	var out lockedBuffer
	v, _ := testBrandView(&out, 100)
	fmt.Fprint(v.logWriter(), "> GET /x\n")
	v.fail(false)
	v.stop()
	if got := out.String(); got != "> GET /x\n" {
		t.Errorf("early failure printed %q", got)
	}
}

func TestBrandFinish(t *testing.T) {
	cases := []struct {
		name string
		out  hosting.Outcome
		want []string
		not  []string
	}{
		{"live", hosting.Outcome{URL: "https://example.com", Target: "example.com", Reachable: true, NewLink: true},
			[]string{"Your website is live", "https://example.com", "Linked", "hostinger link", "● Live"}, []string{"Website", "Ctrl+C"}},
		{"no dns", hosting.Outcome{URL: "https://x.hostingersite.com", Target: "example.com", ReachableNote: "Point DNS " + hostile},
			[]string{"Deployed", "Website", "example.com", "● Point DNS"}, []string{"Live", "live", "Linked"}},
		{"importing", hosting.Outcome{URL: "https://example.com", Target: "example.com", Reachable: true, Note: "the import finishes in a few minutes"},
			[]string{"still finishing", "● Not finished yet: the import"}, []string{"Live"}},
		// The Website row compares hosts: a.com is not the site at banana.com.
		{"other host", hosting.Outcome{URL: "https://banana.com", Target: "a.com", Reachable: true},
			[]string{"Website   a.com"}, nil},
		{"same host, other case", hosting.Outcome{URL: "https://Example.com/", Target: "example.com", Reachable: true},
			nil, []string{"Website"}},
	}
	for _, c := range cases {
		var out lockedBuffer
		v, _ := testBrandView(&out, 100)
		v.finish(&c.out)
		got := screen(out.String())
		if raw := out.String(); strings.Contains(raw, "\x1b]52") || strings.ContainsRune(raw, '\a') {
			t.Errorf("%s: wrote an unsafe escape: %q", c.name, raw)
		}
		for _, w := range c.want {
			if !strings.Contains(got, w) {
				t.Errorf("%s: no %q in %q", c.name, w, got)
			}
		}
		for _, n := range c.not {
			if strings.Contains(got, n) {
				t.Errorf("%s: %q in %q", c.name, n, got)
			}
		}
	}
}

func TestShortDuration(t *testing.T) {
	for d, want := range map[time.Duration]string{
		8 * time.Second: "8s", 65 * time.Second: "1m 05s", 62 * time.Minute: "1h 02m",
	} {
		if got := shortDuration(d); got != want {
			t.Errorf("shortDuration(%v) = %q, want %q", d, got, want)
		}
	}
}
