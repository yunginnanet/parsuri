package protocol

import (
	"testing"
)

func TestEmpty_MainFunctionality(t *testing.T) {
	dhcp := DHCP{}
	if !dhcp.Empty() {
		t.Errorf("Expected DHCP to be empty, but it was not")
	}

	dhcp.ClientMAC = "00:1A:2B:3C:4D:5E"
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty, but it was")
	}
}

func TestEmpty_EdgeCases(t *testing.T) {
	dhcp := DHCP{
		ClientMAC:     "",
		AssignedIP:    "",
		ClientIP:      "",
		RelayIP:       "",
		NextServerIP:  "",
		DHCPType:      "",
		SubnetMask:    "",
		Routers:       []string{},
		Hostname:      "",
		LeaseTime:     0,
		RenewalTime:   0,
		RebindingTime: 0,
		ClientID:      "",
		DnsServers:    []string{},
	}
	if !dhcp.Empty() {
		t.Errorf("Expected DHCP to be empty, but it was not")
	}

	dhcp.ClientMAC = "00:1A:2B:3C:4D:5E"
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty, but it was")
	}
}

func TestEmpty_InputValidation(t *testing.T) {
	dhcp := DHCP{
		ClientMAC:  "",
		AssignedIP: "192.168.1.1",
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to AssignedIP, but it was")
	}

	dhcp = DHCP{
		ClientIP: "192.168.1.1",
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to ClientIP, but it was")
	}

	dhcp = DHCP{
		RelayIP: "192.168.1.1",
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to RelayIP, but it was")
	}

	dhcp = DHCP{
		NextServerIP: "192.168.1.1",
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to NextServerIP, but it was")
	}

	dhcp = DHCP{
		DHCPType: "BOOTP",
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to DHCPType, but it was")
	}

	dhcp = DHCP{
		SubnetMask: "255.255.255.0",
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to SubnetMask, but it was")
	}

	dhcp = DHCP{
		Routers: []string{"192.168.1.1"},
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to Routers, but it was")
	}

	dhcp = DHCP{
		Hostname: "example.com",
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to Hostname, but it was")
	}

	dhcp = DHCP{
		LeaseTime: 3600,
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to LeaseTime, but it was")
	}

	dhcp = DHCP{
		RenewalTime: 1800,
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to RenewalTime, but it was")
	}

	dhcp = DHCP{
		RebindingTime: 2700,
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to RebindingTime, but it was")
	}

	dhcp = DHCP{
		ClientID: "12345",
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to ClientID, but it was")
	}

	dhcp = DHCP{
		DnsServers: []string{"8.8.8.8"},
	}
	if dhcp.Empty() {
		t.Errorf("Expected DHCP not to be empty due to DnsServers, but it was")
	}
}
