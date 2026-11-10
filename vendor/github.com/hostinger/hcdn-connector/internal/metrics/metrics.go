// Package metrics registers the connector's Prometheus metrics.
package metrics

import (
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	Registry     *prometheus.Registry
	Sessions     *prometheus.GaugeVec
	Requests     *prometheus.CounterVec
	ServiceRT    prometheus.Histogram
	CertDaysLeft prometheus.Gauge
	Reconnects   *prometheus.CounterVec
	BuildInfo    *prometheus.GaugeVec
}

func New() *Metrics {
	m := &Metrics{
		Registry: prometheus.NewRegistry(),
		Sessions: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "hcdn_connector_sessions", Help: "Sessions per PoP and state.",
		}, []string{"pop", "state"}),
		Requests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "hcdn_connector_requests_total", Help: "Visitor requests by result.",
		}, []string{"result"}),
		ServiceRT: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name: "hcdn_connector_service_rt_seconds", Help: "Time to first response byte from the service.",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10, 30, 45},
		}),
		CertDaysLeft: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "hcdn_connector_cert_days_left", Help: "Days until the connector certificate expires.",
		}),
		Reconnects: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "hcdn_connector_reconnects_total", Help: "Session ends by reason.",
		}, []string{"reason"}),
		BuildInfo: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "hcdn_connector_build_info", Help: "Build information.",
		}, []string{"version", "commit"}),
	}
	m.Registry.MustRegister(m.Sessions, m.Requests, m.ServiceRT, m.CertDaysLeft, m.Reconnects, m.BuildInfo,
		collectors.NewGoCollector(), collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	for _, r := range []string{"ok", "service_unreachable", "service_timeout", "bad_gateway", "overloaded", "rejected", "http_status", "ws", "canceled"} {
		m.Requests.WithLabelValues(r)
	}
	return m
}

func (m *Metrics) Handler() http.Handler {
	return promhttp.HandlerFor(m.Registry, promhttp.HandlerOpts{})
}

// Serve exposes /metrics on addr (validated as loopback by config).
func (m *Metrics) Serve(addr string) (*http.Server, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.Handle("/metrics", m.Handler())
	srv := &http.Server{Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	go func() { _ = srv.Serve(ln) }()
	return srv, nil
}
