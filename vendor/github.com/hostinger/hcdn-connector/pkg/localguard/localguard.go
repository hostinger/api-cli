// Package localguard keeps a public tunnel away from the sensitive services
// of the machine it runs on: databases, remote shells, debuggers and cluster
// agents. A port on the list is refused when the address is this machine:
// loopback, the unspecified address, one of its interface addresses, or the
// IPv4-mapped IPv6 form of any of these. The same port on another host is
// allowed, since it means nothing here. While the interface addresses cannot
// be read (a sandbox that denies netlink sockets, for one), any host may be
// this machine, so a port on the list is refused on every host.
//
// Control is the rule at dial time, for net.Dialer.Control. It sees the
// address that is actually dialed, after DNS and after any redirect, so it
// is the check that holds. Check is the same rule on a URL before anything
// starts, for a clear message.
package localguard

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/netip"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

// Port is a listed port and the service that usually listens on it.
type Port struct {
	Number  int
	Service string // for messages, for example "Redis" or "the Node inspector"
	// WarnOnly: the port is shared, but the caller should warn. Jupyter's
	// 8888 is also a common generic dev port, and a notebook asks for a token.
	WarnOnly bool
}

// blocked are the usual ports of databases, remote shells, mail, debuggers
// and cluster services. A tunnel is public, so these never reach this
// machine through it, even when they answer HTTP.
var blocked = map[int]Port{
	22:    {Service: "SSH"},
	23:    {Service: "Telnet"},
	25:    {Service: "SMTP"},
	110:   {Service: "POP3"},
	143:   {Service: "IMAP"},
	445:   {Service: "SMB file sharing"},
	631:   {Service: "CUPS printing"},
	2375:  {Service: "the Docker API"},
	2376:  {Service: "the Docker API"},
	2379:  {Service: "etcd"},
	2380:  {Service: "etcd peering"},
	3306:  {Service: "MySQL"},
	4444:  {Service: "Selenium Grid"},
	5432:  {Service: "PostgreSQL"},
	5555:  {Service: "ADB (Android Debug Bridge)"},
	5900:  {Service: "VNC"},
	5984:  {Service: "CouchDB"},
	6060:  {Service: "Go pprof"},
	6379:  {Service: "Redis"},
	6443:  {Service: "the Kubernetes API"},
	8500:  {Service: "Consul"},
	8888:  {Service: "Jupyter", WarnOnly: true},
	9200:  {Service: "Elasticsearch"},
	9222:  {Service: "Chrome DevTools"},
	9229:  {Service: "the Node inspector"},
	10250: {Service: "the Kubernetes kubelet"},
	10255: {Service: "the kubelet read-only API"},
	11211: {Service: "Memcached"},
	27017: {Service: "MongoDB"},
}

func init() {
	for n, p := range blocked {
		p.Number = n
		blocked[n] = p
	}
}

// BlockedPorts returns the list, warn-only ports included, sorted by number.
func BlockedPorts() []Port {
	out := make([]Port, 0, len(blocked))
	for _, p := range blocked {
		out = append(out, p)
	}
	slices.SortFunc(out, func(a, b Port) int { return a.Number - b.Number })
	return out
}

// LookupPort returns the listed port n, warn-only or not.
func LookupPort(n int) (Port, bool) {
	p, ok := blocked[n]
	return p, ok
}

// ErrBlocked matches every *BlockedError.
var ErrBlocked = errors.New("localguard: a sensitive port on this machine")

// BlockedError refuses a listed port on this machine.
type BlockedError struct {
	Port Port
	Addr string // the address dialed (Control) or the URL's host (Check)
	// Err is why the interface addresses cannot be read, when Addr is
	// refused only because it may be one of them; nil otherwise.
	Err error
}

func (e *BlockedError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("port %d is not shared: it is the usual port of %s, and %s may be this machine: its interface addresses cannot be read (%v)",
			e.Port.Number, e.Port.Service, e.Addr, e.Err)
	}
	return fmt.Sprintf("port %d is not shared: it is the usual port of %s, which must not be reachable from the internet",
		e.Port.Number, e.Port.Service)
}

func (e *BlockedError) Is(target error) bool { return target == ErrBlocked }

func (e *BlockedError) Unwrap() error { return e.Err }

// Resolver resolves a host name; *net.Resolver is one.
type Resolver interface {
	LookupNetIP(ctx context.Context, network, host string) ([]netip.Addr, error)
}

// localTTL bounds how long the interface addresses are cached. A failed
// read is not cached.
const localTTL = 30 * time.Second

// checkTimeout bounds the DNS lookup of Check.
const checkTimeout = 5 * time.Second

// Guard applies the rule. The zero value refuses every listed port that is
// not WarnOnly; Control, Check and Match use a shared zero Guard.
type Guard struct {
	// Allow lists ports that pass, for an operator who means to publish one
	// of them from this machine.
	Allow []int
	// Resolver resolves host names in Check and Match; nil: net.DefaultResolver.
	Resolver Resolver

	// interfaceAddrs is net.InterfaceAddrs; tests replace it.
	interfaceAddrs func() ([]net.Addr, error)

	mu      sync.Mutex
	local   []netip.Addr
	localAt time.Time
}

