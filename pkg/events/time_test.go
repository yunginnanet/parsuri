package events

import (
	"encoding/json"
	"testing"
	"time"
)

func TestUnmarshalJSONParsesValidTimestamp(t *testing.T) {
	input := `"2023-03-15T14:30:00.000000-0700"`
	var parsedTime Time
	err := json.Unmarshal([]byte(input), &parsedTime)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	expectedTime, _ := time.Parse(TimestampFormat, "2023-03-15T14:30:00.000000-0700")
	if !parsedTime.Time.Equal(expectedTime) {
		diff := expectedTime.Compare(parsedTime.Time)
		t.Errorf("expected %v, got %v, %d off", expectedTime, parsedTime.Time, diff)
	}
}

func TestUnmarshalJSONReturnsErrorForInvalidTimestamp(t *testing.T) {
	input := `"invalid-timestamp"`
	var parsedTime Time
	err := json.Unmarshal([]byte(input), &parsedTime)
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestMarshalJSONFormatsTimeCorrectly(t *testing.T) {
	parsedTime := Time{Time: time.Date(2023, 3, 15, 14, 30, 0, 0, time.FixedZone("-0700", -7*3600))}
	output, err := json.Marshal(&parsedTime)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	expectedOutput := `"2023-03-15T14:30:00.000000-0700"`
	if string(output) != expectedOutput {
		t.Errorf("expected %s, got %s", expectedOutput, string(output))
	}
}

func TestMarshalJSONHandlesZeroTime(t *testing.T) {
	parsedTime := Time{}
	output, err := json.Marshal(&parsedTime)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	expectedOutput := `"0001-01-01T00:00:00.000000+0000"`
	if string(output) != expectedOutput {
		t.Errorf("expected %s, got %s", expectedOutput, string(output))
	}
}
