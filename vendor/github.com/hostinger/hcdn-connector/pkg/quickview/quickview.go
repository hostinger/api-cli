// Package quickview shows a quick tunnel to a person at a terminal: the
// Hostinger logo in a gradient, a checklist while the tunnel connects, the
// link in a box, and a status line with a pulsing dot while it is live.
//
// The view redraws the lines below the logo in place. Anything else printed
// while it runs must go through [View.LogWriter], which prints above them.
// Use it only when [Supported] reports true; pipes, logs and JSON output
// should keep plain lines.
//
//	v := quickview.New(os.Stdout, quickview.Options{Target: "http://localhost:5173"})
//	v.Start()
//	expired, err := connector.RunQuick(ctx, connector.QuickOptions{
//		URL:         "http://localhost:5173",
//		Logger:      slog.New(slog.NewTextHandler(v.LogWriter(os.Stderr), nil)),
//		OnReady:     func(t connector.QuickTunnel) { v.Ready(t, "") },
//		OnReachable: func(connector.QuickTunnel) { v.Reachable() },
//	})
//	v.End(quickview.Stopped)
package quickview

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/hostinger/hcdn-connector/internal/termsafe"
	"github.com/hostinger/hcdn-connector/pkg/connector"
)

const frameEvery = 80 * time.Millisecond

// Options configure a View.
type Options struct {
	// Target is the origin the tunnel shares, such as http://localhost:5173.
	Target string
	// Detected says Target was found on this computer rather than given.
	Detected bool
	// Now and Location set the clock and the time zone of the expiry.
	// nil: time.Now and time.Local.
	Now      func() time.Time
	Location *time.Location
}

// Ending is how a tunnel ended.
type Ending int

const (
	// Stopped: the person stopped it, for example with Ctrl+C.
	Stopped Ending = iota
	// Expired: its time ran out.
	Expired
	// Failed: it ended with an error. The caller prints the reason after End.
	Failed
)

type stepState int

const (
	stepPending stepState = iota
	stepActive
	stepDone
	stepFailed
)

type step struct {
	state       stepState
	doing, done string
}

// View draws a quick tunnel in place. Its methods are safe to call from
// several goroutines, as the connector's callbacks do.
type View struct {
	mu   sync.Mutex
	out  io.Writer
	pal  palette
	size func() (cols, rows int)
	now  func() time.Time
	loc  *time.Location

	steps     []step
	tunnel    *connector.QuickTunnel
	account   string
	reachable bool
	pause     string // why the tunnel is paused; "" when it is not
	pauseTill time.Time
	ended     bool
	ending    Ending

	began   time.Time
	drawn   []string // the rows on screen, as last written; the cursor is below them
	stop    chan struct{}
	done    chan struct{}
	started bool
	once    sync.Once
}

// New returns a view that writes to out. The terminal size comes from out
// when it is an *os.File; otherwise the view assumes 80 columns.
func New(out io.Writer, o Options) *View {
	size := func() (int, int) { return 80, 24 }
	if f, ok := out.(*os.File); ok {
		size = terminalSize(f)
	}
	if o.Now == nil {
		o.Now = time.Now
	}
	if o.Location == nil {
		o.Location = time.Local
	}
	first := "Sharing " + termsafe.String(o.Target)
	if o.Detected {
		first = "Found your server at " + termsafe.String(o.Target)
	}
	return &View{
		out: out, pal: palette{truecolor: hasTruecolor()}, size: size, now: o.Now, loc: o.Location,
		steps: []step{
			{stepDone, first, first},
			{stepActive, "Opening a secure tunnel", "Opened a secure tunnel"},
			{stepPending, "Getting your link ready", "Your link is live"},
		},
	}
}

// Start prints the logo, draws the checklist and animates the view until
// End. Call it once, before the tunnel starts.
func (v *View) Start() {
	cols, _ := v.size()
	writeLogo(v.out, v.pal, cols)
	v.mu.Lock()
	v.began = v.now()
	v.started = true
	v.draw()
	v.mu.Unlock()
	v.stop, v.done = make(chan struct{}), make(chan struct{})
	go func() {
		defer close(v.done)
		tick := time.NewTicker(frameEvery)
		defer tick.Stop()
		for {
			select {
			case <-v.stop:
				return
			case <-tick.C:
				v.mu.Lock()
				v.draw()
				v.mu.Unlock()
			}
		}
	}()
}

