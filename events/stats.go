package events

type Stats struct {
	Uptime  int       `json:"uptime,omitempty"`
	Capture Capture   `json:"capture,omitempty"`
	Decoder Decoder   `json:"decoder,omitempty"`
	Flow    FlowStats `json:"flow,omitempty"`
	Defrag  Defrag    `json:"defrag,omitempty"`
	Stream  Stream    `json:"stream,omitempty"`
	TCP     TCPStats  `json:"tcp,omitempty"`
	Detect  Detect    `json:"detect,omitempty"`
	FlowMgr FlowMgr   `json:"flow_mgr,omitempty"`
	DNS     DNSStats  `json:"dns,omitempty"`
	HTTP    HTTPStats `json:"http,omitempty"`
}

type Capture struct {
	KernelPackets int `json:"kernel_packets,omitempty"`
	KernelDrops   int `json:"kernel_drops,omitempty"`
}

type Decoder struct {
	Pkts       int    `json:"pkts,omitempty"`
	Bytes      int64  `json:"bytes,omitempty"`
	Invalid    int    `json:"invalid,omitempty"`
	IPv4       int    `json:"IPv4,omitempty"`
	IPv6       int    `json:"IPv6,omitempty"`
	Ethernet   int    `json:"ethernet,omitempty"`
	Raw        int    `json:"raw,omitempty"`
	Null       int    `json:"null,omitempty"`
	Sll        int    `json:"sll,omitempty"`
	TCP        int    `json:"tcp,omitempty"`
	UDP        int    `json:"udp,omitempty"`
	SCTP       int    `json:"sctp,omitempty"`
	ICMPv4     int    `json:"ICMPv4,omitempty"`
	ICMPv6     int    `json:"ICMPv6,omitempty"`
	PPP        int    `json:"ppp,omitempty"`
	PPPoE      int    `json:"pppoe,omitempty"`
	GRE        int    `json:"gre,omitempty"`
	VLAN       int    `json:"vlan,omitempty"`
	VLANQINQ   int    `json:"vlanqinq,omitempty"`
	Teredo     int    `json:"teredo,omitempty"`
	IPv4InIPv6 int    `json:"i_pv_4_in_i_pv_6,omitempty"`
	IPv6InIPv6 int    `json:"i_pv_6_in_i_pv_6,omitempty"`
	MPLS       int    `json:"mpls,omitempty"`
	AvgPktSize int    `json:"avg_pkt_size,omitempty"`
	MaxPktSize int    `json:"max_pkt_size,omitempty"`
	ERSPAN     int    `json:"erspan,omitempty"`
	IPRaw      Ipraw  `json:"ipraw,omitempty"`
	Ltnull     Ltnull `json:"ltnull,omitempty"`
	DCE        Dce    `json:"dce,omitempty"`
}

type Ipraw struct {
	InvalidIPVersion int `json:"invalid_ip_version,omitempty"`
}

type Ltnull struct {
	PktTooSmall     int `json:"pkt_too_small,omitempty"`
	UnsupportedType int `json:"unsupported_type,omitempty"`
}

type Dce struct {
	PktTooSmall int `json:"pkt_too_small,omitempty"`
}

type FlowStats struct {
	Memcap           int `json:"memcap,omitempty"`
	Spare            int `json:"spare,omitempty"`
	EmergModeEntered int `json:"emerg_mode_entered,omitempty"`
	EmergModeOver    int `json:"emerg_mode_over,omitempty"`
	TCPReuse         int `json:"tcp_reuse,omitempty"`
	Memuse           int `json:"memuse,omitempty"`
}

type Defrag struct {
	IPv4        DefragDetails `json:"ipv4,omitempty"`
	IPv6        DefragDetails `json:"ipv6,omitempty"`
	MaxFragHits int           `json:"max_frag_hits,omitempty"`
}

type DefragDetails struct {
	Fragments   int `json:"fragments,omitempty"`
	Reassembled int `json:"reassembled,omitempty"`
	Timeouts    int `json:"timeouts,omitempty"`
}

type Stream struct {
	ThreeWhsAckInWrongDir           int `json:"3whs_ack_in_wrong_dir,omitempty"`
	ThreeWhsAsyncWrongSeq           int `json:"3whs_async_wrong_seq,omitempty"`
	ThreeWhsRightSeqWrongAckEvasion int `json:"3whs_right_seq_wrong_ack_evasion,omitempty"`
}

type TCPStats struct {
	Sessions           int `json:"sessions,omitempty"`
	SsnMemcapDrop      int `json:"ssn_memcap_drop,omitempty"`
	Pseudo             int `json:"pseudo,omitempty"`
	PseudoFailed       int `json:"pseudo_failed,omitempty"`
	InvalidChecksum    int `json:"invalid_checksum,omitempty"`
	NoFlow             int `json:"no_flow,omitempty"`
	Syn                int `json:"syn,omitempty"`
	Synack             int `json:"synack,omitempty"`
	Rst                int `json:"rst,omitempty"`
	SegmentMemcapDrop  int `json:"segment_memcap_drop,omitempty"`
	StreamDepthReached int `json:"stream_depth_reached,omitempty"`
	ReassemblyGap      int `json:"reassembly_gap,omitempty"`
	Memuse             int `json:"memuse,omitempty"`
	ReassemblyMemuse   int `json:"reassembly_memuse,omitempty"`
}

type Detect struct {
	Alert int `json:"alert,omitempty"`
}

type FlowMgr struct {
	ClosedPruned int `json:"closed_pruned,omitempty"`
	NewPruned    int `json:"new_pruned,omitempty"`
	EstPruned    int `json:"est_pruned,omitempty"`
}

type DNSStats struct {
	Memuse       int `json:"memuse,omitempty"`
	MemcapState  int `json:"memcap_state,omitempty"`
	MemcapGlobal int `json:"memcap_global,omitempty"`
}

type HTTPStats struct {
	Memuse int `json:"memuse,omitempty"`
	Memcap int `json:"memcap,omitempty"`
}
