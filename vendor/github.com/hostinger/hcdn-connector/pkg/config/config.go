// Package config parses and validates config.yml, or builds the same Config
// in code from Default.
package config

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/hostinger/hcdn-connector/internal/fileguard"
	"github.com/hostinger/hcdn-connector/internal/proto"
)

const (
	ModeAuto   = "auto"
	ModeManual = "manual"

	ServiceStatusPrefix = "http_status:"

	DefaultMetricsAddr = "127.0.0.1:9112"
)

type Config struct {
	TunnelID        string        `yaml:"tunnel_id"`
	CredentialsFile string        `yaml:"credentials_file"`
	HomePops        int           `yaml:"home_pops"`
	SessionsPerPop  int           `yaml:"sessions_per_pop"`
	PopSelection    PopSelection  `yaml:"pop_selection"`
	Ports           []int         `yaml:"ports"`
	Ingress         []IngressRule `yaml:"ingress"`
	AdminSocket     string        `yaml:"admin_socket"`
	MetricsAddr     *string       `yaml:"metrics_addr"`
	Log             Log           `yaml:"log"`
	ServerSPKIPins  []string      `yaml:"server_spki_pins"`
	ServerCAFile    string        `yaml:"server_ca_file"`
	Transport       string        `yaml:"transport"`
	// AllowLocalPorts opts these ports out of the local port guard
	// (localguard.BlockedPorts on this machine). The run and check commands
	// pass it as connector.Options.AllowLocalPorts.
	AllowLocalPorts []int `yaml:"allow_local_ports"`
}

type PopSelection struct {
	Mode             string        `yaml:"mode"`
	Pops             []string      `yaml:"pops"`
	ReprobeInterval  time.Duration `yaml:"reprobe_interval"`
	SwitchMinGainPct int           `yaml:"switch_min_gain_pct"`
	SwitchMinGainMs  int           `yaml:"switch_min_gain_ms"`
	Holddown         time.Duration `yaml:"holddown"`
}

type IngressRule struct {
	Hostname      string        `yaml:"hostname"`
	Service       string        `yaml:"service"`
	OriginRequest OriginRequest `yaml:"origin_request"`
}

type OriginRequest struct {
	TLSServerName         string        `yaml:"tls_server_name"`
	CAFile                string        `yaml:"ca_file"`
	NoTLSVerify           bool          `yaml:"no_tls_verify"`
	ConnectTimeout        time.Duration `yaml:"connect_timeout"`
	ResponseHeaderTimeout time.Duration `yaml:"response_header_timeout"`
	KeepaliveConnections  int           `yaml:"keepalive_connections"`
	HTTPHostHeader        string        `yaml:"http_host_header"`
	HTTP2                 bool          `yaml:"http2"`
}

type Log struct {
	Format string `yaml:"format"`
	Level  string `yaml:"level"`
}

var uuidRe = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
var popCodeRe = regexp.MustCompile(`^[a-z0-9]{2,16}$`)
var labelRe = regexp.MustCompile(`^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`)

func IsUUID(s string) bool { return uuidRe.MatchString(strings.ToLower(s)) }

// IsHostname reports whether s is a lowercase host name without a wildcard.
func IsHostname(s string) bool {
	if s == "" || len(s) > 253 {
		return false
	}
	for _, l := range strings.Split(s, ".") {
		if !labelRe.MatchString(l) {
			return false
		}
	}
	return true
}

// Load reads, defaults and validates a config file. Unknown keys are errors.
// The file chooses the services, the CA files and the credentials file, so it
// is refused when another account can change it or its directory: on Windows
// an owner or a write ACE for anyone but SYSTEM, Administrators and the
// current user, elsewhere a world-writable file or directory.
func Load(path string) (*Config, error) {
	b, err := fileguard.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(b)
}

func Parse(b []byte) (*Config, error) {
	var c Config
	dec := yaml.NewDecoder(strings.NewReader(string(b)))
	dec.KnownFields(true)
	if err := dec.Decode(&c); err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("config: %w", err)
	}
	var set explicitKeys
	if err := yaml.Unmarshal(b, &set); err != nil {
		return nil, fmt.Errorf("config: %w", err)
	}
	c.applyDefaults(&set)
	if err := c.validate(true); err != nil {
		return nil, err
	}
	return &c, nil
}

// Default returns every default and one ingress rule, the catch-all
// http_status:404. Set TunnelID and add rules with AddRule. Fields set to
// zero after the call keep zero (for example Holddown: 0).
func Default() *Config {
	metrics := DefaultMetricsAddr
	return &Config{
		HomePops:       2,
		SessionsPerPop: 2,
		PopSelection: PopSelection{
			Mode:             ModeAuto,
			ReprobeInterval:  10 * time.Minute,
			SwitchMinGainPct: 20,
			SwitchMinGainMs:  15,
			Holddown:         30 * time.Minute,
		},
		Ports:       []int{proto.PortDirect, proto.PortFallback},
		Ingress:     []IngressRule{NewIngressRule("", ServiceStatusPrefix+"404")},
		AdminSocket: DefaultAdminSocket,
		MetricsAddr: &metrics,
		Log:         Log{Format: "json", Level: "info"},
		Transport:   "h2",
	}
}