// Ready shows the link (connector.QuickOptions.OnReady). account is a line
// about the account the tunnel is linked to; "" hides it.
func (v *View) Ready(t connector.QuickTunnel, account string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.tunnel, v.account = &t, account
	v.steps[1].state = stepDone
	v.steps[2].state = stepActive
	if t.Reachable {
		v.steps[2].state = stepDone
		v.reachable = true
	}
	v.draw()
}

// Reachable marks the link as answering (connector.QuickOptions.OnReachable).
func (v *View) Reachable() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.steps[2].state = stepDone
	v.reachable = true
	v.draw()
}

// Paused shows why the tunnel waits, until it retries at till.
func (v *View) Paused(why string, till time.Time) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.pause, v.pauseTill = termsafe.String(why), till
	v.draw()
}

// End stops the animation and draws how the tunnel ended. Call it once
// RunQuick has returned.
func (v *View) End(e Ending) {
	v.once.Do(func() {
		if v.stop != nil {
			close(v.stop)
			<-v.done
		}
	})
	v.mu.Lock()
	defer v.mu.Unlock()
	v.ended, v.ending = true, e
	for i := range v.steps {
		if v.steps[i].state != stepActive {
			continue
		}
		v.steps[i].state = stepPending
		if e == Failed {
			v.steps[i].state = stepFailed
		}
	}
	v.draw()
}

// LogWriter returns a writer that prints to w above the view, then draws the
// view again below. Use it for the connector's Logger and for any other
// output while the view runs.
func (v *View) LogWriter(w io.Writer) io.Writer {
	return writerFunc(func(p []byte) (int, error) {
		v.mu.Lock()
		defer v.mu.Unlock()
		if n := len(v.drawn); n > 0 {
			fmt.Fprintf(v.out, "\x1b[%dA\r\x1b[J", n)
			v.drawn = nil
		}
		n, err := w.Write(p)
		v.draw()
		return n, err
	})
}

type writerFunc func([]byte) (int, error)

func (f writerFunc) Write(p []byte) (int, error) { return f(p) }

// draw writes the rows that changed since the last draw. v.mu is held.
func (v *View) draw() {
	if !v.started {
		return
	}
	cols, _ := v.size()
	ls := v.lines(v.now())
	rows := make([]string, len(ls))
	for i, l := range ls {
		rows[i] = l.render(cols - 1)
	}
	i := 0
	for i < len(rows) && i < len(v.drawn) && rows[i] == v.drawn[i] {
		i++
	}
	if i == len(rows) && i == len(v.drawn) {
		return
	}
	var b strings.Builder
	if up := len(v.drawn) - i; up > 0 {
		fmt.Fprintf(&b, "\x1b[%dA", up)
	}
	b.WriteString("\r\x1b[J")
	for _, r := range rows[i:] {
		b.WriteString(r + "\n")
	}
	_, _ = io.WriteString(v.out, b.String())
	v.drawn = rows
}

// lines is the view at now.
func (v *View) lines(now time.Time) []line {
	t := now.Sub(v.began)
	ls := []line{}
	for _, s := range v.steps {
		ls = append(ls, v.stepLine(s, t))
	}
	if v.tunnel != nil {
		ls = append(ls, nil)
		ls = append(ls, v.box()...)
		if !v.ended {
			ls = append(ls, nil)
			ls = append(ls, v.details(now)...)
			ls = append(ls, nil, line{
				sp("  ! ", v.pal.fg(colAmber)),
				sp("Anyone with the link can open this server, including debug consoles and dev tools.", v.pal.fg(colGray)),
			})
		}
	}
	ls = append(ls, nil)
	ls = append(ls, v.status(now, t)...)
	if v.ended {
		ls = append(ls, nil) // room before the shell prompt
	}
	return ls
}

var spinner = []rune("⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏")

func spin(t time.Duration) string { return string(spinner[int(t/frameEvery)%len(spinner)]) }

func (v *View) stepLine(s step, t time.Duration) line {
	switch s.state {
	case stepActive:
		return line{sp("  "+spin(t)+" ", v.pal.fg(colLilac)), sp(s.doing, "")}
	case stepDone:
		return line{sp("  ✓ ", v.pal.fg(colGreen)), sp(s.done, "")}
	case stepFailed:
		return line{sp("  ✕ ", v.pal.fg(colRed)), sp(s.doing, v.pal.fg(colRed))}
	}
	return line{sp("  ○ ", v.pal.fg(colGray)), sp(s.doing, v.pal.fg(colGray))}
}

