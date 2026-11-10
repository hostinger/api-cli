package connector

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/hostinger/hcdn-connector/internal/proto"
	"github.com/hostinger/hcdn-connector/pkg/config"
	"github.com/hostinger/hcdn-connector/pkg/creds"
	"github.com/hostinger/hcdn-connector/pkg/localguard"
)

// DefaultQuickBootstrap returns the bootstrap endpoints of a quick tunnel:
// tunnel.hostinger.io, a GeoIP name that answers with a nearby PoP.
func DefaultQuickBootstrap() []string { return []string{"tunnel.hostinger.io"} }

// PauseCapExceeded is the QuickPause reason of goaway cap_exceeded: too many
// quick tunnels from this network.
const PauseCapExceeded = proto.ReasonCapExceeded

var (
	// ErrEdgeUnreachable: the whoami dial failed on every port, or timed out
	// (20 s). Outbound TCP to the edge is probably blocked.
	ErrEdgeUnreachable = errors.New("connector: the edge is unreachable")
	// ErrQuickRefused: the edge answered but refused the quick tunnel, for
	// example a handshake that rejects the self-signed leaf, or a failed
	// whoami with the owner ticket in require mode.
	ErrQuickRefused = errors.New("connector: the edge refused the quick tunnel")
	// ErrQuickNotOffered: whoami answered without class quick or without a
	// hostname (an edge older than tunneld 0.3.0).
	ErrQuickNotOffered = errors.New("connector: the edge does not offer quick tunnels")
	// ErrTakenDown: goaway revoked; the tunnel was taken down (*ExitError Code 78).
	ErrTakenDown = errors.New("connector: the quick tunnel was taken down")
	// ErrCertRejected: goaway unknown_serial; the edge no longer accepts the
	// tunnel certificate (*ExitError Code 78). A new run gets a new URL.
	ErrCertRejected = errors.New("connector: the edge no longer accepts the tunnel certificate")
)

// QuickError is a RunQuick failure with a known cause. errors.Is matches
// Kind and errors.As reaches Err, for example the *ExitError of a goaway.
type QuickError struct {
	// Kind is ErrEdgeUnreachable, ErrQuickRefused, ErrQuickNotOffered,
	// ErrOwnerRequired, ErrTakenDown or ErrCertRejected.
	Kind error
	// Err is the cause: the whoami or LinkOwner error, the class the edge
	// reported, or the *ExitError of the goaway.
	Err error
	// Owner is the outcome of LinkOwner when Kind is ErrOwnerRequired, for
	// example OtherEnvironment for the message.
	Owner OwnerLink
}

func (e *QuickError) Error() string {
	if e.Err == nil {
		return e.Kind.Error()
	}
	if msg := e.Err.Error(); strings.HasPrefix(msg, e.Kind.Error()) {
		return msg
	}
	return e.Kind.Error() + ": " + e.Err.Error()
}

func (e *QuickError) Unwrap() []error {
	if e.Err == nil {
		return []error{e.Kind}
	}
	return []error{e.Kind, e.Err}
}

// QuickTunnel is the tunnel that RunQuick passes to OnReady and OnReachable.
type QuickTunnel struct {
	TunnelID  string
	Hostname  string    // the random public host name, from whoami
	URL       string    // https:// and Hostname
	ExpiresAt time.Time // NotAfter of the quick certificate; RunQuick stops 1 min before
	Owner     OwnerLink // the outcome of LinkOwner
	// Reachable: visitors can reach the tunnel. False in an OnReady after
	// ReadyTimeout (or a negative ReadyTimeout); OnReachable follows.
	Reachable bool
	// Scope says who confirmed it (see WaitReachable): ScopePop, ScopeNode
	// or ScopeSession; "" while not reachable.
	Scope string
	// Pop is the edge PoP of the session that confirmed it; "" while not reachable.
	Pop string
}

// DefaultQuickReadyTimeout is the default of QuickOptions.ReadyTimeout.
const DefaultQuickReadyTimeout = 10 * time.Second

// QuickPause is a pause that RunQuick passes to OnPause.
type QuickPause struct {
	Reason  string    // PauseCapExceeded
	RetryAt time.Time // the connector redials after this time
}

