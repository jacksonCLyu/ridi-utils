package base64util

import (
	"testing"
)

func TestEncodingAndDecoding(t *testing.T) {
	b := []byte("hello world")
	s := EncodeBytes(b)
	if s != "aGVsbG8gd29ybGQ=" {
		t.Errorf("EncodeBytes failed")
	}
	b2, err := DecodeString(s)
	if err != nil {
		t.Errorf("DecodeString failed")
	}
	if string(b2) != "hello world" {
		t.Errorf("DecodeString failed")
	}
}
