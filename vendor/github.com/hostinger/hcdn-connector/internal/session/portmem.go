package session

import (
	"net/netip"
	"sync"
	"time"
)

const (
	PortMemoryTTL = 24 * time.Hour
	// Primary re-test after a fallback: PrimaryRetestMin first, doubling on every
	// failed re-test up to PrimaryRetestTime. A blocked 7844 costs one connect attempt
	// per interval; a transient failure returns to 7844 within a minute.
	PrimaryRetestMin  = time.Minute
	PrimaryRetestTime = 24 * time.Hour
)

type portEntry struct {
	working      int
	since        time.Time
	lastPrimaryT time.Time
	// primaryDown: the primary failed on every resolved address, so the fallback goes first.
	primaryDown bool
	// retest: current re-test interval; 0 until the first primary failure after a primary success.
	retest time.Duration
	failed map[netip.Addr]time.Time
}

func (e *portEntry) retestAfter() time.Duration {
	if e.retest == 0 {
		return PrimaryRetestMin
	}
	return e.retest
}

// markPrimaryDown records a primary failure on every address; a failed re-test
// (a failure after the interval elapsed) doubles the interval.
func (e *portEntry) markPrimaryDown(now time.Time) {
	switch {
	// Not !primaryDown: a partial re-test failure clears primaryDown without a primary success.
	case e.retest == 0:
		e.retest = PrimaryRetestMin
	case now.Sub(e.lastPrimaryT) >= e.retestAfter():
		e.retest = min(2*e.retestAfter(), PrimaryRetestTime)
	}
	e.primaryDown, e.lastPrimaryT, e.failed = true, now, nil
}

// PortMemory remembers the working port per endpoint for 24 h and re-tests the
// primary port on a growing interval while it is down.
type PortMemory struct {
	mu      sync.Mutex
	ports   []int
	entries map[string]*portEntry
	now     func() time.Time
}

func NewPortMemory(ports []int) *PortMemory {
	return &PortMemory{ports: ports, entries: map[string]*portEntry{}, now: time.Now}
}

// Primary is the configured first port (7844); a failure on it starts the re-test clock.
func (p *PortMemory) Primary() int {
	if len(p.ports) == 0 {
		return 0
	}
	return p.ports[0]
}

// Order returns the ports to try for endpoint, most likely first.
func (p *PortMemory) Order(endpoint string) []int {
	p.mu.Lock()
	defer p.mu.Unlock()
	out := append([]int(nil), p.ports...)
	e := p.entries[endpoint]
	if e == nil || len(out) < 2 || !e.primaryDown {
		return out
	}
	now := p.now()
	if now.Sub(e.since) > PortMemoryTTL || e.working == out[0] {
		return out
	}
	if now.Sub(e.lastPrimaryT) >= e.retestAfter() {
		return out
	}
	for i, port := range out {
		if port == e.working {
			out = append([]int{port}, append(out[:i:i], out[i+1:]...)...)
			break
		}
	}
	return out
}

func (p *PortMemory) entry(endpoint string) *portEntry {
	e := p.entries[endpoint]
	if e == nil {
		e = &portEntry{}
		p.entries[endpoint] = e
	}
	return e
}

// RecordSuccess notes that port worked for endpoint.
func (p *PortMemory) RecordSuccess(endpoint string, port int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	e := p.entry(endpoint)
	now := p.now()
	e.since = now
	e.working = port
	if len(p.ports) > 0 && port == p.ports[0] {
		e.lastPrimaryT = now
		e.primaryDown = false
		e.retest = 0
		e.failed = nil
	}
}

// RecordPrimaryFailure notes that the primary port failed on the failed
// addresses. The endpoint counts as down on the primary only when every
// resolved address failed; nil resolved means the whole endpoint was tried.
func (p *PortMemory) RecordPrimaryFailure(endpoint string, failed, resolved []netip.Addr) {
	p.mu.Lock()
	defer p.mu.Unlock()
	e := p.entry(endpoint)
	now := p.now()
	if resolved == nil {
		e.markPrimaryDown(now)
		return
	}
	// The failure set restarts empty once the primary is down: only a success
	// or the next re-test may move the endpoint back to the primary.
	if e.primaryDown && now.Sub(e.lastPrimaryT) < e.retestAfter() {
		return
	}
	for a, at := range e.failed {
		if now.Sub(at) >= PrimaryRetestTime {
			delete(e.failed, a)
		}
	}
	if e.failed == nil {
		e.failed = map[netip.Addr]time.Time{}
	}
	for _, a := range failed {
		e.failed[a] = now
	}
	for _, a := range resolved {
		if _, ok := e.failed[a]; !ok {
			e.primaryDown = false
			return
		}
	}
	e.markPrimaryDown(now)
}

// Working returns the remembered working port (0 if none).
func (p *PortMemory) Working(endpoint string) int {
	p.mu.Lock()
	defer p.mu.Unlock()
	if e := p.entries[endpoint]; e != nil && p.now().Sub(e.since) <= PortMemoryTTL {
		return e.working
	}
	return 0
}
