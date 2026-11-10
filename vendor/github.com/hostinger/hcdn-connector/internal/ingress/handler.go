package ingress

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/hostinger/hcdn-connector/internal/metrics"
	"github.com/hostinger/hcdn-connector/internal/proto"
)

// Handler serves visitor requests already vetted by the session layer
// (control routes matched, hcdn-request-id present, inflight cap applied).
type Handler struct {
	Table    *Table
	TunnelID string
	Log      *slog.Logger
	Metrics  *metrics.Metrics
}

func NewHandler(t *Table, tunnelID string, log *slog.Logger, m *metrics.Metrics) *Handler {
	h := &Handler{Table: t, TunnelID: tunnelID, Log: log, Metrics: m}
	if log == nil {
		h.Log = slog.Default()
	}
	errLog := slog.NewLogLogger(h.Log.Handler(), slog.LevelWarn)
	for _, r := range t.Rules {
		r.Service.observe = h.observe
		r.Service.log = h.Log
		if r.Service.Proxy != nil {
			r.Service.Proxy.ErrorLog = errLog
		}
	}
	return h
}

func (h *Handler) observe(result string, rt time.Duration) {
	if h.Metrics == nil {
		return
	}
	h.Metrics.Requests.WithLabelValues(result).Inc()
	if result == "ok" {
		h.Metrics.ServiceRT.Observe(rt.Seconds())
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		h.serveConnect(w, r)
		return
	}
	rule := h.Table.Match(r.Host)
	if rule == nil {
		WriteError(w, proto.ErrBadGateway)
		return
	}
	if rule.Service.Status != 0 {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(rule.Service.Status)
		h.observe("http_status", 0)
		return
	}
	rule.Service.Proxy.ServeHTTP(w, WithStart(r, time.Now()))
}
