package events

import (
	"testing"
	"time"

	"github.com/yunginnanet/parsuri/pkg/protocol"
)

func TestAlert_Empty(t *testing.T) {
	alert := Alert{}

	if !alert.Empty() {
		t.Errorf("Expected Alert to be empty, but it was not")
	}

	alert.Action = "test"
	if alert.Empty() {
		t.Errorf("Expected Alert to not be empty, but it was")
	}
}

func TestFileInfo_Empty(t *testing.T) {
	fileInfo := FileInfo{}

	if !fileInfo.Empty() {
		t.Errorf("Expected FileInfo to be empty, but it was not")
	}

	fileInfo.Filename = "test"
	if fileInfo.Empty() {
		t.Errorf("Expected FileInfo to not be empty, but it was")
	}
}

func TestFlow_Empty(t *testing.T) {
	flow := Flow{}

	if !flow.Empty() {
		t.Errorf("Expected Flow to be empty, but it was not")
	}

	flow.PktsToserver = 1
	if flow.Empty() {
		t.Errorf("Expected Flow to not be empty, but it was")
	}
}

func TestEveEvent_Empty(t *testing.T) {
	event := EveEvent{}

	if !event.Empty() {
		t.Errorf("Expected EveEvent to be empty, but it was not")
	}

	event.Timestamp = &Time{time.Now()}
	event.SrcIP = "127.0.0.1"

	if event.Empty() {
		t.Errorf("Expected EveEvent to not be empty, but it was")
	}
}

func TestEveEvent_EmptyWithAlert(t *testing.T) {
	event := EveEvent{
		Timestamp: &Time{time.Now()},
		Alert:     &Alert{},
	}

	if !event.Empty() {
		t.Errorf("Expected EveEvent to be empty with Alert, but it was not")
	}

	event.Alert.Action = "test"

	if event.Empty() {
		t.Errorf("Expected EveEvent to not be empty with non-empty Alert, but it was")
	}
}

func TestEveEvent_EmptyWithFileInfo(t *testing.T) {
	event := EveEvent{
		Timestamp: &Time{time.Now()},
		Fileinfo:  &FileInfo{},
	}

	if !event.Empty() {
		t.Errorf("Expected EveEvent to be empty with FileInfo, but it was not")
	}

	event.Fileinfo.Filename = "test"
	if event.Empty() {
		t.Errorf("Expected EveEvent to not be empty with non-empty FileInfo, but it was")
	}
}

func TestEveEvent_EmptyWithFlow(t *testing.T) {
	event := EveEvent{
		Timestamp: &Time{time.Now()},
		Flow:      &Flow{},
	}

	if !event.Empty() {
		t.Errorf("Expected EveEvent to be empty with Flow, but it was not")
	}

	event.Flow.PktsToserver = 1
	if event.Empty() {
		t.Errorf("Expected EveEvent to not be empty with non-empty Flow, but it was")
	}
}

func TestEveEvent_EmptyWithSSH(t *testing.T) {
	event := EveEvent{
		Timestamp: &Time{time.Now()},
		SSH:       &protocol.SSH{},
	}

	if !event.Empty() {
		t.Errorf("Expected EveEvent to be empty with SSH, but it was not")
	}

	event.SSH.Client.SoftwareVersion = "test"

	if event.Empty() {
		t.Errorf("Expected EveEvent to not be empty with non-empty SSH, but it was")
	}
}

func TestEveEvent_EmptyWithTLS(t *testing.T) {
	event := EveEvent{
		Timestamp: &Time{time.Now()},
		TLS:       &protocol.TLS{},
	}

	if !event.Empty() {
		t.Errorf("Expected EveEvent to be empty with TLS, but it was not")
	}

	event.TLS.Version = "test"
	if event.Empty() {
		t.Errorf("Expected EveEvent to not be empty with non-empty TLS, but it was")
	}
}
