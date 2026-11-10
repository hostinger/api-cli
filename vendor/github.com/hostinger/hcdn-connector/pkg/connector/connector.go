// Package connector runs the hCDN Tunnel connector inside a Go program: PoP
// discovery and selection, the outbound mTLS sessions to the CDN edge and the
// reverse proxy to the configured services. The hcdn-connector binary is a
// thin wrapper over this package.
//
// The package opens no listener and installs no signal handler. The caller
// owns the process: the library never calls os.Exit.
//
// # Text from the edge
//
// Error text (QuickError, ExitError, the errors of Run and Whoami) and the
// error and reason fields of Status (ErrorEntry.Message, PopInfo.LastError,
// PopInfo.PauseReason, ProbeRow.Error) can hold text from the edge as
// received, control characters included. Escape it before you write it to a
// terminal; the hcdn-connector binary does. PoP codes, endpoints and regions,
// and the fields of WhoamiReply and WhoamiOwner, are checked to be printable.
package connector

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"runtime/debug"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hostinger/hcdn-connector/internal/fileguard"
	"github.com/hostinger/hcdn-connector/internal/ingress"
	"github.com/hostinger/hcdn-connector/internal/logging"
	"github.com/hostinger/hcdn-connector/internal/metrics"
	"github.com/hostinger/hcdn-connector/internal/proto"
	"github.com/hostinger/hcdn-connector/internal/session"
	"github.com/hostinger/hcdn-connector/pkg/config"
	"github.com/hostinger/hcdn-connector/pkg/creds"
	"github.com/hostinger/hcdn-connector/pkg/localguard"
)

type (
	// Status is the snapshot that `hcdn-connector status` prints. Fields:
	// Version, ConnectorID, TunnelID, StartedAt (zero before Run),
	// CertSerial, CertDaysLeft, GlobalPauseUntil (set while the tunnel is
	// disabled), Selection, Pops (a PopInfo per home PoP) and LastErrors.
	// Status marshals to the JSON of `hcdn-connector status --json`.
	//
	// Status and the types below are aliases of internal types: their
	// fields are documented here and on the internal package, and they are
	// part of this package's API.
	Status = session.Status
	// Selection is Status.Selection, the last PoP selection: Pops, Reason,
	// At, Trigger, PopsGen, Source and Probes.
	Selection = session.Selection
	// ProbeRow is one PoP in Selection.Probes: Pop, Endpoint, Region,
	// Allowed, Pinned, Selected, Port, RTTms and Error.
	ProbeRow = session.ProbeRow
	// PopInfo is one home PoP in Status.Pops: Pop, Endpoint, Region,
	// Pinned, Since, PausedUntil, PauseReason, Target, Ready, Draining,
	// Sessions, LastError and LastErrorAt.
	PopInfo = session.SupervisorInfo
	// SessionInfo is one session slot in PopInfo.Sessions: Slot, State,
	// Addr, Port, Transport, EdgeNode, Since, Inflight, MaxInflight and
	// Backoff.
	SessionInfo = session.SessionInfo
	// SessionState is SessionInfo.State: SessionConnecting,
	// SessionHandshaking, SessionReady, SessionDraining or SessionPaused.
	SessionState = session.State
	// ErrorEntry is one warning or error in Status.LastErrors.
	ErrorEntry = session.ErrorEntry
	// ExitError carries the exit code that a goaway decided and its Reason
	// (78: revoked, or unknown_serial when the edge rejects the certificate;
	// 77: owner_required, the edge wants a linked Hostinger account).
	// errors.Is(err, ErrOwnerRequired) matches the 77 exit.
	ExitError = session.ExitError
	// Metrics holds the Prometheus registry and collectors.
	Metrics = metrics.Metrics
	// Reach is what WaitReachable returns: Scope (ScopePop, ScopeNode or
	// ScopeSession), the Pop and EdgeNode of the session that decided it,
	// and At.
	Reach = session.Reach
)

