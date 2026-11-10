package session

import (
	"context"
	"log/slog"
	"net/netip"
	"sync"
	"time"

	"github.com/hostinger/hcdn-connector/internal/h2srv"
	"github.com/hostinger/hcdn-connector/internal/proto"
)

// GoawayEvent is a goaway received on one session.
type GoawayEvent struct {
	Pop    string
	Addr   netip.Addr
	Serial string // client certificate serial the session dialed with
	Goaway h2srv.Goaway
}

// Controller applies the process-level goaway reasons (docs/protocol.md, section 6);
// every dependency is injectable so the state machine is unit-testable.
type Controller struct {
	Log           *slog.Logger
	Now           func() time.Time
	Exit          func(code int)
	ReloadCreds   func() (changed bool, err error)
	CanRefresh    func() bool                     // a memory store with a RefreshFunc
	Refresh       func(ctx context.Context) error // network I/O: runs off the Handle goroutine
	Context       func() context.Context          // bounds Refresh; nil or a nil result: Background
	CurrentSerial func() string
	Reselect      func(refreshPops bool)
	Supervisors   func() []*Supervisor
	StopAll       func()
	After         func(d time.Duration, f func()) // time.AfterFunc
	RevokedHint   string

	once            sync.Once
	mu              sync.Mutex
	serialMu        sync.Mutex // one unknown_serial decision at a time: every session gets the goaway
	refreshing      bool       // under serialMu
	refreshedSerial string     // under serialMu: the serial the last Refresh installed
	refreshedAt     time.Time  // under serialMu: start of the last Refresh
	stopped         bool       // under serialMu: WaitRefresh ran, so no refreshWG.Add may follow
	refreshWG       sync.WaitGroup
	globalPauseTill time.Time
	exiting         bool
	exitReason      string // under mu: the goaway reason that decided the exit
}

const (
	refreshTimeout = 30 * time.Second
	// A rejection of the refreshed certificate within this window exits 78;
	// a later one (the edge forgot an old certificate) gets a new refresh.
	refreshHoldoff = 10 * time.Minute
)

// Action names what the controller did, for logs and tests.
type Action string

const (
	ActionNone         Action = "none"
	ActionPausePop     Action = "pause_pop"
	ActionExcludeAddr  Action = "exclude_addr"
	ActionReselect     Action = "reselect"
	ActionPauseAll     Action = "pause_all"
	ActionExit         Action = "exit"
	ActionReloadCreds  Action = "reload_creds"
	ActionBackoff      Action = "backoff"
	ActionAlreadyExits Action = "exiting"
)

func (c *Controller) defaults() {
	if c.Log == nil {
		c.Log = slog.Default()
	}
	if c.Now == nil {
		c.Now = time.Now
	}
	if c.After == nil {
		c.After = func(d time.Duration, f func()) { time.AfterFunc(d, f) }
	}
	if c.Reselect == nil {
		c.Reselect = func(bool) {}
	}
	if c.StopAll == nil {
		c.StopAll = func() {}
	}
	if c.Supervisors == nil {
		c.Supervisors = func() []*Supervisor { return nil }
	}
	if c.Exit == nil {
		c.Exit = func(int) {}
	}
}

// GlobalPauseUntil is set by a `disabled` goaway.
func (c *Controller) GlobalPauseUntil() time.Time {
	c.once.Do(c.defaults)
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.globalPauseTill
}

func (c *Controller) supervisor(pop string) *Supervisor {
	for _, s := range c.Supervisors() {
		if s.Pop == pop {
			return s
		}
	}
	return nil
}

