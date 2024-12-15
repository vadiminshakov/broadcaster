// Package broadcaster provides some concurrency
// tools to block and unblock execution across
// many goroutines based on broadcasting a signal.
package broadcaster

import (
	"sync"
)

type Broadcaster struct {
	mu       *sync.Mutex
	cond     *sync.Cond
	signaled bool
}

// NewBroadcaster creates a new broadcaster.
func NewBroadcaster() *Broadcaster {
	var mu sync.Mutex
	return &Broadcaster{
		mu:       &mu,
		cond:     sync.NewCond(&mu),
		signaled: false,
	}
}

// Go waits until something is broadcasted,
// and runs the given function in a new
// goroutine.
func (b *Broadcaster) Go(fn func()) {
	go func() {
		b.cond.L.Lock()
		defer b.cond.L.Unlock()

		for !b.signaled {
			b.cond.Wait()
		}
		fn()
	}()
}

// Broadcast broadcasts a signal to all
// waiting function and unblocks them.
func (b *Broadcaster) Broadcast() {
	b.cond.L.Lock()
	b.signaled = true
	b.cond.L.Unlock()

	b.cond.Broadcast()
}
