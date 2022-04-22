package rescueutil

import "testing"

func TestRecover(t *testing.T) {
	defer Recover(func(err any) {
		t.Logf("Recover: %v", err)
	})
	panic("test panic")
}
