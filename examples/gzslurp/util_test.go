package gzslurp

import (
	"testing"
)

func TestHeaderPool(t *testing.T) {
	h := getHeader()

	t.Run("capacity", func(t *testing.T) {
		if cap(h) != 2 {
			t.Errorf("expected header capacity to be 2, got %d", cap(h))
		}
	})

	t.Run("put", func(t *testing.T) {
		h[0] = 0x1f
		h[1] = 0x8b

		putHeader(h)
	})

	t.Run("get after put", func(t *testing.T) {
		h = getHeader()

		if h[0] != 0 || h[1] != 0 {
			t.Errorf("header values should be reset after putHeader")
		}

		if len(h) != 2 {
			t.Errorf("expected header length to be 2, got %d", len(h))
		}

		if cap(h) != 2 {
			t.Errorf("expected header capacity to be 2 after putHeader, got %d", cap(h))
		}
	})
}
