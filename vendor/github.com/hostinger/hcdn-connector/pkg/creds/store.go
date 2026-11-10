package creds

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// RefreshFunc fetches new credentials, for example by a new CSR enrollment.
type RefreshFunc func(ctx context.Context) (*Credentials, error)

// Store serves the current credentials. A file store reloads them on a
// signal or an mtime change; a memory store changes only through Replace or Refresh.
type Store struct {
	path     string // "" for a memory store
	tunnelID string
	refresh  RefreshFunc
	log      atomic.Pointer[slog.Logger]
	cur      atomic.Pointer[Credentials]
	gen      atomic.Uint64
	reloadMu sync.Mutex // concurrent reloads must not store an older certificate last
	mu       sync.Mutex
	nextID   uint64
	onChange []hook
	gauge    hook
	hint     string
	warnDays int
	now      func() time.Time
}

type hook struct {
	id uint64
	f  func(*Credentials)
}

// NewStore loads credentials.json from path.
func NewStore(path, tunnelID string, log *slog.Logger) (*Store, error) {
	c, err := Load(path, tunnelID)
	if err != nil {
		return nil, err
	}
	if err := checkValidity(c, time.Now()); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	s := &Store{hint: "issue a new one in hPanel", path: path, tunnelID: c.TunnelID, warnDays: WarnDaysLeft, now: time.Now}
	s.SetLogger(log)
	s.cur.Store(c)
	return s, nil
}

// NewMemoryStore serves c without a file. refresh runs when the edge rejects
// the certificate serial (unknown_serial); with a nil refresh that goaway
// stops the connector with exit code 78, the same as a file store without a new file.
// connector.New replaces log with the connector logger.
func NewMemoryStore(c *Credentials, log *slog.Logger, refresh RefreshFunc) (*Store, error) {
	if err := checkValidity(c, time.Now()); err != nil {
		return nil, err
	}
	s := &Store{tunnelID: c.TunnelID, refresh: refresh, warnDays: WarnDaysLeft, now: time.Now}
	s.SetLogger(log)
	s.cur.Store(c)
	return s, nil
}

// SetLogger sets the logger for reloads and expiry warnings; nil: slog.Default().
func (s *Store) SetLogger(log *slog.Logger) {
	if log == nil {
		log = slog.Default()
	}
	s.log.Store(log)
}

func (s *Store) logger() *slog.Logger { return s.log.Load() }

// SetExpiryHint sets the advice at the end of the "expires soon" warning.
// A file store starts with "issue a new one in hPanel", a memory store with "".
func (s *Store) SetExpiryHint(h string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.hint = h
}

func (s *Store) ExpiryHint() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.hint
}

// SetWarnDays sets the "expires soon" threshold in days (default WarnDaysLeft).
// 0 turns the warning off, for example for quick tunnels that live 24 h.
func (s *Store) SetWarnDays(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.warnDays = n
}

func (s *Store) warnDaysLeft() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.warnDays
}

// checkValidity refuses a certificate outside its validity window: at start
// every handshake would fail, on reload it would replace a working certificate.
func checkValidity(c *Credentials, now time.Time) error {
	if err := checkParsed(c); err != nil {
		return err
	}
	if now.After(c.Leaf.NotAfter) {
		return fmt.Errorf("certificate serial %s expired at %s", c.Serial, c.Leaf.NotAfter.UTC().Format(time.RFC3339))
	}
	if now.Add(clockSkew).Before(c.Leaf.NotBefore) {
		return fmt.Errorf("certificate serial %s is not valid before %s", c.Serial, c.Leaf.NotBefore.UTC().Format(time.RFC3339))
	}
	return nil
}

func checkParsed(c *Credentials) error {
	if c == nil {
		return errors.New("credentials: nil")
	}
	if c.Leaf == nil {
		return errors.New("credentials: no parsed certificate; build them with FromPEM, Parse or Load")
	}
	return nil
}

func (s *Store) Current() *Credentials { return s.cur.Load() }
func (s *Store) Generation() uint64    { return s.gen.Load() }
func (s *Store) TunnelID() string      { return s.tunnelID }

// Path is the credentials file, or "" for a memory store.
func (s *Store) Path() string { return s.path }

// OnChange registers a callback for successful reloads that changed the
// certificate; the returned func removes it. f runs under the reload lock:
// it must not call Reload, Replace or Refresh.
func (s *Store) OnChange(f func(*Credentials)) (remove func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextID++
	id := s.nextID
	s.onChange = append(s.onChange, hook{id: id, f: f})
	return func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		for i, h := range s.onChange {
			if h.id == id {
				s.onChange = append(s.onChange[:i:i], s.onChange[i+1:]...)
				return
			}
		}
	}
}

// SetGauge wires the hcdn_connector_cert_days_left setter; a later call
// replaces it. The returned func removes f if it is still the current setter.
func (s *Store) SetGauge(f func(float64)) (remove func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextID++
	id := s.nextID
	s.gauge = hook{id: id, f: func(c *Credentials) { f(c.DaysLeft(time.Now())) }}
	s.gauge.f(s.Current())
	return func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		if s.gauge.id == id {
			s.gauge = hook{}
		}
	}
}

// Reload re-reads the file. It returns whether the certificate changed.
// On a memory store it returns (false, nil).
func (s *Store) Reload() (changed bool, err error) {
	if s.path == "" {
		return false, nil
	}
	s.reloadMu.Lock()
	defer s.reloadMu.Unlock()
	c, err := Load(s.path, s.tunnelID)
	if err != nil {
		return false, err
	}
	return s.swap(c, "credentials reloaded")
}

