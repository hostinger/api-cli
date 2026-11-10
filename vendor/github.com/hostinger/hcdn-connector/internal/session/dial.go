package session

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"math/rand/v2"
	"net"
	"net/http"
	"net/netip"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hostinger/hcdn-connector/internal/ingress"
	"github.com/hostinger/hcdn-connector/internal/proto"
	"github.com/hostinger/hcdn-connector/pkg/creds"
)

var (
	ErrNoAddress = errors.New("no usable address")
	// ErrNotNetwork wraps failures that must not trigger the port fallback
	// (certificate verification, rejected upgrade, wrong ALPN, oversized
	// response header).
	ErrNotNetwork = errors.New("non-network failure")
)

// After a `drain` the primary port is retried for DrainRetryWindow in DrainRetryStep
// steps before the fallback port is tried (see DialAfterDrain).
var (
	DrainRetryWindow = 5 * time.Second
	DrainRetryStep   = 200 * time.Millisecond
)

// AddrError is a TCP or TLS failure on Failed, a subset of the endpoint's Resolved addresses.
type AddrError struct {
	Err              error
	Failed, Resolved []netip.Addr
}

func (e *AddrError) Error() string { return e.Err.Error() }
func (e *AddrError) Unwrap() error { return e.Err }

// addrErr blames the TLS peer when there is one, else every address tried.
func addrErr(err error, peer netip.Addr, tried []netip.Addr) error {
	failed := tried
	if peer.IsValid() {
		failed = []netip.Addr{peer}
	}
	return &AddrError{Err: err, Failed: failed, Resolved: tried}
}

type Resolver interface {
	LookupNetIP(ctx context.Context, network, host string) ([]netip.Addr, error)
}

// Dialer opens session and API connections to a PoP endpoint.
type Dialer struct {
	Creds              *creds.Store
	RootCAs            *x509.CertPool // nil = system roots
	Pins               [][]byte
	Ports              *PortMemory
	Resolver           Resolver
	ConnectTimeout     time.Duration
	HappyEyeballsDelay time.Duration
	FallbackPort       int
	Version            string
	Log                *slog.Logger

	once sync.Once // slots dial concurrently; defaults must be written once
}

// Conn is an established session transport.
type Conn struct {
	net.Conn
	Addr      netip.Addr
	Port      int
	Transport string
	Endpoint  string
	// CertSerial is the client certificate sent in this handshake.
	CertSerial string
}

// ConnectionState exposes the TLS state to x/net (version check, r.TLS);
// an embedded interface does not promote it.
func (c *Conn) ConnectionState() tls.ConnectionState {
	if cs, ok := c.Conn.(interface{ ConnectionState() tls.ConnectionState }); ok {
		return cs.ConnectionState()
	}
	return tls.ConnectionState{}
}

func (d *Dialer) defaults() { d.once.Do(d.setDefaults) }

func (d *Dialer) setDefaults() {
	if d.ConnectTimeout == 0 {
		d.ConnectTimeout = 5 * time.Second
	}
	if d.HappyEyeballsDelay == 0 {
		d.HappyEyeballsDelay = 250 * time.Millisecond
	}
	if d.FallbackPort == 0 {
		d.FallbackPort = proto.PortFallback
	}
	if d.Resolver == nil {
		d.Resolver = net.DefaultResolver
	}
	if d.Log == nil {
		d.Log = slog.Default()
	}
}

// IsFallbackPort reports whether port goes through the nginx 443 path.
func (d *Dialer) IsFallbackPort(port int) bool {
	d.defaults()
	return port == d.FallbackPort
}

// ALPN returns the protocol list to offer on port. On the fallback port the
// peer is nginx, which aborts the handshake unless http/1.1 or h2 is offered.
func (d *Dialer) ALPN(port int, api bool) string {
	if d.IsFallbackPort(port) {
		return proto.ALPNHTTP1
	}
	if api {
		return proto.ALPNAPI
	}
	return proto.ALPNSession
}

