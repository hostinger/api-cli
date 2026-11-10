package session

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log/slog"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hostinger/hcdn-connector/internal/h2srv"
	"github.com/hostinger/hcdn-connector/internal/metrics"
	"github.com/hostinger/hcdn-connector/internal/proto"
	"github.com/hostinger/hcdn-connector/pkg/config"
	"github.com/hostinger/hcdn-connector/pkg/creds"
)

// ProbeRow is one line of the `status` probe table.
type ProbeRow struct {
	Pop      string  `json:"pop"`
	Endpoint string  `json:"endpoint"`
	Region   string  `json:"region,omitempty"`
	Allowed  bool    `json:"allowed"`
	Pinned   bool    `json:"pinned"`
	Selected bool    `json:"selected"`
	Port     int     `json:"port,omitempty"`
	RTTms    float64 `json:"rtt_ms,omitempty"`
	Error    string  `json:"error,omitempty"`
}

// Selection is the last selection outcome.
type Selection struct {
	Pops    []string   `json:"pops"`
	Reason  string     `json:"reason"`
	At      time.Time  `json:"at"`
	Trigger string     `json:"trigger"`
	PopsGen int64      `json:"pops_gen"`
	Source  string     `json:"pops_source"`
	Probes  []ProbeRow `json:"probes"`
}

// Status is the admin-socket snapshot.
type Status struct {
	Version          string           `json:"version"`
	ConnectorID      string           `json:"connector_id"`
	TunnelID         string           `json:"tunnel_id"`
	StartedAt        time.Time        `json:"started_at"`
	CertSerial       string           `json:"cert_serial"`
	CertDaysLeft     int              `json:"cert_days_left"`
	GlobalPauseUntil time.Time        `json:"global_pause_until,omitempty"`
	Selection        Selection        `json:"selection"`
	Pops             []SupervisorInfo `json:"pops"`
	LastErrors       []ErrorEntry     `json:"last_errors"`
}

type ErrorEntry struct {
	Time    time.Time `json:"time"`
	Level   string    `json:"level"`
	Message string    `json:"message"`
}

// ErrorRing records the last N warnings/errors for `status`.
type ErrorRing struct {
	mu   sync.Mutex
	buf  []ErrorEntry
	size int
}

func NewErrorRing(n int) *ErrorRing { return &ErrorRing{size: n} }

func (r *ErrorRing) Add(e ErrorEntry) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.buf = append(r.buf, e)
	if len(r.buf) > r.size {
		r.buf = r.buf[len(r.buf)-r.size:]
	}
}

func (r *ErrorRing) List() []ErrorEntry {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]ErrorEntry(nil), r.buf...)
}

// Manager runs discovery, selection and the per-PoP supervisors.
type Manager struct {
	Config      *config.Config
	Creds       *creds.Store
	Dialer      *Dialer
	Discovery   *Discovery
	Handler     http.Handler
	Metrics     *metrics.Metrics
	Log         *slog.Logger
	Version     string
	ConnectorID string
	// Exit runs once when a goaway decides the exit code, before any drain. nil: no-op.
	Exit   func(code int)
	Errors *ErrorRing
	// DrainTimeout bounds the drain of in-flight requests after ctx ends. 0: 15 s.
	DrainTimeout time.Duration
	// RevokedHint ends the "tunnel revoked" log line.
	RevokedHint string
	// OwnerTicket and OwnerAbsent go out in every hello answer (h2srv.Options).
	OwnerTicket string
	OwnerAbsent string

	initOnce   sync.Once
	ctrl       *Controller
	ctx        context.Context
	cancel     context.CancelFunc
	unsubCreds func()

	mu         sync.Mutex
	startedAt  time.Time
	sups       map[string]*Supervisor
	selection  Selection
	pops       *PopsReply
	reselectCh chan string
	selecting  bool
	pending    string
	ends       []time.Time
	lastStorm  time.Time
	lastDown   time.Time
	wg         sync.WaitGroup
	reach      *reachLatch
}

