package protocol

type DHCP struct {
	Type          string   `json:"type"`
	ID            int64    `json:"id"`
	ClientMAC     string   `json:"client_mac"`
	AssignedIP    string   `json:"assigned_ip"`
	ClientIP      string   `json:"client_ip"`
	RelayIP       string   `json:"relay_ip"`
	NextServerIP  string   `json:"next_server_ip"`
	DHCPType      string   `json:"dhcp_type"`
	SubnetMask    string   `json:"subnet_mask"`
	Routers       []string `json:"routers"`
	Hostname      string   `json:"hostname"`
	LeaseTime     int      `json:"lease_time"`
	RenewalTime   int      `json:"renewal_time"`
	RebindingTime int      `json:"rebinding_time"`
	ClientID      string   `json:"client_id"`
	DnsServers    []string `json:"dns_servers"`
}

func (dhcp DHCP) Empty() bool {
	return dhcp.ClientMAC == "" && dhcp.AssignedIP == "" && dhcp.ClientIP == "" && dhcp.RelayIP == "" &&
		dhcp.NextServerIP == "" && dhcp.DHCPType == "" && dhcp.SubnetMask == "" && len(dhcp.Routers) == 0 &&
		dhcp.Hostname == "" && dhcp.LeaseTime == 0 && dhcp.RenewalTime == 0 && dhcp.RebindingTime == 0 &&
		dhcp.ClientID == "" && len(dhcp.DnsServers) == 0
}
