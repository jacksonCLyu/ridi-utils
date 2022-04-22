package compressutil

import (
	"fmt"
	"testing"

	"github.com/jacksonCLyu/ridi-utils/utils/convutil"
)

func TestGzipAndUnzip(t *testing.T) {
	s := "hello world"
	sb, err := Gzip(convutil.Str2bytes(s))
	if err != nil {
		t.Errorf("Gzip failed: %v", err)
	}
	b, err := Gunzip(sb)
	if err != nil {
		t.Errorf("Gunzip failed: %v", err)
	}
	if string(b) != "hello world" {
		t.Errorf("Gunzip failed")
	}
}

func TestGzipStringAndUnzip(t *testing.T) {
	s := "hello world"
	sb, err := GzipString(s)
	if err != nil {
		t.Errorf("GzipString failed: %v", err)
	}
	b, err := GunzipString(sb)
	if err != nil {
		t.Errorf("Gunzip failed: %v", err)
	}
	if string(b) != "hello world" {
		t.Errorf("Gunzip failed")
	}
}

func TestGzipThenBase64AndUnzip(t *testing.T) {
	s := "hello world"
	sb, err := GzipThenBase64(s)
	if err != nil {
		t.Errorf("GzipThenBase64 failed: %v", err)
	} else {
		fmt.Println(sb)
	}
	b, err := GunzipWithBase64(sb)
	if err != nil {
		t.Errorf("GunzipWithBase64 failed: %v", err)
	} else {
		fmt.Println(string(b))
	}
	if string(b) != "hello world" {
		t.Errorf("GunzipWithBase64 failed")
	}
}
