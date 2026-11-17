// Package brand is the Hostinger terminal look of hcdn-connector's quickview.
// Spans are written as given: make outside text (file names, API errors)
// safe before it gets here.
package brand

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
	"golang.org/x/term"
)

type RGB struct{ R, G, B uint8 }

var (
	Purple  = RGB{0x67, 0x3d, 0xe6}
	Pink    = RGB{0xa8, 0x55, 0xf7}
	Lilac   = RGB{0x8c, 0x85, 0xff}
	Gray    = RGB{0x6b, 0x72, 0x80}
	Green   = RGB{0x1f, 0xe0, 0xa8}
	Amber   = RGB{0xff, 0xb0, 0x20}
	Red     = RGB{0xfc, 0x51, 0x85}
	SoftRed = RGB{0xd9, 0x5f, 0x6e}
	White   = RGB{0xff, 0xff, 0xff}
)

const (
	Reset = "\x1b[0m"
	Bold  = "\x1b[1m"
)

func Mix(a, b RGB, t float64) RGB {
	l := func(x, y uint8) uint8 { return uint8(math.Round(float64(x) + (float64(y)-float64(x))*t)) }
	return RGB{l(a.R, b.R), l(a.G, b.G), l(a.B, b.B)}
}

// Palette falls back to the 256 xterm colours when the terminal does not
// announce 24-bit colour.
type Palette struct{ Truecolor bool }

func NewPalette() Palette { return Palette{Truecolor: hasTruecolor()} }

func (p Palette) Fg(c RGB) string { return p.code(38, c) }
func (p Palette) Bg(c RGB) string { return p.code(48, c) }

func (p Palette) code(base int, c RGB) string {
	if p.Truecolor {
		return fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", base, c.R, c.G, c.B)
	}
	return fmt.Sprintf("\x1b[%d;5;%dm", base, xterm256(c))
}

var cubeLevels = [6]int{0, 0x5f, 0x87, 0xaf, 0xd7, 0xff}

// xterm256 also considers the grey ramp (232-255): in the cube alone Gray
// lands on a blue-lilac (103) that reads like Lilac.
func xterm256(c RGB) int {
	level := func(x uint8) int {
		switch {
		case x < 48:
			return 0
		case x < 115:
			return 1
		}
		return (int(x) - 35) / 40
	}
	dist := func(r, g, b int) int {
		dr, dg, db := int(c.R)-r, int(c.G)-g, int(c.B)-b
		return dr*dr + dg*dg + db*db
	}
	ri, gi, bi := level(c.R), level(c.G), level(c.B)
	cube := 16 + 36*ri + 6*gi + bi
	cubeDist := dist(cubeLevels[ri], cubeLevels[gi], cubeLevels[bi])

	avg := (int(c.R) + int(c.G) + int(c.B)) / 3
	step := min(max((avg-3)/10, 0), 23) // ramp step i is 8+10i
	v := 8 + 10*step
	if dist(v, v, v) < cubeDist {
		return 232 + step
	}
	return cube
}

func hasTruecolor() bool {
	switch strings.ToLower(os.Getenv("COLORTERM")) {
	case "truecolor", "24bit":
		return true
	}
	switch os.Getenv("TERM_PROGRAM") {
	case "iTerm.app", "WezTerm", "ghostty", "vscode", "WarpTerminal":
		return true
	}
	return false
}

// Supported also turns on escape sequences in a Windows console that can take them.
func Supported(f *os.File) bool {
	if f == nil || !term.IsTerminal(int(f.Fd())) {
		return false
	}
	if _, noColor := os.LookupEnv("NO_COLOR"); noColor || os.Getenv("TERM") == "dumb" || agentShell() {
		return false
	}
	return enableEscapes(f)
}

// agentShell reports a shell started by an AI agent, which may still have a
// pseudo-terminal: a redrawn view there only costs it tokens.
func agentShell() bool {
	return os.Getenv("AI_AGENT") != "" || os.Getenv("CLAUDECODE") != ""
}

func Width(f *os.File) int {
	w, _, err := term.GetSize(int(f.Fd()))
	if err != nil || w <= 0 {
		return 80
	}
	return w
}

