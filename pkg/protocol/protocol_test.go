package protocol

import (
	"testing"
)

func TestDNSEmpty(t *testing.T) {
	t.Run("DNSReturnsTrueForEmptyValues", func(t *testing.T) {
		dns := DNS{}
		if !dns.Empty() {
			t.Errorf("expected Empty to return true for DNS with empty values, got false")
		}
	})

	t.Run("DNSReturnsFalseForNonEmptyValues", func(t *testing.T) {
		dns := DNS{Answers: []DNSAnswer{{Rdata: "data", Rrname: "name", Rrtype: "type", TTL: 1}}}
		if dns.Empty() {
			t.Errorf("expected Empty to return false for DNS with non-empty values, got true")
		}
	})
}

func TestHTTPEmpty(t *testing.T) {
	t.Run("HTTPReturnsTrueForEmptyValues", func(t *testing.T) {
		http := HTTP{}
		if !http.Empty() {
			t.Errorf("expected Empty to return true for HTTP with empty values, got false")
		}
	})

	t.Run("HTTPReturnsFalseForNonEmptyValues", func(t *testing.T) {
		http := HTTP{Hostname: "example.com", URL: "/path", Status: HTTPStatus(200), Length: 123}
		if http.Empty() {
			t.Errorf("expected Empty to return false for HTTP with non-empty values, got true")
		}
	})
}

func TestTLSEmpty(t *testing.T) {
	t.Run("TLSReturnsTrueForEmptyValues", func(t *testing.T) {
		tls := TLS{}
		if !tls.Empty() {
			t.Errorf("expected Empty to return true for TLS with empty values, got false")
		}
	})

	t.Run("TLSReturnsFalseForNonEmptyValues", func(t *testing.T) {
		tls := TLS{Subject: "subject", Fingerprint: "fingerprint"}
		if tls.Empty() {
			t.Errorf("expected Empty to return false for TLS with non-empty values, got true")
		}
	})
}

func TestSSHEmpty(t *testing.T) {
	t.Run("SSHReturnsTrueForEmptyValues", func(t *testing.T) {
		ssh := SSH{}
		if !ssh.Empty() {
			t.Errorf("expected Empty to return true for SSH with empty values, got false")
		}
	})

	t.Run("SSHReturnsFalseForNonEmptyValues", func(t *testing.T) {
		ssh := SSH{Client: SSHDetails{ProtoVersion: "2.0", SoftwareVersion: "OpenSSH"}}
		if ssh.Empty() {
			t.Errorf("expected Empty to return false for SSH with non-empty values, got true")
		}
	})
}

func TestSMTPEmpty(t *testing.T) {
	t.Run("SMTPReturnsTrueForEmptyValues", func(t *testing.T) {
		smtp := SMTP{}
		if !smtp.Empty() {
			t.Errorf("expected Empty to return true for SMTP with empty values, got false")
		}
	})

	t.Run("SMTPReturnsFalseForNonEmptyValues", func(t *testing.T) {
		smtp := SMTP{Helo: "helo", MailFrom: "from@example.com", RcptTo: []string{"to@example.com"}}
		if smtp.Empty() {
			t.Errorf("expected Empty to return false for SMTP with non-empty values, got true")
		}
	})
}

func TestEmailEmpty(t *testing.T) {
	t.Run("EmailReturnsTrueForEmptyValues", func(t *testing.T) {
		email := Email{}
		if !email.Empty() {
			t.Errorf("expected Empty to return true for Email with empty values, got false")
		}
	})

	t.Run("EmailReturnsFalseForNonEmptyValues", func(t *testing.T) {
		email := Email{Status: "delivered"}
		if email.Empty() {
			t.Errorf("expected Empty to return false for Email with non-empty values, got true")
		}
	})
}
