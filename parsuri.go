package surevego

import (
	"bufio"
	"errors"
	"fmt"
	sanic "github.com/bytedance/sonic"
	"github.com/yunginnanet/parsuri/buffer"
	"github.com/yunginnanet/parsuri/events"
	"io"
	"os"
	"reflect"
	"sync"
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
	mu      sync.Mutex
	running sync.Mutex
}

func NewLoader() *Loader {
	return &Loader{
		queue: buffer.NewQueue[events.EveEvent](100000),
		errs:  make([]error, 0),
	}
}

// More returns true if there are more events to process.
func (l *Loader) More() bool {
	if l.queue.Len() > 0 {
		return true
	}
	if !l.running.TryLock() {
		return true
	}
	l.running.Unlock()
	return false
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

// Err clears the error slice and returns a single error.
func (l *Loader) Err() error {
	l.mu.Lock()
	err := errors.Join(l.errs...)
	clear(l.errs)
	l.mu.Unlock()
	return err
}

func (l *Loader) addErr(err error) {
	l.mu.Lock()
	l.errs = append(l.errs, err)
	l.mu.Unlock()
}

func (l *Loader) LoadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	return l.ParseReaderAsynchronously(file)
}

func (l *Loader) ParseReaderAsynchronously(r io.ReadCloser) error {

	l.running.Lock()

	go func() {
		var err error
		scanner := bufio.NewScanner(r)
		ct := 0

		for scanner.Scan() {
			if e := scanner.Err(); e != nil {
				l.addErr(e)
			}

			ct++

			ev := events.EveEvent{}

			if err = sonic.Unmarshal(scanner.Bytes(), &ev); err == nil {
				l.queue.Push(ev)
				continue
			}

			l.addErr(fmt.Errorf("error unmarshaling eve.json line %d: %w", ct, err))
		}

		if err = scanner.Err(); err != nil {
			l.addErr(err)
		}

		if err = r.Close(); err != nil {
			l.addErr(err)
		}

		l.running.Unlock()
	}()

	return nil
}
