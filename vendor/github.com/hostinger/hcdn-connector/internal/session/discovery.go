package session

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"net"
	"net/http"
	"net/netip"
	"net/url"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/publicsuffix"

	"github.com/hostinger/hcdn-connector/internal/proto"
	"github.com/hostinger/hcdn-connector/internal/termsafe"
	"github.com/hostinger/hcdn-connector/pkg/config"
)

const (
	PopsCacheTTL   = time.Hour
	ProbeTimeout   = 3 * time.Second
	ProbeSamples   = 3
	apiTimeout     = 5 * time.Second
	maxAPIBody     = 1 << 20
	stormWindow    = 60 * time.Second
	stormThreshold = 3
)

// refreshMinGap rate-limits forced pop-list refreshes; a var so tests can shorten it.
var refreshMinGap = 60 * time.Second

// Pop is one entry of GET /_hcdn/v1/pops.
type Pop struct {
	Code     string `json:"code"`
	Endpoint string `json:"endpoint"`
	Ports    []int  `json:"ports"`
	Region   string `json:"region"`
	Allowed  bool   `json:"allowed"`
}

type PopsReply struct {
	Gen        int64    `json:"gen"`
	Pops       []Pop    `json:"pops"`
	PinnedPops []string `json:"pinned_pops"`
	FetchedAt  time.Time
	Source     string // endpoint that answered, or "manual"
}

func (r *PopsReply) IsPinned(code string) bool {
	for _, p := range r.PinnedPops {
		if p == code {
			return true
		}
	}
	return false
}

type Whoami struct {
	TunnelID   string    `json:"tunnel_id"`
	CertSerial string    `json:"cert_serial"`
	DaysLeft   int       `json:"days_left"`
	Pop        string    `json:"pop"`
	Node       string    `json:"node"`
	Class      string    `json:"class"`      // named | quick; "" from tunneld < 0.3.0
	Hostname   string    `json:"hostname"`   // quick only
	ExpiresAt  time.Time `json:"expires_at"` // zero from tunneld < 0.3.0
	// Owner is set when the edge asks for an owner ticket (quick class, owner
	// mode log or require). tunneld before 0.5.0, and mode off, omit it.
	Owner *WhoamiOwner `json:"owner,omitempty"`
}

// WhoamiOwner is the whoami owner object. Mode and Issuer are always set;
// Status, Reason and ClientID answer a whoami that carried hcdn-owner-ticket.
type WhoamiOwner struct {
	Mode     string `json:"mode"`                // log | require
	Issuer   string `json:"issuer"`              // the iss the edge accepts
	Status   string `json:"status,omitempty"`    // ok | missing | invalid
	Reason   string `json:"reason,omitempty"`    // the edge's verification label
	ClientID string `json:"client_id,omitempty"` // the linked owner, when Status is ok
}

// OwnerMode is the mode that asks for a ticket (proto.OwnerModeLog or
// proto.OwnerModeRequire), or "" when the edge did not ask. Another value
// (off, or a mode this connector does not know) counts as not asked.
func (w *Whoami) OwnerMode() string {
	if w == nil || w.Owner == nil {
		return ""
	}
	switch m := strings.ToLower(strings.TrimSpace(w.Owner.Mode)); m {
	case proto.OwnerModeLog, proto.OwnerModeRequire:
		return m
	}
	return ""
}

const (
	ClassNamed = "named"
	ClassQuick = "quick"
)

// ProbeResult is the median TLS handshake RTT to one PoP.
type ProbeResult struct {
	Pop     Pop
	Port    int
	RTT     time.Duration
	Samples []time.Duration
	Err     error
	At      time.Time
}

// Discovery talks to the discovery API and caches the PoP list.
type Discovery struct {
	Dialer    *Dialer
	Bootstrap func() []string
	Log       *slog.Logger

	mu          sync.Mutex
	cached      *PopsReply
	lastRefresh time.Time
}

func (d *Discovery) log() *slog.Logger {
	if d.Log == nil {
		return slog.Default()
	}
	return d.Log
}

// Cached returns the cached PoP list, if any.
func (d *Discovery) Cached() *PopsReply {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.cached
}

