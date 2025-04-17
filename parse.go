package parsuri

import (
	"bufio"
	"fmt"
	"github.com/yunginnanet/parsuri/events"
	"io"
)

func (l *Loader) parseLine(scanner *bufio.Scanner) (shouldContinue bool) {
	if e := scanner.Err(); e != nil {
		l.addErr(e)
	}

	if !l.isRunning() {
		return false
	}

	ev := events.EveEvent{}

	err := sonic.Unmarshal(scanner.Bytes(), &ev)
	if err == nil {
		l.queue.Push(ev)

		return true
	}

	l.addErr(fmt.Errorf("error unmarshaling eve.json line: %w", err))

	return true
}

// Parse parses the input stream synchronously.
func (l *Loader) Parse(r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if !l.parseLine(scanner) {
			break
		}
	}
}

// ParseAsync parses the input stream asynchronously.
func (l *Loader) ParseAsync(r io.Reader) {
	go l.Parse(r)
}

// ParseAndCloseAsync parses the input stream and closes it asynchronously.
// It also calls [Loader.Close] when finished, causing [Loader.More] to return false.
func (l *Loader) ParseAndCloseAsync(r io.ReadCloser) {
	go func() {
		l.Parse(r)

		if err := r.Close(); err != nil {
			l.addErr(err)
		}

		_ = l.Close()
	}()
}
