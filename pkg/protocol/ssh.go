package protocol

type SSHDetails struct {
	ProtoVersion    string `json:"proto_version"`
	SoftwareVersion string `json:"software_version"`
}

type SSH struct {
	Client SSHDetails `json:"client"`
	Server SSHDetails `json:"server"`
}

func (s SSH) Empty() bool {
	return s.Client.ProtoVersion == "" && s.Client.SoftwareVersion == "" && s.Server.ProtoVersion == "" && s.Server.SoftwareVersion == ""
}
