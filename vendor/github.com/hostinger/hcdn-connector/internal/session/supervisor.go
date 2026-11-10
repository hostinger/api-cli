package session

import (
	"context"
	"errors"
	"log/slog"
	"net/netip"
	"sort"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/hostinger/hcdn-connector/internal/h2srv"
	"github.com/hostinger/hcdn-connector/internal/metrics"
	"github.com/hostinger/hcdn-connector/internal/proto"
)

type State string

const (
	StateConnecting  State = "connecting"
	StateHandshaking State = "handshaking"
	StateReady       State = "ready"
	StateDraining    State = "draining"
	StatePaused      State = "paused"
	stateNone        State = ""
)

var allStates = []State{StateConnecting, StateHandshaking, StateReady, StateDraining, StatePaused}

const (
	dialBudget      = 30 * time.Second
	readyWaitBudget = 60 * time.Second
	readIdleTimeout = 60 * time.Second
	// A drain sooner than this after dial redials with backoff, not at once (drain loop guard).
	drainMinLife = 5 * time.Second

	// Session end reason when the connector closed or drained the session itself.
	reasonClosedLocal = "closed_local"
)

// OptionsFunc builds the h2srv options for a new session of this supervisor.
type OptionsFunc func(sup *Supervisor, slot int, conn *Conn) h2srv.Options

type SessionInfo struct {
	Slot        int       `json:"slot"`
	State       State     `json:"state"`
	Addr        string    `json:"addr,omitempty"`
	Port        int       `json:"port,omitempty"`
	Transport   string    `json:"transport,omitempty"`
	EdgeNode    string    `json:"edge_node,omitempty"`
	Since       time.Time `json:"since,omitempty"`
	Inflight    int64     `json:"inflight"`
	MaxInflight int64     `json:"max_inflight,omitempty"`
	Backoff     int       `json:"backoff_attempt,omitempty"`
}

type SupervisorInfo struct {
	Pop         string        `json:"pop"`
	Endpoint    string        `json:"endpoint"`
	Region      string        `json:"region,omitempty"`
	Pinned      bool          `json:"pinned"`
	Since       time.Time     `json:"since"`
	PausedUntil time.Time     `json:"paused_until,omitempty"`
	PauseReason string        `json:"pause_reason,omitempty"`
	Target      int           `json:"target"`
	Ready       int           `json:"ready"`
	Draining    int           `json:"draining,omitempty"` // replaced sessions still finishing streams
	Sessions    []SessionInfo `json:"sessions"`
	LastError   string        `json:"last_error,omitempty"`
	LastErrorAt time.Time     `json:"last_error_at,omitempty"`
}

type slot struct {
	idx      int
	state    State
	sess     *h2srv.Session
	conn     *Conn
	since    time.Time
	edgeNode string
	backoff  *Backoff
	avoid    netip.Addr
	intent   netip.Addr // first address of an in-progress dial (spread)
	// drainedAt: the last session on the primary port ended on a `drain`; the next dial
	// retries the primary through the edge's restart gap instead of falling back to 443.
	drainedAt time.Time
}

// Supervisor keeps `target` sessions to one PoP.
type Supervisor struct {
	Pop      string
	Endpoint string
	Region   string
	Pinned   bool
	Since    time.Time

	target  int
	dialer  *Dialer
	options OptionsFunc
	metrics *metrics.Metrics
	log     *slog.Logger
	onEnd   func(pop, reason string)
	now     func() time.Time

	readyWait time.Duration // per-slot budget in Rotate

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	mu             sync.Mutex
	slots          []*slot
	draining       []*h2srv.Session // detached on `drain`; the edge closes them
	pausedUntil    time.Time
	pauseReason    string
	badAddrs       map[netip.Addr]time.Time
	lastRetryAfter time.Duration
	stopping       bool
	lastErr        string
	lastErrAt      time.Time
	readyCh        chan struct{}
	downSince      time.Time // zero while a slot is ready
}

