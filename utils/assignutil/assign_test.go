package assignutil

import (
	"errors"
	"reflect"
	"testing"

	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

func TestAssignAndRecover(t *testing.T) {
	defer rescueutil.Recover(func(err any) {
		t.Logf("Recover: %v", err)
	})
	var a int = 1
	var b string = "hello world"
	var c bool = true
	var d float64 = 1.1
	var e []byte = []byte("hello world")
	var f map[string]string = map[string]string{"a": "b"}
	var err error = errors.New("test error")
	a1 := Assign(a, nil)
	if a1 != a {
		t.Errorf("Assign failed")
	}
	b1 := Assign(b, nil)
	if b1 != b {
		t.Errorf("Assign failed")
	}
	c1 := Assign(c, nil)
	if c1 != c {
		t.Errorf("Assign failed")
	}
	d1 := Assign(d, nil)
	if d1 != d {
		t.Errorf("Assign failed")
	}
	e1 := Assign(e, nil)
	if !reflect.DeepEqual(e1, e) {
		t.Errorf("Assign failed")
	}
	f1 := Assign(f, nil)
	if !reflect.DeepEqual(f1, f) {
		t.Errorf("Assign failed")
	}
	// panic recover
	a2 := Assign(a, err)
	if a2 != a {
		t.Errorf("Assign failed")
	}
}
