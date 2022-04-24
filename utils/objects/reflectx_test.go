package objects

import (
	"testing"

	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

func TestRequireNonNil(t *testing.T) {
	defer rescueutil.Recover(func(err any) {
		if err != "value is nil" {
			t.Errorf("err = %v, want %v", err, "value is nil")
		}
	})
	var nilable interface{}
	v := RequireNonNil(nilable)
	if v == nil {
		t.Error("test failed")
	}
}

func TestIsNil(t *testing.T) {
	var nilable interface{}
	if !IsNil(nilable) {
		t.Error("test failed")
	}
	if IsNil(1) || IsNil(1.0) || IsNil("") || IsNil([]int{}) || IsNil(map[string]int{}) || IsNil(func() {}) {
		t.Error("test failed")
	}
}
