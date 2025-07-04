package parsuri

import (
	"encoding/json"
	"errors"
	"github.com/yunginnanet/parsuri/pkg/events"
	"log"
	"sync"
	"testing"
	"time"
)

func ExampleLoader_LoadOneFile() {
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

func ExampleLoader_LoadFile() {
	loader := NewLoader()

	var errs = make(chan error, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		time.Sleep(100 * time.Millisecond)
		// use LoadFile instead of LoadOneFile
		errs <- loader.LoadFile("pathto/eve1.json")
		wg.Done()
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		// use LoadFile instead of LoadOneFile
		errs <- loader.LoadFile("pathto/eve2.json")
		wg.Done()
	}()

	go func() {
		time.Sleep(600 * time.Millisecond)
		// use LoadFile instead of LoadOneFile
		errs <- loader.LoadFile("pathto/eve3.json")
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(errs)
		_ = loader.Close()
	}()

	var err error
	for e := range errs {
		err = errors.Join(err, e)
	}

	if err != nil {
		log.Println("incomplete load with errors: ", err.Error())
	}

	for loader.More() {
		if err = loader.Err(); err != nil {
			log.Println("error processing data:", err.Error())
			break
		}
		log.Println(loader.Event())
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
