package events

type Alert struct {
	Action      string `json:"action,omitempty"`
	Gid         int    `json:"gid,omitempty"`
	SignatureID int    `json:"signature_id,omitempty"`
	Rev         int    `json:"rev,omitempty"`
	Signature   string `json:"signature,omitempty"`
	Category    string `json:"category"`
	Severity    int    `json:"severity"`
}

func (a Alert) Empty() bool {
	return a.Action == "" && a.Gid == 0 && a.SignatureID == 0 && a.Rev == 0 && a.Signature == "" && a.Category == "" && a.Severity == 0
}

type FileInfo struct {
	Filename string `json:"filename"`
	Magic    string `json:"magic"`
	State    string `json:"state,omitempty"`
	Md5      string `json:"md5,omitempty"`
	Stored   bool   `json:"stored,omitempty"`
	Size     int    `json:"size"`
	TxID     int    `json:"tx_id,omitempty"`
}

func (f FileInfo) Empty() bool {
	return f.Filename == "" && f.Magic == "" && f.State == "" && f.Md5 == "" && f.Stored == false && f.Size == 0
}

type Bypassed struct {
	PktsToserver  int `json:"pkts_toserver"`
	PktsToclient  int `json:"pkts_toclient"`
	BytesToserver int `json:"bytes_toserver"`
	BytesToclient int `json:"bytes_toclient"`
}

type ExceptionPolicy struct {
	Target string `json:"target"`
	Policy string `json:"policy"`
}

type Ether struct {
	DestMacs []string `json:"dest_macs"`
	SrcMacs  []string `json:"src_macs"`
}

type Flow struct {
	PktsToserver    int               `json:"pkts_toserver,omitempty"`
	PktsToclient    int               `json:"pkts_toclient,omitempty"`
	BytesToserver   int               `json:"bytes_toserver,omitempty"`
	BytesToclient   int               `json:"bytes_toclient,omitempty"`
	Bypassed        Bypassed          `json:"bypassed,omitempty"`
	Start           string            `json:"start,omitempty"`
	End             string            `json:"end,omitempty"`
	Age             int               `json:"age,omitempty"`
	Bypass          string            `json:"bypass,omitempty"`
	State           string            `json:"state,omitempty"`
	Reason          string            `json:"reason,omitempty"`
	Alerted         bool              `json:"alerted,omitempty"`
	Action          string            `json:"action,omitempty"`
	ExceptionPolicy []ExceptionPolicy `json:"exception_policy,omitempty"`
}

func (f Flow) Empty() bool {
	return f.PktsToserver == 0 && f.PktsToclient == 0 && f.BytesToserver == 0 && f.BytesToclient == 0 &&
		f.Start == "" && f.End == "" && f.Age == 0 && f.State == "" && f.Reason == ""
}

type PacketInfo struct {
	Linktype int `json:"linktype"`
}

// EveEvent is the huge struct which can contain a parsed suricata eve.json
// log event.
type EveEvent struct {
	Timestamp         *Time  `json:"timestamp"`
	EventType         string `json:"event_type"`
	FlowID            int64  `json:"flow_id,omitempty"`
	InIface           string `json:"in_iface,omitempty"`
	VLAN              VLAN   `json:"vlan,omitempty"`
	SrcIP             string `json:"src_ip,omitempty"`
	SrcPort           int    `json:"src_port,omitempty"`
	DestIP            string `json:"dest_ip,omitempty"`
	DestPort          int    `json:"dest_port,omitempty"`
	Proto             string `json:"proto,omitempty"`
	AppProto          string `json:"app_proto,omitempty"`
	TxID              int    `json:"tx_id,omitempty"`
	TCP               *TCP   `json:"tcp,omitempty"`
	EtherParticipants *Ether `json:"ether,omitempty"`

	PacketDetails *PacketInfo `json:"packet_info,omitempty"`

	// Alert Events have some additional high level attributes to the json model
	Alert            *Alert `json:"alert,omitempty"`
	Payload          string `json:"payload,omitempty"`
	PayloadPrintable string `json:"payload_printable,omitempty"`
	Stream           int    `json:"stream,omitempty"`
	Packet           string `json:"packet,omitempty"`

	// SMTP Events have some additional high level attributes to the json model
	SMTP *SMTP `json:"smtp,omitempty"`

	// Other sub event_types
	Email    *Email    `json:"email,omitempty"`
	DNS      *DNS      `json:"dns,omitempty"`
	HTTP     *HTTP     `json:"http,omitempty"`
	Fileinfo *FileInfo `json:"fileinfo,omitempty"`
	Flow     *Flow     `json:"flow,omitempty"`
	SSH      *SSH      `json:"ssh,omitempty"`
	TLS      *TLS      `json:"tls,omitempty"`
	Stats    *Stats    `json:"stats,omitempty"`
}