// ErrOwnerRequired: the edge runs this tunnel only with a valid owner ticket
// (goaway owner_required, or a require-mode whoami without an ok verdict).
var ErrOwnerRequired = errors.New("the edge requires a linked Hostinger account")

// ExitError carries a process exit code out of Run, and the goaway reason
// that decided it (revoked, unknown_serial or owner_required).
type ExitError struct {
	Code   int
	Reason string
}

func (e *ExitError) Error() string {
	if e.Reason == "" {
		return "exit " + strconv.Itoa(e.Code)
	}
	return "exit " + strconv.Itoa(e.Code) + " (" + e.Reason + ")"
}

// Is matches ErrOwnerRequired for an owner_required exit.
func (e *ExitError) Is(target error) bool {
	return target == ErrOwnerRequired && e.Reason == proto.ReasonOwnerRequired
}

const (
	// A home PoP with no ready session this long is replaced.
	popDownAfter = 30 * time.Second
	popDownCheck = 10 * time.Second
)

func NewConnectorID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// Init wires the controller and defaults; call it before Status is exposed.
func (m *Manager) Init() { m.initOnce.Do(m.init) }

func (m *Manager) init() {
	if m.Log == nil {
		m.Log = slog.Default()
	}
	if m.Exit == nil {
		m.Exit = func(int) {}
	}
	if m.ConnectorID == "" {
		m.ConnectorID = NewConnectorID()
	}
	if m.Errors == nil {
		m.Errors = NewErrorRing(20)
	}
	m.sups = map[string]*Supervisor{}
	m.reselectCh = make(chan string, 4)
	m.reach = newReachLatch()
	m.ctrl = &Controller{
		Log:           m.Log,
		Exit:          m.Exit,
		ReloadCreds:   m.Creds.Reload,
		CanRefresh:    m.Creds.CanRefresh,
		Refresh:       m.Creds.Refresh,
		Context:       m.runContext,
		CurrentSerial: func() string { return m.Creds.Current().Serial },
		Reselect:      func(refresh bool) { m.requestReselect(pick(refresh, "goaway_refresh", "goaway")) },
		Supervisors:   m.Supervisors,
		StopAll:       m.stopAll,
		RevokedHint:   m.RevokedHint,
	}
	m.unsubCreds = m.Creds.OnChange(func(*creds.Credentials) {
		go m.rotateSessions()
	})
}

// Close stops the credential rotation hook; a shared store outlives the Manager.
func (m *Manager) Close() {
	m.Init()
	m.unsubCreds()
}

func (m *Manager) drainTimeout() time.Duration {
	if m.DrainTimeout > 0 {
		return m.DrainTimeout
	}
	return drainBudget
}

// runContext is the Run context, or nil before Run.
func (m *Manager) runContext() context.Context {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.ctx
}

func pick(b bool, t, f string) string {
	if b {
		return t
	}
	return f
}

// Run blocks until ctx ends, then drains every session. It returns an
// *ExitError when a goaway decided the process exit code.
func (m *Manager) Run(ctx context.Context) error {
	m.Init()
	rctx, cancel := context.WithCancel(ctx)
	m.mu.Lock()
	m.ctx, m.cancel = rctx, cancel
	m.startedAt = time.Now()
	m.mu.Unlock()
	defer cancel()
	// Every return follows the end of rctx, so a refresh in flight stops soon.
	defer m.ctrl.WaitRefresh()

	bo := NewBackoff()
	for !m.reselect("startup") {
		d := bo.Next()
		m.Log.Warn("initial pop selection failed; retrying", "retry_in", d.Round(time.Millisecond).String())
		select {
		case <-time.After(d):
		case <-rctx.Done():
			return m.exitErr()
		}
	}

	tick := time.NewTicker(m.Config.PopSelection.ReprobeInterval)
	defer tick.Stop()
	health := time.NewTicker(popDownCheck)
	defer health.Stop()
	for {
		select {
		case <-rctx.Done():
			m.drainAll()
			return m.exitErr()
		case <-tick.C:
			m.requestReselect("reprobe")
		case <-health.C:
			m.checkPopsDown(time.Now())
		case trigger := <-m.reselectCh:
			m.mu.Lock()
			if m.selecting {
				m.pending = mergeTrigger(m.pending, trigger)
				m.mu.Unlock()
				continue
			}
			m.selecting = true
			m.mu.Unlock()
			m.wg.Add(1)
			go m.selectLoop(trigger)
		}
	}
}

