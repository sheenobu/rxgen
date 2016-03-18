// generated by genrx -name RxInt32 -type int32 .; DO NOT EDIT

package lib

import (
	"sync"
)

// RxInt32 is the reactive wrapper for int32
type RxInt32 struct {
	value int32
	lock  sync.RWMutex

	handles     chan int
	subscribers []chan<- int32
}

func NewRxInt32(v int32) *RxInt32 {
	return &RxInt32{
		value:   v,
		handles: make(chan int, 1),
	}
}

// Get gets the int32
func (rx *RxInt32) Get() int32 {
	rx.lock.RLock()
	defer rx.lock.RUnlock()
	return rx.value
}

// Set sets the int32 and notifies subscribers
func (rx *RxInt32) Set(v int32) {
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
func (rx *RxInt32) Subscribe() *RxInt32Subscriber {

	c := make(chan int32)

	s := &RxInt32Subscriber{
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

// RxInt32Subscriber allows subscribing to the reactive int32
type RxInt32Subscriber struct {
	C      <-chan int32
	handle int
	parent *RxInt32
}

// Close closes the subscription
func (s *RxInt32Subscriber) Close() {
	// remove from parent and close channel
	s.parent.lock.Lock()
	close(s.parent.subscribers[s.handle])
	s.parent.subscribers[s.handle] = nil
	s.parent.lock.Unlock()

	go func() {
		s.parent.handles <- s.handle
	}()
}