func defaultOriginRequest() OriginRequest {
	return OriginRequest{ConnectTimeout: 5 * time.Second, ResponseHeaderTimeout: 45 * time.Second, KeepaliveConnections: 16}
}

// NewIngressRule returns a rule with the origin_request defaults.
func NewIngressRule(hostname, service string) IngressRule {
	return IngressRule{Hostname: hostname, Service: service, OriginRequest: defaultOriginRequest()}
}

// AddRule inserts r before the catch-all rule, or appends it when there is none.
func (c *Config) AddRule(r IngressRule) {
	if n := len(c.Ingress); n > 0 && c.Ingress[n-1].Hostname == "" {
		c.Ingress = append(c.Ingress[:n-1], r, c.Ingress[n-1])
		return
	}
	c.Ingress = append(c.Ingress, r)
}

// Clone returns a deep copy.
func (c *Config) Clone() *Config {
	n := *c
	n.Ports = slices.Clone(c.Ports)
	n.Ingress = slices.Clone(c.Ingress)
	n.ServerSPKIPins = slices.Clone(c.ServerSPKIPins)
	n.AllowLocalPorts = slices.Clone(c.AllowLocalPorts)
	n.PopSelection.Pops = slices.Clone(c.PopSelection.Pops)
	if c.MetricsAddr != nil {
		addr := *c.MetricsAddr
		n.MetricsAddr = &addr
	}
	return &n
}

// Normalize lowercases the tunnel ID, the host names and the PoP codes.
func (c *Config) Normalize() {
	c.TunnelID = strings.ToLower(c.TunnelID)
	for i := range c.Ingress {
		c.Ingress[i].Hostname = strings.ToLower(strings.TrimSuffix(c.Ingress[i].Hostname, "."))
	}
	for i, p := range c.PopSelection.Pops {
		c.PopSelection.Pops[i] = strings.ToLower(p)
	}
}

// explicitKeys records keys whose zero value is a valid setting, so a default
// does not replace an explicit 0.
type explicitKeys struct {
	AdminSocket  any `yaml:"admin_socket"`
	PopSelection struct {
		SwitchMinGainPct any `yaml:"switch_min_gain_pct"`
		SwitchMinGainMs  any `yaml:"switch_min_gain_ms"`
		Holddown         any `yaml:"holddown"`
	} `yaml:"pop_selection"`
	Ingress []struct {
		OriginRequest struct {
			KeepaliveConnections any `yaml:"keepalive_connections"`
		} `yaml:"origin_request"`
	} `yaml:"ingress"`
}

// applyDefaults fills a file config from Default(). The catch-all rule is not
// added: a file must declare it.
func (c *Config) applyDefaults(set *explicitKeys) {
	d := Default()
	if c.HomePops == 0 {
		c.HomePops = d.HomePops
	}
	if c.SessionsPerPop == 0 {
		c.SessionsPerPop = d.SessionsPerPop
	}
	if c.PopSelection.Mode == "" {
		c.PopSelection.Mode = d.PopSelection.Mode
	}
	if c.PopSelection.ReprobeInterval == 0 {
		c.PopSelection.ReprobeInterval = d.PopSelection.ReprobeInterval
	}
	if set.PopSelection.SwitchMinGainPct == nil {
		c.PopSelection.SwitchMinGainPct = d.PopSelection.SwitchMinGainPct
	}
	if set.PopSelection.SwitchMinGainMs == nil {
		c.PopSelection.SwitchMinGainMs = d.PopSelection.SwitchMinGainMs
	}
	if set.PopSelection.Holddown == nil {
		c.PopSelection.Holddown = d.PopSelection.Holddown
	}
	if len(c.Ports) == 0 {
		c.Ports = d.Ports
	}
	if set.AdminSocket == nil {
		c.AdminSocket = d.AdminSocket
	}
	if c.MetricsAddr == nil {
		c.MetricsAddr = d.MetricsAddr
	}
	if c.Log.Format == "" {
		c.Log.Format = d.Log.Format
	}
	if c.Log.Level == "" {
		c.Log.Level = d.Log.Level
	}
	if c.Transport == "" {
		c.Transport = d.Transport
	}
	for i := range c.Ingress {
		keepaliveSet := i < len(set.Ingress) && set.Ingress[i].OriginRequest.KeepaliveConnections != nil
		c.Ingress[i].OriginRequest.applyDefaults(keepaliveSet)
	}
	c.Normalize()
}

