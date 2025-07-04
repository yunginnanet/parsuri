package events

import (
	"encoding/json"
	"testing"
)

func TestVLANUnmarshalJSON(t *testing.T) {
	t.Run("UnmarshalJSONParsesSingleVLANValueCorrectly", func(t *testing.T) {
		input := `42`
		var vlan VLAN
		err := json.Unmarshal([]byte(input), &vlan)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(vlan) != 1 || vlan[0] != 42 {
			t.Errorf("expected VLAN to contain [42], got %v", vlan)
		}
	})

	t.Run("UnmarshalJSONParsesVLANArrayCorrectly", func(t *testing.T) {
		input := `[42, 43, 44]`
		var vlan VLAN
		err := json.Unmarshal([]byte(input), &vlan)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		expected := VLAN{42, 43, 44}
		for i, v := range expected {
			if vlan[i] != v {
				t.Errorf("expected VLAN[%d] to be %d, got %d", i, v, vlan[i])
			}
		}
	})

	t.Run("UnmarshalJSONReturnsErrorForInvalidVLANData", func(t *testing.T) {
		input := `"invalid"`
		var vlan VLAN
		err := json.Unmarshal([]byte(input), &vlan)
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})

	t.Run("UnmarshalJSONHandlesEmptyInputGracefully", func(t *testing.T) {
		input := `[]`
		var vlan VLAN
		err := json.Unmarshal([]byte(input), &vlan)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(vlan) != 0 {
			t.Errorf("expected VLAN to be empty, got %v", vlan)
		}
	})
}