// Pops returns the PoP list, from cache unless it is older than 1 h or force
// is set. A forced refresh is rate-limited to one per minute.
func (d *Discovery) Pops(ctx context.Context, force bool) (*PopsReply, error) {
	d.mu.Lock()
	c := d.cached
	last := d.lastRefresh
	d.mu.Unlock()
	if c != nil {
		fresh := time.Since(c.FetchedAt) < PopsCacheTTL
		if !force && fresh {
			return c, nil
		}
		if force && time.Since(last) < refreshMinGap {
			return c, nil
		}
	}
	r, err := d.fetchPops(ctx)
	d.mu.Lock()
	d.lastRefresh = time.Now()
	if err == nil {
		d.cached = r
	} else if c != nil {
		r = c
		err = nil
	}
	d.mu.Unlock()
	return r, err
}

func (d *Discovery) fetchPops(ctx context.Context) (*PopsReply, error) {
	var errs []error
	for _, ep := range d.shuffledBootstrap() {
		var r PopsReply
		port, err := d.get(ctx, ep, proto.PathPops, nil, &r)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", ep, err))
			continue
		}
		if err := validatePops(&r, d.Bootstrap(), d.Dialer != nil && d.Dialer.RootCAs != nil); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", ep, err))
			continue
		}
		r.FetchedAt = time.Now()
		r.Source = ep
		d.log().Info("pop list refreshed", "source", ep, "port", port, "gen", r.Gen, "pops", len(r.Pops))
		return &r, nil
	}
	return nil, fmt.Errorf("discovery failed on every bootstrap endpoint: %w", errors.Join(errs...))
}

// validatePops checks a /pops reply. Each endpoint must be in the
// registrable domain (eTLD+1) of a bootstrap endpoint: a session sends the
// hello (with the owner ticket) to it and serves the requests it sends, so
// a reply must not point sessions at an unrelated host that has a public
// certificate. With a private CA (privateCA) only that CA's hosts can
// answer, and any endpoint is accepted.
func validatePops(r *PopsReply, bootstrap []string, privateCA bool) error {
	if len(r.Pops) == 0 {
		return errors.New("empty pop list")
	}
	var domains []string
	for _, b := range bootstrap {
		if d := registrableDomain(b); d != "" {
			domains = append(domains, d)
		}
	}
	seen := map[string]bool{}
	for i := range r.Pops {
		p := &r.Pops[i]
		p.Code = strings.ToLower(p.Code)
		if p.Code == "" || p.Endpoint == "" || strings.ContainsAny(p.Endpoint, " /:") {
			return fmt.Errorf("pop %d: invalid code/endpoint", i)
		}
		if !termsafe.Printable(p.Code) || !termsafe.Printable(p.Endpoint) || !termsafe.Printable(p.Region) {
			return fmt.Errorf("pop %d: non-printable code/endpoint/region", i)
		}
		if d := registrableDomain(p.Endpoint); !privateCA && (d == "" || !slices.Contains(domains, d)) {
			return fmt.Errorf("pop %s: endpoint %s is outside the domains of the bootstrap endpoints", p.Code, p.Endpoint)
		}
		if seen[p.Code] {
			return fmt.Errorf("duplicate pop %s", p.Code)
		}
		seen[p.Code] = true
	}
	for i, p := range r.PinnedPops {
		r.PinnedPops[i] = strings.ToLower(p)
	}
	return nil
}

// registrableDomain is the eTLD+1 of host per the public suffix list
// ("hostinger.io" for fra.tunnel.hostinger.io), or "" for an IP address or a
// name without one.
func registrableDomain(host string) string {
	host = strings.TrimSuffix(strings.ToLower(strings.Trim(host, "[]")), ".")
	if _, err := netip.ParseAddr(host); err == nil {
		return ""
	}
	d, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return ""
	}
	return d
}

func (d *Discovery) shuffledBootstrap() []string {
	eps := append([]string(nil), d.Bootstrap()...)
	rand.Shuffle(len(eps), func(i, j int) { eps[i], eps[j] = eps[j], eps[i] })
	return eps
}

// Whoami calls GET /_hcdn/v1/whoami on the first bootstrap endpoint that
// answers with a valid reply (validateWhoami). A non-empty ownerTicket goes
// out as hcdn-owner-ticket.
func (d *Discovery) Whoami(ctx context.Context, ownerTicket string) (*Whoami, string, error) {
	var hdr http.Header
	if ownerTicket != "" {
		hdr = http.Header{proto.HdrOwnerTicket: {ownerTicket}}
	}
	var errs []error
	for _, ep := range d.shuffledBootstrap() {
		var w Whoami
		if _, err := d.get(ctx, ep, proto.PathWhoami, hdr, &w); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", ep, err))
			continue
		}
		if err := validateWhoami(&w); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", ep, err))
			continue
		}
		return &w, ep, nil
	}
	return nil, "", errors.Join(errs...)
}

