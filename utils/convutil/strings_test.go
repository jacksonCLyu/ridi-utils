package convutil

import "testing"

func TestBytes2Str(t *testing.T) {
	b := []byte("hello world")
	s := Bytes2str(b)
	if s != "hello world" {
		t.Errorf("Str2bytes failed")
	}
}

func TestStr2Bytes(t *testing.T) {
	s := "hello world"
	b := Str2bytes(s)
	if string(b) != "hello world" {
		t.Errorf("Str2bytes failed")
	}
}