func NewSupervisor(pop, endpoint, region string, pinned bool, target int, d *Dialer, opts OptionsFunc,
	m *metrics.Metrics, log *slog.Logger, onEnd func(pop, reason string)) *Supervisor {
	if log == nil {
		log = slog.Default()
	}
	if onEnd == nil {
		onEnd = func(string, string) {}
	}
	s := &Supervisor{
		Pop: pop, Endpoint: endpoint, Region: region, Pinned: pinned, Since: time.Now(),
		target: target, dialer: d, options: opts, metrics: m, log: log.With("pop", pop), onEnd: onEnd,
		now:       time.Now,
		readyWait: readyWaitBudget,
		badAddrs:  map[netip.Addr]time.Time{},
		readyCh:   make(chan struct{}),
	}
	for i := 0; i < target; i++ {
		s.slots = append(s.slots, &slot{idx: i, backoff: NewBackoff()})
	}
	return s
}

// Start launches one goroutine per session slot; it is a no-op after Stop or Drain.
func (s *Supervisor) Start(ctx context.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.stopping {
		return
	}
	s.ctx, s.cancel = context.WithCancel(ctx)
	s.downSince = s.now()
	for _, sl := range s.slots {
		s.wg.Add(1)
		go s.runSlot(sl)
	}
}

// Stop closes every session immediately and waits for the slots to exit.
func (s *Supervisor) Stop() {
	cancel := s.markStopping()
	if cancel != nil {
		cancel()
	}
	s.CloseSessions()
	s.wg.Wait()
	s.clearGauge()
}

// Drain sends GOAWAY on every session, waits for in-flight streams, then stops.
func (s *Supervisor) Drain(ctx context.Context) {
	cancel := s.markStopping()
	s.DrainSessions(ctx)
	if cancel != nil {
		cancel()
	}
	s.wg.Wait()
	s.clearGauge()
}

func (s *Supervisor) markStopping() context.CancelFunc {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stopping = true
	return s.cancel
}

// DrainSessions sends GOAWAY on every session and waits for in-flight streams
// (bounded by ctx); slots redial per pause/backoff afterwards.
func (s *Supervisor) DrainSessions(ctx context.Context) {
	s.mu.Lock()
	sessions := s.liveSessions()
	for _, sl := range s.slots {
		if sl.sess != nil {
			sl.state = StateDraining
		}
	}
	s.mu.Unlock()
	s.updateGauge()
	var wg sync.WaitGroup
	for _, sess := range sessions {
		wg.Add(1)
		go func(sess *h2srv.Session) { defer wg.Done(); sess.Drain(ctx) }(sess)
	}
	wg.Wait()
}

// CloseSessions terminates all sessions now; slots redial per pause/backoff.
func (s *Supervisor) CloseSessions() {
	for _, sess := range s.liveSessionsLocked() {
		sess.Close()
	}
}

func (s *Supervisor) liveSessionsLocked() []*h2srv.Session {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.liveSessions()
}

func (s *Supervisor) liveSessions() []*h2srv.Session {
	var out []*h2srv.Session
	for _, sl := range s.slots {
		if sl.sess != nil {
			out = append(out, sl.sess)
		}
	}
	return append(out, s.draining...)
}

// Pause blocks new dials until now+d (never shortens an existing pause).
func (s *Supervisor) Pause(d time.Duration, reason string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	d = min(d, proto.MaxPause)
	until := s.now().Add(d)
	if until.After(s.pausedUntil) {
		s.pausedUntil = until
		s.pauseReason = reason
	}
	s.log.Warn("pop paused", "until", until.UTC().Format(time.RFC3339), "reason", reason)
}

func (s *Supervisor) PausedUntil() time.Time {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.pausedUntil
}

// Paused reports whether dials are blocked right now.
func (s *Supervisor) Paused() bool { return s.PausedUntil().After(s.now()) }

// MarkAddrBad excludes addr from dials for d (auth_unavailable, section 6.4).
func (s *Supervisor) MarkAddrBad(addr netip.Addr, d time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if d <= 0 {
		d = proto.DefaultRetryAfterAuthUnavailable
	}
	d = min(d, proto.MaxRetryAfterShort)
	s.badAddrs[addr] = s.now().Add(d)
	s.lastRetryAfter = d
}

