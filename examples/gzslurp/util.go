package gzslurp

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var headers = &sync.Pool{New: func() any { return make([]byte, 2) }}

func getHeader() []byte {
	h := headers.Get().([]byte)
	clear(h)
	return h
}

func putHeader(h []byte) {
	headers.Put(h)
}

// twoByteErrCheck checks if the error is nil and n is not equal to 2. file is also  closed in this function.
func twoByteErrCheck(file io.Closer, n int, err error) error {
	if err == nil && n == 2 {
		panic("unreachable")
	}

	if err == nil {
		err = fmt.Errorf("expected to read 2 bytes but got %d", n)
	}

	return errors.Join(fmt.Errorf("error reading file header: %w", err), file.Close())
}
