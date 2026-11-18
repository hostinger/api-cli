package quickview

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

type rgb struct{ r, g, b uint8 }

var (
	colPurple   = rgb{0x67, 0x3d, 0xe6} // Hostinger purple
	colPink     = rgb{0xa8, 0x55, 0xf7}
	colLilac    = rgb{0x8c, 0x85, 0xff}
	colGray     = rgb{0x6b, 0x72, 0x80}
	colDimGreen = rgb{0x0b, 0x5c, 0x4a}
	colGreen    = rgb{0x1f, 0xe0, 0xa8}
	colAmber    = rgb{0xff, 0xb0, 0x20}
	colRed      = rgb{0xfc, 0x51, 0x85}
	colSoftRed  = rgb{0xd9, 0x5f, 0x6e} // the Ctrl+C key: red, but calm
	colWhite    = rgb{0xff, 0xff, 0xff}
)

const (
	reset  = "\x1b[0m"
	bold   = "\x1b[1m"
	strike = "\x1b[9m"
)

func mix(a, b rgb, t float64) rgb {
	l := func(x, y uint8) uint8 { return uint8(math.Round(float64(x) + (float64(y)-float64(x))*t)) }
	return rgb{l(a.r, b.r), l(a.g, b.g), l(a.b, b.b)}
}

// palette writes colours as 24-bit escapes, or as the nearest of the 256
// xterm colours when the terminal does not announce 24-bit colour.
type palette struct{ truecolor bool }

func (p palette) fg(c rgb) string { return p.code(38, c) }
func (p palette) bg(c rgb) string { return p.code(48, c) }

func (p palette) code(base int, c rgb) string {
	if p.truecolor {
		return fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", base, c.r, c.g, c.b)
	}
	six := func(x uint8) int { return int(math.Round(float64(x) / 255 * 5)) }
	return fmt.Sprintf("\x1b[%d;5;%dm", base, 16+36*six(c.r)+6*six(c.g)+six(c.b))
}

// hasTruecolor reports a terminal that announces 24-bit colour.
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

// span is text in one style. A line is a list of spans, so its width is known
// without parsing escapes.
type span struct {
	text  string
	style string // escapes written before text; "" for plain text
}

func sp(text, style string) span { return span{text: text, style: style} }

type line []span

// render writes l cut to cols columns, so that no line wraps and the view
// knows how many rows it takes.
func (l line) render(cols int) string {
	var b strings.Builder
	left := cols
	for _, s := range l {
		if left <= 0 {
			break
		}
		t := s.text
		if n := utf8.RuneCountInString(t); n > left {
			t = string([]rune(t)[:left])
		}
		left -= utf8.RuneCountInString(t)
		if s.style == "" {
			b.WriteString(t)
			continue
		}
		b.WriteString(s.style + t + reset)
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

// writeLogo prints the logo with a diagonal purple to pink gradient, or the
// word alone when the terminal is too narrow for it.
func writeLogo(w io.Writer, pal palette, cols int) {
	art := logoArt
	width := 0
	for _, row := range art {
		width = max(width, utf8.RuneCountInString(row))
	}
	if cols < width+2 {
		art = []string{"  H O S T I N G E R"}
		width = utf8.RuneCountInString(art[0])
	}
	diag := float64(width + 2*len(art))
	fmt.Fprintln(w)
	for y, row := range art {
		var l line
		for x, r := range []rune(row) {
			if r == ' ' {
				l = append(l, sp(" ", ""))
				continue
			}
			c := mix(colPurple, colPink, (float64(x)+2*float64(y))/diag)
			l = append(l, sp(string(r), bold+pal.fg(c)))
		}
		fmt.Fprintln(w, l.render(cols))
	}
	sub := "P R E V I E W"
	pad := max(width-utf8.RuneCountInString(sub), 2)
	fmt.Fprintln(w, line{sp(strings.Repeat(" ", pad), ""), sp(sub, pal.fg(colLilac))}.render(cols))
	fmt.Fprintln(w)
}
