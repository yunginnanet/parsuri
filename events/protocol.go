package events

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