// Handle reacts to one goaway. It never blocks on network or session I/O.
// Goaway callbacks run concurrently, so defaults are applied exactly once.
func (c *Controller) Handle(ev GoawayEvent) Action {
	c.once.Do(c.defaults)
	if c.isExiting() {
		return ActionAlreadyExits
	}
	g := ev.Goaway
	log := c.Log.With("pop", ev.Pop, "reason", g.Reason)
	switch g.Reason {
	case proto.ReasonDrain:
		return ActionNone
	case proto.ReasonProtocolError:
		log.Warn("edge reported a protocol error; the session will back off")
		return ActionBackoff
	case proto.ReasonCapExceeded:
		if s := c.supervisor(ev.Pop); s != nil {
			s.Pause(g.RetryAfter, proto.ReasonCapExceeded)
		}
		return ActionPausePop
	case proto.ReasonAuthUnavailable:
		if s := c.supervisor(ev.Pop); s != nil && ev.Addr.IsValid() {
			s.MarkAddrBad(ev.Addr, g.RetryAfter)
		}
		return ActionExcludeAddr
	case proto.ReasonReselect:
		log.Info("edge asked for reselection; refreshing the pop list")
		c.Reselect(true)
		return ActionReselect
	case proto.ReasonDisabled:
		d := g.RetryAfter
		if d <= 0 {
			d = proto.DefaultRetryAfterDisabled
		}
		d = min(d, proto.MaxPause)
		until := c.Now().Add(d)
		c.mu.Lock()
		c.globalPauseTill = until
		c.mu.Unlock()
		log.Warn("tunnel disabled in hPanel; pausing all sessions", "resume_at", until.UTC().Format(time.RFC3339))
		sups := c.Supervisors()
		for _, s := range sups {
			s.Pause(d, proto.ReasonDisabled)
		}
		go drainSessions(sups, drainBudget)
		c.After(d, func() {
			c.Log.Info("disabled pause over; re-running pop selection")
			c.Reselect(true)
		})
		return ActionPauseAll
	case proto.ReasonRevoked:
		msg := "tunnel revoked; stopping"
		if c.RevokedHint != "" {
			msg += ". " + c.RevokedHint
		}
		log.Error(msg)
		return c.exit(g.Reason)
	case proto.ReasonOwnerRequired:
		log.Error("the edge requires a linked Hostinger account for this tunnel; stopping")
		return c.exit(g.Reason)
	case proto.ReasonUnknownSerial:
		c.serialMu.Lock()
		defer c.serialMu.Unlock()
		if c.isExiting() {
			return ActionAlreadyExits
		}
		// The slots re-dial with the old certificate until the refresh ends.
		if c.refreshing {
			return ActionReloadCreds
		}
		if ev.Serial != "" && c.CurrentSerial != nil && ev.Serial != c.CurrentSerial() {
			log.Info("edge rejected an older certificate; the session re-dials with the current one", "serial", ev.Serial)
			return ActionNone
		}
		if c.CanRefresh != nil && c.Refresh != nil && c.CanRefresh() {
			base := c.baseContext()
			if base.Err() != nil || c.stopped {
				return ActionNone
			}
			// Same rule as the file path: one new certificate per rejection, then exit 78.
			if c.refreshedSerial != "" && c.CurrentSerial != nil && c.CurrentSerial() == c.refreshedSerial &&
				c.Now().Sub(c.refreshedAt) < refreshHoldoff {
				log.Error("certificate serial unknown to the edge; the refreshed certificate was rejected too; stopping",
					"serial", c.refreshedSerial)
				return c.exit(g.Reason)
			}
			c.refreshing = true
			c.refreshedAt = c.Now()
			log.Warn("certificate serial unknown to the edge; refreshing credentials, sessions will re-dial")
			c.refreshWG.Add(1)
			go c.refresh(base, log)
			return ActionReloadCreds
		}
		if c.ReloadCreds == nil {
			log.Error("certificate serial unknown to the edge; no credentials reload available; stopping")
			return c.exit(g.Reason)
		}
		changed, err := c.ReloadCreds()
		// The file watcher can load the new file between the serial check and this reload.
		if err == nil && !changed && ev.Serial != "" && c.CurrentSerial != nil && ev.Serial != c.CurrentSerial() {
			changed = true
		}
		if err != nil || !changed {
			log.Error("certificate serial unknown to the edge; reload did not yield a new certificate; stopping", "err", err)
			return c.exit(g.Reason)
		}
		log.Warn("certificate serial unknown to the edge; reloaded new credentials, sessions will re-dial")
		return ActionReloadCreds
	}
	return ActionNone
}

func (c *Controller) baseContext() context.Context {
	if c.Context != nil {
		if b := c.Context(); b != nil {
			return b
		}
	}
	return context.Background()
}

// WaitRefresh waits for a running refresh; it returns soon after the Context ends.
// A later unknown_serial goaway starts no refresh.
func (c *Controller) WaitRefresh() {
	c.serialMu.Lock()
	c.stopped = true
	c.serialMu.Unlock()
	c.refreshWG.Wait()
}

// refresh runs Refresh; on success the store's OnChange rotates the sessions.
func (c *Controller) refresh(base context.Context, log *slog.Logger) {
	defer c.refreshWG.Done()
	ctx, cancel := context.WithTimeout(base, refreshTimeout)
	defer cancel()
	done := make(chan error, 1)
	go func() { done <- c.Refresh(ctx) }()
	var err error
	select {
	case err = <-done:
	case <-ctx.Done():
		// A RefreshFunc that ignores ctx must not hold the refresh state forever.
		select {
		case err = <-done:
		default:
			err = ctx.Err()
		}
	}
	if err != nil && base.Err() == nil {
		log.Error("certificate serial unknown to the edge; credentials refresh failed; stopping", "err", err)
		c.exit(proto.ReasonUnknownSerial)
	} else if err == nil {
		log.Info("credentials refreshed after unknown_serial")
	}
	c.serialMu.Lock()
	c.refreshing = false
	if err == nil && c.CurrentSerial != nil {
		c.refreshedSerial = c.CurrentSerial()
	}
	c.serialMu.Unlock()
}

// exit calls Exit before StopAll, so a caller backstop arms before the drain.
func (c *Controller) exit(reason string) Action {
	c.mu.Lock()
	if c.exiting {
		c.mu.Unlock()
		return ActionAlreadyExits
	}
	c.exiting = true
	c.exitReason = reason
	c.mu.Unlock()
	c.Exit(ExitCode(reason))
	c.StopAll()
	return ActionExit
}

// ExitCode is the process exit code for a goaway reason that stops the
// connector: 77 for owner_required (EX_NOPERM), 78 otherwise.
func ExitCode(reason string) int {
	if reason == proto.ReasonOwnerRequired {
		return proto.ExitOwnerRequired
	}
	return proto.ExitRevoked
}

// ExitReason is the goaway reason that decided the exit, or "" before one.
func (c *Controller) ExitReason() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.exitReason
}

// Exiting reports whether an exit decision was made (set before StopAll runs).
func (c *Controller) Exiting() bool {
	c.once.Do(c.defaults)
	return c.isExiting()
}

func (c *Controller) isExiting() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.exiting
}

const (
	drainBudget     = 15 * time.Second
	exitDrainBudget = time.Second
)

// drainSessions drains every supervisor's sessions in parallel, bounded by d.
// Sessions stay alive until in-flight streams finish or d elapses.
func drainSessions(sups []*Supervisor, d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	var wg sync.WaitGroup
	for _, s := range sups {
		wg.Add(1)
		go func(s *Supervisor) { defer wg.Done(); s.DrainSessions(ctx) }(s)
	}
	wg.Wait()
}
