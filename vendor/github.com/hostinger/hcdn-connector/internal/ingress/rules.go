// Package ingress routes edge requests to local services.
package ingress

import (
	"net"
	"sort"
	"strings"

	"github.com/hostinger/hcdn-connector/pkg/config"
	"github.com/hostinger/hcdn-connector/pkg/localguard"
)

// Rule is one ingress entry. Hostname "" is the catch-all.
type Rule struct {
	Hostname string
	Service  *Service
	suffix   string // ".example.com" for "*.example.com"
}

// Table matches exact hostnames first, then the longest wildcard suffix,
// then the catch-all. "*.example.com" matches any depth below example.com.
type Table struct {
	exact     map[string]*Rule
	wildcards []*Rule
	catchAll  *Rule
	Rules     []*Rule
}

func NewTable(rules []*Rule) *Table {
	t := &Table{exact: map[string]*Rule{}, Rules: rules}
	for _, r := range rules {
		switch {
		case r.Hostname == "":
			t.catchAll = r
		case strings.HasPrefix(r.Hostname, "*."):
			r.suffix = r.Hostname[1:]
			t.wildcards = append(t.wildcards, r)
		default:
			t.exact[r.Hostname] = r
		}
	}
	sort.SliceStable(t.wildcards, func(i, j int) bool { return len(t.wildcards[i].suffix) > len(t.wildcards[j].suffix) })
	return t
}

// NormalizeHost strips an optional numeric port and one trailing dot and
// lowercases ASCII. It returns "" (the catch-all) for any Host an origin could
// read differently from the result: brackets around anything but an IPv6
// literal, a non-numeric port, or bytes outside [A-Za-z0-9._-].
func NormalizeHost(h string) string {
	if strings.HasPrefix(h, "[") {
		end := strings.IndexByte(h, ']')
		if end < 0 || !validPortSuffix(h[end+1:]) {
			return ""
		}
		ip := net.ParseIP(h[1:end])
		if ip == nil || ip.To4() != nil || strings.Contains(h[1:end], "%") {
			return ""
		}
		return strings.ToLower(h[1:end])
	}
	if i := strings.IndexByte(h, ':'); i >= 0 {
		if !validPortSuffix(h[i:]) {
			return ""
		}
		h = h[:i]
	}
	h = strings.TrimSuffix(h, ".")
	if strings.HasPrefix(h, ".") || strings.Contains(h, "..") {
		return ""
	}
	b := make([]byte, len(h))
	for i := 0; i < len(h); i++ {
		c := h[i]
		switch {
		case c >= 'A' && c <= 'Z':
			b[i] = c + ('a' - 'A')
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9', c == '.', c == '-', c == '_':
			b[i] = c
		default:
			return ""
		}
	}
	return string(b)
}

// validPortSuffix accepts "" or ":" followed by 1-5 digits.
func validPortSuffix(s string) bool {
	if s == "" {
		return true
	}
	if s[0] != ':' || len(s) < 2 || len(s) > 6 {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func (t *Table) Match(host string) *Rule {
	h := NormalizeHost(host)
	if r, ok := t.exact[h]; ok {
		return r
	}
	for _, r := range t.wildcards {
		if len(h) > len(r.suffix) && strings.HasSuffix(h, r.suffix) {
			return r
		}
	}
	return t.catchAll
}

// Build creates the table from config; it opens ca_file for https services.
// guard applies to every service dial; nil dials anything.
func Build(rules []config.IngressRule, tunnelID string, guard *localguard.Guard) (*Table, error) {
	out := make([]*Rule, 0, len(rules))
	for _, rc := range rules {
		svc, err := NewService(rc.Service, rc.OriginRequest, tunnelID, guard)
		if err != nil {
			return nil, err
		}
		out = append(out, &Rule{Hostname: rc.Hostname, Service: svc})
	}
	return NewTable(out), nil
}

// Close releases idle service connections.
func (t *Table) Close() {
	for _, r := range t.Rules {
		if r.Service.Transport != nil {
			r.Service.Transport.CloseIdleConnections()
		}
	}
}
