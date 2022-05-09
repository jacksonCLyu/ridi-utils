package fileutil

import (
	"os"
	"testing"
)

func TestGetFileModTime(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Errorf("Getwd failed")
	}
	fileName := pwd + "/testdata/test.txt"
	modTime, err := GetFileModTime(fileName)
	if err != nil {
		t.Errorf("GetFileModTime error: %v", err)
	}
	t.Logf("modTime: %v", modTime)
}