func (m *Manager) exitErr() error {
	if m.ctrl.Exiting() {
		r := m.ctrl.ExitReason()
		return &ExitError{Code: ExitCode(r), Reason: r}
	}
	return nil
}

// selectLoop runs trigger, then any trigger that arrived meanwhile.
func (m *Manager) selectLoop(trigger string) {
	defer m.wg.Done()
	for {
		m.reselect(trigger)
		m.mu.Lock()
		trigger, m.pending = m.pending, ""
		if trigger == "" || m.ctx.Err() != nil {
			m.selecting = false
			m.mu.Unlock()
			return
		}
		m.mu.Unlock()
	}
}

func forcesRefresh(trigger string) bool { return trigger != "startup" && trigger != "reprobe" }

// mergeTrigger keeps one pending trigger; a pop-list refresh wins over a plain reprobe.
func mergeTrigger(pending, next string) string {
	if pending == "" || (!forcesRefresh(pending) && forcesRefresh(next)) {
		return next
	}
	return pending
}

// checkPopsDown asks for a reselection when a home PoP has had no ready session
// for popDownAfter outside a pause; dial errors alone never reach the storm detector.
func (m *Manager) checkPopsDown(now time.Time) {
	for _, s := range m.Supervisors() {
		if down := s.DownFor(now); down >= popDownAfter {
			m.mu.Lock()
			fire := now.Sub(m.lastDown) >= time.Minute
			if fire {
				m.lastDown = now
			}
			m.mu.Unlock()
			if fire {
				m.Log.Warn("home pop has no ready session; re-running selection", "pop", s.Pop, "down_for", down.Round(time.Second).String())
				m.requestReselect("pop_down")
			}
			return
		}
	}
}

func (m *Manager) requestReselect(trigger string) {
	select {
	case m.reselectCh <- trigger:
	default:
	}
}

// Supervisors returns the live supervisors sorted by PoP code.
func (m *Manager) Supervisors() []*Supervisor {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]*Supervisor, 0, len(m.sups))
	for _, s := range m.sups {
		out = append(out, s)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Pop < out[j].Pop })
	return out
}

// ConnectedPops lists PoPs with at least one ready session.
func (m *Manager) ConnectedPops() []string { return m.connectedPops("") }

// connectedPops also lists `with`: the PoP answering hello is not ready yet.
func (m *Manager) connectedPops(with string) []string {
	var out []string
	for _, s := range m.Supervisors() {
		if s.Pop == with || s.ReadyCount() > 0 {
			out = append(out, s.Pop)
		}
	}
	return out
}

