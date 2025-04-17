package events

type Stats struct {
	Uptime  int `json:"uptime"`
	Capture struct {
		KernelPackets int `json:"kernel_packets"`
		KernelDrops   int `json:"kernel_drops"`
	} `json:"capture"`
	Decoder struct {
		Pkts       int   `json:"pkts"`
		Bytes      int64 `json:"bytes"`
		Invalid    int   `json:"invalid"`
		Ipv4       int   `json:"ipv4"`
		Ipv6       int   `json:"ipv6"`
		Ethernet   int   `json:"ethernet"`
		Raw        int   `json:"raw"`
		Null       int   `json:"null"`
		Sll        int   `json:"sll"`
		TCP        int   `json:"tcp"`
		UDP        int   `json:"udp"`
		Sctp       int   `json:"sctp"`
		Icmpv4     int   `json:"icmpv4"`
		Icmpv6     int   `json:"icmpv6"`
		Ppp        int   `json:"ppp"`
		Pppoe      int   `json:"pppoe"`
		Gre        int   `json:"gre"`
		Vlan       int   `json:"vlan"`
		VlanQinq   int   `json:"vlan_qinq"`
		Teredo     int   `json:"teredo"`
		Ipv4InIpv6 int   `json:"ipv4_in_ipv6"`
		Ipv6InIpv6 int   `json:"ipv6_in_ipv6"`
		Mpls       int   `json:"mpls"`
		AvgPktSize int   `json:"avg_pkt_size"`
		MaxPktSize int   `json:"max_pkt_size"`
		Erspan     int   `json:"erspan"`
		Ipraw      struct {
			InvalidIPVersion int `json:"invalid_ip_version"`
		} `json:"ipraw"`
		Ltnull struct {
			PktTooSmall     int `json:"pkt_too_small"`
			UnsupportedType int `json:"unsupported_type"`
		} `json:"ltnull"`
		Dce struct {
			PktTooSmall int `json:"pkt_too_small"`
		} `json:"dce"`
	} `json:"decoder"`
	Flow struct {
		Memcap           int `json:"memcap"`
		Spare            int `json:"spare"`
		EmergModeEntered int `json:"emerg_mode_entered"`
		EmergModeOver    int `json:"emerg_mode_over"`
		TCPReuse         int `json:"tcp_reuse"`
		Memuse           int `json:"memuse"`
	} `json:"flow"`
	Defrag struct {
		Ipv4 struct {
			Fragments   int `json:"fragments"`
			Reassembled int `json:"reassembled"`
			Timeouts    int `json:"timeouts"`
		} `json:"ipv4"`
		Ipv6 struct {
			Fragments   int `json:"fragments"`
			Reassembled int `json:"reassembled"`
			Timeouts    int `json:"timeouts"`
		} `json:"ipv6"`
		MaxFragHits int `json:"max_frag_hits"`
	} `json:"defrag"`
	Stream struct {
		ThreeWhsAckInWrongDir           int `json:"3whs_ack_in_wrong_dir"`
		ThreeWhsAsyncWrongSeq           int `json:"3whs_async_wrong_seq"`
		ThreeWhsRightSeqWrongAckEvasion int `json:"3whs_right_seq_wrong_ack_evasion"`
	} `json:"stream"`
	TCP struct {
		Sessions           int `json:"sessions"`
		SsnMemcapDrop      int `json:"ssn_memcap_drop"`
		Pseudo             int `json:"pseudo"`
		PseudoFailed       int `json:"pseudo_failed"`
		InvalidChecksum    int `json:"invalid_checksum"`
		NoFlow             int `json:"no_flow"`
		Syn                int `json:"syn"`
		Synack             int `json:"synack"`
		Rst                int `json:"rst"`
		SegmentMemcapDrop  int `json:"segment_memcap_drop"`
		StreamDepthReached int `json:"stream_depth_reached"`
		ReassemblyGap      int `json:"reassembly_gap"`
		Memuse             int `json:"memuse"`
		ReassemblyMemuse   int `json:"reassembly_memuse"`
	} `json:"tcp"`
	Detect struct {
		Alert int `json:"alert"`
	} `json:"detect"`
	FlowMgr struct {
		ClosedPruned int `json:"closed_pruned"`
		NewPruned    int `json:"new_pruned"`
		EstPruned    int `json:"est_pruned"`
	} `json:"flow_mgr"`
	DNS struct {
		Memuse       int `json:"memuse"`
		MemcapState  int `json:"memcap_state"`
		MemcapGlobal int `json:"memcap_global"`
	} `json:"dns"`
	HTTP struct {
		Memuse int `json:"memuse"`
		Memcap int `json:"memcap"`
	} `json:"http"`
}
