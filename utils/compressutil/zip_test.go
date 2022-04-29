package compressutil

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jacksonCLyu/ridi-utils/utils/assignutil"
	"github.com/jacksonCLyu/ridi-utils/utils/fileutil"
	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

func TestZip(t *testing.T) {
	defer rescueutil.Recover(func(err any) {
		t.Errorf("TestZip err: %v \n", err)
	})
	pwd := assignutil.Assign(os.Getwd())
	t.Logf("pwd: %v", pwd)
	zipFilePath := filepath.Join(pwd, "testdata", "zipped", "test.zip")
	srcDirPath := filepath.Join(pwd, "testdata", "unzip")
	srcFileName := filepath.Join(srcDirPath, "test.txt")
	if fileutil.IsNotExists(srcFileName) {
		createTestFile(srcFileName)
	}
	Zip(zipFilePath, srcDirPath)
}

func TestUnzip(t *testing.T) {
	defer rescueutil.Recover(func(err any) {
		t.Errorf("TestUnzip err: %v \n", err)
	})
	pwd := assignutil.Assign(os.Getwd())
	t.Logf("pwd: %v", pwd)
	unzipFilePath := filepath.Join(pwd, "testdata", "unzip")
	zipFile := filepath.Join(pwd, "testdata", "zipped", "test.zip")
	// 若测试待解压文件不存在，则创建
	if fileutil.IsNotExists(zipFile) {
		// 若测试待压缩文件不存在，则创建
		unzipfile := filepath.Join(unzipFilePath, "test.txt")
		if fileutil.IsNotExists(unzipfile) {
			createTestFile(unzipfile)
		}
		Zip(zipFile, unzipFilePath)
	}
	UnZip(unzipFilePath, zipFile)
}

func createTestFile(testFile string) {
	srcFile := assignutil.Assign(os.Create(testFile))
	defer srcFile.Close()
	_ = assignutil.Assign(srcFile.WriteString("hello world"))
}