// tlsConfig stores the serial of the client certificate it sends in *serial (may be nil).
func (d *Dialer) tlsConfig(endpoint, alpn string, serial *string) *tls.Config {
	cfg := &tls.Config{
		MinVersion: tls.VersionTLS13,
		ServerName: endpoint,
		RootCAs:    d.RootCAs,
		NextProtos: []string{alpn},
	}
	if d.Creds != nil {
		cfg.GetClientCertificate = func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			c := d.Creds.Current()
			if c == nil {
				return nil, errors.New("credentials not loaded")
			}
			if serial != nil {
				*serial = c.Serial
			}
			return &c.Cert, nil
		}
	}
	if len(d.Pins) > 0 {
		pins := d.Pins
		cfg.VerifyConnection = func(cs tls.ConnectionState) error {
			for _, chain := range cs.VerifiedChains {
				for _, c := range chain {
					sum := sha256.Sum256(c.RawSubjectPublicKeyInfo)
					for _, p := range pins {
						if bytes.Equal(sum[:], p) {
							return nil
						}
					}
				}
			}
			return fmt.Errorf("%w: no certificate in the chain matches server_spki_pins", ErrNotNetwork)
		}
	}
	return cfg
}

// Resolve returns the endpoint's addresses, IPv6 first, shuffled within each family.
func (d *Dialer) Resolve(ctx context.Context, endpoint string) ([]netip.Addr, error) {
	d.defaults()
	if a, err := netip.ParseAddr(strings.Trim(endpoint, "[]")); err == nil {
		return []netip.Addr{a}, nil
	}
	addrs, err := d.Resolver.LookupNetIP(ctx, "ip", endpoint)
	if err != nil {
		return nil, err
	}
	var v6, v4 []netip.Addr
	for _, a := range addrs {
		a = a.Unmap()
		if a.Is6() {
			v6 = append(v6, a)
		} else {
			v4 = append(v4, a)
		}
	}
	rand.Shuffle(len(v6), func(i, j int) { v6[i], v6[j] = v6[j], v6[i] })
	rand.Shuffle(len(v4), func(i, j int) { v4[i], v4[j] = v4[j], v4[i] })
	out := make([]netip.Addr, 0, len(addrs))
	for len(v6) > 0 || len(v4) > 0 {
		if len(v6) > 0 {
			out = append(out, v6[0])
			v6 = v6[1:]
		}
		if len(v4) > 0 {
			out = append(out, v4[0])
			v4 = v4[1:]
		}
	}
	if len(out) == 0 {
		return nil, ErrNoAddress
	}
	return out, nil
}

// filterAddrs drops hard exclusions and, when possible, soft ones (spread).
func filterAddrs(addrs []netip.Addr, hard, soft map[netip.Addr]bool) []netip.Addr {
	var out, fallback []netip.Addr
	for _, a := range addrs {
		if hard[a] {
			continue
		}
		fallback = append(fallback, a)
		if !soft[a] {
			out = append(out, a)
		}
	}
	if len(out) == 0 {
		return fallback
	}
	return out
}

// spread rotates each address family by slot and re-interleaves (IPv6 still first).
func spread(addrs []netip.Addr, slot int) []netip.Addr {
	if slot <= 0 || len(addrs) < 2 {
		return addrs
	}
	var v6, v4 []netip.Addr
	for _, a := range addrs {
		if a.Is6() {
			v6 = append(v6, a)
		} else {
			v4 = append(v4, a)
		}
	}
	rot := func(s []netip.Addr) []netip.Addr {
		if len(s) < 2 {
			return s
		}
		k := slot % len(s)
		return append(append([]netip.Addr(nil), s[k:]...), s[:k]...)
	}
	v6, v4 = rot(v6), rot(v4)
	out := make([]netip.Addr, 0, len(addrs))
	for len(v6) > 0 || len(v4) > 0 {
		if len(v6) > 0 {
			out = append(out, v6[0])
			v6 = v6[1:]
		}
		if len(v4) > 0 {
			out = append(out, v4[0])
			v4 = v4[1:]
		}
	}
	return out
}

