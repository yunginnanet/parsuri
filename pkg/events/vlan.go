package events

import (
	"bytes"
	json "github.com/bytedance/sonic"
)

type VLAN []int64

func (v *VLAN) UnmarshalJSON(data []byte) error {
	if !bytes.Contains(data, []byte{'['}) {
		var vlan int64
		if err := json.Unmarshal(data, &vlan); err != nil {
			return err
		}
		*v = VLAN{vlan}
		return nil
	}
	var vlans []int64
	if err := json.Unmarshal(data, &vlans); err != nil {
		return err
	}
	*v = vlans
	return nil
}
