package fileutil

import (
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// GetFileFromURL get file from url
func GetFileFromURL(url *url.URL) (*os.File, error) {
	if url == nil || url.Scheme != "file" {
		return nil, errors.New("file URL is nil or URL scheme is not `file`")
	}
	path := url.Path
	path = strings.ReplaceAll(path, "/", string(filepath.Separator))
	open, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return open, nil
}