// Dial opens a session transport to endpoint: TCP with Happy Eyeballs, TLS 1.3
// with the connector certificate, and on the fallback port the hand-written
// Upgrade through nginx. Ports are tried in PortMemory order; a network failure
// on one port falls through to the next.
func (d *Dialer) Dial(ctx context.Context, endpoint string, hard, soft map[netip.Addr]bool) (*Conn, error) {
	d.defaults()
	addrs, err := d.Resolve(ctx, endpoint)
	if err != nil {
		return nil, err
	}
	return d.DialAddrs(ctx, endpoint, filterAddrs(addrs, hard, soft), addrs)
}

// DialAddrs dials addrs, an already filtered subset of resolved (the
// supervisor picks the list under its lock so concurrent slots spread).
func (d *Dialer) DialAddrs(ctx context.Context, endpoint string, addrs, resolved []netip.Addr) (*Conn, error) {
	d.defaults()
	if len(addrs) == 0 {
		return nil, ErrNoAddress
	}
	var ports []int
	if d.Ports != nil {
		ports = d.Ports.Order(endpoint)
	} else {
		ports = []int{proto.PortDirect, proto.PortFallback}
	}
	var lastErr error
	for _, port := range ports {
		c, failed, err := d.dialPort(ctx, endpoint, addrs, port, false)
		if err == nil {
			if d.Ports != nil {
				d.Ports.RecordSuccess(endpoint, port)
			}
			return c, nil
		}
		lastErr = err
		if errors.Is(err, ErrNotNetwork) || ctx.Err() != nil {
			return nil, err
		}
		if d.Ports != nil && port == d.Ports.Primary() {
			d.Ports.RecordPrimaryFailure(endpoint, failed, resolved)
		}
		d.Log.Debug("port failed, trying next", "endpoint", endpoint, "port", port, "err", err)
	}
	return nil, lastErr
}

// DialAfterDrain is DialAddrs for the redial that follows a `drain`: the edge
// announced a planned restart, and its listener is back within milliseconds. The
// primary port is retried every DrainRetryStep for window, and those refusals
// are not primary-port failures (they would park the endpoint on the fallback). After
// the window the normal port order applies.
func (d *Dialer) DialAfterDrain(ctx context.Context, endpoint string, addrs, resolved []netip.Addr, window time.Duration) (*Conn, error) {
	d.defaults()
	if len(addrs) == 0 {
		return nil, ErrNoAddress
	}
	primary := proto.PortDirect
	if d.Ports != nil {
		primary = d.Ports.Primary()
	}
	deadline := time.Now().Add(window)
	for time.Now().Before(deadline) {
		c, _, err := d.dialPort(ctx, endpoint, addrs, primary, false)
		if err == nil {
			if d.Ports != nil {
				d.Ports.RecordSuccess(endpoint, primary)
			}
			return c, nil
		}
		if errors.Is(err, ErrNotNetwork) || ctx.Err() != nil {
			return nil, err
		}
		d.Log.Debug("primary refused after drain, retrying", "endpoint", endpoint, "port", primary, "err", err)
		t := time.NewTimer(DrainRetryStep)
		select {
		case <-t.C:
		case <-ctx.Done():
			t.Stop()
			return nil, ctx.Err()
		}
	}
	return d.DialAddrs(ctx, endpoint, addrs, resolved)
}

// DialAPI opens a discovery-API connection (ALPN hcdn-tunnel-api/1, or
// http/1.1 on the fallback port). The returned conn speaks HTTP/1.1.
func (d *Dialer) DialAPI(ctx context.Context, endpoint string, port int) (*tls.Conn, netip.Addr, error) {
	d.defaults()
	addrs, err := d.Resolve(ctx, endpoint)
	if err != nil {
		return nil, netip.Addr{}, err
	}
	tc, addr, _, err := d.handshake(ctx, endpoint, addrs, port, d.ALPN(port, true), nil)
	if err != nil {
		return nil, addr, addrErr(err, addr, addrs)
	}
	return tc, addr, nil
}