const (
	titleLive  = " Your preview link "
	titleEnded = " Link closed "
)

// box is the link in a frame, crossed out once the tunnel ends.
func (v *View) box() []line {
	url := termsafe.String(v.tunnel.URL)
	frame, link := v.pal.fg(colLilac), bold+v.pal.fg(colWhite)
	title := titleLive
	if v.ended {
		frame, link = v.pal.fg(colGray), strike+v.pal.fg(colGray)
		title = titleEnded
	}
	// At least as wide as the longer title, so a short link never
	// leaves the top border with a negative length.
	inner := max(utf8.RuneCountInString(url)+6, utf8.RuneCountInString(titleLive)+2)
	if cols, _ := v.size(); inner+4 >= cols {
		return []line{{sp("  "+strings.TrimSpace(title)+"  ", v.pal.fg(colGray)), sp(url, link)}}
	}
	blank := line{sp("  │"+strings.Repeat(" ", inner)+"│", frame)}
	return []line{
		{sp("  ╭─", frame), sp(title, frame), sp(strings.Repeat("─", inner-1-utf8.RuneCountInString(title))+"╮", frame)},
		blank,
		{sp("  │   ", frame), sp(url, link), sp(strings.Repeat(" ", inner-3-utf8.RuneCountInString(url))+"│", frame)},
		blank,
		{sp("  ╰"+strings.Repeat("─", inner)+"╯", frame)},
	}
}

func (v *View) details(now time.Time) []line {
	key := v.pal.fg(colGray)
	exp := v.tunnel.ExpiresAt
	ls := []line{{
		sp("  Expires    ", key),
		sp("in "+roughDuration(exp.Sub(now)), ""),
		sp(", at "+exp.In(v.loc).Format("15:04 MST, 2 Jan"), key),
	}}
	if v.account != "" {
		ls = append(ls, line{sp("  Account    ", key), sp(termsafe.String(v.account), "")})
	}
	return ls
}

// roughDuration is d to the minute: "23h 59m", "45m", "under a minute".
func roughDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	switch {
	case d < time.Minute:
		return "under a minute"
	case d < time.Hour:
		return fmt.Sprintf("%dm", int(d.Minutes()))
	case d%time.Hour == 0:
		return fmt.Sprintf("%dh", int(d.Hours()))
	}
	return fmt.Sprintf("%dh %dm", int(d.Hours()), int(d.Minutes())%60)
}

// pulse goes from 0 to 1 and back every 1.6 seconds.
func pulse(t time.Duration) float64 {
	const period = 1600 * time.Millisecond
	return (1 - math.Cos(2*math.Pi*float64(t%period)/float64(period))) / 2
}

func (v *View) status(now time.Time, t time.Duration) []line {
	gray := v.pal.fg(colGray)
	keycap := bold + v.pal.fg(colWhite) + v.pal.bg(colSoftRed)
	stop := []span{sp("Press ", gray), sp(" Ctrl+C ", keycap), sp(" to stop sharing", gray)}
	switch {
	case v.ended && v.ending == Stopped:
		return []line{{sp("  ○ ", gray), sp("Preview ended. The link no longer works.", gray)}}
	case v.ended && v.ending == Expired:
		return []line{{sp("  ○ ", gray), sp("Preview expired. The link no longer works.", gray)}}
	case v.ended:
		return []line{{sp("  ✕ ", v.pal.fg(colRed)), sp("Preview stopped", v.pal.fg(colRed))}}
	case v.pause != "" && now.Before(v.pauseTill):
		dot := v.pal.fg(mix(colGray, colAmber, pulse(t)))
		return []line{{sp("  ● ", dot), sp(v.pause, v.pal.fg(colAmber))}, append(line{sp("  ", "")}, stop...)}
	case v.tunnel == nil:
		return []line{append(line{sp("  ", "")}, stop[0], stop[1], sp(" to cancel", gray))}
	case !v.reachable:
		return []line{
			{sp("  "+spin(t)+" ", v.pal.fg(colAmber)), sp("Starting up. The link may show an error for a few seconds.", v.pal.fg(colAmber))},
			append(line{sp("  ", "")}, stop...),
		}
	}
	dot := v.pal.fg(mix(colDimGreen, colGreen, pulse(t)))
	return []line{append(line{sp("  ● ", dot), sp("Live", bold+v.pal.fg(colGreen)), sp("     ", "")}, stop...)}
}
