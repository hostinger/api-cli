package session

import (
	"context"
	"sync"
	"time"

	"github.com/hostinger/hcdn-connector/internal/h2srv"
	"github.com/hostinger/hcdn-connector/internal/proto"
)

// Reach scopes beyond the edge's own (proto.VisibleScopePop and
// proto.VisibleScopeNode).
const (
	// ScopeSession: a session is ready, but no edge confirmed presence (an
	// edge without the visible feature, or the visible wait ran out).
	ScopeSession = "session"
)

// DefaultVisibleWait is how long WaitReachable waits for a visible signal
// after a session is ready on an edge that advertises the feature.
const DefaultVisibleWait = 3 * time.Second

// Reach says how the tunnel became reachable.
type Reach struct {
	Scope    string // proto.VisibleScopePop, proto.VisibleScopeNode or ScopeSession
	Pop      string // PoP of the session that decided it
	EdgeNode string
	At       time.Time
}

// reachLatch keeps the first event of each kind for the run. Supervisors
// come and go with reselection, so waiters watch the Manager, not them.
type reachLatch struct {
	mu       sync.Mutex
	visible  *Reach        // first visible signal
	plain    *Reach        // first ready session on an edge without the feature
	featured *Reach        // first ready session on an edge with the feature
	changed  chan struct{} // closed and replaced on every first event
}

func newReachLatch() *reachLatch { return &reachLatch{changed: make(chan struct{})} }

func (l *reachLatch) set(slot **Reach, r Reach) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if *slot != nil {
		return
	}
	*slot = &r
	close(l.changed)
	l.changed = make(chan struct{})
}

// noteHello records a ready session.
func (m *Manager) noteHello(pop string, h h2srv.Hello) {
	r := Reach{Scope: ScopeSession, Pop: pop, EdgeNode: h.EdgeNode, At: time.Now()}
	if h.HasEdgeFeature(proto.EdgeFeatureVisible) {
		m.reach.set(&m.reach.featured, r)
	} else {
		m.reach.set(&m.reach.plain, r)
	}
}

// noteVisible records an edge's visible signal.
func (m *Manager) noteVisible(pop string, v h2srv.Visible) {
	m.reach.set(&m.reach.visible, Reach{Scope: v.Scope, Pop: pop, EdgeNode: v.EdgeNode, At: v.At})
}

// WaitReachable blocks until the tunnel is reachable, at the first of: a
// visible signal on any session; a ready session on an edge without the
// visible feature; or visibleWait after the first ready session on an edge
// with it. It returns ctx.Err() when ctx ends first. It works before Run.
func (m *Manager) WaitReachable(ctx context.Context, visibleWait time.Duration) (Reach, error) {
	m.Init()
	if visibleWait <= 0 {
		visibleWait = DefaultVisibleWait
	}
	l := m.reach
	for {
		l.mu.Lock()
		visible, plain, featured, ch := l.visible, l.plain, l.featured, l.changed
		l.mu.Unlock()
		switch {
		case visible != nil:
			return *visible, nil
		case plain != nil:
			return *plain, nil
		}
		var timeout <-chan time.Time
		stop := func() bool { return false }
		if featured != nil {
			left := time.Until(featured.At.Add(visibleWait))
			if left <= 0 {
				m.Log.Warn("edge did not confirm presence", "pop", featured.Pop, "edge_node", featured.EdgeNode,
					"visible_wait", visibleWait.String())
				return *featured, nil
			}
			t := time.NewTimer(left)
			timeout, stop = t.C, t.Stop
		}
		select {
		case <-ch:
		case <-timeout:
		case <-ctx.Done():
			stop()
			return Reach{}, ctx.Err()
		}
		stop()
	}
}