func (m *Manager) sessionOptions(sup *Supervisor, _ int, conn *Conn) h2srv.Options {
	return h2srv.Options{
		Handler:       m.Handler,
		Version:       m.Version,
		ConnectorID:   m.ConnectorID,
		Transport:     conn.Transport,
		ConnectedPops: func() []string { return m.connectedPops(sup.Pop) },
		Log:           m.Log.With("pop", sup.Pop),
		// Never below the h2srv default: other drains (pop removal, rotation) keep their budget.
		DrainTimeout: max(m.drainTimeout(), drainBudget),
		OwnerTicket:  m.OwnerTicket,
		OwnerAbsent:  m.OwnerAbsent,
		OnHello: func(h h2srv.Hello) {
			m.noteHello(sup.Pop, h)
			if h.PopsGen != "" {
				if gen, err := strconv.ParseInt(h.PopsGen, 10, 64); err == nil {
					m.mu.Lock()
					cur := m.pops
					m.mu.Unlock()
					if cur != nil && cur.Source != "manual" && cur.Gen != gen {
						m.Log.Info("pop list generation changed", "have", cur.Gen, "edge", gen)
						m.requestReselect("pops_gen")
					}
				}
			}
		},
		OnVisible: func(v h2srv.Visible) { m.noteVisible(sup.Pop, v) },
		OnGoaway: func(g h2srv.Goaway) {
			m.ctrl.Handle(GoawayEvent{Pop: sup.Pop, Addr: conn.Addr, Serial: conn.CertSerial, Goaway: g})
		},
		OnRequestEnd: func(result string) {
			if m.Metrics != nil {
				m.Metrics.Requests.WithLabelValues(result).Inc()
			}
		},
	}
}

// onSessionEnd detects reconnect storms: >= 3 established sessions lost in
// 60 s -> refresh the pop list. Dial failures, local closes and goaways (the
// controller already handles each reason) do not count.
func (m *Manager) onSessionEnd(_ string, reason string) {
	if strings.HasPrefix(reason, "goaway_") || reason == "dial_error" || reason == reasonClosedLocal {
		return
	}
	now := time.Now()
	m.mu.Lock()
	keep := m.ends[:0]
	for _, t := range m.ends {
		if now.Sub(t) < stormWindow {
			keep = append(keep, t)
		}
	}
	m.ends = append(keep, now)
	storm := len(m.ends) >= stormThreshold && now.Sub(m.lastStorm) > 5*time.Minute
	if storm {
		m.lastStorm = now
	}
	m.mu.Unlock()
	if storm {
		m.Log.Warn("reconnect storm; refreshing the pop list")
		m.requestReselect("storm")
	}
}

// reselect runs discovery + probing + selection and applies the result.
// It returns false when no PoP could be selected.
func (m *Manager) reselect(trigger string) bool {
	ctx, cancel := context.WithTimeout(m.ctx, 2*time.Minute)
	defer cancel()
	pops, err := m.Discovery.Pops(ctx, forcesRefresh(trigger))
	if err != nil {
		if m.Config.PopSelection.Mode == config.ModeManual {
			m.Log.Warn("discovery unavailable; using manual pop list", "err", err)
			pops = ManualPops(m.Config.PopSelection.Pops, m.Creds.Current().BootstrapEndpoints)
		} else {
			m.Log.Error("discovery failed", "err", err)
			return false
		}
	}
	m.mu.Lock()
	m.pops = pops
	m.mu.Unlock()

	var allowed []Pop
	for _, p := range pops.Pops {
		if p.Allowed {
			allowed = append(allowed, p)
		}
	}
	probes := m.Discovery.Probe(ctx, allowed)

	cands := make([]Candidate, 0, len(pops.Pops))
	rows := make([]ProbeRow, 0, len(pops.Pops))
	probeBy := map[string]ProbeResult{}
	for _, pr := range probes {
		probeBy[pr.Pop.Code] = pr
	}
	for _, p := range pops.Pops {
		c := Candidate{Code: p.Code, Region: p.Region, Allowed: p.Allowed, Pinned: pops.IsPinned(p.Code)}
		row := ProbeRow{Pop: p.Code, Endpoint: p.Endpoint, Region: p.Region, Allowed: p.Allowed, Pinned: c.Pinned}
		if pr, ok := probeBy[p.Code]; ok {
			if pr.Err != nil {
				c.Err = pr.Err.Error()
				row.Error = c.Err
			} else {
				c.RTT = pr.RTT
				row.RTTms, row.Port = ms(pr.RTT), pr.Port
			}
		} else if !p.Allowed {
			c.Err = "not allowed"
		}
		cands = append(cands, c)
		rows = append(rows, row)
	}

	var current []Current
	now := time.Now()
	for _, s := range m.Supervisors() {
		current = append(current, Current{Code: s.Pop, Since: s.Since, Ready: s.ReadyCount() > 0, Down: s.DownFor(now) >= popDownAfter})
	}
	pol := Policy{
		HomePops:   m.Config.HomePops,
		MinGainPct: m.Config.PopSelection.SwitchMinGainPct,
		MinGainMs:  m.Config.PopSelection.SwitchMinGainMs,
		Holddown:   m.Config.PopSelection.Holddown,
		Mode:       m.Config.PopSelection.Mode,
		ManualPops: m.Config.PopSelection.Pops,
	}
	selected, reason := Select(cands, current, pol, now)
	for i := range rows {
		for _, s := range selected {
			if rows[i].Pop == s {
				rows[i].Selected = true
			}
		}
	}
	sel := Selection{Pops: selected, Reason: reason, At: time.Now(), Trigger: trigger, PopsGen: pops.Gen, Source: pops.Source, Probes: rows}
	m.mu.Lock()
	m.selection = sel
	m.mu.Unlock()
	m.Log.Info("pop selection", "trigger", trigger, "selected", selected, "reason", reason)
	if len(selected) == 0 {
		m.Log.Error("no pop selectable", "reason", reason)
		return false
	}
	m.apply(ctx, pops, selected)
	return true
}

