// Package h2srv serves reverse HTTP/2 on a dialed connection: the edge is
// the h2 client, the connector the server. It owns the hello, goaway and
// visible control routes and the per-session request checks.
package h2srv

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hostinger/hcdn-connector/internal/ingress"
	"github.com/hostinger/hcdn-connector/internal/proto"
	"github.com/hostinger/hcdn-connector/internal/termsafe"
)

// Hello is what the edge announced on the first stream.
type Hello struct {
	EdgeNode     string
	EdgePop      string
	CertDaysLeft int
	MaxInflight  int
	MaxStreams   int
	PopsGen      string
	// EdgeFeatures lists hcdn-edge-features, lower case (nil: an edge
	// without the header).
	EdgeFeatures []string
	ReceivedAt   time.Time
}

// HasEdgeFeature reports whether the edge advertised name in hcdn-edge-features.
func (h Hello) HasEdgeFeature(name string) bool {
	for _, f := range h.EdgeFeatures {
		if f == name {
			return true
		}
	}
	return false
}

// Visible is a decoded visible control request: the edge has the tunnel in
// its presence, on every node of its PoP (proto.VisibleScopePop) or on its
// own node (proto.VisibleScopeNode).
type Visible struct {
	Scope    string
	EdgeNode string // from the hello of this session
	At       time.Time
}

// Goaway is a decoded goaway control request.
type Goaway struct {
	Reason     string
	RetryAfter time.Duration
}

// Options wires a Session to the rest of the process.
type Options struct {
	Handler       http.Handler // visitor requests
	Version       string
	ConnectorID   string
	Transport     string // proto.TransportDirect | proto.TransportFallback
	ConnectedPops func() []string
	OnHello       func(Hello)
	OnGoaway      func(Goaway)
	// OnVisible runs once per session, for the first valid visible request,
	// on the handler goroutine after the 204.
	OnVisible      func(Visible)
	OnRequestEnd   func(result string)
	Log            *slog.Logger
	HelloTimeout   time.Duration
	ReadIdle       time.Duration
	DrainTimeout   time.Duration
	MaxInflightDef int
	// OwnerTicket goes out in every hello answer as hcdn-owner-ticket. Without
	// one, OwnerAbsent goes out as hcdn-owner-absent ("": not_requested).
	OwnerTicket string
	OwnerAbsent string
}

// Session is one reverse-h2 connection to an edge node.
type Session struct {
	opts     Options
	conn     net.Conn
	rc       *readErrConn
	hs       *http.Server
	h2       h2server
	helloCh  chan struct{}
	helloOnc sync.Once
	hello    atomic.Pointer[Hello]
	inflight atomic.Int64
	maxInfl  atomic.Int64
	closed   atomic.Bool
	local    atomic.Bool
	closeErr atomic.Pointer[error]
	done     chan struct{}
	mu       sync.Mutex
	goaway   *Goaway
	visible  atomic.Bool
}

var ErrHelloTimeout = errors.New("hello not received in time")
var ErrBadHello = errors.New("hello rejected")

func New(conn net.Conn, opts Options) *Session {
	if opts.HelloTimeout == 0 {
		opts.HelloTimeout = proto.HelloTimeout
	}
	if opts.DrainTimeout == 0 {
		opts.DrainTimeout = 15 * time.Second
	}
	if opts.MaxInflightDef == 0 {
		opts.MaxInflightDef = proto.DefaultMaxInflight
	}
	if opts.Log == nil {
		opts.Log = slog.Default()
	}
	if opts.ConnectedPops == nil {
		opts.ConnectedPops = func() []string { return nil }
	}
	if opts.OwnerTicket == "" && opts.OwnerAbsent == "" {
		opts.OwnerAbsent = proto.OwnerAbsentNotRequested
	}
	wrapped, rc := wrapConn(conn)
	s := &Session{
		opts:    opts,
		conn:    wrapped,
		rc:      rc,
		helloCh: make(chan struct{}),
		done:    make(chan struct{}),
	}
	s.maxInfl.Store(int64(opts.MaxInflightDef))
	s.hs = &http.Server{
		Handler:        s,
		MaxHeaderBytes: proto.MaxHeaderBytes,
		IdleTimeout:    0,
		ErrorLog:       slog.NewLogLogger(opts.Log.Handler(), slog.LevelWarn),
		HTTP2: &http.HTTP2Config{
			MaxConcurrentStreams:          proto.MaxConcurrentStreams,
			MaxReadFrameSize:              proto.MaxFrameSize,
			SendPingTimeout:               opts.ReadIdle,
			PingTimeout:                   15 * time.Second,
			MaxReceiveBufferPerConnection: 4 << 20,
			MaxReceiveBufferPerStream:     1 << 20,
		},
	}
	s.h2 = newH2Server(s.hs)
	return s
}

