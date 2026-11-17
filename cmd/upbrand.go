package cmd

import (
	"fmt"
	"io"
	"math"
	"net/url"
	"os"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/hostinger/api-cli/brand"
	"github.com/hostinger/api-cli/deploy/hosting"
	"github.com/hostinger/api-cli/utils"
)

const brandFrameEvery = 80 * time.Millisecond

// brandView is `hostinger up` at a terminal. Finished steps scroll up; the
// live block under them (running step, "Press Ctrl+C") is redrawn in place,
// so anything else printed meanwhile must go through printAbove or logWriter.
type brandView struct {
	out   io.Writer
	pal   brand.Palette
	width func() int
	now   func() time.Time

	mu     sync.Mutex
	active *brandStep
	drawn  []string // the live block as on screen; the cursor is below it
	shown  bool     // the logo is printed
	hidden bool     // a prompt is showing
	ended  bool
	blank  bool // the last printed row is empty, so a prompt adds no second one
}

func newBrandView(f *os.File) *brandView {
	return &brandView{
		out: f, pal: brand.NewPalette(), now: time.Now,
		width: func() int { return brand.Width(f) },
	}
}

// showLogo waits for the first output, so a deploy that fails before it
// starts (a mistyped path) prints only its error. v.mu is held.
func (v *brandView) showLogo() {
	if v.shown || v.ended {
		return
	}
	brand.Logo(v.out, v.pal, v.width(), "U P")
	v.shown = true
	v.blank = true
}

func (v *brandView) block() []brand.Line {
	gray := v.pal.Fg(brand.Gray)
	var ls []brand.Line
	if v.active != nil {
		ls = append(ls, v.active.line(v.now()))
	}
	if !v.blank || v.active != nil {
		ls = append(ls, nil)
	}
	return append(ls, brand.Line{
		brand.S("  Press ", gray),
		brand.S(" Ctrl+C ", brand.Bold+v.pal.Fg(brand.White)+v.pal.Bg(brand.SoftRed)),
		brand.S(" to cancel", gray),
	})
}

// redraw and clear: v.mu is held.
func (v *brandView) redraw() {
	if !v.shown || v.hidden || v.ended {
		return
	}
	cols := v.width() - 1
	ls := v.block()
	rows := make([]string, len(ls))
	for i, l := range ls {
		rows[i] = l.Render(cols)
	}
	if slices.Equal(rows, v.drawn) {
		return
	}
	var b strings.Builder
	if n := len(v.drawn); n > 0 {
		fmt.Fprintf(&b, "\x1b[%dA", n)
	}
	b.WriteString("\r\x1b[J")
	for _, r := range rows {
		b.WriteString(r + "\n")
	}
	io.WriteString(v.out, b.String())
	v.drawn = rows
}

func (v *brandView) clear() {
	if n := len(v.drawn); n > 0 {
		fmt.Fprintf(v.out, "\x1b[%dA\r\x1b[J", n)
	}
	v.drawn = nil
}

// printAbove lets lines wrap rather than cutting them: a hint is no use with
// its end missing. v.mu is held.
func (v *brandView) printAbove(ls ...brand.Line) {
	v.clear()
	for _, l := range ls {
		fmt.Fprintln(v.out, l.Render(math.MaxInt))
	}
	if len(ls) > 0 {
		v.blank = len(ls[len(ls)-1]) == 0
	}
	v.redraw()
}

func (v *brandView) logWriter() io.Writer {
	return writerFunc(func(p []byte) (int, error) {
		v.mu.Lock()
		defer v.mu.Unlock()
		v.clear()
		n, err := v.out.Write(p)
		if len(p) > 0 && p[len(p)-1] != '\n' {
			io.WriteString(v.out, "\n") // the redraw counts whole rows
		}
		v.blank = strings.HasSuffix(string(p), "\n\n")
		v.redraw()
		return n, err
	})
}

type writerFunc func([]byte) (int, error)

func (f writerFunc) Write(p []byte) (int, error) { return f(p) }

func (v *brandView) info(msg string) {
	var ls []brand.Line
	for i, row := range strings.Split(utils.SafeText(msg), "\n") {
		mark := brand.S("    ", "")
		if i == 0 {
			mark = brand.S("  › ", v.pal.Fg(brand.Lilac))
		}
		ls = append(ls, brand.Line{mark, brand.S(row, "")})
	}
	v.mu.Lock()
	defer v.mu.Unlock()
	v.showLogo()
	v.printAbove(ls...)
}

// hide and show wrap a prompt, with a blank line on each side of it.
func (v *brandView) hide() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.showLogo()
	v.clear()
	if !v.blank {
		io.WriteString(v.out, "\n")
	}
	v.hidden = true
}

func (v *brandView) show() {
	v.mu.Lock()
	defer v.mu.Unlock()
	io.WriteString(v.out, "\n")
	v.blank = true
	v.hidden = false
	v.redraw()
}

func (v *brandView) start(msg string) *brandStep {
	s := &brandStep{v: v, msg: oneLine(msg), began: v.now(), stop: make(chan struct{}), done: make(chan struct{})}
	v.mu.Lock()
	v.showLogo()
	v.active = s
	v.redraw()
	v.mu.Unlock()
	go func() {
		defer close(s.done)
		tick := time.NewTicker(brandFrameEvery)
		defer tick.Stop()
		for {
			select {
			case <-s.stop:
				return
			case <-tick.C:
				v.mu.Lock()
				v.redraw()
				v.mu.Unlock()
			}
		}
	}()
	return s
}

type brandStep struct {
	v     *brandView
	msg   string // v.mu
	began time.Time
	stop  chan struct{}
	done  chan struct{}
	once  sync.Once // a second Done or Fail does nothing
}

