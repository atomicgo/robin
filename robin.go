package robin

import "sync"

// Loadbalancer is a simple, generic round-robin load balancer for Go.
type Loadbalancer[T any] struct {
	Items      []T
	ThreadSafe bool

	idx int
	mu  sync.Mutex
}

// NewLoadbalancer creates a new Loadbalancer.
// For maximum speed, this is not thread-safe. Use NewThreadSafeLoadbalancer if you need thread-safety.
// If two goroutines call Loadbalancer.Next at the exact same time, it can happen that they both return the same item.
func NewLoadbalancer[T any](items []T) *Loadbalancer[T] {
	return &Loadbalancer[T]{
		Items: items,
	}
}

// NewThreadSafeLoadbalancer creates a new Loadbalancer.
// This is thread-safe, but slower than NewLoadbalancer.
// It is guaranteed that two concurrent calls to Loadbalancer.Next will not return the same item, if the slice contains more than one item.
func NewThreadSafeLoadbalancer[T any](items []T) *Loadbalancer[T] {
	return &Loadbalancer[T]{
		Items:      items,
		ThreadSafe: true,
	}
}

// Current returns the current item in the slice, without advancing the Loadbalancer.
func (l *Loadbalancer[T]) Current() T {
	if l.ThreadSafe {
		l.mu.Lock()
		defer l.mu.Unlock()
	}
	return l.Items[l.idx]
}

// Next returns the next item in the slice. When the end of the slice is reached, it starts again from the beginning.
func (l *Loadbalancer[T]) Next() T {
	var item T
	if l.ThreadSafe {
		l.mu.Lock()
		item = l.Items[l.idx]
		l.idx = (l.idx + 1) % len(l.Items)
		l.mu.Unlock()
	} else {
		item = l.Items[l.idx]
		l.idx = (l.idx + 1) % len(l.Items)
	}
	return item
}

// Reset resets the Loadbalancer to its initial state.
func (l *Loadbalancer[T]) Reset() {
	if l.ThreadSafe {
		l.mu.Lock()
		defer l.mu.Unlock()
	}
	l.idx = 0
}

// AddItems adds items to the Loadbalancer.
func (l *Loadbalancer[T]) AddItems(items ...T) {
	if l.ThreadSafe {
		l.mu.Lock()
		defer l.mu.Unlock()
	}
	l.Items = append(l.Items, items...)
}
