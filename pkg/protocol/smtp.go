package protocol

type SMTP struct {
	Helo     string   `json:"helo"`
	MailFrom string   `json:"mail_from"`
	RcptTo   []string `json:"rcpt_to"`
}

func (s SMTP) Empty() bool {
	return s.Helo == "" && s.MailFrom == "" && len(s.RcptTo) == 0
}
