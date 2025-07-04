package protocol

type Email struct {
	Status string `json:"status"`
}

func (e Email) Empty() bool {
	return e.Status == ""
}