// BadAddrs returns the currently excluded addresses.
func (s *Supervisor) BadAddrs() []netip.Addr {
	m := s.hardExcludes()
	out := make([]netip.Addr, 0, len(m))
	for a := range m {
		out = append(out, a)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Less(out[j]) })
	return out
}

func (s *Supervisor) hardExcludes() map[netip.Addr]bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := s.now()
	out := map[netip.Addr]bool{}
	for a, until := range s.badAddrs {
		if until.After(now) {
			out[a] = true
		} else {
			delete(s.badAddrs, a)
		}
	}
	return out
}

// OnNoAddress is what a slot does when every address is excluded: pause the
// PoP for the last retry-after and let the other home PoP carry traffic.
func (s *Supervisor) OnNoAddress() {
	s.mu.Lock()
	d := s.lastRetryAfter
	s.badAddrs = map[netip.Addr]time.Time{}
	s.mu.Unlock()
	if d <= 0 {
		d = proto.DefaultRetryAfterAuthUnavailable
	}
	s.Pause(d, "auth_unavailable on every address")
}

// pickAddrs excludes the other slots' live and intended addresses and reserves
// the first choice under s.mu, so concurrent dials spread across nodes.
func (s *Supervisor) pickAddrs(sl *slot, resolved []netip.Addr, hard map[netip.Addr]bool) []netip.Addr {
	s.mu.Lock()
	defer s.mu.Unlock()
	soft := map[netip.Addr]bool{}
	for _, o := range s.slots {
		if o == sl {
			continue
		}
		if o.conn != nil {
			soft[o.conn.Addr] = true
		}
		if o.intent.IsValid() {
			soft[o.intent] = true
		}
	}
	if sl.avoid.IsValid() {
		soft[sl.avoid] = true
	}
	addrs := spread(filterAddrs(resolved, hard, soft), sl.idx)
	sl.intent = netip.Addr{}
	if len(addrs) > 0 {
		sl.intent = addrs[0]
	}
	return addrs
}

func (s *Supervisor) isStopping() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stopping
}

func (s *Supervisor) ReadyCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := 0
	for _, sl := range s.slots {
		if sl.state == StateReady {
			n++
		}
	}
	return n
}

// WaitReady blocks until a session is ready or ctx ends.
func (s *Supervisor) WaitReady(ctx context.Context) bool {
	return s.waitUntil(ctx, func() bool {
		for _, sl := range s.slots {
			if sl.state == StateReady {
				return true
			}
		}
		return false
	})
}

// waitUntil re-evaluates cond (under s.mu) on every ready transition.
func (s *Supervisor) waitUntil(ctx context.Context, cond func() bool) bool {
	for {
		s.mu.Lock()
		ch := s.readyCh
		ok := cond()
		s.mu.Unlock()
		if ok {
			return true
		}
		select {
		case <-ch:
		case <-ctx.Done():
			return false
		}
	}
}

func (s *Supervisor) Snapshot() SupervisorInfo {
	s.mu.Lock()
	defer s.mu.Unlock()
	info := SupervisorInfo{
		Pop: s.Pop, Endpoint: s.Endpoint, Region: s.Region, Pinned: s.Pinned, Since: s.Since,
		PauseReason: s.pauseReason, Target: s.target, Draining: len(s.draining),
		LastError: s.lastErr, LastErrorAt: s.lastErrAt,
	}
	if s.pausedUntil.After(s.now()) {
		info.PausedUntil = s.pausedUntil
	}
	for _, sl := range s.slots {
		si := SessionInfo{Slot: sl.idx, State: sl.state, Since: sl.since, EdgeNode: sl.edgeNode, Backoff: sl.backoff.Attempt()}
		if sl.conn != nil {
			si.Addr, si.Port, si.Transport = sl.conn.Addr.String(), sl.conn.Port, sl.conn.Transport
		}
		if sl.sess != nil {
			si.Inflight, si.MaxInflight = sl.sess.Inflight(), sl.sess.MaxInflight()
		}
		if sl.state == StateReady {
			info.Ready++
		}
		info.Sessions = append(info.Sessions, si)
	}
	return info
}

