package optional

import (
	"testing"

	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

type testStruct struct {
	code  int
	value string
}

func TestOf(t *testing.T) {
	o := Of(1)
	if o.Get() != 1 {
		t.Errorf("expected 1, got %d", o.Get())
	}
}

func TestOfNilable(t *testing.T) {
	o := OfNilable(1)
	if o.Get() != 1 {
		t.Errorf("expected 1, got %d", o.Get())
	}
	var nilable *testStruct
	n := OfNilable(nilable)
	if n.IsPresent() {
		t.Errorf("expected nilable to be nil")
	}
}

func TestFilter(t *testing.T) {
	o := Of(1)
	if o.Filter(func(i int) bool { return i == 1 }).Get() != 1 {
		t.Errorf("expected 1, got %d", o.Filter(func(i int) bool { return i == 1 }).Get())
	}
}

func TestOrElse(t *testing.T) {
	o := Of(1)
	if o.OrElse(2) != 1 {
		t.Errorf("expected 1, got %d", o.OrElse(2))
	}
	var nilable *testStruct
	n := OfNilable(nilable)
	orElse := n.OrElse(&testStruct{1, "test"})
	if orElse.code != 1 || orElse.value != "test" {
		t.Errorf("expected code: 1 and msg: `test`, got %d, %s", orElse.code, orElse.value)
	}
}

func TestOrElseGet(t *testing.T) {
	o := Of(1)
	if o.OrElseGet(func() int { return 2 }) != 1 {
		t.Errorf("expected 1, got %d", o.OrElseGet(func() int { return 2 }))
	}
	var nilable *testStruct
	n := OfNilable(nilable)
	orElse := n.OrElseGet(func() *testStruct { return &testStruct{1, "test"} })
	if orElse.code != 1 || orElse.value != "test" {
		t.Errorf("expected code: 1 and msg: `test`, got %d, %s", orElse.code, orElse.value)
	}
}

func TestOrElsePanic(t *testing.T) {
	defer rescueutil.Recover(func(e any) {
		t.Errorf("expected panic, got %v", e)
	})
	var nilable *testStruct
	n := OfNilable(nilable)
	n.OrElsePanic(func() any { return "oh! the value is nil" })
}

func TestIsPresent(t *testing.T) {
	o := Of(1)
	if !o.IsPresent() {
		t.Errorf("expected true, got false")
	}
	var nilable *testStruct
	n := OfNilable(nilable)
	if n.IsPresent() {
		t.Errorf("expected false, got true")
	}
}

func TestIfPresent(t *testing.T) {
	o := Of(1)
	called := false
	o.IfPresent(func(_ int) {
		called = true
	})
	if !called {
		t.Errorf("expected true, got false")
	}
	var nilable *testStruct
	n := OfNilable(nilable)
	n.IfPresent(func(_ *testStruct) {
		t.Errorf("expected panic, got true")
	})
}
