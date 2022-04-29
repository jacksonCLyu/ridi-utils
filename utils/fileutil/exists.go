package fileutil

import (
	"os"
)

// IsExists checks if a file or directory exists.
func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsNotExists checks if a file or directory does not exist.
func IsNotExists(path string) bool {
	return !IsExists(path)
}
