package compressutil

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/jacksonCLyu/ridi-utils/utils/assignutil"
	"github.com/jacksonCLyu/ridi-utils/utils/errcheck"
	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

// Zip zip src directory to zipFilePath file
func Zip(zipFilePath string, srcDirPath string) (err error) {
	defer rescueutil.Recover(func(e any) {
		if underlineErr, ok := e.(error); ok {
			err = underlineErr
		} else {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	})
	// if zipFile exists, remove it.
	errcheck.CheckAndPanic(os.RemoveAll(zipFilePath))

	// create a new file for zip
	zipFile := assignutil.Assign(os.Create(zipFilePath))
	defer func() {
		errcheck.CheckAndPanic(zipFile.Close())
	}()

	// create a new zip writer by the file
	zw := zip.NewWriter(zipFile)
	defer func() {
		errcheck.CheckAndPanic(zw.Close())
	}()

	// walk dir and write zip file
	return filepath.Walk(srcDirPath, func(path string, fi os.FileInfo, errBack error) (err error) {
		defer rescueutil.Recover(func(e any) {
			if underlineErr, ok := e.(error); ok {
				err = underlineErr
			} else {
				err = errors.New(fmt.Sprintf("%v", e))
			}
		})

		errcheck.CheckAndPanic(errBack)

		if path == srcDirPath {
			return
		}

		// create zip header
		fh := assignutil.Assign(zip.FileInfoHeader(fi))
		// trim prefix
		fh.Name = strings.TrimPrefix(path, srcDirPath+string(filepath.Separator))

		// if dir add filepath.Separator
		if fi.IsDir() {
			fh.Name += "/"
		} else {
			fh.Method = zip.Deflate
		}

		// create zip writer
		w := assignutil.Assign(zw.CreateHeader(fh))

		// if not standard file like: directory, just return
		if !fh.Mode().IsRegular() {
			return
		}

		// open file
		fr := assignutil.Assign(os.Open(path))
		defer fr.Close()

		// Copy file to zip writer
		n := assignutil.Assign(io.Copy(w, fr))

		// print file name and size
		fmt.Printf("成功压缩文件： %s, 共写入了 %d 个字符的数据\n", path, n)
		return
	})
}

// UnZip 解压缩，将 src 文件解压缩到 dst 目录中，如果 dst 不存在，则创建。默认传空则 dst 为当前目录
func UnZip(unzipDirPath string, zipFilePath string) (err error) {
	defer rescueutil.Recover(func(e any) {
		err = fmt.Errorf("%v", e)
	})

	// if unzipDirPath exists, remove it.
	errcheck.CheckAndPanic(os.RemoveAll(unzipDirPath))

	// 打开压缩文件，这个 zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	zr := assignutil.Assign(zip.OpenReader(zipFilePath))
	defer zr.Close()

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if unzipDirPath != "" {
		errcheck.CheckAndPanic(os.MkdirAll(unzipDirPath, 0755))
	}

	// 遍历 zr ，将文件写入到磁盘
	var decodeName string
	for _, file := range zr.File {
		//如果标致位是0  则是默认的本地编码   默认为gbk
		i := bytes.NewReader([]byte(file.Name))
		decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
		content, err := ioutil.ReadAll(decoder)
		if err != nil {
			//如果标志为是 1 << 11也就是 2048  则是utf-8编码
			decodeName = file.Name
		} else {
			decodeName = string(content)
		}
		path := filepath.Join(unzipDirPath, decodeName)

		// 如果是目录，就创建目录
		if file.FileInfo().IsDir() {
			errcheck.CheckAndPanic(os.MkdirAll(path, file.Mode()))
			// 因为是目录，跳过当前循环，因为后面都是文件的处理
			continue
		}

		doUnzipCopy(file, path)
	}
	return
}

func doUnzipCopy(file *zip.File, path string) {
	// 获取到 Reader
	fr := assignutil.Assign(file.Open())
	defer fr.Close()

	// 创建要写出的文件对应的 Write
	fw := assignutil.Assign(os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode()))
	defer fw.Close()

	_ = assignutil.Assign(io.Copy(fw, fr))
}
