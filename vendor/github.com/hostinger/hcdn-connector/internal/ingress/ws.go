package ingress

import (
	"bufio"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hostinger/hcdn-connector/internal/proto"
)

var hopByHop = []string{"Connection", "Keep-Alive", "Proxy-Authenticate", "Proxy-Authorization", "Te", "Trailer", "Transfer-Encoding", "Upgrade"}

func removeHopByHop(h http.Header) {
	for _, f := range h.Values("Connection") {
		for _, sf := range strings.Split(f, ",") {
			if sf = strings.TrimSpace(sf); sf != "" {
				h.Del(sf)
			}
		}
	}
	for _, k := range hopByHop {
		h.Del(k)
	}
}

// serveConnect bridges a CONNECT + hcdn-upgrade: websocket stream to an
// HTTP/1.1 Upgrade against the service (docs/protocol.md, section 5).
func (h *Handler) serveConnect(w http.ResponseWriter, r *http.Request) {
	if !strings.EqualFold(r.Header.Get(proto.HdrUpgrade), "websocket") {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusNotImplemented)
		h.observe("rejected", 0)
		return
	}
	rule := h.Table.Match(r.Host)
	if rule == nil {
		WriteError(w, proto.ErrBadGateway)
		return
	}
	svc := rule.Service
	if svc.Status != 0 {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(svc.Status)
		h.observe("http_status", 0)
		return
	}
	target, err := wsTarget(svc.URL, r.Header.Get(proto.HdrPath))
	if err != nil {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusBadRequest)
		h.observe("rejected", 0)
		return
	}
	ctx := r.Context()
	start := time.Now()
	conn, err := svc.dialRaw(ctx)
	if err != nil {
		svc.logBlocked(err)
		h.fail(ctx, w, Classify(err))
		return
	}
	defer conn.Close()
	// An edge RST_STREAM cancels ctx; close the service conn so no copy below can stall.
	finished := make(chan struct{})
	defer close(finished)
	go func() {
		select {
		case <-ctx.Done():
			conn.Close()
		case <-finished:
		}
	}()

	req := &http.Request{
		Method:     http.MethodGet,
		URL:        target,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     r.Header.Clone(),
		Host:       svc.HostHeader(r.Host),
	}
	removeHopByHop(req.Header)
	PrepareServiceHeaders(req.Header, r.Header, h.TunnelID)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")
	// Extensions are not echoed to the visitor (no hcdn-ws-extensions yet), so
	// the service must not negotiate one (permessage-deflate would corrupt frames).
	req.Header.Del("Sec-WebSocket-Extensions")
	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", "") // suppress Go's default
	}

	_ = conn.SetDeadline(time.Now().Add(svc.Origin.ResponseHeaderTimeout))
	if err := req.Write(conn); err != nil {
		h.fail(ctx, w, Classify(err))
		return
	}
	// The limit covers the response head (with any 1xx) plus bufio read-ahead;
	// it lifts before the body or the WebSocket bytes flow.
	head := &io.LimitedReader{R: conn, N: proto.MaxHeaderBytes}
	br := bufio.NewReaderSize(head, 32<<10)
	res, err := readUpgradeResponse(br, req)
	if err != nil {
		reason := proto.ErrBadGateway
		if ne, ok := err.(net.Error); ok && ne.Timeout() {
			reason = proto.ErrServiceTimeout
		}
		h.fail(ctx, w, reason)
		return
	}
	head.N = math.MaxInt64
	_ = conn.SetDeadline(time.Time{})
	rt := time.Since(start)

	if res.StatusCode != http.StatusSwitchingProtocols {
		defer res.Body.Close()
		StripControlHeaders(res.Header)
		removeHopByHop(res.Header)
		out := w.Header()
		for k, v := range res.Header {
			out[k] = v
		}
		out.Set(proto.HdrConnectorRT, strconv.FormatInt(rt.Milliseconds(), 10))
		w.WriteHeader(res.StatusCode)
		_, _ = io.Copy(w, res.Body)
		h.observe("ok", rt)
		return
	}

	// tunneld treats 200 + hcdn-ws-accept as success; a 101 without it would open an unkeyed socket.
	accept := res.Header.Get("Sec-WebSocket-Accept")
	if accept == "" || !strings.EqualFold(res.Header.Get("Upgrade"), "websocket") {
		h.observe(proto.ErrBadGateway, rt)
		WriteError(w, proto.ErrBadGateway)
		return
	}
	out := w.Header()
	out.Set(proto.HdrWSAccept, accept)
	if p := res.Header.Get("Sec-WebSocket-Protocol"); p != "" {
		out.Set(proto.HdrWSProtocol, p)
	}
	out.Set(proto.HdrConnectorRT, strconv.FormatInt(rt.Milliseconds(), 10))
	w.WriteHeader(http.StatusOK)
	fw := &flushWriter{w: w}
	fw.Flush()
	h.observe("ws", rt)

	go func() {
		_, err := io.Copy(conn, r.Body)
		if cw, ok := conn.(interface{ CloseWrite() error }); ok && err == nil {
			_ = cw.CloseWrite()
		} else {
			conn.Close()
		}
	}()
	_, _ = io.Copy(fw, br)
}

// fail reports an edge cancel (RST_STREAM) as canceled, not as a service error.
func (h *Handler) fail(ctx context.Context, w http.ResponseWriter, reason string) {
	if ctx.Err() != nil {
		h.observe("canceled", 0)
		return
	}
	h.observe(reason, 0)
	WriteError(w, reason)
}

// readUpgradeResponse skips interim 1xx responses (103 Early Hints) before the 101 or final status.
func readUpgradeResponse(br *bufio.Reader, req *http.Request) (*http.Response, error) {
	for {
		res, err := http.ReadResponse(br, req)
		if err != nil || res.StatusCode < 100 || res.StatusCode >= 200 || res.StatusCode == http.StatusSwitchingProtocols {
			return res, err
		}
	}
}

// wsTarget joins the service path prefix and the byte-exact hcdn-path (RawPath
// kept, so %2F survives).
func wsTarget(svc *url.URL, path string) (*url.URL, error) {
	if path == "" {
		path = "/"
	}
	if !strings.HasPrefix(path, "/") {
		return nil, errors.New("hcdn-path must be absolute")
	}
	target, err := url.ParseRequestURI(path)
	if err != nil {
		return nil, err
	}
	if p := svc.EscapedPath(); p != "" && p != "/" {
		prefix := strings.TrimSuffix(p, "/")
		target.RawPath = prefix + target.EscapedPath()
		target.Path = strings.TrimSuffix(svc.Path, "/") + target.Path
	}
	return target, nil
}

func (s *Service) dialRaw(ctx context.Context) (net.Conn, error) {
	host := s.URL.Host
	if s.URL.Port() == "" {
		if s.URL.Scheme == "https" {
			host = net.JoinHostPort(s.URL.Hostname(), "443")
		} else {
			host = net.JoinHostPort(s.URL.Hostname(), "80")
		}
	}
	d := s.Dialer()
	if s.URL.Scheme == "https" {
		return (&tls.Dialer{NetDialer: d, Config: s.wsTLS}).DialContext(ctx, "tcp", host)
	}
	return d.DialContext(ctx, "tcp", host)
}

type flushWriter struct{ w http.ResponseWriter }

func (f *flushWriter) Write(p []byte) (int, error) {
	n, err := f.w.Write(p)
	f.Flush()
	return n, err
}

func (f *flushWriter) Flush() {
	if fl, ok := f.w.(http.Flusher); ok {
		fl.Flush()
	}
}
