package parsuri

import "os"

// LoadSTDIN loads from stdin and parses it asynchronously.
// It does NOT call [Loader.Close] when finished, so [Loader.More] will return true.
func (l *Loader) LoadSTDIN() {
	l.ParseAsync(os.Stdin)
}

// LoadFile loads a file, parses it, and closes it asynchronously.
// It does NOT call [Loader.Close] when finished, so [Loader.More] will return true.
func (l *Loader) LoadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	go func() {
		l.Parse(file)
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
