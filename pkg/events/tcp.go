package events

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
