package ip

import "testing"

func TestMyIp_GetExternalIp(t *testing.T) {
	ipe := &myIp{}
	_, err := ipe.GetExternalIp()
	if err != nil {
		t.Error(err)
	}
}
