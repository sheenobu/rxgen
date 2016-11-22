// Generated by: main
// TypeWriter: rx
// Directive: +gen on Int32_

package rx

import (
	"context"
	sync "sync"
)

// Int32 is the reactive wrapper for int32
type Int32 struct {
	value       int32
	lock        sync.RWMutex
	handles     chan int
	subscribers []chan<- int32
}

// NewInt32 creates a new reactive object for the initial value of int32
func NewInt32(v int32) *Int32 {
	return &Int32{
		value:   v,
		handles: make(chan int, 10),
	}
}

// Get gets the int32
func (rx *Int32) Get() int32 {
	rx.lock.RLock()
	defer rx.lock.RUnlock()
	return rx.value
}

// Set sets the int32 and notifies subscribers
func (rx *Int32) Set(v int32) {
	rx.lock.Lock()
	defer rx.lock.Unlock()
	rx.value = v
	for _, s := range rx.subscribers {
		if s != nil {
			s <- v
		}
	}
}

// Subscribe subscribes to changes on the int32
func (rx *Int32) Subscribe() *Int32Subscriber {
	c := make(chan int32)
	s := &Int32Subscriber{
		C:      c,
		parent: rx,
	}
	rx.lock.Lock()
	select {
	case handle := <-rx.handles:
		s.handle = handle
		rx.subscribers[handle] = c
	default:
		rx.subscribers = append(rx.subscribers, c)
		s.handle = len(rx.subscribers) - 1
	}
	rx.lock.Unlock()
	return s
}

// Int32Subscriber allows subscribing to the reactive int32
type Int32Subscriber struct {
	C      <-chan int32
	handle int
	parent *Int32
}

// Close closes the subscription
func (s *Int32Subscriber) Close() {
	// remove from parent and close channel
	s.parent.lock.Lock()
	close(s.parent.subscribers[s.handle])
	s.parent.subscribers[s.handle] = nil
	s.parent.lock.Unlock()
	go func() {
		select {
		case s.parent.handles <- s.handle:
		default:
		}
	}()
}

// Bind applies all changes to the 'other' Int32 to this name and returns a CancelFunc for
// closure
func (s *Int32) Bind(ctx context.Context, other *Int32) context.CancelFunc {
	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		s2 := other.Subscribe()
		defer s2.Close()
		for {
			select {
			case o := <-s2.C:
				s.Set(o)
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	return cancel
}
