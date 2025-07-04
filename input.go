package parsuri

import (
	"bufio"
	"os"
	"time"
)

// LoadSTDIN loads from stdin and parses it asynchronously.
// It does NOT call [Loader.Close] when finished, so [Loader.More] will return true.
func (l *Loader) LoadSTDIN() {
	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				l.addErr(err)
				time.Sleep(1 * time.Millisecond)
				continue
			}
			l.parseLine(line)
		}
	}()
}

// LoadFile loads a file, parses it, and closes it asynchronously.
// does NOT call [Loader.Close] when finished, [Loader.More] will return true until the [Loader] is explicitly closed.
func (l *Loader) LoadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	go func() {
		l.Scan(file)
		if err = file.Close(); err != nil {
			l.addErr(err)
		}
	}()
	return nil
}

// LoadOneFile loads a file, parses it, and closes it asynchronously.
// It also calls [Loader.Close] when finished, causing [Loader.More] to return false.
func (l *Loader) LoadOneFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	l.ParseAndCloseAsync(file)
	return nil
}
