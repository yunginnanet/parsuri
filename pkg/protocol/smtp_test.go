package protocol

import (
	"testing"
)

func TestSMTP_Empty(t *testing.T) {
	tests := []struct {
		name     string
		smtp     SMTP
		expected bool
	}{
		{
			name:     "All fields empty",
			smtp:     SMTP{},
			expected: true,
		},
		{
			name:     "Helo set, others empty",
			smtp:     SMTP{Helo: "test"},
			expected: false,
		},
		{
			name:     "MailFrom set, others empty",
			smtp:     SMTP{MailFrom: "test@example.com"},
			expected: false,
		},
		{
			name:     "RcptTo set, others empty",
			smtp:     SMTP{RcptTo: []string{"test@example.com"}},
			expected: false,
		},
		{
			name:     "All fields set",
			smtp:     SMTP{Helo: "test", MailFrom: "test@example.com", RcptTo: []string{"test@example.com"}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.smtp.Empty()
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
