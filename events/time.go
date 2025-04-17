package events

import (
	"strconv"
	"time"
)

const TimestampFormat = "2006-01-02T15:04:05.000000-0700"

type Time struct{ time.Time }

func (t *Time) UnmarshalJSON(b []byte) error {
	data, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	t.Time, err = time.Parse(TimestampFormat, data)
	return err
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.Time.Format(TimestampFormat) + "\""), nil
}
