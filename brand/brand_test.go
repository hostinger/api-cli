package brand

import (
	"bytes"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestRenderCutsToWidth(t *testing.T) {
	l := Line{S("abc", ""), S("défg", "\x1b[1m")}
	if got := l.Render(5); got != "abc\x1b[1md…"+Reset {
		t.Errorf("Render(5) = %q", got)
	}
	if got := l.Render(2); got != "a…" {
		t.Errorf("Render(2) = %q", got)
	}
	if got := l.Render(7); got != "abc\x1b[1mdéfg"+Reset {
		t.Errorf("Render(7) = %q", got)
	}
}

// width is the columns of an unstyled line.
func width(l Line) int { return TextWidth(l.Render(1 << 30)) }

// Wide characters take two columns: a row that counted runes would wrap and
// break the live block's redraw.
func TestRenderMeasuresDisplayColumns(t *testing.T) {
	l := Line{S("  ⠋ ", ""), S("Packing ./プロジェクト", "")}
	if width(l) != 4+10+12 {
		t.Errorf("width = %d, want 26", width(l))
	}
	got := l.Render(16)
	if w := TextWidth(got); w > 16 {
		t.Errorf("Render(16) is %d columns: %q", w, got)
	}
	if got != "  ⠋ Packing ./…" {
		t.Errorf("Render(16) = %q", got)
	}
	box := Box("Your website is live", "https://例え.jp", "", "", 80)
	for i, r := range box {
		if width(r) != width(box[0]) {
			t.Errorf("box row %d is %d columns, want %d", i, width(r), width(box[0]))
		}
	}
}

func TestPaletteFallsBackTo256Colours(t *testing.T) {
	if got := (Palette{Truecolor: true}).Fg(Purple); got != "\x1b[38;2;103;61;230m" {
		t.Errorf("truecolor = %q", got)
	}
	if got := (Palette{}).Bg(White); got != "\x1b[48;5;231m" {
		t.Errorf("256 = %q", got)
	}
	// Gray takes the grey ramp, not the blue-lilac cube colour 103.
	if got := (Palette{}).Fg(Gray); got != "\x1b[38;5;243m" {
		t.Errorf("gray = %q", got)
	}
	if got := xterm256(Lilac); got < 16 || got > 231 {
		t.Errorf("lilac = %d, want a cube colour", got)
	}
}

func TestLogo(t *testing.T) {
	var b bytes.Buffer
	Logo(&b, Palette{}, 120, "U P")
	out := b.String()
	if !strings.Contains(out, "█") || !strings.Contains(out, "U P") {
		t.Errorf("wide logo = %q", out)
	}
	b.Reset()
	Logo(&b, Palette{}, 40, "U P")
	if strings.Contains(b.String(), "█") || !strings.Contains(b.String(), "mR") {
		t.Errorf("narrow logo = %q", b.String())
	}
}

func TestBox(t *testing.T) {
	rows := Box("Your website is live", "https://example.com", "", "", 80)
	if len(rows) != 5 {
		t.Fatalf("rows = %d", len(rows))
	}
	want := width(rows[0])
	for i, r := range rows {
		if width(r) != want {
			t.Errorf("row %d is %d wide, want %d: %q", i, width(r), want, r.Render(200))
		}
	}
	if !strings.Contains(rows[2].Render(200), "https://example.com") {
		t.Errorf("link row = %q", rows[2].Render(200))
	}
	// A short link still fits the title.
	if rows := Box("Your website is live", "a.io", "", "", 80); width(rows[0]) != width(rows[4]) {
		t.Errorf("short link box is ragged")
	}
	narrow := Box("Your website is live", "https://example.com", "", "", 25)
	if len(narrow) != 1 || !strings.Contains(narrow[0].Render(200), "https://example.com") {
		t.Errorf("narrow box = %q", narrow)
	}
}

func TestAgentShell(t *testing.T) {
	t.Setenv("AI_AGENT", "")
	t.Setenv("CLAUDECODE", "")
	if AgentShell() {
		t.Error("agent shell without the variables")
	}
	t.Setenv("AI_AGENT", "claude-code_2-1-326_agent")
	if !AgentShell() {
		t.Error("AI_AGENT not seen")
	}
	t.Setenv("AI_AGENT", "")
	t.Setenv("CLAUDECODE", "1")
	if !AgentShell() {
		t.Error("CLAUDECODE not seen")
	}
}

func TestSpinnerCycles(t *testing.T) {
	if Spinner(0) != Spinner(10) || utf8.RuneCountInString(Spinner(3)) != 1 {
		t.Errorf("spinner frames are off")
	}
}
