package ingress

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/hostinger/hcdn-connector/internal/fileguard"
	"github.com/hostinger/hcdn-connector/internal/proto"
	"github.com/hostinger/hcdn-connector/pkg/config"
	"github.com/hostinger/hcdn-connector/pkg/localguard"
)

// Service is the target of an ingress rule: an http(s) URL or a fixed status.
type Service struct {
	Raw       string
	URL       *url.URL
	Status    int
	Origin    config.OriginRequest
	Transport *http.Transport
	TLS       *tls.Config
	Proxy     *httputil.ReverseProxy
	wsTLS     *tls.Config // HTTP/1.1-only ALPN for the WebSocket Upgrade
	tunnelID  string
	observe   func(result string, rt time.Duration)
	log       *slog.Logger
	// control is the localguard rule for every dial to the service; nil: none.
	control func(network, address string, c syscall.RawConn) error
}

type ctxKey int

const startKey ctxKey = 1

// NewService builds the service of one rule. guard refuses every dial to a
// sensitive port on this machine; nil dials anything.
func NewService(raw string, o config.OriginRequest, tunnelID string, guard *localguard.Guard) (*Service, error) {
	s := &Service{Raw: raw, Origin: o, tunnelID: tunnelID, observe: func(string, time.Duration) {}, log: slog.Default()}
	if guard != nil {
		s.control = guard.Control
	}
	if code, ok := config.ParseStatusService(raw); ok {
		s.Status = code
		return s, nil
	}
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	s.URL = u
	if u.Scheme == "https" {
		s.TLS = &tls.Config{MinVersion: tls.VersionTLS12, ServerName: o.TLSServerName, InsecureSkipVerify: o.NoTLSVerify}
		if s.TLS.ServerName == "" {
			s.TLS.ServerName = u.Hostname()
		}
		if o.CAFile != "" {
			pem, err := fileguard.ReadFile(o.CAFile)
			if err != nil {
				return nil, fmt.Errorf("service %s: ca_file: %w", raw, err)
			}
			pool := x509.NewCertPool()
			if !pool.AppendCertsFromPEM(pem) {
				return nil, fmt.Errorf("service %s: ca_file %s: no certificates found", raw, o.CAFile)
			}
			s.TLS.RootCAs = pool
		}
		s.wsTLS = s.TLS.Clone()
		s.wsTLS.NextProtos = []string{"http/1.1"}
	}
	var transportTLS *tls.Config
	if s.TLS != nil {
		// http.Transport adds "h2" to NextProtos in place; a shared config would leak it into dialRaw.
		transportTLS = s.TLS.Clone()
	}
	dialer := s.Dialer()
	dialer.KeepAlive = 30 * time.Second
	s.Transport = &http.Transport{
		Proxy:                  nil,
		DialContext:            dialer.DialContext,
		TLSClientConfig:        transportTLS,
		ForceAttemptHTTP2:      o.HTTP2,
		MaxIdleConns:           o.KeepaliveConnections,
		MaxIdleConnsPerHost:    o.KeepaliveConnections,
		IdleConnTimeout:        90 * time.Second,
		TLSHandshakeTimeout:    o.ConnectTimeout,
		ResponseHeaderTimeout:  o.ResponseHeaderTimeout,
		ExpectContinueTimeout:  time.Second,
		MaxResponseHeaderBytes: proto.MaxHeaderBytes,
		DisableCompression:     true,
	}
	if o.KeepaliveConnections == 0 {
		s.Transport.DisableKeepAlives = true
	}
	s.Proxy = &httputil.ReverseProxy{
		Rewrite:        s.rewrite,
		Transport:      s.Transport,
		ModifyResponse: s.modifyResponse,
		ErrorHandler:   s.errorHandler,
		FlushInterval:  -1,
		ErrorLog:       nil,
	}
	return s, nil
}

// Dialer returns a dialer for the service: connect_timeout, and the guard
// of NewService at dial time, after DNS.
func (s *Service) Dialer() *net.Dialer {
	return &net.Dialer{Timeout: s.Origin.ConnectTimeout, Control: s.control}
}

// HostHeader returns the Host to send: the visitor Host unless http_host_header overrides it.
func (s *Service) HostHeader(visitorHost string) string {
	if s.Origin.HTTPHostHeader != "" {
		return s.Origin.HTTPHostHeader
	}
	return visitorHost
}

