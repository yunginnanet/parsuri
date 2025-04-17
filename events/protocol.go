package events

type DNS struct {
	Type   string `json:"type"`
	ID     int    `json:"id,omitempty"`
	Rcode  string `json:"rcode"`
	Rrname string `json:"rrname"`
	Rrtype string `json:"rrtype"`
	TTL    int    `json:"ttl,omitempty"`
	Rdata  string `json:"rdata,omitempty"`
	TxID   int    `json:"tx_id,omitempty"`
}

func (d DNS) Empty() bool {
	return d.Rdata == "" && d.Rrname == "" && d.Rrtype == "" && d.TTL == 0
}

type HTTP struct {
	Hostname        string `json:"hostname"`
	URL             string `json:"url"`
	HTTPUserAgent   string `json:"http_user_agent,omitempty"`
	HTTPContentType string `json:"http_content_type,omitempty"`
	HTTPMethod      string `json:"http_method,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	Status          int    `json:"status"`
	Length          int    `json:"length"`
}

func (h HTTP) Empty() bool {
	return h.Hostname == "" && h.URL == "" && h.HTTPUserAgent == "" && h.HTTPContentType == "" && h.HTTPMethod == "" && h.Protocol == "" && h.Status == 0 && h.Length == 0
}

type TLS struct {
	Subject     string `json:"subject"`
	Issuerdn    string `json:"issuerdn,omitempty"`
	Fingerprint string `json:"fingerprint"`
	Sni         string `json:"sni,omitempty"`
	Version     string `json:"version,omitempty"`
}

func (t TLS) Empty() bool {
	return t.Subject == "" && t.Issuerdn == "" && t.Fingerprint == "" && t.Sni == "" && t.Version == ""
}

type SSHDetails struct {
	ProtoVersion    string `json:"proto_version"`
	SoftwareVersion string `json:"software_version"`
}

type SSH struct {
	Client SSHDetails `json:"client"`
	Server SSHDetails `json:"server"`
}

func (s SSH) Empty() bool {
	return s.Client.ProtoVersion == "" && s.Client.SoftwareVersion == "" && s.Server.ProtoVersion == "" && s.Server.SoftwareVersion == ""
}

type SMTP struct {
	Helo     string   `json:"helo"`
	MailFrom string   `json:"mail_from"`
	RcptTo   []string `json:"rcpt_to"`
}

func (s SMTP) Empty() bool {
	return s.Helo == "" && s.MailFrom == "" && len(s.RcptTo) == 0
}

type TCP struct {
	State      string `json:"state"`
	Syn        bool   `json:"syn"`
	TCPflags   string `json:"tcp_flags"`
	TCPflagsTc string `json:"tcp_flags_tc"`
	TCPflagsTs string `json:"tcp_flags_ts"`
}

func (t TCP) Empty() bool {
	return t.State == "" && t.Syn == false && t.TCPflags == "" && t.TCPflagsTc == "" && t.TCPflagsTs == ""
}

type Email struct {
	Status string `json:"status"`
}

func (e Email) Empty() bool {
	return e.Status == ""
}