// Handshake performs one TLS handshake for RTT probing and returns the conn.
func (d *Dialer) Handshake(ctx context.Context, endpoint string, port int, alpn string) (*tls.Conn, time.Duration, error) {
	d.defaults()
	addrs, err := d.Resolve(ctx, endpoint)
	if err != nil {
		return nil, 0, err
	}
	tc, addr, rtt, err := d.handshake(ctx, endpoint, addrs, port, alpn, nil)
	if err != nil {
		return nil, 0, addrErr(err, addr, addrs)
	}
	return tc, rtt, nil
}

// dialPort also returns the addresses that failed: the TLS peer, or all of
// addrs when no TCP connect succeeded.
func (d *Dialer) dialPort(ctx context.Context, endpoint string, addrs []netip.Addr, port int, api bool) (*Conn, []netip.Addr, error) {
	alpn := d.ALPN(port, api)
	var serial string
	tc, addr, _, err := d.handshake(ctx, endpoint, addrs, port, alpn, &serial)
	if err != nil {
		if addr.IsValid() {
			return nil, []netip.Addr{addr}, err
		}
		return nil, addrs, err
	}
	c := &Conn{Conn: tc, Addr: addr, Port: port, Endpoint: endpoint, Transport: proto.TransportDirect, CertSerial: serial}
	if d.IsFallbackPort(port) {
		c.Transport = proto.TransportFallback
		up, err := d.upgrade(ctx, tc, endpoint)
		if err != nil {
			tc.Close()
			return nil, []netip.Addr{addr}, err
		}
		c.Conn = up
	}
	return c, nil, nil
}

// handshake dials TCP across addrs (Happy Eyeballs) and completes TLS.
func (d *Dialer) handshake(ctx context.Context, endpoint string, addrs []netip.Addr, port int, alpn string, serial *string) (*tls.Conn, netip.Addr, time.Duration, error) {
	start := time.Now()
	raw, addr, err := d.dialTCP(ctx, addrs, port)
	if err != nil {
		return nil, netip.Addr{}, 0, err
	}
	hctx, cancel := context.WithTimeout(ctx, d.ConnectTimeout)
	defer cancel()
	tc := tls.Client(raw, d.tlsConfig(endpoint, alpn, serial))
	if err := tc.HandshakeContext(hctx); err != nil {
		raw.Close()
		return nil, addr, 0, classifyTLSErr(err)
	}
	if got := tc.ConnectionState().NegotiatedProtocol; got != alpn {
		tc.Close()
		return nil, addr, 0, fmt.Errorf("%w: server negotiated ALPN %q, want %q", ErrNotNetwork, got, alpn)
	}
	return tc, addr, time.Since(start), nil
}

// A remote alert (client cert refused, ALPN refused) is a policy failure that
// nginx on 443 would repeat, so it must not trigger the port fallback.
func classifyTLSErr(err error) error {
	var cv *tls.CertificateVerificationError
	var ua x509.UnknownAuthorityError
	var he x509.HostnameError
	var ci x509.CertificateInvalidError
	if errors.As(err, &cv) || errors.As(err, &ua) || errors.As(err, &he) || errors.As(err, &ci) ||
		ingress.IsTLSAlert(err) || errors.Is(err, ErrNotNetwork) {
		return fmt.Errorf("%w: %v", ErrNotNetwork, err)
	}
	return err
}

type dialResult struct {
	c    net.Conn
	addr netip.Addr
	err  error
}

// dialTCP starts connection attempts HappyEyeballsDelay apart and keeps the first success.
func (d *Dialer) dialTCP(ctx context.Context, addrs []netip.Addr, port int) (net.Conn, netip.Addr, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	results := make(chan dialResult, len(addrs))
	nd := &net.Dialer{Timeout: d.ConnectTimeout}
	started, pending := 0, 0
	var lastErr error
	timer := time.NewTimer(0)
	defer timer.Stop()
	for started < len(addrs) || pending > 0 {
		select {
		case <-timer.C:
			if started < len(addrs) {
				a := addrs[started]
				started++
				pending++
				go func() {
					c, err := nd.DialContext(ctx, "tcp", netip.AddrPortFrom(a, uint16(port)).String())
					results <- dialResult{c, a, err}
				}()
				if started < len(addrs) {
					timer.Reset(d.HappyEyeballsDelay)
				}
			}
		case r := <-results:
			pending--
			if r.err == nil {
				cancel()
				go drainDials(results, pending)
				return r.c, r.addr, nil
			}
			lastErr = r.err
			if started < len(addrs) {
				timer.Reset(0)
			}
		case <-ctx.Done():
			go drainDials(results, pending)
			return nil, netip.Addr{}, ctx.Err()
		}
	}
	if lastErr == nil {
		lastErr = ErrNoAddress
	}
	return nil, netip.Addr{}, lastErr
}

