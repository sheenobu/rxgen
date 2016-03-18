// generated by genrx -name RxString -type string .; DO NOT EDIT

package lib

import (
	"sync"
)

// RxString is the reactive wrapper for string
type RxString struct {
	value string
	lock  sync.RWMutex

	handles     chan int
	subscribers []chan<- string
}

// NewRxString creates a new reactive object for the initial value of string
func NewRxString(v string) *RxString {
	return &RxString{
		value:   v,
		handles: make(chan int, 1),
	}
}

// Get gets the string
func (rx *RxString) Get() string {
	rx.lock.RLock()
	defer rx.lock.RUnlock()
	return rx.value
}

// Set sets the string and notifies subscribers
func (rx *RxString) Set(v string) {
	rx.lock.Lock()
	defer rx.lock.Unlock()
	rx.value = v

	for _, s := range rx.subscribers {
		if s != nil {
			s <- v
		}
	}
}

// Subscribe subscribes to changes on the string
func (rx *RxString) Subscribe() *RxStringSubscriber {

	c := make(chan string)

	s := &RxStringSubscriber{
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

// RxStringSubscriber allows subscribing to the reactive string
type RxStringSubscriber struct {
	C      <-chan string
	handle int
	parent *RxString
}

// Close closes the subscription
func (s *RxStringSubscriber) Close() {
	// remove from parent and close channel
	s.parent.lock.Lock()
	close(s.parent.subscribers[s.handle])
	s.parent.subscribers[s.handle] = nil
	s.parent.lock.Unlock()

	go func() {
		s.parent.handles <- s.handle
	}()
}