// validateWhoami checks the strings that check and quick print: tunnel_id
// is a UUID, cert_serial is hex, hostname (quick only) is a host name and
// every other string is printable. Fields older tunneld versions omit may be
// empty.
func validateWhoami(w *Whoami) error {
	if !config.IsUUID(w.TunnelID) {
		return errors.New("whoami: invalid tunnel_id")
	}
	if w.CertSerial == "" || strings.Trim(strings.ToLower(w.CertSerial), "0123456789abcdef") != "" {
		return errors.New("whoami: invalid cert_serial")
	}
	if w.Hostname != "" && !config.IsHostname(strings.ToLower(w.Hostname)) {
		return errors.New("whoami: invalid hostname")
	}
	if !termsafe.Printable(w.Pop) || !termsafe.Printable(w.Node) || !termsafe.Printable(w.Class) {
		return errors.New("whoami: non-printable pop/node/class")
	}
	if o := w.Owner; o != nil {
		for _, s := range []string{o.Mode, o.Issuer, o.Status, o.Reason, o.ClientID} {
			if !termsafe.Printable(s) {
				return errors.New("whoami: non-printable owner field")
			}
		}
	}
	return nil
}

// get performs a JSON GET against endpoint, trying ports in memory order.
// hdr adds request headers (nil: none).
func (d *Discovery) get(ctx context.Context, endpoint, path string, hdr http.Header, out any) (int, error) {
	var lastErr error
	for _, port := range d.Dialer.Ports.Order(endpoint) {
		err := d.getOnPort(ctx, endpoint, port, path, hdr, out)
		if err == nil {
			d.Dialer.Ports.RecordSuccess(endpoint, port)
			return port, nil
		}
		lastErr = err
		if errors.Is(err, ErrNotNetwork) || ctx.Err() != nil {
			return 0, err
		}
		if port == d.Dialer.Ports.Primary() {
			recordPrimaryFailure(d.Dialer.Ports, endpoint, err)
		}
	}
	return 0, lastErr
}

// recordPrimaryFailure counts dial failures only: a node that completed TLS on
// the primary and then stalled has a node problem, not a blocked port.
func recordPrimaryFailure(pm *PortMemory, endpoint string, err error) {
	var ae *AddrError
	if errors.As(err, &ae) {
		pm.RecordPrimaryFailure(endpoint, ae.Failed, ae.Resolved)
	}
}

func (d *Discovery) getOnPort(ctx context.Context, endpoint string, port int, path string, hdr http.Header, out any) error {
	ctx, cancel := context.WithTimeout(ctx, apiTimeout)
	defer cancel()
	// The transport dials on its own goroutine, which can outlive RoundTrip.
	var (
		dialMu  sync.Mutex
		dialErr error
	)
	tr := &http.Transport{
		DialTLSContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
			tc, _, err := d.Dialer.DialAPI(ctx, endpoint, port)
			dialMu.Lock()
			dialErr = err
			dialMu.Unlock()
			return tc, err
		},
		DisableKeepAlives: true,
	}
	defer tr.CloseIdleConnections()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.Dialer.APIBaseURL(endpoint, port)+path, nil)
	if err != nil {
		return err
	}
	for k, v := range hdr {
		req.Header[k] = v
	}
	req.Header.Set("User-Agent", "hcdn-connector/"+d.Dialer.Version)
	req.Header.Set("Accept", "application/json")
	res, err := tr.RoundTrip(req)
	if err != nil {
		dialMu.Lock()
		defer dialMu.Unlock()
		if dialErr != nil {
			return dialErr
		}
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(io.LimitReader(res.Body, maxAPIBody))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %s answered %d", ErrNotNetwork, path, res.StatusCode)
	}
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("%w: %s: bad JSON: %v", ErrNotNetwork, path, err)
	}
	return nil
}