var std = &Guard{}

// Control refuses a dial to a listed port on this machine with the shared
// Guard. Its signature fits net.Dialer.Control.
func Control(network, address string, c syscall.RawConn) error {
	return std.Control(network, address, c)
}

// Check refuses a URL that reaches a listed port on this machine, with the
// shared Guard.
func Check(rawURL string) error { return std.Check(rawURL) }

// CheckContext is Check with a context for the DNS lookup.
func CheckContext(ctx context.Context, rawURL string) error { return std.CheckContext(ctx, rawURL) }

// Match returns the listed port, warn-only or not, that rawURL reaches on
// this machine, with the shared Guard.
func Match(ctx context.Context, rawURL string) (*Port, error) { return std.Match(ctx, rawURL) }

// Control is for net.Dialer.Control: address is the resolved ip:port about
// to be dialed. It returns a *BlockedError for a listed port, other than a
// WarnOnly or allowed one, on this machine, or on any address but loopback
// while the interface addresses cannot be read. Networks other than TCP
// and UDP pass.
func (g *Guard) Control(network, address string, _ syscall.RawConn) error {
	if !strings.HasPrefix(network, "tcp") && !strings.HasPrefix(network, "udp") {
		return nil
	}
	host, portS, err := net.SplitHostPort(address)
	if err != nil {
		return fmt.Errorf("localguard: dial address %q: %w", address, err)
	}
	n, err := strconv.Atoi(portS)
	if err != nil {
		return fmt.Errorf("localguard: dial address %q: bad port", address)
	}
	p, ok := g.listed(n)
	if !ok || p.WarnOnly {
		return nil
	}
	ip, err := netip.ParseAddr(host)
	if err != nil {
		return fmt.Errorf("localguard: dial address %q: %w", address, err)
	}
	if local, err := g.isLocal(ip); local {
		return &BlockedError{Port: p, Addr: address, Err: err}
	}
	return nil
}

// Probe reads the interface addresses, as Control does, and returns why
// they cannot be read. Call it at start to report a sandbox that hides
// them: until they can be read, every listed port is refused on every host.
func (g *Guard) Probe() error {
	if _, err := g.localAddrs(); err != nil {
		return fmt.Errorf("localguard: read interface addresses: %w", err)
	}
	return nil
}

// Check refuses rawURL with a *BlockedError when it reaches a listed port,
// other than a WarnOnly or allowed one, on this machine, or on any host
// while the interface addresses cannot be read. It resolves the host once,
// within 5 s. A host that does not resolve passes: Control still
// applies at dial time. A URL that does not parse is an error.
func (g *Guard) Check(rawURL string) error { return g.CheckContext(context.Background(), rawURL) }

// CheckContext is Check with a context for the DNS lookup, still within 5 s.
func (g *Guard) CheckContext(ctx context.Context, rawURL string) error {
	ctx, cancel := context.WithTimeout(ctx, checkTimeout)
	defer cancel()
	p, unknown, err := g.match(ctx, rawURL)
	if err != nil {
		return err
	}
	if p != nil && !p.WarnOnly {
		u, _ := url.Parse(rawURL)
		return &BlockedError{Port: *p, Addr: u.Host, Err: unknown}
	}
	return nil
}

// Match returns the listed port, warn-only or not, that rawURL reaches on
// this machine, or nil. Allowed ports never match. It resolves the host
// once; a host that does not resolve does not match. While the interface
// addresses cannot be read, a listed port matches on any host that
// resolves.
func (g *Guard) Match(ctx context.Context, rawURL string) (*Port, error) {
	p, _, err := g.match(ctx, rawURL)
	return p, err
}

// match is Match, with why the interface addresses cannot be read when
// the port matched only because the host may be one of them.
func (g *Guard) match(ctx context.Context, rawURL string) (_ *Port, unknown, err error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, nil, err
	}
	n, err := urlPort(u)
	if err != nil {
		return nil, nil, err
	}
	p, ok := g.listed(n)
	if !ok {
		return nil, nil, nil
	}
	local, unknown := g.hostIsLocal(ctx, u.Hostname())
	if !local {
		return nil, nil, nil
	}
	return &p, unknown, nil
}

// urlPort is the URL's port, or its scheme's default; 0 for an unknown
// scheme without a port.
func urlPort(u *url.URL) (int, error) {
	if s := u.Port(); s != "" {
		n, err := strconv.Atoi(s)
		if err != nil || n < 1 || n > 65535 {
			return 0, fmt.Errorf("localguard: %q: invalid port %q", u.Redacted(), s)
		}
		return n, nil
	}
	switch strings.ToLower(u.Scheme) {
	case "http", "ws":
		return 80, nil
	case "https", "wss":
		return 443, nil
	}
	return 0, nil
}

