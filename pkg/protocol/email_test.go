package protocol

import (
	"testing"
)

func TestEmail_Empty(t *testing.T) {
	tests := []struct {
		name     string
		email    Email
		expected bool
	}{
		{
			name:     "Empty status",
			email:    Email{Status: ""},
			expected: true,
		},
		{
			name:     "Non-empty status",
			email:    Email{Status: "sent"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.email.Empty()
			if result != tt.expected {
				t.Errorf("Email.Empty() = %v; want %v", result, tt.expected)
			}
		})
	}
}