// Replace installs c on a memory store, for example after a scheduled
// re-enrollment. It refuses a file store (the next poll would load the file
// again), another tunnel and a certificate outside its validity window.
func (s *Store) Replace(c *Credentials) (changed bool, err error) {
	s.reloadMu.Lock()
	defer s.reloadMu.Unlock()
	return s.replaceLocked(c, "credentials replaced")
}

func (s *Store) replaceLocked(c *Credentials, msg string) (bool, error) {
	if s.path != "" {
		return false, errors.New("credentials: Replace needs a memory store; update " + s.path + " instead")
	}
	if err := checkParsed(c); err != nil {
		return false, err
	}
	if c.TunnelID != s.tunnelID {
		return false, ErrTunnelIDMismatch
	}
	return s.swap(c, msg)
}

// CanRefresh reports whether Refresh can fetch new credentials.
func (s *Store) CanRefresh() bool { return s.refresh != nil }

// Refresh calls the RefreshFunc and installs the result. It fails when the
// result is the current certificate: the edge already rejected that one.
// It succeeds without a change when a Replace installed new credentials meanwhile.
func (s *Store) Refresh(ctx context.Context) error {
	if s.refresh == nil {
		return errors.New("credentials: no refresh function")
	}
	gen := s.Generation()
	c, err := s.refresh(ctx)
	if err != nil {
		return err
	}
	s.reloadMu.Lock()
	defer s.reloadMu.Unlock()
	// The caller gave up (timeout, shutdown): it must not see a late rotation.
	if err := ctx.Err(); err != nil {
		return err
	}
	if s.Generation() != gen {
		return nil
	}
	changed, err := s.replaceLocked(c, "credentials refreshed")
	if err != nil {
		return err
	}
	if !changed {
		return errors.New("credentials refresh returned the current certificate")
	}
	return nil
}

// swap needs reloadMu.
func (s *Store) swap(c *Credentials, msg string) (bool, error) {
	old := s.cur.Load()
	changed := old == nil || old.Serial != c.Serial || string(old.Leaf.Raw) != string(c.Leaf.Raw)
	if changed {
		if err := checkValidity(c, s.now()); err != nil {
			return false, fmt.Errorf("%w; keeping the current one", err)
		}
	}
	s.cur.Store(c)
	if changed {
		s.gen.Add(1)
		s.mu.Lock()
		cbs := append([]hook{}, s.onChange...)
		if s.gauge.f != nil {
			s.gauge.f(c)
		}
		s.mu.Unlock()
		for _, h := range cbs {
			h.f(c)
		}
		s.logger().Info(msg, "serial", c.Serial, "days_left", int(c.DaysLeft(time.Now())))
	}
	return changed, nil
}

// GetClientCertificate is the tls.Config hook: every handshake uses the current cert.
func (s *Store) GetClientCertificate(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	c := s.cur.Load()
	if c == nil {
		return nil, errors.New("credentials not loaded")
	}
	return &c.Cert, nil
}

// Watch warns daily when the certificate expires in < SetWarnDays days and refreshes
// the gauge hourly. A file store also reloads on reload and on an mtime change.
func (s *Store) Watch(stop <-chan struct{}, reload <-chan os.Signal) {
	var pollC <-chan time.Time
	if s.path != "" {
		poll := time.NewTicker(pollInterval)
		defer poll.Stop()
		pollC = poll.C
	} else {
		reload = nil
	}
	gauge := time.NewTicker(gaugeInterval)
	defer gauge.Stop()
	warn := time.NewTicker(warnInterval)
	defer warn.Stop()
	s.warnIfExpiring()
	s.tickGauge()
	var lastErr string
	for {
		select {
		case <-stop:
			return
		case <-reload:
			if _, err := s.Reload(); err != nil {
				s.logger().Error("credentials reload failed", "err", err)
			}
		case <-pollC:
			lastErr = s.pollOnce(lastErr)
		case <-gauge.C:
			s.tickGauge()
		case <-warn.C:
			s.warnIfExpiring()
		}
	}
}

// The gauge writes happen under mu, so this tick cannot overwrite a newer swap.
func (s *Store) tickGauge() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.gauge.f != nil {
		s.gauge.f(s.Current())
	}
}

// pollOnce retries a failed file every poll: chmod, chown or the clock can fix
// it without an mtime change. It logs an error only when it differs from lastErr.
func (s *Store) pollOnce(lastErr string) string {
	st, err := os.Stat(s.path)
	if err != nil {
		return lastErr
	}
	if cur := s.cur.Load(); cur == nil || st.ModTime().Equal(cur.ModTime) {
		return lastErr
	}
	if _, err := s.Reload(); err != nil {
		if err.Error() != lastErr {
			s.logger().Error("credentials reload failed", "err", err)
		}
		return err.Error()
	}
	return ""
}

func (s *Store) warnIfExpiring() {
	c := s.Current()
	d := c.DaysLeft(time.Now())
	if d < float64(s.warnDaysLeft()) {
		msg := "connector certificate expires soon"
		if h := s.ExpiryHint(); h != "" {
			msg += "; " + h
		}
		s.logger().Warn(msg, "days_left", int(d), "not_after", c.Leaf.NotAfter.UTC().Format(time.RFC3339))
	}
}
