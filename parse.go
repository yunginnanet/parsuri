package parsuri

import (
	"bufio"
	"fmt"
	"github.com/yunginnanet/parsuri/pkg/events"
	"io"
)

func (l *Loader) parseLine(dat []byte) (shouldContinue bool) {

	if !l.isRunning() {
		return false
	}

	ev := events.EveEvent{}

	err := sonic.Unmarshal(dat, &ev)
	if err == nil {
		l.queue.Push(ev)

		return true
	}

	if len(string(dat)) > 0 {
		l.addErr(fmt.Errorf("error unmarshaling eve.json line: %w\n---\n\t%s\n---", err, string(dat)))
		return true
	}

	l.addErr(fmt.Errorf("error unmarshaling eve.json line: %w", err))

	return true
}

// Scan parses the input stream synchronously.
func (l *Loader) Scan(r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if e := scanner.Err(); e != nil {
			l.addErr(e)
		}
		if !l.parseLine(scanner.Bytes()) {
			break
		}
	}
}

// ParseAndCloseAsync parses the input stream and closes it asynchronously.
// It also calls [Loader.Close] when finished, causing [Loader.More] to return false.
func (l *Loader) ParseAndCloseAsync(r io.ReadCloser) {
	go func() {
		l.Scan(r)

		if err := r.Close(); err != nil {
			l.addErr(err)
		}

		_ = l.Close()
	}()
}
