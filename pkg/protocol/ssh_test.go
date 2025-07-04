package protocol

import (
	"testing"
)

func TestEmpty_FullStruct(t *testing.T) {
	ssh := SSH{
		Client: SSHDetails{
			ProtoVersion:    "2.0",
			SoftwareVersion: "OpenSSH_8.2p1",
		},
		Server: SSHDetails{
			ProtoVersion:    "2.0",
			SoftwareVersion: "OpenSSH_8.2p1",
		},
	}
	if ssh.Empty() {
		t.Errorf("Expected SSH struct not to be empty, but got true")
	}
}

func TestEmpty_EmptyStruct(t *testing.T) {
	ssh := SSH{}
	if !ssh.Empty() {
		t.Errorf("Expected SSH struct to be empty, but got false")
	}
}

func TestEmpty_PartiallyFilledStruct(t *testing.T) {
	ssh := SSH{
		Client: SSHDetails{
			ProtoVersion:    "2.0",
			SoftwareVersion: "",
		},
		Server: SSHDetails{
			ProtoVersion:    "",
			SoftwareVersion: "OpenSSH_8.2p1",
		},
	}
	if ssh.Empty() {
		t.Errorf("Expected SSH struct not to be empty, but got true")
	}
}

func TestEmpty_AllFieldsEmpty(t *testing.T) {
	ssh := SSH{
		Client: SSHDetails{
			ProtoVersion:    "",
			SoftwareVersion: "",
		},
		Server: SSHDetails{
			ProtoVersion:    "",
			SoftwareVersion: "",
		},
	}
	if !ssh.Empty() {
		t.Errorf("Expected SSH struct to be empty, but got false")
	}
}
