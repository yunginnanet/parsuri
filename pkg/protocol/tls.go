package protocol

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