// Reach scopes.
const (
	// ScopePop: the edge confirmed that every node of its PoP routes the tunnel.
	ScopePop = proto.VisibleScopePop
	// ScopeNode: the edge confirmed only its own node.
	ScopeNode = proto.VisibleScopeNode
	// ScopeSession: a session is ready but no edge confirmed presence (an
	// edge without the feature, or Options.VisibleWait ran out).
	ScopeSession = session.ScopeSession
)

// Session states of SessionInfo.State.
const (
	SessionConnecting  = session.StateConnecting
	SessionHandshaking = session.StateHandshaking
	SessionReady       = session.StateReady
	SessionDraining    = session.StateDraining
	SessionPaused      = session.StatePaused
)

// DefaultVisibleWait is the default of Options.VisibleWait.
const DefaultVisibleWait = session.DefaultVisibleWait

// DefaultDrainTimeout bounds the drain of in-flight requests after the Run context ends.
const DefaultDrainTimeout = 15 * time.Second

const modulePath = "github.com/hostinger/hcdn-connector"

// Hints replace the hPanel advice in operator log lines; an empty field keeps the binary text.
type Hints struct{ CertExpiry, Revoked string }

const defaultRevokedHint = "Create a new tunnel in hPanel and install its credentials."

type Options struct {
	Config      *config.Config   // New copies it, then runs Normalize and Validate on the copy
	Credentials *creds.Store     // nil: open Config.CredentialsFile; set: CredentialsFile must be "", New sets its logger
	Reload      <-chan os.Signal // nil: no signal-driven reload; the file store still polls mtime
	Version     string           // hello and build_info; "": the module version
	Commit      string
	Log         *slog.Logger // nil: format and level from Config.Log, written to LogOutput
	LogOutput   io.Writer    // nil: os.Stderr
	ConnectorID string       // "": random
	// OnExit runs once, when a goaway decides the exit code, before any drain.
	OnExit       func(code int)
	DrainTimeout time.Duration // 0: DefaultDrainTimeout
	Hints        *Hints
	// Owner is the outcome of LinkOwner. Every hello answer carries its ticket,
	// or its no-ticket reason. The zero value sends not_requested.
	Owner OwnerLink
	// AllowLocalPorts opts these ports out of the local port guard. Without
	// it, a dial from an ingress service to a port of localguard.BlockedPorts
	// on this machine (SSH, databases, debuggers, cluster agents) is refused
	// and answers service_unreachable. Set it only to publish such a port on
	// purpose. RunQuick never sets it.
	AllowLocalPorts []int
	// VisibleWait bounds how long WaitReachable waits for the edge's visible
	// signal after a session is ready on an edge that advertises it.
	// 0: DefaultVisibleWait.
	VisibleWait time.Duration
}

// Connector is one tunnel connector. Create it with New.
type Connector struct {
	cfg    *config.Config
	creds  *creds.Store
	table  *ingress.Table
	dialer *session.Dialer
	disc   *session.Discovery
	log    *slog.Logger
	mgr    *session.Manager
	m      *metrics.Metrics
	reload <-chan os.Signal
	ran    atomic.Bool

	visibleWait time.Duration

	unsetGauge func()
	closeOnce  sync.Once
}