// QuickOptions configure RunQuick. Only URL is required.
type QuickOptions struct {
	URL          string   // the local service, for example http://localhost:8080 (required)
	HostHeader   string   // "": the visitor's Host; dev servers such as Vite refuse unknown hosts
	Bootstrap    []string // nil: DefaultQuickBootstrap()
	Ports        []int    // nil: 7844 only, as `hcdn-connector quick`
	ServerCAFile string   // "": the system roots
	// Version goes into hcdn-connector-version of every hello and the
	// User-Agent hcdn-connector/<Version>; "": the module version.
	Version string
	Logger  *slog.Logger // nil: discard
	// Owner gets the owner ticket when the edge asks for one (see
	// LinkOwner). nil: the run has no way to get one (unsupported).
	Owner OwnerTicketFunc
	// ReadyTimeout bounds the wait for a reachable tunnel before OnReady.
	// 0: DefaultQuickReadyTimeout. Negative: OnReady runs before the first
	// session, with Reachable false.
	ReadyTimeout time.Duration
	// VisibleWait bounds the wait for the edge's visible signal after a
	// session is ready on an edge that advertises it (Options.VisibleWait).
	// 0: DefaultVisibleWait.
	VisibleWait time.Duration
	// OnReady runs once, from the RunQuick goroutine, when the tunnel is
	// reachable or after ReadyTimeout; RunQuick keeps running the sessions
	// meanwhile. It does not run when the run ends first, for example on a
	// goaway revoked or when ctx ends.
	OnReady func(QuickTunnel)
	// OnReachable runs once, from the RunQuick goroutine, when the tunnel
	// becomes reachable after an OnReady with Reachable false.
	OnReachable func(QuickTunnel)
	// OnPause runs once per cap_exceeded pause (one call when several PoPs
	// pause until the same time), from another goroutine. It may run before
	// OnReady. No call follows the return of RunQuick.
	OnPause func(QuickPause)
}

// quickStopAt: tunneld aborts quick sessions at NotAfter; stop cleanly first.
// A var so tests can shorten it.
var quickStopAt = func(notAfter time.Time) time.Time { return notAfter.Add(-time.Minute) }

// quickWhoamiTimeout bounds the first whoami and LinkOwner.
const quickWhoamiTimeout = 20 * time.Second

// RunQuick runs one anonymous quick tunnel: a key in memory, a random host
// name, at most 24 h, no file, no admin socket and no metrics. It calls
// whoami (20 s for whoami and LinkOwner together), links the owner and runs
// up to 3 home PoPs until ctx ends or 1 min before the certificate expires.
// It calls OnReady once the tunnel is reachable (Connector.WaitReachable),
// or after ReadyTimeout with Reachable false and then OnReachable.
// RunQuick writes nothing but log records to Logger.
//
// It returns expired true and a nil error when the tunnel reached its stop
// time, and false with a nil error when ctx ended. A *QuickError tells the
// known failures apart (errors.Is with ErrEdgeUnreachable, ErrQuickRefused,
// ErrQuickNotOffered, ErrOwnerRequired, ErrTakenDown or ErrCertRejected);
// another goaway exit returns the *ExitError as is. A URL that reaches a
// sensitive port on this machine (localguard.CheckContext) returns a
// *localguard.BlockedError before any edge call, and every dial to the
// service is checked again (localguard.Control). Any other error is a bad
// option, such as an invalid URL or ServerCAFile.
func RunQuick(ctx context.Context, o QuickOptions) (expired bool, err error) {
	if o.URL == "" {
		return false, errors.New("connector: QuickOptions.URL is required")
	}
	if err := localguard.CheckContext(ctx, o.URL); err != nil {
		return false, err
	}
	log := o.Logger
	if log == nil {
		log = slog.New(slog.DiscardHandler)
	}
	bootstrap := o.Bootstrap
	if len(bootstrap) == 0 {
		bootstrap = DefaultQuickBootstrap()
	}
	ports := o.Ports
	if len(ports) == 0 {
		ports = []int{proto.PortDirect}
	}
	c, err := creds.NewQuick(bootstrap, time.Now())
	if err != nil {
		return false, err
	}
	store, err := creds.NewMemoryStore(c, log, nil)
	if err != nil {
		return false, err
	}
	store.SetWarnDays(0)

	wctx, wcancel := context.WithTimeout(ctx, quickWhoamiTimeout)
	defer wcancel()
	wopts := WhoamiOptions{Credentials: store, Ports: ports, ServerCAFile: o.ServerCAFile, Version: o.Version, Logger: log}
	w, err := Whoami(wctx, wopts)
	switch {
	case err != nil && ctx.Err() != nil:
		return false, nil
	case err != nil && isDialFailure(err):
		return false, &QuickError{Kind: ErrEdgeUnreachable, Err: err}
	case err != nil:
		return false, &QuickError{Kind: ErrQuickRefused, Err: err}
	case w.Class != ClassQuick || w.Hostname == "":
		return false, &QuickError{Kind: ErrQuickNotOffered, Err: fmt.Errorf("class %q", w.Class)}
	}
	link, err := LinkOwner(wctx, wopts, w, o.Owner)
	wcancel()
	linked := time.Now()
	switch {
	case err != nil && ctx.Err() != nil:
		return false, nil
	case errors.Is(err, ErrOwnerRequired):
		return false, &QuickError{Kind: ErrOwnerRequired, Err: err, Owner: link}
	case err != nil:
		return false, &QuickError{Kind: ErrQuickRefused, Err: err}
	}

	cfg := config.Default()
	cfg.TunnelID = c.TunnelID
	rule := config.NewIngressRule(w.Hostname, o.URL)
	rule.OriginRequest.HTTPHostHeader = o.HostHeader
	cfg.AddRule(rule)
	// The selector caps this at the PoPs the edge allows for quick tunnels.
	cfg.HomePops = 3
	cfg.AdminSocket = ""
	noMetrics := ""
	cfg.MetricsAddr = &noMetrics
	cfg.Ports = ports
	cfg.ServerCAFile = o.ServerCAFile
	conn, err := New(Options{
		Config: cfg, Credentials: store, Version: o.Version, Log: log,
		Hints: &Hints{Revoked: "This tunnel was taken down."}, Owner: link, VisibleWait: o.VisibleWait,
	})
	if err != nil {
		return false, err
	}
	defer conn.Close()

	notAfter := c.Leaf.NotAfter
	tunnel := QuickTunnel{TunnelID: c.TunnelID, Hostname: w.Hostname, URL: "https://" + w.Hostname,
		ExpiresAt: notAfter, Owner: link}
	rctx, rcancel := context.WithDeadline(ctx, quickStopAt(notAfter))
	defer rcancel()
	stopWatch := func() {}
	if o.OnPause != nil {
		stopWatch = watchCapExceeded(conn, o.OnPause)
	}
	err = runAnnounced(rctx, conn, tunnel, o, log, linked)
	stopWatch()
	var ee *ExitError
	switch {
	case errors.As(err, &ee):
		switch ee.Reason {
		case proto.ReasonRevoked:
			return false, &QuickError{Kind: ErrTakenDown, Err: ee}
		case proto.ReasonOwnerRequired:
			return false, &QuickError{Kind: ErrOwnerRequired, Err: ee, Owner: link}
		case proto.ReasonUnknownSerial:
			return false, &QuickError{Kind: ErrCertRejected, Err: ee}
		}
		return false, err
	case err != nil:
		return false, err
	}
	return ctx.Err() == nil && errors.Is(rctx.Err(), context.DeadlineExceeded), nil
}