func (s *brandStep) line(now time.Time) brand.Line {
	d := now.Sub(s.began)
	l := brand.Line{brand.S("  "+brand.Spinner(int(d/brandFrameEvery))+" ", s.v.pal.Fg(brand.Lilac)), brand.S(s.msg, "")}
	if d >= time.Second {
		l = append(l, brand.S("  "+shortDuration(d), s.v.pal.Fg(brand.Gray)))
	}
	return l
}

func (s *brandStep) Update(msg string) {
	s.v.mu.Lock()
	defer s.v.mu.Unlock()
	s.msg = oneLine(msg)
	s.v.redraw()
}

func (s *brandStep) Done(msg string) {
	s.end(brand.S("  ✓ ", s.v.pal.Fg(brand.Green)), brand.S(oneLine(msg), ""))
}

func (s *brandStep) Fail(msg string) {
	red := s.v.pal.Fg(brand.Red)
	s.end(brand.S("  ✕ ", red), brand.S(oneLine(msg), red))
}

// end must not hold v.mu: it waits for the spinner goroutine, which takes it.
func (s *brandStep) end(mark, text brand.Span) {
	s.once.Do(func() {
		close(s.stop)
		<-s.done
		v := s.v
		v.mu.Lock()
		defer v.mu.Unlock()
		l := brand.Line{mark, text}
		if d := v.now().Sub(s.began); d >= time.Second {
			l = append(l, brand.S("  "+shortDuration(d), v.pal.Fg(brand.Gray)))
		}
		if v.active == s {
			v.active = nil
		}
		v.printAbove(l)
	})
}

func (v *brandView) finish(out *hosting.Outcome) {
	v.endActive(false)
	p := v.pal
	gray, amber := p.Fg(brand.Gray), p.Fg(brand.Amber)
	link := oneLine(out.URL)
	if link == "" {
		link = oneLine(out.Target)
	}
	title, frame := "Your website is live", p.Fg(brand.Lilac)
	status := brand.Line{brand.S("  ● ", p.Fg(brand.Green)), brand.S("Live", brand.Bold+p.Fg(brand.Green))}
	switch {
	case out.Note != "":
		title, frame = "Deployed, still finishing", amber
		status = brand.Line{brand.S("  ● ", amber), brand.S("Not finished yet: "+oneLine(out.Note), amber)}
	case !out.Reachable:
		title, frame = "Deployed", amber
		status = brand.Line{brand.S("  ● ", amber), brand.S(oneLine(out.ReachableNote), amber)}
	}
	ls := []brand.Line{nil}
	ls = append(ls, brand.Box(title, link, frame, brand.Bold+p.Fg(brand.White), v.width()-1)...)
	var details []brand.Line
	key := func(k string) brand.Span { return brand.S(fmt.Sprintf("  %-10s", k), gray) }
	if target := oneLine(out.Target); target != "" && out.URL != "" && !strings.EqualFold(urlHost(out.URL), target) {
		details = append(details, brand.Line{key("Website"), brand.S(target, "")})
	}
	if out.NewLink {
		details = append(details, brand.Line{key("Linked"), brand.S("this directory; `hostinger up` now redeploys here. "+hosting.RelinkHint, "")})
	}
	if len(details) > 0 {
		ls = append(append(ls, nil), details...)
	}
	ls = append(ls, nil, status, nil)
	v.mu.Lock()
	defer v.mu.Unlock()
	v.showLogo()
	v.clear()
	v.ended = true
	for _, l := range ls {
		fmt.Fprintln(v.out, l.Render(math.MaxInt))
	}
}

func (v *brandView) fail(cancelled bool) {
	v.endActive(!cancelled)
	gray, red := v.pal.Fg(brand.Gray), v.pal.Fg(brand.Red)
	status := brand.Line{brand.S("  ✕ ", red), brand.S("Deploy failed", red)}
	if cancelled {
		status = brand.Line{brand.S("  ○ ", gray), brand.S("Deploy cancelled", gray)}
	}
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.ended {
		return
	}
	v.clear()
	v.ended = true
	if !v.shown { // nothing was shown: the error alone says it
		return
	}
	for _, l := range []brand.Line{nil, status, nil} {
		fmt.Fprintln(v.out, l.Render(math.MaxInt))
	}
}

// stop prints nothing, for output that is not a deploy result (JSON).
func (v *brandView) stop() {
	v.endActive(false)
	v.mu.Lock()
	defer v.mu.Unlock()
	v.clear()
	v.ended = true
}

func (v *brandView) endActive(failed bool) {
	v.mu.Lock()
	s := v.active
	msg := ""
	if s != nil {
		msg = s.msg
	}
	v.mu.Unlock()
	if s == nil {
		return
	}
	if failed {
		s.Fail(msg)
		return
	}
	gray := v.pal.Fg(brand.Gray)
	s.end(brand.S("  ○ ", gray), brand.S(msg, gray))
}

func urlHost(u string) string {
	p, err := url.Parse(u)
	if err != nil {
		return ""
	}
	return p.Hostname()
}

// oneLine: a redrawn row must not contain a newline.
func oneLine(s string) string {
	return strings.Join(strings.Fields(utils.SafeText(s)), " ")
}

func shortDuration(d time.Duration) string {
	d = d.Round(time.Second)
	switch {
	case d < time.Minute:
		return fmt.Sprintf("%ds", int(d.Seconds()))
	case d < time.Hour:
		return fmt.Sprintf("%dm %02ds", int(d.Minutes()), int(d.Seconds())%60)
	}
	return fmt.Sprintf("%dh %02dm", int(d.Hours()), int(d.Minutes())%60)
}
