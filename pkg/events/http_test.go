package events

import (
	"bytes"
	"github.com/bytedance/sonic"
	"github.com/yunginnanet/parsuri/pkg/protocol"
	"testing"
)

func TestHTTP(t *testing.T) {
	t.Run("ParseStatusIntAndString", func(t *testing.T) {
		var (
			dat  = `{"timestamp":"2017-01-31T16:12:36.103826+0100","flow_id":1501577875823823,"pcap_cnt":10451,"event_type":"fileinfo","src_ip":"54.230.202.37","src_port":80,"dest_ip":"172.17.0.2","dest_port":41402,"proto":"TCP","http":{"hostname":"deb.debian.org","url":"\/debian\/pool\/main\/p\/perl\/perl-base_5.24.1-1_amd64.deb","http_user_agent":"Debian APT-HTTP\/1.3 (1.4~beta3)","http_content_type":"application\/x-debian-package","http_method":"GET","protocol":"HTTP\/1.1","status":"200","length":105175},"app_proto":"http","fileinfo":{"filename":"\/debian\/pool\/main\/p\/perl\/perl-base_5.24.1-1_amd64.deb","state":"TRUNCATED","stored":false,"size":919,"tx_id":0}}`
			dat1 = `{"timestamp":"2017-01-31T16:12:36.103826+0100","flow_id":1501577875823823,"pcap_cnt":10451,"event_type":"fileinfo","src_ip":"54.230.202.37","src_port":80,"dest_ip":"172.17.0.2","dest_port":41402,"proto":"TCP","http":{"hostname":"deb.debian.org","url":"\/debian\/pool\/main\/p\/perl\/perl-base_5.24.1-1_amd64.deb","http_user_agent":"Debian APT-HTTP\/1.3 (1.4~beta3)","http_content_type":"application\/x-debian-package","http_method":"GET","protocol":"HTTP\/1.1","status":200,"length":105175},"app_proto":"http","fileinfo":{"filename":"\/debian\/pool\/main\/p\/perl\/perl-base_5.24.1-1_amd64.deb","state":"TRUNCATED","stored":false,"size":919,"tx_id":0}}`
		)

		b := bytes.NewBufferString(dat)
		b1 := bytes.NewBufferString(dat1)

		var httpEvent EveEvent

		if err := sonic.Unmarshal(b.Bytes(), &httpEvent); err != nil {
			t.Fatalf("failed to unmarshal JSON: %v", err)
		}

		if httpEvent.HTTP.Hostname != "deb.debian.org" {
			t.Errorf("expected Hostname to be 'deb.debian.org', got '%s'", httpEvent.HTTP.Hostname)
		}

		if httpEvent.HTTP.Status != protocol.HTTPStatus(200) {
			t.Errorf("expected Status to be 200, got %d", httpEvent.HTTP.Status)
		}

		if err := sonic.Unmarshal(b1.Bytes(), &httpEvent); err != nil {
			t.Fatalf("failed to unmarshal JSON: %v", err)
		}

		if httpEvent.HTTP.Hostname != "deb.debian.org" {
			t.Errorf("expected Hostname to be 'deb.debian.org', got '%s'", httpEvent.HTTP.Hostname)
		}

		if httpEvent.HTTP.Status != protocol.HTTPStatus(200) {
			t.Errorf("expected Status to be 200, got %d", httpEvent.HTTP.Status)
		}
	})
}