func (s *Supervisor) setState(sl *slot, st State) {
	s.mu.Lock()
	s.setStateLocked(sl, st)
	s.mu.Unlock()
	s.updateGauge()
}

func (s *Supervisor) setStateLocked(sl *slot, st State) {
	sl.state = st
	if st == StateReady {
		close(s.readyCh)
		s.readyCh = make(chan struct{})
	}
	ready := false
	for _, o := range s.slots {
		ready = ready || o.state == StateReady
	}
	if ready {
		s.downSince = time.Time{}
	} else if s.downSince.IsZero() {
		s.downSince = s.now()
	}
}

// DownSince is when the last ready session ended (or Start); zero while one is ready.
func (s *Supervisor) DownSince() time.Time {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.downSince
}

// DownFor is how long the PoP has had no ready session, counted from the end
// of the last pause; 0 while a session is ready or the PoP is paused.
func (s *Supervisor) DownFor(now time.Time) time.Duration {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.downSince.IsZero() || s.pausedUntil.After(now) {
		return 0
	}
	return now.Sub(maxTime(s.downSince, s.pausedUntil))
}

func maxTime(a, b time.Time) time.Time {
	if b.After(a) {
		return b
	}
	return a
}

func (s *Supervisor) updateGauge() {
	if s.metrics == nil {
		return
	}
	counts := map[State]int{}
	s.mu.Lock()
	for _, sl := range s.slots {
		counts[sl.state]++
	}
	counts[StateDraining] += len(s.draining)
	s.mu.Unlock()
	for _, st := range allStates {
		s.metrics.Sessions.WithLabelValues(s.Pop, string(st)).Set(float64(counts[st]))
	}
}

func (s *Supervisor) clearGauge() {
	if s.metrics != nil {
		s.metrics.Sessions.DeletePartialMatch(prometheus.Labels{"pop": s.Pop})
	}
}

func (s *Supervisor) noteErr(err error) {
	s.mu.Lock()
	s.lastErr, s.lastErrAt = err.Error(), s.now()
	s.mu.Unlock()
}

func (s *Supervisor) countReconnect(reason string) {
	if s.metrics != nil {
		s.metrics.Reconnects.WithLabelValues(reason).Inc()
	}
	s.onEnd(s.Pop, reason)
}

func (s *Supervisor) sleep(d time.Duration) bool {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return true
	case <-s.ctx.Done():
		return false
	}
}

func (s *Supervisor) waitUnpaused(sl *slot) bool {
	for {
		until := s.PausedUntil()
		d := until.Sub(s.now())
		if d <= 0 {
			return s.ctx.Err() == nil
		}
		s.setState(sl, StatePaused)
		if !s.sleep(min(d, 5*time.Second)) {
			return false
		}
	}
}

func (s *Supervisor) runSlot(sl *slot) {
	defer s.wg.Done()
	log := s.log.With("slot", sl.idx)
	for {
		if !s.waitUnpaused(sl) || s.isStopping() {
			return
		}
		hard := s.hardExcludes()
		s.setState(sl, StateConnecting)
		dctx, cancel := context.WithTimeout(s.ctx, dialBudget)
		resolved, err := s.dialer.Resolve(dctx, s.Endpoint)
		var conn *Conn
		if err == nil {
			addrs := s.pickAddrs(sl, resolved, hard)
			if window := sl.drainedAt.Add(DrainRetryWindow).Sub(s.now()); !sl.drainedAt.IsZero() && window > 0 {
				conn, err = s.dialer.DialAfterDrain(dctx, s.Endpoint, addrs, resolved, window)
			} else {
				conn, err = s.dialer.DialAddrs(dctx, s.Endpoint, addrs, resolved)
			}
		}
		cancel()
		sl.avoid, sl.drainedAt = netip.Addr{}, time.Time{}
		if err != nil {
			s.mu.Lock()
			sl.intent = netip.Addr{}
			s.mu.Unlock()
			if s.ctx.Err() != nil {
				return
			}
			if errors.Is(err, ErrNoAddress) && len(hard) > 0 {
				s.OnNoAddress()
				continue
			}
			s.noteErr(err)
			s.countReconnect("dial_error")
			d := sl.backoff.Next()
			log.Warn("dial failed", "err", err, "retry_in", d.Round(time.Millisecond).String())
			if !s.sleep(d) {
				return
			}
			continue
		}
		s.serveSession(sl, conn, log)
		if s.ctx.Err() != nil || s.isStopping() {
			return
		}
	}
}

