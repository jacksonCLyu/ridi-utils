package compressutil

import (
	"archive/zip"
	"bytes"
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

func Zip(dst, src string) (err error) {
	defer rescueutil.Recover(func(e any) {
		err = fmt.Errorf("%v", e)
	})
	// 创建准备写入的文件
	fw := assignutil.Assign(os.Create(dst))
	defer fw.Close()

	// 通过 fw 来创建 zip.Write
	zw := zip.NewWriter(fw)
	// 检测一下是否成功关闭
	defer errcheck.CheckAndPanic(zw.Close())

	// 下面来将文件写入 zw ，因为有可能会有很多个目录及文件，所以递归处理
	errcheck.CheckAndPanic(filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		defer rescueutil.Recover(func(e any) {
			err = fmt.Errorf("%v", e)
		})
		if errBack != nil {
			return errBack
		}

		// 通过文件信息，创建 zip 的文件信息
		fh := assignutil.Assign(zip.FileInfoHeader(fi))

		// 替换文件信息中的文件名
		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))

		// 这步开始没有加，会发现解压的时候说它不是个目录
		if fi.IsDir() {
			fh.Name += "/"
		}

		// 写入文件信息，并返回一个 Write 结构
		w := assignutil.Assign(zw.CreateHeader(fh))

		// 检测，如果不是标准文件就只写入头信息，不写入文件数据到 w
		// 如目录，也没有数据需要写
		if !fh.Mode().IsRegular() {
			return nil
		}

		// 打开要压缩的文件
		fr := assignutil.Assign(os.Open(path))
		defer fr.Close()

		// 将打开的文件 Copy 到 w
		n := assignutil.Assign(io.Copy(w, fr))

		// 输出压缩的内容
		fmt.Printf("成功压缩文件： %s, 共写入了 %d 个字符的数据\n", path, n)
		return
	}))
	return
}

func UnZip(dst, src string) (err error) {
	defer rescueutil.Recover(func(e any) {
		err = fmt.Errorf("%v", e)
	})
	// 打开压缩文件，这个 zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	zr := assignutil.Assign(zip.OpenReader(src))
	defer zr.Close()

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if dst != "" {
		errcheck.CheckAndPanic(os.MkdirAll(dst, 0755))
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
		path := filepath.Join(dst, decodeName)

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
	defer rescueutil.Recover(func(err any) {
		fmt.Printf("%v", err)
	})
	// 获取到 Reader
	fr := assignutil.Assign(file.Open())
	defer fr.Close()

	// 创建要写出的文件对应的 Write
	fw := assignutil.Assign(os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode()))
	defer fw.Close()

	_ = assignutil.Assign(io.Copy(fw, fr))
}
