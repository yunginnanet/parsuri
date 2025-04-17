package events

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type HTTPHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HTTPStatus int

func (h *HTTPStatus) UnmarshalJSON(data []byte) error {
	var status int
	var err error
	if !bytes.Contains(data, []byte{'"'}) {
		return json.Unmarshal(data, &status)
	}
	var statusString string
	if err = json.Unmarshal(data, &statusString); err != nil {
		return err
	}
	if status, err = strconv.Atoi(statusString); err != nil {
		return err
	}
	*h = HTTPStatus(status)
	return nil
}

type HTTP struct {
	Hostname        string       `json:"hostname,omitempty"`
	URL             string       `json:"url,omitempty"`
	UserAgent       string       `json:"user_agent,omitempty"`
	ContentType     string       `json:"content_type,omitempty"`
	Refer           string       `json:"refer,omitempty"`
	Method          string       `json:"method,omitempty"`
	Protocol        string       `json:"protocol,omitempty"`
	Status          HTTPStatus   `json:"status,omitempty"`
	Length          int          `json:"length,omitempty"`
	RequestHeaders  []HTTPHeader `json:"request_headers,omitempty"`
	ResponseHeaders []HTTPHeader `json:"response_headers,omitempty"`
}

func (h HTTP) Empty() bool {
	return h.Hostname == "" && h.URL == "" &&
		h.UserAgent == "" && h.ContentType == "" &&
		h.Refer == "" && h.Method == "" &&
		h.Protocol == "" && h.Status == 0 && h.Length == 0
}