func (s *Supervisor) serveSession(sl *slot, conn *Conn, log *slog.Logger) {
	opts := s.options(s, sl.idx, conn)
	userHello, userGoaway, userVisible := opts.OnHello, opts.OnGoaway, opts.OnVisible
	var sess *h2srv.Session
	opts.OnHello = func(h h2srv.Hello) {
		s.mu.Lock()
		// ServeConn does not wait for handlers: a hello can finish after this session left the slot.
		if sl.sess != sess {
			s.mu.Unlock()
			return
		}
		sl.edgeNode = h.EdgeNode
		s.setStateLocked(sl, StateReady)
		s.mu.Unlock()
		s.updateGauge()
		// edge_node is a header value, but from an mTLS-authenticated edge and needed for support.
		log.Info("session ready", "edge_node", h.EdgeNode, "addr", conn.Addr.String(), "port", conn.Port,
			"transport", conn.Transport, "max_inflight", h.MaxInflight)
		if userHello != nil {
			userHello(h)
		}
	}
	// A visible signal counts only from the session in the slot: a replaced
	// or draining session is leaving the edge's presence. The slot may not
	// be ready yet: the edge can send visible as soon as it reads the hello
	// answer, before OnHello runs.
	opts.OnVisible = func(v h2srv.Visible) {
		s.mu.Lock()
		live := sl.sess == sess
		s.mu.Unlock()
		if !live {
			return
		}
		log.Debug("edge confirmed presence", "scope", v.Scope, "edge_node", v.EdgeNode)
		if userVisible != nil {
			userVisible(v)
		}
	}
	drainCh := make(chan struct{})
	var drainOnce sync.Once
	opts.OnGoaway = func(g h2srv.Goaway) {
		log.Info("goaway received", "reason", g.Reason, "retry_after", g.RetryAfter.String())
		if g.Reason == proto.ReasonDrain {
			drainOnce.Do(func() { close(drainCh) })
		}
		if userGoaway != nil {
			userGoaway(g)
		}
	}
	if opts.Transport == "" {
		opts.Transport = conn.Transport
	}
	if opts.ReadIdle == 0 {
		opts.ReadIdle = readIdleTimeout
	}
	sess = h2srv.New(conn, opts)
	s.mu.Lock()
	// Drain snapshots sessions after it sets stopping, and after hello ctx no
	// longer ends a session: a dial that finishes now would outlive Drain.
	if s.stopping {
		sl.intent = netip.Addr{}
		s.mu.Unlock()
		_ = conn.Close()
		return
	}
	sl.sess, sl.conn, sl.since, sl.intent = sess, conn, s.now(), netip.Addr{}
	s.mu.Unlock()
	s.setState(sl, StateHandshaking)
	if conn.Transport == proto.TransportFallback {
		log.Info("session on 443 fallback", "hcdn-transport", proto.TransportFallback, "addr", conn.Addr.String())
	}
	started := s.now()
	done := make(chan error, 1)
	go func() { done <- sess.Serve(s.ctx) }()
	var err error
	select {
	case err = <-done:
	case <-drainCh:
		// The edge finishes in-flight streams and closes; this slot redials now
		// on another address (section 6.4 drain).
		lived := s.now().Sub(started)
		s.detach(sl, sess, started, done, log)
		sl.avoid = conn.Addr
		if conn.Transport == proto.TransportDirect {
			sl.drainedAt = s.now()
		}
		if lived > StableSession {
			sl.backoff.Reset()
		}
		if lived < drainMinLife {
			s.sleep(sl.backoff.Next())
		}
		return
	}
	lived := s.now().Sub(started)
	g := sess.LastGoaway()
	reason := endReason(sess, err, g)
	s.countReconnect(reason)
	s.mu.Lock()
	sl.sess, sl.conn, sl.edgeNode = nil, nil, ""
	s.mu.Unlock()
	s.setState(sl, stateNone)
	logEnd(log, reason, lived, err)
	if lived > StableSession {
		sl.backoff.Reset()
	}
	if s.ctx.Err() != nil || s.isStopping() {
		return
	}
	skip := false
	if g != nil {
		switch g.Reason {
		case proto.ReasonDrain:
			skip = lived >= drainMinLife
			sl.avoid = conn.Addr
			if conn.Transport == proto.TransportDirect {
				sl.drainedAt = s.now()
			}
		case proto.ReasonAuthUnavailable:
			s.MarkAddrBad(conn.Addr, g.RetryAfter)
			skip = true
		case proto.ReasonCapExceeded:
			s.Pause(g.RetryAfter, proto.ReasonCapExceeded)
		}
	}
	if err != nil {
		s.noteErr(err)
	}
	if !skip {
		s.sleep(sl.backoff.Next())
	}
}