// New builds the connector. It loads the credentials, the ingress table and
// the dialer, but it does not dial: Status works before Run.
func New(o Options) (*Connector, error) {
	if o.Config == nil {
		return nil, errors.New("connector: Options.Config is required")
	}
	cfg := o.Config.Clone()
	cfg.Normalize()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	if err := o.Owner.validate(); err != nil {
		return nil, err
	}
	for _, p := range o.AllowLocalPorts {
		if p < 1 || p > 65535 {
			return nil, fmt.Errorf("connector: Options.AllowLocalPorts: %d is not a port", p)
		}
	}
	switch {
	case o.Credentials != nil && cfg.CredentialsFile != "":
		return nil, errors.New("connector: set Options.Credentials or Config.CredentialsFile, not both")
	case o.Credentials == nil && cfg.CredentialsFile == "":
		return nil, errors.New("connector: no credentials: set Options.Credentials or Config.CredentialsFile")
	}

	ring := session.NewErrorRing(20)
	var log *slog.Logger
	if o.Log != nil {
		log = slog.New(logging.Tee(o.Log.Handler(), ring))
	} else {
		out := o.LogOutput
		if out == nil {
			out = os.Stderr
		}
		log = logging.New(out, cfg.Log, ring)
	}

	cs := o.Credentials
	if cs == nil {
		var err error
		if cs, err = creds.NewStore(cfg.CredentialsFile, cfg.TunnelID, log); err != nil {
			return nil, err
		}
	} else if cs.TunnelID() != cfg.TunnelID {
		return nil, creds.ErrTunnelIDMismatch
	}
	revokedHint := defaultRevokedHint
	if o.Hints != nil && o.Hints.Revoked != "" {
		revokedHint = o.Hints.Revoked
	}

	version := o.Version
	if version == "" {
		version = moduleVersion()
	}
	d, err := newDialer(cs, cfg.Ports, cfg.ServerCAFile, cfg.ServerSPKIPins, version, log)
	if err != nil {
		return nil, err
	}
	guard := &localguard.Guard{Allow: slices.Clone(o.AllowLocalPorts)}
	if err := guard.Probe(); err != nil {
		// Not fatal: the guard fails closed, and only listed ports are refused.
		log.Warn("local port guard cannot read the interface addresses; sensitive ports not in allow_local_ports are refused on every host", "err", err)
	}
	table, err := ingress.Build(cfg.Ingress, cfg.TunnelID, guard)
	if err != nil {
		return nil, fmt.Errorf("ingress: %w", err)
	}

	// No failure below: a caller store changes only for a connector that exists.
	if o.Credentials != nil {
		cs.SetLogger(log)
	}
	if o.Hints != nil && o.Hints.CertExpiry != "" {
		cs.SetExpiryHint(o.Hints.CertExpiry)
	}
	m := metrics.New()
	m.BuildInfo.WithLabelValues(version, o.Commit).Set(1)
	unsetGauge := cs.SetGauge(m.CertDaysLeft.Set)
	disc := &session.Discovery{Dialer: d, Bootstrap: func() []string { return cs.Current().BootstrapEndpoints }, Log: log}
	mgr := &session.Manager{
		Config: cfg, Creds: cs, Dialer: d, Discovery: disc, Handler: ingress.NewHandler(table, cfg.TunnelID, log, m),
		Metrics: m, Log: log, Version: version, ConnectorID: o.ConnectorID, Errors: ring,
		DrainTimeout: o.DrainTimeout, RevokedHint: revokedHint,
		OwnerTicket: o.Owner.Ticket, OwnerAbsent: o.Owner.Absent,
		Exit: func(code int) {
			log.Error("exiting", "code", code)
			if o.OnExit != nil {
				o.OnExit(code)
			}
		},
	}
	mgr.Init()
	return &Connector{cfg: cfg, creds: cs, table: table, dialer: d, disc: disc, log: log, mgr: mgr, m: m, reload: o.Reload,
		unsetGauge: unsetGauge, visibleWait: o.VisibleWait}, nil
}