// runAnnounced runs conn until rctx ends or a goaway stops it, and calls
// OnReady (and OnReachable) on this goroutine as the tunnel becomes reachable.
func runAnnounced(rctx context.Context, conn *Connector, tunnel QuickTunnel, o QuickOptions, log *slog.Logger, since time.Time) error {
	timeoutAfter := o.ReadyTimeout
	if timeoutAfter == 0 {
		timeoutAfter = DefaultQuickReadyTimeout
	}
	announced := false
	announce := func(q QuickTunnel) {
		switch {
		case !announced && o.OnReady != nil:
			o.OnReady(q)
		case announced && q.Reachable && o.OnReachable != nil:
			o.OnReachable(q)
		}
		announced = true
	}
	if timeoutAfter < 0 {
		announce(tunnel)
	}

	runDone := make(chan error, 1)
	go func() { runDone <- conn.Run(rctx) }()
	wctx, wcancel := context.WithCancel(rctx)
	defer wcancel()
	reached := make(chan Reach, 1)
	go func() {
		if r, err := conn.WaitReachable(wctx); err == nil {
			reached <- r
		}
	}()
	var timeout <-chan time.Time
	if timeoutAfter > 0 {
		t := time.NewTimer(timeoutAfter)
		defer t.Stop()
		timeout = t.C
	}
	for {
		select {
		case r := <-reached:
			reached = nil
			log.Info("tunnel reachable", "scope", r.Scope, "pop", r.Pop, "edge_node", r.EdgeNode,
				"waited_ms", time.Since(since).Milliseconds())
			q := tunnel
			q.Reachable, q.Scope, q.Pop = true, r.Scope, r.Pop
			announce(q)
		case <-timeout:
			timeout = nil
			if !announced {
				log.Warn("tunnel not reachable yet; announcing the URL", "waited_ms", time.Since(since).Milliseconds())
				announce(tunnel)
			}
		case err := <-runDone:
			return err
		}
	}
}

// watchCapExceeded calls fn once per cap_exceeded pause; stop waits for the
// last call.
func watchCapExceeded(c *Connector, fn func(QuickPause)) (stop func()) {
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		t := time.NewTicker(time.Second)
		defer t.Stop()
		var shown time.Time
		for {
			select {
			case <-done:
				return
			case <-t.C:
			}
			for _, p := range c.Status().Pops {
				if p.PauseReason == proto.ReasonCapExceeded && p.PausedUntil.After(shown) {
					shown = p.PausedUntil
					fn(QuickPause{Reason: PauseCapExceeded, RetryAt: p.PausedUntil})
				}
			}
		}
	}()
	return func() { close(done); wg.Wait() }
}

// isDialFailure: a TCP dial error, or a timeout (a firewall that drops SYN).
func isDialFailure(err error) bool {
	var op *net.OpError
	if errors.As(err, &op) && op.Op == "dial" {
		return true
	}
	return errors.Is(err, context.DeadlineExceeded) || errors.Is(err, os.ErrDeadlineExceeded)
}