func logEnd(log *slog.Logger, reason string, lived time.Duration, err error) {
	if err != nil {
		log.Warn("session ended", "reason", reason, "lived", lived.Round(time.Second).String(), "err", err.Error())
		return
	}
	log.Info("session ended", "reason", reason, "lived", lived.Round(time.Second).String())
}

// detach moves a draining session out of its slot and reaps it in the background.
func (s *Supervisor) detach(sl *slot, sess *h2srv.Session, started time.Time, done <-chan error, log *slog.Logger) {
	s.mu.Lock()
	sl.sess, sl.conn, sl.edgeNode = nil, nil, ""
	s.draining = append(s.draining, sess)
	s.mu.Unlock()
	s.setState(sl, stateNone)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		err := <-done
		lived := s.now().Sub(started)
		s.countReconnect(endReason(sess, err, sess.LastGoaway()))
		s.mu.Lock()
		for i, d := range s.draining {
			if d == sess {
				s.draining = append(s.draining[:i], s.draining[i+1:]...)
				break
			}
		}
		s.mu.Unlock()
		s.updateGauge()
		logEnd(log, "goaway_"+proto.ReasonDrain, lived, err)
	}()
}

func endReason(sess *h2srv.Session, err error, g *h2srv.Goaway) string {
	switch {
	case g != nil:
		return "goaway_" + g.Reason
	case sess.ClosedLocally():
		return reasonClosedLocal
	case errors.Is(err, h2srv.ErrHelloTimeout):
		return "hello_timeout"
	case errors.Is(err, h2srv.ErrBadHello):
		return "bad_hello"
	case sess.Hello() == nil:
		return "closed_before_hello"
	}
	return "closed"
}

// Rotate drains and re-dials sessions one at a time (credential change). It
// returns false and keeps the remaining old sessions (valid through the grace)
// when a slot is not ready within readyWait, e.g. the edge rejects the new certificate.
func (s *Supervisor) Rotate(ctx context.Context) bool {
	for _, sl := range s.slots {
		s.mu.Lock()
		sess := sl.sess
		s.mu.Unlock()
		if sess == nil {
			continue
		}
		if len(s.slots) > 1 && !s.waitFor(ctx, func() bool { return s.readyExcept(sl) }) {
			return false
		}
		s.setState(sl, StateDraining)
		sess.Drain(ctx)
		if !s.waitFor(ctx, func() bool { return sl.state == StateReady && sl.sess != sess }) {
			return false
		}
	}
	return true
}

func (s *Supervisor) waitFor(ctx context.Context, cond func() bool) bool {
	wctx, cancel := context.WithTimeout(ctx, s.readyWait)
	defer cancel()
	return s.waitUntil(wctx, cond)
}

// readyExcept reports whether a slot other than sl is ready (caller holds s.mu).
func (s *Supervisor) readyExcept(sl *slot) bool {
	for _, o := range s.slots {
		if o != sl && o.state == StateReady {
			return true
		}
	}
	return false
}