// Probe measures every PoP in parallel: three TLS handshakes each followed by
// GET /_hcdn/v1/ping, median RTT, 3 s timeout per sample.
func (d *Discovery) Probe(ctx context.Context, pops []Pop) []ProbeResult {
	out := make([]ProbeResult, len(pops))
	var wg sync.WaitGroup
	for i, p := range pops {
		wg.Add(1)
		go func(i int, p Pop) {
			defer wg.Done()
			out[i] = d.probeOne(ctx, p)
		}(i, p)
	}
	wg.Wait()
	return out
}

func (d *Discovery) probeOne(ctx context.Context, p Pop) ProbeResult {
	r := ProbeResult{Pop: p, At: time.Now()}
	var lastErr error
	for _, port := range d.Dialer.Ports.Order(p.Endpoint) {
		samples, err := d.samplePort(ctx, p.Endpoint, port)
		if err == nil {
			d.Dialer.Ports.RecordSuccess(p.Endpoint, port)
			sort.Slice(samples, func(a, b int) bool { return samples[a] < samples[b] })
			r.Port, r.Samples, r.RTT = port, samples, samples[len(samples)/2]
			return r
		}
		lastErr = err
		if errors.Is(err, ErrNotNetwork) || ctx.Err() != nil {
			break
		}
		if port == d.Dialer.Ports.Primary() {
			recordPrimaryFailure(d.Dialer.Ports, p.Endpoint, err)
		}
	}
	r.Err = lastErr
	return r
}

func (d *Discovery) samplePort(ctx context.Context, endpoint string, port int) ([]time.Duration, error) {
	samples := make([]time.Duration, 0, ProbeSamples)
	for i := 0; i < ProbeSamples; i++ {
		rtt, err := d.sample(ctx, endpoint, port)
		if err != nil {
			return nil, err
		}
		samples = append(samples, rtt)
	}
	return samples, nil
}

func (d *Discovery) sample(ctx context.Context, endpoint string, port int) (time.Duration, error) {
	ctx, cancel := context.WithTimeout(ctx, ProbeTimeout)
	defer cancel()
	tc, rtt, err := d.Dialer.Handshake(ctx, endpoint, port, d.Dialer.ALPN(port, true))
	if err != nil {
		return 0, err
	}
	defer tc.Close()
	if err := ping(ctx, tc, endpoint, port, d.Dialer.Version); err != nil {
		return 0, err
	}
	return rtt, nil
}

// ping sends GET /_hcdn/v1/ping over an established API conn and expects 2xx.
func ping(ctx context.Context, tc *tls.Conn, endpoint string, port int, version string) error {
	if dl, ok := ctx.Deadline(); ok {
		_ = tc.SetDeadline(dl)
	}
	req := &http.Request{
		Method: http.MethodGet, URL: &url.URL{Path: proto.PathPing}, Proto: "HTTP/1.1", ProtoMajor: 1, ProtoMinor: 1,
		Host: hostFor(endpoint, port), Header: http.Header{"User-Agent": {"hcdn-connector/" + version}},
	}
	if err := req.Write(tc); err != nil {
		return err
	}
	res, _, err := readResponse(tc, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, io.LimitReader(res.Body, 4096))
	res.Body.Close()
	if res.StatusCode/100 != 2 {
		return fmt.Errorf("%w: ping answered %d", ErrNotNetwork, res.StatusCode)
	}
	return nil
}

func hostFor(endpoint string, port int) string {
	if port == 443 {
		return endpoint
	}
	return net.JoinHostPort(endpoint, fmt.Sprint(port))
}

// ManualPops builds a PoP list without the discovery API (tests and manual setups).
// A bootstrap entry named <code>.<suffix> wins; otherwise <code> + the suffix
// of the first bootstrap name, or the bootstrap address itself when it is an IP.
func ManualPops(codes, bootstrap []string) *PopsReply {
	r := &PopsReply{FetchedAt: time.Now(), Source: "manual"}
	suffix, literal := "", ""
	if len(bootstrap) > 0 {
		if _, err := netip.ParseAddr(strings.Trim(bootstrap[0], "[]")); err == nil {
			literal = bootstrap[0]
		} else if i := strings.IndexByte(bootstrap[0], '.'); i > 0 {
			suffix = bootstrap[0][i:]
		}
	}
	for _, c := range codes {
		ep := c + suffix
		if literal != "" {
			ep = literal
		}
		for _, b := range bootstrap {
			if strings.HasPrefix(b, c+".") {
				ep = b
			}
		}
		r.Pops = append(r.Pops, Pop{Code: c, Endpoint: ep, Allowed: true})
	}
	return r
}
