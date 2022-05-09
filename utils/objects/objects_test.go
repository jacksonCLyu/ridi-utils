package objects

import (
	"testing"

	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

type testStruct struct {
	code  int
	value string
}

func TestIsNil(t *testing.T) {
	var a *int
	if !IsNil(a) {
		t.Error("a is not nil")
	}
	var nilable *testStruct
	if !IsNil(nilable) {
		t.Error("nilable is not nil")
	}
	var b testStruct
	if IsNil(b) {
		t.Error("t is nil")
	}
}

func TestRequireNonNil(t *testing.T) {
	defer rescueutil.Recover(func(err any) {
		t.Errorf("expected panic, got `%v`", err)
	})
	var a *int
	RequireNonNil(a)
}

func TestRequireNonNilM(t *testing.T) {
	defer rescueutil.Recover(func(err any) {
		t.Errorf("expected panic, got `%v`", err)
	})
	var a *testStruct
	RequireNonNilM(a, "a is nil")
}