func newDialer(cs *creds.Store, ports []int, serverCAFile string, pins []string, version string, log *slog.Logger) (*session.Dialer, error) {
	d := &session.Dialer{Creds: cs, Ports: session.NewPortMemory(ports), Version: version, Log: log}
	if serverCAFile != "" {
		pem, err := fileguard.ReadFile(serverCAFile)
		if err != nil {
			return nil, fmt.Errorf("server_ca_file: %w", err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("server_ca_file %s: no certificates found", serverCAFile)
		}
		d.RootCAs = pool
	}
	for _, p := range pins {
		b, err := config.DecodePin(p)
		if err != nil {
			return nil, err
		}
		d.Pins = append(d.Pins, b)
	}
	return d, nil
}

// WhoamiReply is the answer of GET /_hcdn/v1/whoami: TunnelID, CertSerial,
// DaysLeft, Pop, Node, and from tunneld 0.3.0 Class (ClassNamed or
// ClassQuick), Hostname (quick only) and ExpiresAt. From tunneld 0.5.0 Owner
// is set when the edge asks a quick tunnel for an owner ticket; OwnerMode()
// returns OwnerModeLog, OwnerModeRequire or "" (not asked).
type WhoamiReply = session.Whoami

const (
	ClassNamed = session.ClassNamed
	ClassQuick = session.ClassQuick
)

type WhoamiOptions struct {
	Credentials  *creds.Store // required
	Ports        []int        // nil: 7844, then 443
	ServerCAFile string
	Version      string       // "": the module version
	Logger       *slog.Logger // nil: discard
}

// Whoami calls GET /_hcdn/v1/whoami on the first bootstrap endpoint that
// answers, without a config. A quick tunnel learns its hostname this way, and
// whether the edge asks for an owner ticket (WhoamiReply.Owner, see LinkOwner).
func Whoami(ctx context.Context, o WhoamiOptions) (*WhoamiReply, error) {
	return whoami(ctx, o, "")
}

// whoami sends ownerTicket as hcdn-owner-ticket when it is not empty.
func whoami(ctx context.Context, o WhoamiOptions, ownerTicket string) (*WhoamiReply, error) {
	if o.Credentials == nil {
		return nil, errors.New("connector: WhoamiOptions.Credentials is required")
	}
	ports := o.Ports
	if len(ports) == 0 {
		ports = config.Default().Ports
	}
	version := o.Version
	if version == "" {
		version = moduleVersion()
	}
	d, err := newDialer(o.Credentials, ports, o.ServerCAFile, nil, version, o.logger())
	if err != nil {
		return nil, err
	}
	disc := &session.Discovery{Dialer: d, Bootstrap: func() []string { return o.Credentials.Current().BootstrapEndpoints }, Log: o.logger()}
	w, _, err := disc.Whoami(ctx, ownerTicket)
	return w, err
}

func (o WhoamiOptions) logger() *slog.Logger {
	if o.Logger == nil {
		return slog.New(slog.DiscardHandler)
	}
	return o.Logger
}

func moduleVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Path == modulePath && info.Main.Version != "" && info.Main.Version != "(devel)" {
			return strings.TrimPrefix(info.Main.Version, "v")
		}
		for _, dep := range info.Deps {
			if dep.Path == modulePath {
				return strings.TrimPrefix(dep.Version, "v")
			}
		}
	}
	return "devel"
}

// Run blocks until ctx ends or a goaway stops the connector. After ctx ends
// it drains in-flight requests for up to DrainTimeout; it returns at once when
// no request is in flight. A revoked tunnel or a failed credentials refresh
// returns *ExitError with Code 78, goaway owner_required Code 77. Call Run
// one time per Connector.
func (c *Connector) Run(ctx context.Context) error {
	if !c.ran.CompareAndSwap(false, true) {
		return errors.New("connector: Run called twice")
	}
	stop, watched := make(chan struct{}), make(chan struct{})
	go func() { defer close(watched); c.creds.Watch(stop, c.reload) }()
	// No reload or rotation may start after Run returns.
	defer func() { close(stop); <-watched }()
	cur := c.creds.Current()
	c.log.Info("hcdn-connector starting", "version", c.mgr.Version, "connector_id", c.mgr.ConnectorID,
		"tunnel_id", c.cfg.TunnelID, "cert_serial", cur.Serial, "cert_days_left", int(cur.DaysLeft(time.Now())),
		"home_pops", c.cfg.HomePops, "sessions_per_pop", c.cfg.SessionsPerPop, "ports", c.cfg.Ports,
		"ingress_rules", len(c.cfg.Ingress))
	err := c.mgr.Run(ctx)
	if err == nil {
		c.log.Info("hcdn-connector stopped")
	}
	return err
}

// Status returns the current snapshot. It works before Run.
func (c *Connector) Status() Status { return c.mgr.Status() }

// Ready reports at least one ready session to any PoP.
func (c *Connector) Ready() bool { return len(c.mgr.ConnectedPops()) > 0 }