// apply starts supervisors for new PoPs first, waits for them to become ready,
// then drains the PoPs that left the selection, one ready old PoP per ready new PoP.
func (m *Manager) apply(ctx context.Context, pops *PopsReply, selected []string) {
	want := map[string]bool{}
	for _, s := range selected {
		want[s] = true
	}
	var added []*Supervisor
	m.mu.Lock()
	for _, code := range selected {
		if _, ok := m.sups[code]; ok {
			continue
		}
		p := popByCode(pops, code)
		s := NewSupervisor(code, p.Endpoint, p.Region, pops.IsPinned(code), m.Config.SessionsPerPop,
			m.Dialer, m.sessionOptions, m.Metrics, m.Log, m.onSessionEnd)
		m.sups[code] = s
		added = append(added, s)
	}
	var removed []*Supervisor
	for code, s := range m.sups {
		if !want[code] {
			removed = append(removed, s)
		}
	}
	m.mu.Unlock()
	for _, s := range added {
		if until := m.ctrl.GlobalPauseUntil(); until.After(time.Now()) {
			s.Pause(time.Until(until), proto.ReasonDisabled)
		}
		s.Start(m.ctx)
		m.Log.Info("pop added", "pop", s.Pop, "endpoint", s.Endpoint)
	}
	if len(removed) == 0 {
		return
	}
	drain := removed
	if len(added) > 0 {
		wctx, cancel := context.WithTimeout(ctx, readyWaitBudget)
		var notReady []*Supervisor
		for _, s := range added {
			if !s.WaitReady(wctx) {
				notReady = append(notReady, s)
			}
		}
		cancel()
		if len(notReady) > 0 {
			m.Log.Warn("new pops did not become ready; keeping current ones in their place", "not_ready", popCodes(notReady))
			m.mu.Lock()
			for _, s := range notReady {
				delete(m.sups, s.Pop)
			}
			m.mu.Unlock()
			for _, s := range notReady {
				s.Stop()
			}
		}
		readyAdded := len(added) - len(notReady)
		if readyAdded == 0 {
			return
		}
		drain = drainable(removed, readyAdded)
	}
	m.mu.Lock()
	for _, s := range drain {
		delete(m.sups, s.Pop)
	}
	m.mu.Unlock()
	var wg sync.WaitGroup
	for _, s := range drain {
		m.Log.Info("pop removed; draining", "pop", s.Pop)
		wg.Add(1)
		go func(s *Supervisor) { defer wg.Done(); s.Drain(ctx) }(s)
	}
	wg.Wait()
}

