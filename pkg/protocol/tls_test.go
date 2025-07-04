package protocol

import (
	"testing"
)

func TestTLS_Empty(t *testing.T) {
	tests := []struct {
		name     string
		tls      TLS
		expected bool
	}{
		{
			name: "All fields empty",
			tls: TLS{
				Subject:     "",
				Issuerdn:    "",
				Fingerprint: "",
				Sni:         "",
				Version:     "",
			},
			expected: true,
		},
		{
			name: "Only Subject is set",
			tls: TLS{
				Subject:     "example.com",
				Issuerdn:    "",
				Fingerprint: "",
				Sni:         "",
				Version:     "",
			},
			expected: false,
		},
		{
			name: "Only Issuerdn is set",
			tls: TLS{
				Subject:     "",
				Issuerdn:    "issuer.com",
				Fingerprint: "",
				Sni:         "",
				Version:     "",
			},
			expected: false,
		},
		{
			name: "Only Fingerprint is set",
			tls: TLS{
				Subject:     "",
				Issuerdn:    "",
				Fingerprint: "fingerprint123",
				Sni:         "",
				Version:     "",
			},
			expected: false,
		},
		{
			name: "Only Sni is set",
			tls: TLS{
				Subject:     "",
				Issuerdn:    "",
				Fingerprint: "",
				Sni:         "sni123",
				Version:     "",
			},
			expected: false,
		},
		{
			name: "Only Version is set",
			tls: TLS{
				Subject:     "",
				Issuerdn:    "",
				Fingerprint: "",
				Sni:         "",
				Version:     "TLSv1.2",
			},
			expected: false,
		},
		{
			name: "Multiple fields are set",
			tls: TLS{
				Subject:     "example.com",
				Issuerdn:    "issuer.com",
				Fingerprint: "fingerprint123",
				Sni:         "sni123",
				Version:     "TLSv1.2",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.tls.Empty()
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