func (g *Guard) listed(n int) (Port, bool) {
	p, ok := blocked[n]
	if !ok || slices.Contains(g.Allow, n) {
		return Port{}, false
	}
	return p, true
}

// localNames always reach this machine or its container host, whatever DNS
// says: RFC 6761 localhost names and the names container runtimes give the host.
var localNames = []string{
	"localhost",
	"host.docker.internal",
	"gateway.docker.internal",
	"host.containers.internal",
	"host.lima.internal",
}

// hostIsLocal reports whether host reaches this machine. "" is the local
// system for a dial. Numeric forms such as 127.1, 0 and 0x7f000001 are read
// as the C resolver reads them, before any DNS. A host that does not resolve
// is false with the DNS error; one that may be this machine because the
// interface addresses cannot be read is true with that error.
func (g *Guard) hostIsLocal(ctx context.Context, host string) (bool, error) {
	h := strings.ToLower(strings.TrimSuffix(host, "."))
	if h == "" {
		return true, nil
	}
	for _, name := range localNames {
		if h == name || (name == "localhost" && strings.HasSuffix(h, ".localhost")) {
			return true, nil
		}
	}
	if ip, ok := parseLooseIPv4(h); ok {
		return g.isLocal(ip)
	}
	if ip, err := netip.ParseAddr(h); err == nil {
		return g.isLocal(ip)
	}
	var r Resolver = net.DefaultResolver
	if g.Resolver != nil {
		r = g.Resolver
	}
	ips, err := r.LookupNetIP(ctx, "ip", h)
	if err != nil {
		return false, err
	}
	for _, ip := range ips {
		if local, err := g.isLocal(ip); local {
			return true, err
		}
	}
	return false, nil
}

// parseLooseIPv4 reads the inet_aton forms that getaddrinfo accepts: one to
// four parts, each decimal, octal (leading 0) or hex (0x), the last part
// filling the remaining bytes. "127.1", "0", "2130706433" and "0x7f000001"
// are all IPv4 addresses to the C resolver.
func parseLooseIPv4(s string) (netip.Addr, bool) {
	parts := strings.Split(s, ".")
	if len(parts) > 4 {
		return netip.Addr{}, false
	}
	vals := make([]uint64, len(parts))
	for i, part := range parts {
		base, digits := 10, part
		switch {
		case len(part) > 1 && (part[:2] == "0x" || part[:2] == "0X"):
			base, digits = 16, part[2:]
		case len(part) > 1 && part[0] == '0':
			base, digits = 8, part[1:]
		}
		if digits == "" && base != 16 {
			return netip.Addr{}, false
		}
		if digits == "" {
			vals[i] = 0 // "0x" alone is 0 to inet_aton
			continue
		}
		v, err := strconv.ParseUint(digits, base, 32)
		if err != nil {
			return netip.Addr{}, false
		}
		vals[i] = v
	}
	var n uint64
	for i, v := range vals[:len(vals)-1] {
		if v > 0xff {
			return netip.Addr{}, false
		}
		n |= v << (24 - 8*i)
	}
	last := vals[len(vals)-1]
	if last >= 1<<(8*(5-len(vals))) {
		return netip.Addr{}, false
	}
	n |= last
	return netip.AddrFrom4([4]byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)}), true
}

// isLocal: loopback, unspecified (0.0.0.0/8 dials this machine on Linux),
// or one of this machine's interface addresses. An IPv4-mapped address
// counts as its IPv4 form; a zone is ignored. When the interface addresses
// cannot be read, any other address may be one of them: true, with the
// read error.
func (g *Guard) isLocal(ip netip.Addr) (bool, error) {
	ip = ip.Unmap().WithZone("")
	if ip.IsLoopback() || ip.IsUnspecified() || (ip.Is4() && ip.As4()[0] == 0) {
		return true, nil
	}
	local, err := g.localAddrs()
	if err != nil {
		return true, err
	}
	return slices.Contains(local, ip), nil
}

// localAddrs returns the interface addresses, cached for 30 s. An error is
// not cached: the next check reads them again. A failure is cheap (a sandbox
// refuses the socket at once) and is reached only for a listed port, which
// is then refused, so retrying costs little and recovers at once.
func (g *Guard) localAddrs() ([]netip.Addr, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.local != nil && time.Since(g.localAt) < localTTL {
		return g.local, nil
	}
	list := g.interfaceAddrs
	if list == nil {
		list = net.InterfaceAddrs
	}
	addrs, err := list()
	if err != nil {
		return nil, err
	}
	out := make([]netip.Addr, 0, len(addrs))
	for _, a := range addrs {
		var ip net.IP
		switch v := a.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if na, ok := netip.AddrFromSlice(ip); ok {
			out = append(out, na.Unmap().WithZone(""))
		}
	}
	g.local, g.localAt = out, time.Now()
	return out, nil
}
