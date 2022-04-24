package optional

import (
	"errors"
	"testing"

	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

func TestOf(t *testing.T) {
	o := Of(1)
	if o.IsPresent() {
		if o.Get() != 1 {
			t.Errorf("o.Get() = %v, want %v", o.Get(), 1)
		}
	} else {
		t.Errorf("o.IsPresent() = %v, want %v", o.IsPresent(), true)
	}
}

func TestOfNilable(t *testing.T) {
	o := OfNilable(1)
	if o.IsPresent() {
		if o.Get() != 1 {
			t.Errorf("o.Get() = %v, want %v", o.Get(), 1)
		}
	} else {
		t.Errorf("o.IsPresent() = %v, want %v", o.IsPresent(), true)
	}
	var nilable interface{}
	v := OfNilable(nilable)
	if v.IsPresent() {
		t.Errorf("o.IsPresent() = %v, want %v", v.IsPresent(), false)
	} else {
		t.Logf("o.IsPresent() = %v, want %v", v.IsPresent(), false)
	}
}

func TestIsPresent(t *testing.T) {
	o := Of(1)
	if !o.IsPresent() {
		t.Errorf("o.IsPresent() = %v, want %v", o.IsPresent(), true)
	}
}

func TestIfPresent(t *testing.T) {
	o := Of(1)
	o.IfPresent(func(value int) {
		if value != 1 {
			t.Errorf("value = %v, want %v", value, 1)
		}
	})
}

func TestFilter(t *testing.T) {
	o := Of(1)
	o.filter(func(value int) bool {
		return value == 1
	}).IfPresent(func(value int) {
		if value != 1 {
			t.Errorf("value = %v, want %v", value, 1)
		}
	})
	o.filter(func(value int) bool {
		return value == 2
	}).IfPresent(func(value int) {
		// if filtered return zero value of T
		if value != 0 {
			t.Errorf("value = %v, want %v", value, 1)
		}
	})
}

func TestGet(t *testing.T) {
	o := Of(1)
	if o.Get() != 1 {
		t.Errorf("o.Get() = %v, want %v", o.Get(), 1)
	}
}

func TestOrElse(t *testing.T) {
	var nilable interface{}
	o := OfNilable(nilable)
	if o.OrElse(2) != 2 {
		t.Errorf("o.OrElse(2) = %v, want %v", o.OrElse(2), 2)
	}
}

func TestOrElseThrow(t *testing.T) {
	var nilable interface{}
	o := OfNilable(nilable)
	defer rescueutil.Recover(func(e any) {
		if e.(error).Error() != "no value present" {
			t.Errorf("e.Error() = %v, want %v", e.(error).Error(), "no value present")
		}
	})
	o.OrElseThrow(func() error {
		return errors.New("no value present")
	})
}
