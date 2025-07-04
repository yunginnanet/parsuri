package parsuri

import (
	"encoding/json"
	"github.com/yunginnanet/parsuri/pkg/events"
	"log"
	"testing"
)

func ExampleNewLoader() {
	loader := NewLoader()

	// Load the eve.json file asynchronously
	if err := loader.LoadOneFile("pathto/eve.json"); err != nil {
		log.Fatal(err)
	}

	// Range over the events and print dns answers to stdout
	for loader.More() {
		if err := loader.Err(); err != nil {
			log.Fatal(err)
		}

		event := loader.Event()

		if event.DNS == nil || event.DNS.Empty() {
			continue
		}

		if event.DNS.Type == "answer" {
			log.Println(event.DNS)
		}
	}

	if err := loader.Err(); err != nil {
		log.Fatal(err)
	}
}

func TestAlertEvent(t *testing.T) {
	loader := NewLoader()

	if err := loader.LoadOneFile("testdata/eve.json"); err != nil {
		t.Fatal(err)
	}

	for loader.More() {
		if err := loader.Err(); err != nil {
			t.Error(err)
		}

		event := loader.Event()

		if event.Type != "alert" {
			continue
		}

		if event.Alert.Action == "allowed" {
			return
		}
	}

	t.Error("No allowed alert found")
}

func TestFileinfoEventTruncated(t *testing.T) {
	loader := NewLoader()

	if err := loader.LoadOneFile("testdata/eve.json"); err != nil {
		t.Fatal(err)
	}

	for loader.More() {
		if err := loader.Err(); err != nil {
			t.Error(err)
		}

		event := loader.Event()

		if event.Type != "fileinfo" {
			continue
		}

		// spew.Dump(event)

		if event.Fileinfo.State == "TRUNCATED" {
			t.Logf("Found truncated fileinfo: %v", event.Fileinfo)
			return
		}
	}

	t.Error("No truncated fileinfo found")
}

func TestFileinfoEventClosed(t *testing.T) {
	loader := NewLoader()

	if err := loader.LoadOneFile("testdata/eve.json"); err != nil {
		t.Fatal(err)
	}

	for loader.More() {
		if err := loader.Err(); err != nil {
			t.Error(err)
		}

		event := loader.Event()

		if event.Type != "fileinfo" {
			continue
		}

		// spew.Dump(event)

		if event.Fileinfo.State == "CLOSED" {
			return
		}
	}

	t.Error("No closed fileinfo found")
}

func TestFileinfoEventStored(t *testing.T) {
	loader := NewLoader()

	if err := loader.LoadOneFile("testdata/eve.json"); err != nil {
		t.Fatal(err)
	}

	for loader.More() {
		if err := loader.Err(); err != nil {
			t.Error(err)
		}

		event := loader.Event()

		if event.Type != "fileinfo" {
			continue
		}

		// spew.Dump(event)

		if event.Fileinfo.Stored {
			return
		}
	}

	t.Error("No stored fileinfo found")
}

func TestFileinfoEventNotStored(t *testing.T) {
	loader := NewLoader()

	if err := loader.LoadOneFile("testdata/eve.json"); err != nil {
		t.Fatal(err)
	}

	for loader.More() {
		if err := loader.Err(); err != nil {
			t.Error(err)
		}

		event := loader.Event()
		if event.Type == "fileinfo" && !event.Fileinfo.Stored {
			t.Logf("Found not stored fileinfo: %v", event.Fileinfo)
			return
		}
	}

	t.Error("No not stored fileinfo found")
}

func TestHTTPEventGET(t *testing.T) {
	loader := NewLoader()

	if err := loader.LoadOneFile("testdata/eve.json"); err != nil {
		t.Fatal(err)
	}

	for loader.More() {
		if err := loader.Err(); err != nil {
			t.Error(err)
		}

		event := loader.Event()

		if event.HTTP.Empty() {
			continue
		}

		// spew.Dump(event)

		if event.HTTP.Method == "GET" {
			return
		}

	}

	t.Error("No GET http event found")
}

func TestHTTPEventStatus200(t *testing.T) {
	loader := NewLoader()

	if err := loader.LoadOneFile("testdata/eve.json"); err != nil {
		t.Fatal(err)
	}

	for loader.More() {
		if err := loader.Err(); err != nil {
			t.Error(err)
		}

		event := loader.Event()

		if event.Type != "http" {
			continue
		}

		t.Logf("%v", event.HTTP)

		if event.HTTP.Status == 200 {
			return
		}
	}

	t.Error("No HTTP status 200 found")
}

func TestLoadEveJSONFile(t *testing.T) {
	var countTotal int
	var countDNS int
	var countFlow int

	loader := NewLoader()

	// Load the eve.json file
	if err := loader.LoadOneFile("testdata/eve.json"); err != nil {
		t.Fatal(err)
	}

	for loader.More() {
		if err := loader.Err(); err != nil {
			t.Error(err)
		}

		event := loader.Event()

		t.Logf("%v", event)

		if event.DNS != nil && !event.DNS.Empty() {
			countDNS++
		}
		if event.Flow != nil && !event.Flow.Empty() {
			countFlow++
		}
		if event.Type == "" || event.Timestamp.IsZero() {
			t.Error("Mandatory field missing")
		}

		countTotal++
	}

	if err := loader.Err(); err != nil {
		t.Error(err)
	}

	if countDNS != 48 {
		t.Errorf("DNS count mismatch: %d != 48", countDNS)
	}
	if countFlow != 13 {
		t.Errorf("Flow count mismatch: %d != 13", countFlow)
	}
	if countTotal != 266 {
		t.Errorf("Total count mismatch: %d != 266", countTotal)
	}
}

func TestLoadBrokenEveJSONFile(t *testing.T) {
	var countErrors int

	loader := NewLoader()

	if err := loader.LoadOneFile("testdata/eve_broken.json"); err != nil {
		t.Fatal(err)
	}

	for loader.More() {
		if err := loader.Err(); err != nil {
			t.Log(err)
			countErrors++
		}
		event := loader.Event()
		if event.Type == "" || event.Timestamp.IsZero() {
			t.Error("Mandatory field missing")
		}
	}

	if err := loader.Err(); err != nil {
		t.Log(err)
		countErrors++
	}

	if countErrors < 1 {
		t.Errorf("Expected at least one error, got %d", countErrors)
	}
}

func TestMissingJSONFile(t *testing.T) {
	loader := NewLoader()

	if err := loader.LoadOneFile("nonexistant"); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestMarshalWithTimestamp(t *testing.T) {
	loader := NewLoader()

	if err := loader.LoadOneFile("testdata/eve.json"); err != nil {
		t.Fatal(err)
	}

	if err := loader.Err(); err != nil {
		t.Error(err)
	}

	e := loader.Event()

	out, err := json.Marshal(e)
	if err != nil {
		t.Error(err)
	}

	var inEVE events.EveEvent
	err = json.Unmarshal(out, &inEVE)
	if err != nil {
		t.Error(err)
	}

	if !inEVE.Timestamp.Time.Equal(e.Timestamp.Time) {
		t.Fatalf("timestamp round-trip failed: %v <-> %v", inEVE.Timestamp, e.Timestamp)
	}
}
