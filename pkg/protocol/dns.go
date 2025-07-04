package protocol

type DNS struct {
	Version int         `json:"version,omitempty"`
	Type    string      `json:"type,omitempty"`
	ID      int         `json:"id,omitempty"`
	Flags   string      `json:"flags,omitempty"`
	Qr      bool        `json:"qr,omitempty"`
	Rd      bool        `json:"rd,omitempty"`
	Ra      bool        `json:"ra,omitempty"`
	Rcode   string      `json:"rcode,omitempty"`
	Queries []DNSQuery  `json:"queries,omitempty"`
	Answers []DNSAnswer `json:"answers,omitempty"`
}

type DNSQuery struct {
	Rrname string `json:"rrname"`
	Rrtype string `json:"rrtype"`
}

type DNSAnswer struct {
	Rrname string `json:"rrname"`
	Rrtype string `json:"rrtype"`
	TTL    int    `json:"ttl"`
	Rdata  string `json:"rdata"`
}

func (d DNS) Empty() bool {
	return d.Queries == nil && d.Answers == nil && d.Type == "" && d.ID == 0 &&
		d.Flags == "" && d.Qr == false && d.Rd == false && d.Ra == false && d.Rcode == ""
}
