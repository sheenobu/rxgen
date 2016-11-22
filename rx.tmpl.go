package rx

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	rx,
}

var rx = &typewriter.Template{
	Name: "rx",
	Text: `
// {{.Name}} is the reactive wrapper for {{.Type}}
type {{.Name}} struct {
	value {{.Type}}
	lock  sync.RWMutex
	handles     chan int
	subscribers []chan<- {{.Type}}
}

// New{{.Name}} creates a new reactive object for the initial value of {{.Type}}
func New{{.Name}}(v {{.Type}}) *{{.Name}}{
	return &{{.Name}}{
		value: v,
		handles: make(chan int, 10),
	}
}

// Get gets the {{.Type}}
func (rx *{{.Name}}) Get() {{.Type}} {
	rx.lock.RLock()
	defer rx.lock.RUnlock()
	return rx.value
}

// Set sets the {{.Name}} and notifies subscribers
func (rx *{{.Name}}) Set(v {{.Type}}) {
	rx.lock.Lock()
	defer rx.lock.Unlock()
	rx.value = v
	for _, s := range rx.subscribers {
		if s != nil {
			s <- v
		}
	}
}

// Subscribe subscribes to changes on the {{.Name}}
func (rx *{{.Name}}) Subscribe() *{{.Name}}Subscriber {
	c := make(chan {{.Type}})
	s := &{{.Name}}Subscriber{
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

// {{.Name}}Subscriber allows subscribing to the reactive {{.Type}}
type {{.Name}}Subscriber struct {
	C      <-chan {{.Type}}
	handle int
	parent *{{.Name}}
}

// Close closes the subscription
func (s *{{.Name}}Subscriber) Close() {
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

func (s *{{.Name}}) Bind(ctx context.Context, other *{{.Name}}) context.CancelFunc {
	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		s2 := other.Subscribe()
		defer s2.Close()
		for {
			select{
			case o := <-s2.C:
				s.Set(o)
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	return cancel
}

`,
}
