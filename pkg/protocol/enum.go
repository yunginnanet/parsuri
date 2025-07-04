package protocol

// Protocol represents different types of events.
type Protocol int

const (
	ProtoUnknown Protocol = iota
	ProtoFailed
	ProtoHTTP1
	ProtoFTP
	ProtoSMTP
	ProtoTLS
	ProtoSSH
	ProtoIMAP
	ProtoJABBER
	ProtoSMB
	ProtoDCERPC
	ProtoIRC
	ProtoDNS
	ProtoModbus
	ProtoENIP
	ProtoDNP3
	ProtoNFS
	ProtoNTP
	ProtoFTPDATA
	ProtoTFTP
	ProtoIKE
	ProtoKRB5
	ProtoQUIC
	ProtoDHCP
	ProtoSIP
	ProtoRFB
	ProtoMQTT
	ProtoPGSQL
	ProtoTELNET
	ProtoWEBSOCKET
	ProtoLDAP
	ProtoDOH2
	ProtoTEMPLATE
	ProtoRDP
	ProtoHTTP2
	ProtoBITTORRENT_DHT
	ProtoPOP3
	ProtoMDNS
	ProtoHTTP
	ProtoMaxStatic
)

// String method to convert Protocol to string for easy readability.
func (e Protocol) String() string {
	switch e {
	case ProtoUnknown, ProtoFailed:
		return "Unknown"
	case ProtoHTTP, ProtoHTTP1, ProtoHTTP2:
		return "HTTP"
	case ProtoTLS:
		return "TLS"
	case ProtoDNS:
		return "DNS"
	case ProtoSMTP:
		return "SMTP"
	case ProtoSSH:
		return "SSH"
	case ProtoFTP:
		return "FTP"
	case ProtoSMB:
		return "SMB"
	case ProtoDHCP:
		return "DHCP"
	/* case ProtoARP:
	return "ARP" */
	case ProtoPOP3:
		return "POP3"
	case ProtoModbus:
		return "Modbus"
	case ProtoQUIC:
		return "QUIC"
	default:
		return "Unknown"
	}
}
