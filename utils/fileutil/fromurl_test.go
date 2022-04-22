package fileutil

import (
	"net/url"
	"os"
	"testing"
)

func TestGetFileFromURL(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Errorf("Getwd failed")
	}
	url, _ := url.Parse("file:///" + pwd + "/testdata/test.txt")
	file, err := GetFileFromURL(url)
	if err != nil || file == nil {
		t.Errorf("GetFileFromURL failed")
	}
	t.Logf("GetFileFromURL success: %s", file.Name())
}