func (o *OriginRequest) applyDefaults(keepaliveSet bool) {
	d := defaultOriginRequest()
	if o.ConnectTimeout == 0 {
		o.ConnectTimeout = d.ConnectTimeout
	}
	if o.ResponseHeaderTimeout == 0 {
		o.ResponseHeaderTimeout = d.ResponseHeaderTimeout
	}
	if !keepaliveSet {
		o.KeepaliveConnections = d.KeepaliveConnections
	}
}

// MetricsEnabled reports whether a metrics listener is configured
// (metrics_addr: "" disables it).
func (c *Config) MetricsEnabled() bool { return c.MetricsAddr != nil && *c.MetricsAddr != "" }

// AdminEnabled reports whether the admin socket is configured (admin_socket: "" disables it).
func (c *Config) AdminEnabled() bool { return c.AdminSocket != "" }

// Validate checks every field. credentials_file may be empty: a library
// caller can pass the credentials in memory.
func (c *Config) Validate() error { return c.validate(false) }

func (c *Config) validate(requireFile bool) error {
	var errs []error
	fail := func(format string, a ...any) { errs = append(errs, fmt.Errorf(format, a...)) }

	if !IsUUID(c.TunnelID) {
		fail("tunnel_id: must be a UUID")
	}
	if requireFile && c.CredentialsFile == "" {
		fail("credentials_file: required")
	}
	if c.HomePops < 1 || c.HomePops > 3 {
		fail("home_pops: must be 1..3 (got %d)", c.HomePops)
	}
	if c.SessionsPerPop < 1 || c.SessionsPerPop > 4 {
		fail("sessions_per_pop: must be 1..4 (got %d)", c.SessionsPerPop)
	}
	switch c.PopSelection.Mode {
	case ModeAuto:
		if len(c.PopSelection.Pops) > 0 {
			fail("pop_selection.pops: only valid with mode manual")
		}
	case ModeManual:
		if len(c.PopSelection.Pops) == 0 {
			fail("pop_selection.pops: required with mode manual")
		}
		seen := map[string]bool{}
		for _, p := range c.PopSelection.Pops {
			if !popCodeRe.MatchString(p) {
				fail("pop_selection.pops: %q is not a PoP code", p)
			}
			if seen[p] {
				fail("pop_selection.pops: duplicate %q", p)
			}
			seen[p] = true
		}
	default:
		fail("pop_selection.mode: must be auto or manual (got %q)", c.PopSelection.Mode)
	}
	if c.PopSelection.ReprobeInterval < time.Minute {
		fail("pop_selection.reprobe_interval: must be >= 1m")
	}
	if c.PopSelection.SwitchMinGainPct < 0 || c.PopSelection.SwitchMinGainPct > 100 {
		fail("pop_selection.switch_min_gain_pct: must be 0..100")
	}
	if c.PopSelection.SwitchMinGainMs < 0 {
		fail("pop_selection.switch_min_gain_ms: must be >= 0")
	}
	if h := c.PopSelection.Holddown; h < 0 || (h > 0 && h < time.Second) {
		fail("pop_selection.holddown: must be 0 or >= 1s (use a unit, e.g. 30m)")
	}
	if len(c.Ports) == 0 {
		fail("ports: at least one port is required")
	}
	seenPort := map[int]bool{}
	for _, p := range c.Ports {
		if p < 1 || p > 65535 {
			fail("ports: %d out of range", p)
		}
		// Only the production ports unless a private CA marks a test setup.
		if p != proto.PortDirect && p != proto.PortFallback && c.ServerCAFile == "" {
			fail("ports: %d is not 7844 or 443 (other ports need server_ca_file)", p)
		}
		if seenPort[p] {
			fail("ports: duplicate %d", p)
		}
		seenPort[p] = true
	}
	if err := validateIngress(c.Ingress); err != nil {
		errs = append(errs, err)
	}
	if c.MetricsEnabled() {
		host, _, err := net.SplitHostPort(*c.MetricsAddr)
		if err != nil {
			fail("metrics_addr: %v", err)
		} else if ip := net.ParseIP(host); ip == nil || !ip.IsLoopback() {
			fail("metrics_addr: must bind a loopback address (got %q)", host)
		}
	}
	switch c.Log.Format {
	case "json", "text":
	default:
		fail("log.format: must be json or text")
	}
	switch c.Log.Level {
	case "debug", "info", "warn", "error":
	default:
		fail("log.level: must be debug, info, warn or error")
	}
	for _, pin := range c.ServerSPKIPins {
		if _, err := DecodePin(pin); err != nil {
			fail("server_spki_pins: %v", err)
		}
	}
	for _, p := range c.AllowLocalPorts {
		if p < 1 || p > 65535 {
			fail("allow_local_ports: %d out of range", p)
		}
	}
	if c.Transport != "h2" {
		fail("transport: only h2 is implemented (got %q)", c.Transport)
	}
	return errors.Join(errs...)
}

