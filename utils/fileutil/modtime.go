package fileutil

import (
	"fmt"
	"os"
	"time"
)

// GetFileModTime 获取文件修改时间
func GetFileModTime(fileName string) (int64, error) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Open file %s error: %v \n", fileName, err)
		return time.Now().Local().Unix(), err
	}
	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Printf("Get file stat error %v", err)
		return time.Now().Local().Unix(), err
	}
	return fileInfo.ModTime().Local().Unix(), nil
}
