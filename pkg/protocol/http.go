package protocol

import (
	"bytes"
	"fmt"
	json "github.com/bytedance/sonic"
	"strconv"
)

var (
	httpMethodBytes = []byte(`"http_method"`)
	methodBytes     = []byte(`"method"`)
)

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

func (h *HTTP) UnmarshalJSON(data []byte) error {
	type tmpHTTP HTTP
	var aux = tmpHTTP{}

	if bytes.Contains(data, httpMethodBytes) {
		data = bytes.ReplaceAll(data, httpMethodBytes, methodBytes)
	}

	err := json.Unmarshal(data, &aux)

	if err != nil {
		return fmt.Errorf("failure during http serde: %w", err)
	}

	*h = HTTP(aux)

	return nil
}

type HTTPHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type (
	HTTPStatus int
)

func (hs *HTTPStatus) UnmarshalJSON(data []byte) error {
	var status int
	var err error

	if !bytes.Contains(data, []byte{'"'}) {
		if err = json.Unmarshal(data, &status); err != nil {
			return fmt.Errorf("failure during http status serde: %w", err)
		}
		*hs = HTTPStatus(status)
		return nil
	}

	var statusString string
	if err = json.Unmarshal(data, &statusString); err != nil {
		return fmt.Errorf("failure during http status serde: %w", err)
	}

	if status, err = strconv.Atoi(statusString); err != nil {
		return fmt.Errorf("failure during http status serde: %w", err)
	}

	*hs = HTTPStatus(status)

	return nil
}

func (h *HTTP) Empty() bool {
	if h == nil {
		return true
	}
	return h.Hostname == "" && h.URL == "" &&
		h.UserAgent == "" && h.ContentType == "" &&
		h.Refer == "" && h.Method == "" &&
		h.Protocol == "" && h.Status == 0 && h.Length == 0
}
