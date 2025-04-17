package events

import (
	"bytes"
	"github.com/bytedance/sonic"
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

func TestHTTP(t *testing.T) {
	t.Run("ParseStatusIntAndString", func(t *testing.T) {
		var (
			dat  = `{"timestamp":"2017-01-31T16:12:36.103826+0100","flow_id":1501577875823823,"pcap_cnt":10451,"event_type":"fileinfo","src_ip":"54.230.202.37","src_port":80,"dest_ip":"172.17.0.2","dest_port":41402,"proto":"TCP","http":{"hostname":"deb.debian.org","url":"\/debian\/pool\/main\/p\/perl\/perl-base_5.24.1-1_amd64.deb","http_user_agent":"Debian APT-HTTP\/1.3 (1.4~beta3)","http_content_type":"application\/x-debian-package","http_method":"GET","protocol":"HTTP\/1.1","status":"200","length":105175},"app_proto":"http","fileinfo":{"filename":"\/debian\/pool\/main\/p\/perl\/perl-base_5.24.1-1_amd64.deb","state":"TRUNCATED","stored":false,"size":919,"tx_id":0}}`
			dat1 = `{"timestamp":"2017-01-31T16:12:36.103826+0100","flow_id":1501577875823823,"pcap_cnt":10451,"event_type":"fileinfo","src_ip":"54.230.202.37","src_port":80,"dest_ip":"172.17.0.2","dest_port":41402,"proto":"TCP","http":{"hostname":"deb.debian.org","url":"\/debian\/pool\/main\/p\/perl\/perl-base_5.24.1-1_amd64.deb","http_user_agent":"Debian APT-HTTP\/1.3 (1.4~beta3)","http_content_type":"application\/x-debian-package","http_method":"GET","protocol":"HTTP\/1.1","status":200,"length":105175},"app_proto":"http","fileinfo":{"filename":"\/debian\/pool\/main\/p\/perl\/perl-base_5.24.1-1_amd64.deb","state":"TRUNCATED","stored":false,"size":919,"tx_id":0}}`
		)
		b := bytes.NewBufferString(dat)
		b1 := bytes.NewBufferString(dat1)
		var event EveEvent
		if err := sonic.Unmarshal(b.Bytes(), &event); err != nil {
			t.Fatalf("failed to unmarshal JSON: %v", err)
		}
		if event.HTTP.Hostname != "deb.debian.org" {
			t.Errorf("expected Hostname to be 'deb.debian.org', got '%s'", event.HTTP.Hostname)
		}
		if event.HTTP.Status != HTTPStatus(200) {
			t.Errorf("expected Status to be 200, got %d", event.HTTP.Status)
		}

		if err := sonic.Unmarshal(b1.Bytes(), &event); err != nil {
			t.Fatalf("failed to unmarshal JSON: %v", err)
		}
		if event.HTTP.Hostname != "deb.debian.org" {
			t.Errorf("expected Hostname to be 'deb.debian.org', got '%s'", event.HTTP.Hostname)
		}
		if event.HTTP.Status != HTTPStatus(200) {
			t.Errorf("expected Status to be 200, got %d", event.HTTP.Status)
		}
	})
}
