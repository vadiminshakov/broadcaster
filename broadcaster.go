package broadcaster

import (
	"sync"
)

type Broadcaster struct {
	mu       *sync.Mutex
	cond     *sync.Cond
	signaled bool
}

func NewBroadcaster() *Broadcaster {
	var mu sync.Mutex
	return &Broadcaster{
		mu:       &mu,
		cond:     sync.NewCond(&mu),
		signaled: false,
	}
}

func (b *Broadcaster) Wait(fn func()) {
	go func() {
		b.cond.L.Lock()
		defer b.cond.L.Unlock()

		for !b.signaled {
			b.cond.Wait()
		}
		fn()
	}()
}

func (b *Broadcaster) Broadcast() {
	b.cond.L.Lock()
	b.signaled = true
	b.cond.L.Unlock()

	b.cond.Broadcast()
}