// DecodePin accepts "sha256/<base64>" or bare base64 of a 32-byte SPKI digest.
func DecodePin(pin string) ([]byte, error) {
	s := strings.TrimPrefix(pin, "sha256/")
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("%q: not base64", pin)
	}
	if len(b) != 32 {
		return nil, fmt.Errorf("%q: sha256 digest must be 32 bytes", pin)
	}
	return b, nil
}

func validateIngress(rules []IngressRule) error {
	var errs []error
	fail := func(format string, a ...any) { errs = append(errs, fmt.Errorf("ingress: "+format, a...)) }
	if len(rules) == 0 {
		fail("at least the catch-all rule is required")
		return errors.Join(errs...)
	}
	seen := map[string]bool{}
	catchAll := 0
	for i, r := range rules {
		if r.Hostname == "" {
			catchAll++
			if !strings.HasPrefix(r.Service, ServiceStatusPrefix) {
				fail("rule %d: the catch-all rule (no hostname) must use http_status:<code>", i)
			}
			if i != len(rules)-1 {
				fail("rule %d: the catch-all rule must be last", i)
			}
		} else {
			if err := validateHostname(r.Hostname); err != nil {
				fail("rule %d: %v", i, err)
			}
			if seen[r.Hostname] {
				fail("rule %d: duplicate hostname %q", i, r.Hostname)
			}
			seen[r.Hostname] = true
		}
		if err := validateService(r.Service, r.OriginRequest); err != nil {
			fail("rule %d: %v", i, err)
		}
	}
	if catchAll == 0 {
		fail("a catch-all rule `- service: http_status:<code>` is required so an unknown Host is never proxied")
	}
	if catchAll > 1 {
		fail("only one catch-all rule is allowed")
	}
	return errors.Join(errs...)
}

func validateHostname(h string) error {
	rest := h
	if strings.HasPrefix(h, "*.") {
		rest = h[2:]
	} else if strings.Contains(h, "*") {
		return fmt.Errorf("hostname %q: wildcard only as a leading \"*.\"", h)
	}
	if rest == "" || len(h) > 253 {
		return fmt.Errorf("hostname %q: invalid", h)
	}
	for _, l := range strings.Split(rest, ".") {
		if !labelRe.MatchString(l) {
			return fmt.Errorf("hostname %q: invalid label %q", h, l)
		}
	}
	return nil
}

// ParseStatusService returns the code of an http_status:<code> service.
func ParseStatusService(s string) (int, bool) {
	if !strings.HasPrefix(s, ServiceStatusPrefix) {
		return 0, false
	}
	n, err := strconv.Atoi(strings.TrimPrefix(s, ServiceStatusPrefix))
	if err != nil || n < 200 || n > 599 {
		return 0, false
	}
	return n, true
}

func validateService(s string, o OriginRequest) error {
	if s == "" {
		return errors.New("service: required")
	}
	if strings.HasPrefix(s, ServiceStatusPrefix) {
		if _, ok := ParseStatusService(s); !ok {
			return fmt.Errorf("service %q: http_status code must be 200..599", s)
		}
		return nil
	}
	u, err := url.Parse(s)
	if err != nil {
		// url.Error repeats the raw URL, which can carry a password.
		var ue *url.Error
		if errors.As(err, &ue) {
			err = ue.Err
		}
		return fmt.Errorf("service: invalid URL: %v", err)
	}
	shown := u.Redacted()
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("service %q: scheme must be http, https or http_status", shown)
	}
	if u.Hostname() == "" {
		return fmt.Errorf("service %q: host required", shown)
	}
	if u.RawQuery != "" || u.Fragment != "" || u.User != nil {
		return fmt.Errorf("service %q: query, fragment and userinfo are not allowed", shown)
	}
	// Floors reject values below a sane minimum (yaml.v3 refuses bare integers for durations).
	if o.ConnectTimeout < 100*time.Millisecond || o.ConnectTimeout > time.Minute {
		return errors.New("origin_request.connect_timeout: must be in [100ms, 1m] (use a unit, e.g. 5s)")
	}
	if o.ResponseHeaderTimeout < time.Second || o.ResponseHeaderTimeout > 10*time.Minute {
		return errors.New("origin_request.response_header_timeout: must be in [1s, 10m] (use a unit, e.g. 45s)")
	}
	if o.KeepaliveConnections < 0 || o.KeepaliveConnections > 1024 {
		return errors.New("origin_request.keepalive_connections: must be 0..1024")
	}
	if u.Scheme == "http" && (o.CAFile != "" || o.TLSServerName != "" || o.NoTLSVerify) {
		return fmt.Errorf("service %q: tls options need an https service", s)
	}
	return nil
}