// Serve runs the h2 server until the connection closes (the edge sends the
// preface). ctx bounds only the pre-hello wait; after hello, Drain/Close own the lifetime.
func (s *Session) Serve(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		t := time.NewTimer(s.opts.HelloTimeout)
		defer t.Stop()
		select {
		case <-s.helloCh:
		case <-t.C:
			s.fail(ErrHelloTimeout)
		case <-ctx.Done():
			if s.hello.Load() == nil {
				s.fail(ctx.Err())
			}
		case <-s.done:
		}
	}()
	s.h2.serveConn(s.conn, s.hs, s)
	close(s.done)
	s.closed.Store(true)
	if p := s.closeErr.Load(); p != nil {
		return *p
	}
	if err := s.rc.Err(); err != nil && s.hello.Load() == nil {
		return fmt.Errorf("closed before hello: %w", err)
	}
	return nil
}

func (s *Session) fail(err error) {
	e := err
	s.closeErr.CompareAndSwap(nil, &e)
	s.closed.Store(true)
	_ = s.conn.Close()
}

// Close terminates the connection immediately.
func (s *Session) Close() {
	s.local.Store(true)
	s.fail(nil)
}

// Drain sends GOAWAY, waits for in-flight streams, then closes.
func (s *Session) Drain(ctx context.Context) {
	s.local.Store(true)
	ctx, cancel := context.WithTimeout(ctx, s.opts.DrainTimeout)
	defer cancel()
	_ = s.hs.Shutdown(ctx)
	select {
	case <-s.done:
	case <-ctx.Done():
		_ = s.conn.Close()
	}
}

// ClosedLocally reports that this side called Close or Drain.
func (s *Session) ClosedLocally() bool { return s.local.Load() }

func (s *Session) Hello() *Hello         { return s.hello.Load() }
func (s *Session) Ready() bool           { return s.hello.Load() != nil && !s.closed.Load() }
func (s *Session) Inflight() int64       { return s.inflight.Load() }
func (s *Session) MaxInflight() int64    { return s.maxInfl.Load() }
func (s *Session) Done() <-chan struct{} { return s.done }
func (s *Session) Err() error {
	if p := s.closeErr.Load(); p != nil {
		return *p
	}
	return nil
}
func (s *Session) LastGoaway() *Goaway {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.goaway
}

// ServeHTTP matches control routes on path first; a control path that carries
// hcdn-request-id is a relayed visitor request and gets 404, never proxied.
func (s *Session) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodConnect && isControlPath(r.URL.Path) {
		if r.Header.Get(proto.HdrRequestID) != "" {
			w.Header().Set("Content-Length", "0")
			w.WriteHeader(http.StatusNotFound)
			s.report("rejected")
			return
		}
		switch r.URL.Path {
		case proto.PathHello:
			s.serveHello(w, r)
		case proto.PathGoaway:
			s.serveGoaway(w, r)
		default:
			s.serveVisible(w, r)
		}
		return
	}
	if r.Header.Get(proto.HdrRequestID) == "" {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusBadRequest)
		s.report("rejected")
		return
	}
	n := s.inflight.Add(1)
	defer s.inflight.Add(-1)
	if n > s.maxInfl.Load() {
		ingress.WriteError(w, proto.ErrOverloaded)
		s.report(proto.ErrOverloaded)
		return
	}
	s.opts.Handler.ServeHTTP(w, r)
}

func isControlPath(p string) bool {
	return p == proto.PathHello || p == proto.PathGoaway || p == proto.PathVisible
}

func (s *Session) report(result string) {
	if s.opts.OnRequestEnd != nil {
		s.opts.OnRequestEnd(result)
	}
}

