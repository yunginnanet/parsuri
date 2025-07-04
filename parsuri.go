package parsuri

import (
	sanic "github.com/bytedance/sonic"
	"github.com/yunginnanet/parsuri/pkg/buffer"
	"github.com/yunginnanet/parsuri/pkg/events"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

var sonic = sanic.ConfigStd

func init() {
	var v events.EveEvent
	err := sanic.Pretouch(reflect.TypeOf(v))
	if err != nil {
		panic(err)
	}
}

// Loader is a struct that loads events from a file or stream asynchronously into a queue.
type Loader struct {
	queue   *buffer.Queue[events.EveEvent]
	errs    []error
	errMu   sync.Mutex
	running *atomic.Int32
}

func NewLoader() *Loader {
	running := &atomic.Int32{}
	running.Store(int32(statusRunning))
	return &Loader{
		queue:   buffer.NewQueue[events.EveEvent](100000),
		errs:    make([]error, 0),
		running: running,
	}
}

// More returns true if there are more events to process.
func (l *Loader) More() bool {
	if l.queue.Len() > 0 {
		return true
	}
	return l.isRunning()
}

// Close closes the loader and prevents further processing.
// This will cause [Loader.More] to return false.
func (l *Loader) Close() error {
	_ = l.running.CompareAndSwap(int32(statusRunning), int32(statusClosed))
	return nil
}

// Event removes and returns the next [events.EveEvent] from the queue.
func (l *Loader) Event() events.EveEvent {
	item, ok := l.queue.Pop()
	if ok {
		return item
	}

	for l.More() {
		time.Sleep(20 * time.Millisecond)
		item, ok = l.queue.Pop()
		if ok {
			return item
		}
	}

	return item
}