// WaitReachable blocks until visitors can reach the tunnel, or ctx ends
// (ctx.Err()). It resolves on the first of:
//   - a visible signal from any edge (Scope ScopePop or ScopeNode);
//   - a ready session on an edge that does not advertise the visible
//     feature (ScopeSession);
//   - Options.VisibleWait after the first ready session on an edge that
//     advertises it, without a signal (ScopeSession, logged at warn level).
//
// Call it while Run runs; it also works before Run. After it resolved once,
// later calls return at once.
func (c *Connector) WaitReachable(ctx context.Context) (Reach, error) {
	return c.mgr.WaitReachable(ctx, c.visibleWait)
}

// Metrics returns the Prometheus collectors. Serve them with Metrics().Handler().
func (c *Connector) Metrics() *Metrics { return c.m }

// Creds returns the credentials store, for example to call Replace.
func (c *Connector) Creds() *creds.Store { return c.creds }

// Close closes the transports to the services and detaches the connector
// from the credentials store, which the caller can reuse. Call it after Run returns.
func (c *Connector) Close() {
	c.closeOnce.Do(func() {
		c.mgr.Close()
		c.unsetGauge()
		c.table.Close()
	})
}

type CheckOptions struct{ Whoami bool }

// CheckKind tells the rows of Check apart.
type CheckKind string

const (
	CheckIngress CheckKind = "ingress" // one ingress rule; Target is its service
	CheckWhoami  CheckKind = "whoami"  // GET /_hcdn/v1/whoami; Target is the answer
)

// CheckRow is one line of `hcdn-connector check`. Class is the hcdn-error
// value of a failed dial (for example service_unreachable).
type CheckRow struct {
	Kind                     CheckKind
	Name, Target, Err, Class string
	OK                       bool
}

// CatchAllName is the CheckRow name of the rule without a hostname.
const CatchAllName = "(catch-all)"

// Check dials every ingress service (TCP, and TLS for https). It returns one
// CheckIngress row per rule in config order, then one CheckWhoami row when o.Whoami is set.
func (c *Connector) Check(ctx context.Context, o CheckOptions) ([]CheckRow, error) {
	rows := make([]CheckRow, 0, len(c.table.Rules)+1)
	for _, r := range c.table.Rules {
		row := CheckRow{Kind: CheckIngress, Name: r.Hostname, OK: true}
		if row.Name == "" {
			row.Name = CatchAllName
		}
		if r.Service.Status != 0 {
			row.Target = fmt.Sprintf("%s%d", config.ServiceStatusPrefix, r.Service.Status)
			rows = append(rows, row)
			continue
		}
		row.Target = r.Service.Raw
		if err := dialService(ctx, r.Service); err != nil {
			row.OK, row.Err, row.Class = false, err.Error(), ingress.Classify(err)
		}
		rows = append(rows, row)
	}
	if o.Whoami {
		wctx, cancel := context.WithTimeout(ctx, 20*time.Second)
		w, ep, err := c.disc.Whoami(wctx, "")
		cancel()
		row := CheckRow{Kind: CheckWhoami, Name: "whoami", OK: err == nil}
		if err != nil {
			row.Err = err.Error()
		} else {
			row.Target = fmt.Sprintf("via %s: tunnel %s serial %s days_left %d pop %s node %s",
				ep, w.TunnelID, w.CertSerial, w.DaysLeft, w.Pop, w.Node)
			if w.Class == ClassQuick {
				row.Target += " class quick hostname " + w.Hostname
			}
		}
		rows = append(rows, row)
	}
	return rows, ctx.Err()
}

func dialService(ctx context.Context, svc *ingress.Service) error {
	host := svc.URL.Host
	if svc.URL.Port() == "" {
		port := "80"
		if svc.URL.Scheme == "https" {
			port = "443"
		}
		host = net.JoinHostPort(svc.URL.Hostname(), port)
	}
	ctx, cancel := context.WithTimeout(ctx, svc.Origin.ConnectTimeout)
	defer cancel()
	d := svc.Dialer()
	if svc.URL.Scheme == "https" {
		conn, err := (&tls.Dialer{NetDialer: d, Config: svc.TLS}).DialContext(ctx, "tcp", host)
		if err != nil {
			return err
		}
		return conn.Close()
	}
	conn, err := d.DialContext(ctx, "tcp", host)
	if err != nil {
		return err
	}
	return conn.Close()
}