func drainDials(results chan dialResult, pending int) {
	for i := 0; i < pending; i++ {
		if r := <-results; r.c != nil {
			r.c.Close()
		}
	}
}

// upgrade performs GET /_tunnel/v1/connect with Upgrade: hcdn-tunnel on the
// TLS conn and returns a conn positioned after the 101 (buffered bytes kept).
func (d *Dialer) upgrade(ctx context.Context, tc *tls.Conn, endpoint string) (net.Conn, error) {
	req := &http.Request{
		Method:     http.MethodGet,
		URL:        &url.URL{Path: proto.PathConnect},
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Host:       endpoint,
		Header:     http.Header{},
	}
	req.Header.Set("Upgrade", proto.UpgradeToken)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set(proto.HdrProtocol, proto.Version)
	req.Header.Set("User-Agent", "hcdn-connector/"+d.Version)
	deadline, ok := ctx.Deadline()
	if !ok || time.Until(deadline) > d.ConnectTimeout {
		deadline = time.Now().Add(d.ConnectTimeout)
	}
	_ = tc.SetDeadline(deadline)
	if err := req.Write(tc); err != nil {
		return nil, err
	}
	res, br, err := readResponse(tc, req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusSwitchingProtocols {
		// Still under the deadline: a stalled body must not hang the dial.
		_, _ = io.Copy(io.Discard, io.LimitReader(res.Body, 4096))
		res.Body.Close()
		return nil, fmt.Errorf("%w: fallback upgrade refused with status %d", ErrNotNetwork, res.StatusCode)
	}
	if !strings.EqualFold(res.Header.Get("Upgrade"), proto.UpgradeToken) {
		return nil, fmt.Errorf("%w: fallback upgrade answered an unexpected Upgrade token", ErrNotNetwork)
	}
	_ = tc.SetDeadline(time.Time{})
	if n := br.Buffered(); n > 0 {
		buf, _ := br.Peek(n)
		return &prefixConn{Conn: tc, tls: tc, r: io.MultiReader(bytes.NewReader(append([]byte(nil), buf...)), tc)}, nil
	}
	return tc, nil
}

// readResponse reads the response to req, capping the status line and headers
// (plus bufio read-ahead) at proto.MaxHeaderBytes: http.ReadResponse has no
// limit of its own. The cap lifts once the headers are parsed, so the body or
// upgraded stream read through br is not limited.
func readResponse(r io.Reader, req *http.Request) (*http.Response, *bufio.Reader, error) {
	head := &io.LimitedReader{R: r, N: proto.MaxHeaderBytes}
	br := bufio.NewReader(head)
	res, err := http.ReadResponse(br, req)
	if err != nil {
		if head.N <= 0 {
			return nil, nil, fmt.Errorf("%w: response header exceeds %d bytes", ErrNotNetwork, proto.MaxHeaderBytes)
		}
		return nil, nil, err
	}
	head.N = math.MaxInt64
	return res, br, nil
}

// prefixConn replays bytes the response reader buffered past the 101 and keeps
// the TLS state reachable for Conn.ConnectionState.
type prefixConn struct {
	net.Conn
	tls *tls.Conn
	r   io.Reader
}

func (p *prefixConn) Read(b []byte) (int, error)           { return p.r.Read(b) }
func (p *prefixConn) ConnectionState() tls.ConnectionState { return p.tls.ConnectionState() }

// APIBaseURL is the https URL prefix for the discovery API on endpoint:port.
func (d *Dialer) APIBaseURL(endpoint string, port int) string {
	if port == 443 {
		return "https://" + endpoint
	}
	return "https://" + net.JoinHostPort(endpoint, strconv.Itoa(port))
}
