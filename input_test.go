package parsuri

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

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

func splitTestData(path string, t *testing.T) []string {
	t.Helper()

	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err.Error())
		return []string{}
	}

	buf := &bytes.Buffer{}

	if _, err = buf.ReadFrom(f); err != nil {
		t.Fatal(err.Error())
		return []string{}
	}

	lines := bytes.Count(buf.Bytes(), []byte{'\n'})

	files := []string{
		filepath.Join(t.TempDir(), "eve1.json"),
		filepath.Join(t.TempDir(), "eve2.json"),
		filepath.Join(t.TempDir(), "eve3.json"),
	}

	var osFiles = make([]*os.File, 3)

	for i, file := range files {
		var osf *os.File
		osf, err = os.Create(file)
		if err != nil {
			t.Fatal(err.Error())
			return []string{}
		}
		osFiles[i] = osf
	}

	per := lines / 3
	head := 0

	xerox := bufio.NewScanner(buf)

	for xerox.Scan() {
		if e := xerox.Err(); e != nil {
			t.Fatal(e.Error())
			return []string{}
		}
		head++
		curr := 0
		switch {
		case head < per:
			curr = 0
		case head < 2*per:
			curr = 1
		default:
			curr = 2
		}

		if _, err = osFiles[curr].Write(append(xerox.Bytes(), '\n')); err != nil {
			t.Fatal(err.Error())
			return []string{}
		}
	}

	for _, osf := range osFiles {
		if err = osf.Sync(); err != nil {
			t.Error(err.Error())
		}
		st, _ := osf.Stat()
		t.Logf("created file with %d bytes: %s", st.Size(), osf.Name())
		if err = osf.Close(); err != nil {
			t.Fatal(err.Error())
			return []string{}
		}
	}

	return files
}

func TestLoadMultipeEveJSONFiles(t *testing.T) {
	loader := NewLoader()

	files := splitTestData("testdata/eve.json", t)

	var errs = make(chan error, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		time.Sleep(100 * time.Millisecond)
		// use LoadFile instead of LoadOneFile
		errs <- loader.LoadFile(files[0])
		wg.Done()
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		// use LoadFile instead of LoadOneFile
		errs <- loader.LoadFile(files[1])
		wg.Done()
	}()

	go func() {
		time.Sleep(600 * time.Millisecond)
		// use LoadFile instead of LoadOneFile
		errs <- loader.LoadFile(files[2])
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

	events := 0

	for loader.More() {
		if err = loader.Err(); err != nil {
			log.Println("error processing data:", err.Error())
			break
		}
		_ = loader.Event()
		events++
	}

	t.Logf("processed  %d events", events)

	if events < 175 {
		t.Errorf("expected at least 175 events, got %d", events)
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

func TestLoadSTDIN(t *testing.T) {
	loader := NewLoader()

	loader.LoadSTDIN()

	f, err := os.Open("testdata/eve.json")
	if err != nil {
		t.Fatal(err)
	}

	oldStdin := os.Stdin
	if os.Stdin, err = os.CreateTemp(t.TempDir(), ""); err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.Stdin.Close()
		os.Stdin = oldStdin
	}()

	go func() {
		buf := bufio.NewWriter(os.Stdin)

		if _, err = buf.ReadFrom(f); err != nil {
			t.Error(err)
		}

		if err = loader.Close(); err != nil {
			t.Error(err)
		}

		if err = f.Close(); err != nil {
			t.Error(err)
		}
	}()

	events := 0

	for loader.More() {
		events++
	}

	t.Logf("processed %d events", events)

	if events < 175 {
		t.Errorf("Expected at least 175 events, got %d", events)
	}
}

func TestMissingJSONFile(t *testing.T) {
	loader := NewLoader()

	if err := loader.LoadOneFile("nonexistant"); err == nil {
		t.Fatal("expected error when loading nonexistant file, got nil")
	}

	if err := loader.LoadFile("nonexistant"); err == nil {
		t.Fatal("expected error when loading nonexistant file, got nil")
	}
}
