package events

import "testing"

func TestTCPEmpty(t *testing.T) {
	t.Run("TCPReturnsTrueForEmptyValues", func(t *testing.T) {
		tcp := TCP{}
		if !tcp.Empty() {
			t.Errorf("expected Empty to return true for TCP with empty values, got false")
		}
	})

	t.Run("TCPReturnsFalseForNonEmptyValues", func(t *testing.T) {
		tcp := TCP{State: "ESTABLISHED", Syn: true, TCPflags: "SYN"}
		if tcp.Empty() {
			t.Errorf("expected Empty to return false for TCP with non-empty values, got true")
		}
	})
}