// cells counts ambiguous characters (box drawing, ✓, ●) as one column, as the
// view draws them; runewidth's default reads the locale and makes them two.
var cells = &runewidth.Condition{StrictEmojiNeutral: true}

func TextWidth(s string) int { return cells.StringWidth(s) }

// Span keeps text apart from its escapes, so a line's width is known.
type Span struct {
	Text  string
	Style string
}

func S(text, style string) Span { return Span{Text: text, Style: style} }

type Line []Span

// Render cuts l to cols columns so it never wraps.
func (l Line) Render(cols int) string {
	var b strings.Builder
	left := cols
	for _, s := range l {
		if left <= 0 {
			break
		}
		t := s.Text
		if TextWidth(t) > left {
			t = cells.Truncate(t, left, "…")
			left = 0
		} else {
			left -= TextWidth(t)
		}
		if s.Style == "" {
			b.WriteString(t)
			continue
		}
		b.WriteString(s.Style + t + Reset)
	}
	return b.String()
}

var logoArt = []string{
	" ██╗  ██╗  ██████╗  ███████╗ ████████╗ ██╗ ███╗   ██╗  ██████╗  ███████╗ ██████╗",
	" ██║  ██║ ██╔═══██╗ ██╔════╝ ╚══██╔══╝ ██║ ████╗  ██║ ██╔════╝  ██╔════╝ ██╔══██╗",
	" ███████║ ██║   ██║ ███████╗    ██║    ██║ ██╔██╗ ██║ ██║  ███╗ █████╗   ██████╔╝",
	" ██╔══██║ ██║   ██║ ╚════██║    ██║    ██║ ██║╚██╗██║ ██║   ██║ ██╔══╝   ██╔══██╗",
	" ██║  ██║ ╚██████╔╝ ███████║    ██║    ██║ ██║ ╚████║ ╚██████╔╝ ███████╗ ██║  ██║",
	" ╚═╝  ╚═╝  ╚═════╝  ╚══════╝    ╚═╝    ╚═╝ ╚═╝  ╚═══╝  ╚═════╝  ╚══════╝ ╚═╝  ╚═╝",
}

// Logo falls back to a one-line wordmark when cols is too narrow.
func Logo(w io.Writer, p Palette, cols int, subtitle string) {
	art := logoArt
	width := 0
	for _, row := range art {
		width = max(width, TextWidth(row))
	}
	if cols < width+2 {
		art = []string{"  H O S T I N G E R"}
		width = TextWidth(art[0])
	}
	diag := float64(width + 2*len(art))
	fmt.Fprintln(w)
	for y, row := range art {
		var l Line
		for x, r := range []rune(row) {
			if r == ' ' {
				l = append(l, S(" ", ""))
				continue
			}
			c := Mix(Purple, Pink, (float64(x)+2*float64(y))/diag)
			l = append(l, S(string(r), Bold+p.Fg(c)))
		}
		fmt.Fprintln(w, l.Render(cols))
	}
	if subtitle != "" {
		pad := max(width-TextWidth(subtitle), 2)
		fmt.Fprintln(w, Line{S(strings.Repeat(" ", pad), ""), S(subtitle, p.Fg(Lilac))}.Render(cols))
	}
	fmt.Fprintln(w)
}

// Box falls back to one plain row when cols is too narrow for the frame.
func Box(title, text, frame, textStyle string, cols int) []Line {
	title = " " + title + " "
	inner := max(TextWidth(text)+6, TextWidth(title)+2)
	if inner+4 >= cols {
		return []Line{{S("  "+strings.TrimSpace(title)+"  ", frame), S(text, textStyle)}}
	}
	blank := Line{S("  │"+strings.Repeat(" ", inner)+"│", frame)}
	return []Line{
		{S("  ╭─", frame), S(title, frame), S(strings.Repeat("─", inner-1-TextWidth(title))+"╮", frame)},
		blank,
		{S("  │   ", frame), S(text, textStyle), S(strings.Repeat(" ", inner-3-TextWidth(text))+"│", frame)},
		blank,
		{S("  ╰"+strings.Repeat("─", inner)+"╯", frame)},
	}
}

var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func Spinner(n int) string { return spinnerFrames[n%len(spinnerFrames)] }