func (s *Session) serveHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get(proto.HdrProtocol) != proto.Version {
		s.rejectHello(w, "unsupported hcdn-protocol")
		return
	}
	// These end up in logs and status output; never log the value itself.
	for _, k := range []string{proto.HdrEdgeNode, proto.HdrEdgePop, proto.HdrPopsGen, proto.HdrEdgeFeatures} {
		for _, v := range r.Header.Values(k) {
			if !termsafe.Printable(v) {
				s.rejectHello(w, "non-printable "+strings.ToLower(k))
				return
			}
		}
	}
	h := Hello{
		EdgeNode:     r.Header.Get(proto.HdrEdgeNode),
		EdgePop:      r.Header.Get(proto.HdrEdgePop),
		CertDaysLeft: atoi(r.Header.Get(proto.HdrCertDaysLeft), 0),
		MaxInflight:  atoi(r.Header.Get(proto.HdrMaxInflight), s.opts.MaxInflightDef),
		MaxStreams:   atoi(r.Header.Get(proto.HdrMaxStreams), 0),
		PopsGen:      r.Header.Get(proto.HdrPopsGen),
		EdgeFeatures: parseFeatures(r.Header.Values(proto.HdrEdgeFeatures)),
		ReceivedAt:   time.Now(),
	}
	if h.MaxInflight <= 0 {
		h.MaxInflight = s.opts.MaxInflightDef
	}
	s.maxInfl.Store(int64(h.MaxInflight))
	s.hello.Store(&h)
	s.helloOnc.Do(func() { close(s.helloCh) })

	out := w.Header()
	out.Set(proto.HdrConnVersion, s.opts.Version)
	out.Set(proto.HdrConnID, s.opts.ConnectorID)
	out.Set(proto.HdrFeatures, proto.Features)
	if pops := s.opts.ConnectedPops(); len(pops) > 0 {
		out.Set(proto.HdrConnectedPops, strings.Join(pops, ","))
	}
	out.Set(proto.HdrTransport, s.opts.Transport)
	if s.opts.OwnerTicket != "" {
		out.Set(proto.HdrOwnerTicket, s.opts.OwnerTicket)
	} else {
		out.Set(proto.HdrOwnerAbsent, s.opts.OwnerAbsent)
	}
	out.Set("Content-Length", "0")
	w.WriteHeader(http.StatusOK)
	if s.opts.OnHello != nil {
		s.opts.OnHello(h)
	}
}

// rejectHello answers 400 and closes the session once the answer is out.
func (s *Session) rejectHello(w http.ResponseWriter, why string) {
	w.Header().Set("Content-Length", "0")
	w.WriteHeader(http.StatusBadRequest)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
	s.opts.Log.Warn("hello rejected: " + why)
	time.AfterFunc(100*time.Millisecond, func() { s.fail(ErrBadHello) })
}

func (s *Session) serveGoaway(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	g := Goaway{Reason: strings.ToLower(r.Header.Get(proto.HdrReason))}
	if !proto.KnownReasons[g.Reason] {
		s.opts.Log.Warn("goaway with unknown reason; treating as protocol_error")
		g.Reason = proto.ReasonProtocolError
	}
	var note string
	g.RetryAfter, note = parseRetryAfter(g.Reason, r.Header.Get(proto.HdrRetryAfter))
	if note != "" {
		s.opts.Log.Warn("hcdn-retry-after "+note, "reason", g.Reason, "using", g.RetryAfter.String())
	}
	s.mu.Lock()
	s.goaway = &g
	s.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
	// Off the handler goroutine: the callback may drain this very session, and
	// the control stream must end for that drain to complete.
	if s.opts.OnGoaway != nil {
		go s.opts.OnGoaway(g)
	}
}

// serveVisible takes the edge's word that the tunnel is in its presence. Only
// the first valid request of a session reaches OnVisible; a repeat gets 204
// too, so a retrying edge never sees an error.
func (s *Session) serveVisible(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	h := s.hello.Load()
	scope := strings.ToLower(strings.TrimSpace(r.Header.Get(proto.HdrVisibleScope)))
	if h == nil || (scope != proto.VisibleScopePop && scope != proto.VisibleScopeNode) {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusBadRequest)
		if h == nil {
			s.opts.Log.Warn("visible before hello; ignored")
		} else {
			s.opts.Log.Warn("visible with an unknown scope; ignored")
		}
		return
	}
	first := s.visible.CompareAndSwap(false, true)
	w.WriteHeader(http.StatusNoContent)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
	if first && s.opts.OnVisible != nil {
		s.opts.OnVisible(Visible{Scope: scope, EdgeNode: h.EdgeNode, At: time.Now()})
	}
}

// parseFeatures splits comma lists, trims and lower-cases each item, and
// drops empty ones.
func parseFeatures(values []string) []string {
	var out []string
	for _, v := range values {
		for _, f := range strings.Split(v, ",") {
			if f = strings.ToLower(strings.TrimSpace(f)); f != "" {
				out = append(out, f)
			}
		}
	}
	return out
}

// parseRetryAfter bounds an edge-supplied value: invalid or negative falls back
// to the default for the reason (proto.DefaultRetryAfter), oversized is clamped
// to proto.MaxRetryAfter.
func parseRetryAfter(reason, v string) (time.Duration, string) {
	def := proto.DefaultRetryAfter(reason)
	if v == "" {
		return def, ""
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil || n < 0 {
		return def, "invalid; using the default"
	}
	if n == 0 {
		return def, ""
	}
	if maxSecs := int64(proto.MaxRetryAfter(reason) / time.Second); n > maxSecs {
		return time.Duration(maxSecs) * time.Second, "above the ceiling; clamped"
	}
	return time.Duration(n) * time.Second, ""
}

func atoi(s string, def int) int {
	if s == "" {
		return def
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return n
}
