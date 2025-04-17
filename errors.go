package parsuri

import "errors"

// Err clears the error slice and returns a single error.
func (l *Loader) Err() error {
	l.errMu.Lock()
	err := errors.Join(l.errs...)
	clear(l.errs)
	l.errMu.Unlock()
	return err
}

func (l *Loader) addErr(err error) {
	l.errMu.Lock()
	l.errs = append(l.errs, err)
	l.errMu.Unlock()
}