// drainable keeps the ready PoP count: removed PoPs without a ready session
// always go, ready ones only up to readyAdded.
func drainable(removed []*Supervisor, readyAdded int) []*Supervisor {
	var out, ready []*Supervisor
	for _, s := range removed {
		if s.ReadyCount() == 0 {
			out = append(out, s)
		} else {
			ready = append(ready, s)
		}
	}
	sort.Slice(ready, func(i, j int) bool {
		ri, rj := ready[i].ReadyCount(), ready[j].ReadyCount()
		if ri != rj {
			return ri < rj
		}
		return ready[i].Pop < ready[j].Pop
	})
	return append(out, ready[:min(readyAdded, len(ready))]...)
}

func popCodes(sups []*Supervisor) []string {
	out := make([]string, 0, len(sups))
	for _, s := range sups {
		out = append(out, s.Pop)
	}
	return out
}

func popByCode(r *PopsReply, code string) Pop {
	for _, p := range r.Pops {
		if p.Code == code {
			return p
		}
	}
	return Pop{Code: code, Endpoint: code}
}

func (m *Manager) rotateSessions() {
	m.mu.Lock()
	ctx := m.ctx
	m.mu.Unlock()
	if ctx == nil {
		return
	}
	m.Log.Info("credentials changed; rotating sessions")
	for _, s := range m.Supervisors() {
		if !s.Rotate(ctx) {
			if ctx.Err() == nil {
				m.Log.Warn("rotation stopped: a re-dialed session did not become ready; the other sessions keep the old certificate",
					"pop", s.Pop)
			}
			return
		}
	}
}

// stopAll drains briefly (the goaway 204 must reach the wire before the
// connection closes), then stops every supervisor and ends Run.
func (m *Manager) stopAll() {
	m.mu.Lock()
	sups := make([]*Supervisor, 0, len(m.sups))
	for _, s := range m.sups {
		sups = append(sups, s)
	}
	m.sups = map[string]*Supervisor{}
	m.mu.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), exitDrainBudget)
	defer cancel()
	var wg sync.WaitGroup
	for _, s := range sups {
		wg.Add(1)
		go func(s *Supervisor) { defer wg.Done(); s.Drain(ctx) }(s)
	}
	wg.Wait()
	m.mu.Lock()
	stop := m.cancel
	m.mu.Unlock()
	if stop != nil {
		stop()
	}
}

func (m *Manager) drainAll() {
	ctx, cancel := context.WithTimeout(context.Background(), m.drainTimeout())
	defer cancel()
	var wg sync.WaitGroup
	for _, s := range m.Supervisors() {
		wg.Add(1)
		go func(s *Supervisor) { defer wg.Done(); s.Drain(ctx) }(s)
	}
	wg.Wait()
	m.wg.Wait()
}

// Status builds the admin snapshot.
func (m *Manager) Status() Status {
	m.mu.Lock()
	sel, startedAt := m.selection, m.startedAt
	m.mu.Unlock()
	c := m.Creds.Current()
	st := Status{
		Version: m.Version, ConnectorID: m.ConnectorID, TunnelID: m.Config.TunnelID, StartedAt: startedAt,
		CertSerial: c.Serial, CertDaysLeft: int(c.DaysLeft(time.Now())), Selection: sel,
		LastErrors: m.Errors.List(),
	}
	if m.ctrl != nil {
		if t := m.ctrl.GlobalPauseUntil(); t.After(time.Now()) {
			st.GlobalPauseUntil = t
		}
	}
	for _, s := range m.Supervisors() {
		st.Pops = append(st.Pops, s.Snapshot())
	}
	return st
}