func (s *Service) rewrite(pr *httputil.ProxyRequest) {
	pr.SetURL(s.URL)
	pr.Out.Host = s.HostHeader(pr.In.Host)
	PrepareServiceHeaders(pr.Out.Header, pr.In.Header, s.tunnelID)
}

// PrepareServiceHeaders strips control headers, restores the edge's forwarding
// headers as received (ReverseProxy.Rewrite drops them) and adds Hcdn-Tunnel-Id.
// The edge never sets Forwarded or X-Forwarded-Host: a value there is the visitor's.
func PrepareServiceHeaders(out, in http.Header, tunnelID string) {
	StripControlHeaders(out)
	out.Del("Forwarded")
	out.Del("X-Forwarded-Host")
	for _, k := range []string{"X-Forwarded-For", "X-Forwarded-Proto"} {
		if v, ok := in[k]; ok {
			out[k] = append([]string(nil), v...)
		} else {
			delete(out, k)
		}
	}
	out.Set(proto.HdrTunnelID, tunnelID)
}

// StripControlHeaders removes every hcdn-* and x-hcdn-* header.
func StripControlHeaders(h http.Header) {
	for k := range h {
		if IsControlHeader(k) {
			delete(h, k)
		}
	}
}

func IsControlHeader(k string) bool {
	return hasPrefixFold(k, proto.HdrControlPrefix) || hasPrefixFold(k, proto.HdrXControlPrefix)
}

func hasPrefixFold(s, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[:len(prefix)], prefix)
}

func (s *Service) modifyResponse(res *http.Response) error {
	StripControlHeaders(res.Header)
	var rt time.Duration
	if start, ok := res.Request.Context().Value(startKey).(time.Time); ok {
		rt = time.Since(start)
	}
	res.Header.Set(proto.HdrConnectorRT, strconv.FormatInt(rt.Milliseconds(), 10))
	s.observe("ok", rt)
	return nil
}

func (s *Service) errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, context.Canceled) {
		s.observe("canceled", 0)
		return
	}
	reason := Classify(err)
	s.logBlocked(err)
	s.observe(reason, 0)
	WriteError(w, reason)
}

// logBlocked warns when the guard refused the dial: without the log line the
// edge shows only service_unreachable.
func (s *Service) logBlocked(err error) {
	var be *localguard.BlockedError
	if errors.As(err, &be) {
		s.log.Warn("refused a dial to a sensitive local port", "service", s.Raw, "addr", be.Addr, "err", be)
	}
}

// WriteError answers 502 + hcdn-error with an empty body; tunneld renders the 530 (docs/protocol.md, section 4).
func WriteError(w http.ResponseWriter, reason string) {
	h := w.Header()
	for k := range h {
		delete(h, k)
	}
	h.Set(proto.HdrError, reason)
	h.Set("Content-Length", "0")
	w.WriteHeader(http.StatusBadGateway)
}

// Classify maps a transport error to an hcdn-error value.
func Classify(err error) string {
	var certErr *tls.CertificateVerificationError
	var recErr tls.RecordHeaderError
	var unknownCA x509.UnknownAuthorityError
	var hostErr x509.HostnameError
	var invalid x509.CertificateInvalidError
	if errors.As(err, &certErr) || errors.As(err, &recErr) || IsTLSAlert(err) ||
		errors.As(err, &unknownCA) || errors.As(err, &hostErr) || errors.As(err, &invalid) {
		return proto.ErrServiceUnreachable
	}
	var op *net.OpError
	if errors.As(err, &op) && op.Op == "dial" {
		return proto.ErrServiceUnreachable
	}
	if strings.Contains(err.Error(), "TLS handshake") {
		return proto.ErrServiceUnreachable
	}
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return proto.ErrServiceTimeout
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return proto.ErrServiceTimeout
	}
	return proto.ErrBadGateway
}

// IsTLSAlert reports an alert sent by the peer. tls.AlertError matches only on
// QUIC; over TCP crypto/tls returns net.OpError{Op: "remote error"}.
func IsTLSAlert(err error) bool {
	var alert tls.AlertError
	if errors.As(err, &alert) {
		return true
	}
	var op *net.OpError
	return errors.As(err, &op) && op.Op == "remote error"
}

// WithStart stamps the request context for the hcdn-connector-rt computation.
func WithStart(r *http.Request, start time.Time) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), startKey, start))
}
