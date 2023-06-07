package robin

import (
	"sync/atomic"
)

// Loadbalancer is a simple, generic and thread-safe round-robin load balancer for Go.
type Loadbalancer[T any] struct {
	Items []T

	idx uint64
}

// NewLoadbalancer creates a new Loadbalancer.
// It is guaranteed that two concurrent calls to Loadbalancer.Next will not return the same item, if the slice contains more than one item.
func NewLoadbalancer[T any](items []T) *Loadbalancer[T] {
	return &Loadbalancer[T]{
		Items: items,
	}
}

// Current returns the current item in the slice, without advancing the Loadbalancer.
func (l *Loadbalancer[T]) Current() T {
	idx := atomic.LoadUint64(&l.idx)
	return l.Items[idx%uint64(len(l.Items))]
}

// Next returns the next item in the slice. When the end of the slice is reached, it starts again from the beginning.
func (l *Loadbalancer[T]) Next() T {
	idx := atomic.AddUint64(&l.idx, 1) - 1
	return l.Items[idx%uint64(len(l.Items))]
}

// Reset resets the Loadbalancer to its initial state.
func (l *Loadbalancer[T]) Reset() {
	atomic.StoreUint64(&l.idx, 0)
}

// AddItems adds items to the Loadbalancer.
func (l *Loadbalancer[T]) AddItems(items ...T) {
	// This part is not thread-safe and should be called only from a single goroutine
	l.Items = append(l.Items, items...)
}
