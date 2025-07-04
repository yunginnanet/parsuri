package protocol

import (
	"bytes"
	"github.com/bytedance/sonic"
	"testing"
)

func TestHTTP(t *testing.T) {
	t.Run("ParseStatusIntAndString", func(t *testing.T) {
		var (
			dat  = `{"hostname":"deb.debian.org","url":"\/debian\/pool\/main\/p\/perl\/perl-base_5.24.1-1_amd64.deb","http_user_agent":"Debian APT-HTTP\/1.3 (1.4~beta3)","http_content_type":"application\/x-debian-package","http_method":"GET","protocol":"HTTP\/1.1","status":"200","length":105175}`
			dat1 = `{"hostname":"deb.debian.org","url":"\/debian\/pool\/main\/p\/perl\/perl-base_5.24.1-1_amd64.deb","http_user_agent":"Debian APT-HTTP\/1.3 (1.4~beta3)","http_content_type":"application\/x-debian-package","http_method":"GET","protocol":"HTTP\/1.1","status":200,"length":105175}`
		)

		b := bytes.NewBufferString(dat)
		b1 := bytes.NewBufferString(dat1)

		var httpEvent HTTP

		if err := sonic.Unmarshal(b.Bytes(), &httpEvent); err != nil {
			t.Fatalf("failed to unmarshal JSON: %v", err)
		}

		if httpEvent.Hostname != "deb.debian.org" {
			t.Errorf("expected Hostname to be 'deb.debian.org', got '%s'", httpEvent.Hostname)
		}

		if httpEvent.Status != HTTPStatus(200) {
			t.Errorf("expected Status to be 200, got %d", httpEvent.Status)
		}

		if err := sonic.Unmarshal(b1.Bytes(), &httpEvent); err != nil {
			t.Fatalf("failed to unmarshal JSON: %v", err)
		}

		if httpEvent.Hostname != "deb.debian.org" {
			t.Errorf("expected Hostname to be 'deb.debian.org', got '%s'", httpEvent.Hostname)
		}

		if httpEvent.Status != HTTPStatus(200) {
			t.Errorf("expected Status to be 200, got %d", httpEvent.Status)
		}
	})
}
