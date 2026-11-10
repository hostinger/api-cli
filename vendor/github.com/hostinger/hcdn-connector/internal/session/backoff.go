// Package session owns PoP discovery, selection, dialing and the per-PoP
// session supervisors.
package session

import (
	"math/rand/v2"
	"sync"
	"time"
)

const (
	BackoffBase = time.Second
	BackoffMax  = 60 * time.Second
	// A session that lived longer than this resets the backoff.
	StableSession = 60 * time.Second
)

// Backoff is exponential 1 s -> 60 s with full jitter.
type Backoff struct {
	mu      sync.Mutex
	attempt int
	rnd     func() float64
}

func NewBackoff() *Backoff { return &Backoff{rnd: rand.Float64} }

// Next returns the wait before the next attempt and advances the schedule.
func (b *Backoff) Next() time.Duration {
	b.mu.Lock()
	defer b.mu.Unlock()
	capd := BackoffBase << min(b.attempt, 6)
	if capd > BackoffMax {
		capd = BackoffMax
	}
	if b.attempt < 30 {
		b.attempt++
	}
	return time.Duration(b.rnd() * float64(capd))
}

func (b *Backoff) Reset() {
	b.mu.Lock()
	b.attempt = 0
	b.mu.Unlock()
}

func (b *Backoff) Attempt() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.attempt
}
