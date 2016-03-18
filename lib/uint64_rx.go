// generated by genrx -name RxUint64 -type uint64 .; DO NOT EDIT

package lib

import (
	"sync"
)

// RxUint64 is the reactive wrapper for uint64
type RxUint64 struct {
	value uint64
	lock  sync.RWMutex

	handles     chan int
	subscribers []chan<- uint64
}

// NewRxUint64 creates a new reactive object for the initial value of uint64
func NewRxUint64(v uint64) *RxUint64 {
	return &RxUint64{
		value:   v,
		handles: make(chan int, 1),
	}
}

// Get gets the uint64
func (rx *RxUint64) Get() uint64 {
	rx.lock.RLock()
	defer rx.lock.RUnlock()
	return rx.value
}

// Set sets the uint64 and notifies subscribers
func (rx *RxUint64) Set(v uint64) {
	rx.lock.Lock()
	defer rx.lock.Unlock()
	rx.value = v

	for _, s := range rx.subscribers {
		if s != nil {
			s <- v
		}
	}
}

// Subscribe subscribes to changes on the uint64
func (rx *RxUint64) Subscribe() *RxUint64Subscriber {

	c := make(chan uint64)

	s := &RxUint64Subscriber{
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

// RxUint64Subscriber allows subscribing to the reactive uint64
type RxUint64Subscriber struct {
	C      <-chan uint64
	handle int
	parent *RxUint64
}

// Close closes the subscription
func (s *RxUint64Subscriber) Close() {
	// remove from parent and close channel
	s.parent.lock.Lock()
	close(s.parent.subscribers[s.handle])
	s.parent.subscribers[s.handle] = nil
	s.parent.lock.Unlock()

	go func() {
		s.parent.handles <- s.handle
	}()
}
