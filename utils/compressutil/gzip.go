package compressutil

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/jacksonCLyu/ridi-utils/utils/assignutil"
	"github.com/jacksonCLyu/ridi-utils/utils/base64util"
	"github.com/jacksonCLyu/ridi-utils/utils/convutil"
	"github.com/jacksonCLyu/ridi-utils/utils/errcheck"
	"github.com/jacksonCLyu/ridi-utils/utils/regexutil"
	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

var _bytesBufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

var _gzipWriterPool = sync.Pool{
	New: func() any {
		return gzip.NewWriter(&bytes.Buffer{})
	},
}

func getBytesBuf() *bytes.Buffer {
	return _bytesBufferPool.Get().(*bytes.Buffer)
}

func putBytesBuf(b *bytes.Buffer) {
	b.Reset()
	_bytesBufferPool.Put(b)
}

func getGzipWriter() *gzip.Writer {
	return _gzipWriterPool.Get().(*gzip.Writer)
}

func putGzipWriter(w *gzip.Writer) {
	w.Reset(&bytes.Buffer{})
	_gzipWriterPool.Put(w)
}

// Gzip returns the gzip of the data.
func Gzip(data []byte) (out []byte, e error) {
	defer rescueutil.Recover(func(err any) {
		fmt.Printf("Gzip error: %v \n", err)
		e = err.(error)
	})
	buf := getBytesBuf()
	defer putBytesBuf(buf)
	w := getGzipWriter()
	defer putGzipWriter(w)
	w.Reset(buf)
	assignutil.Assign(w.Write(data))
	errcheck.CheckAndPanic(w.Flush())
	errcheck.CheckAndPanic(w.Close())
	out = buf.Bytes()
	return
}

// GzipString returns the gzip string of the data.
func GzipString(data string) (out string, e error) {
	defer rescueutil.Recover(func(err any) {
		fmt.Printf("GzipString error: %v \n", err)
		e = err.(error)
	})
	buf := getBytesBuf()
	defer putBytesBuf(buf)
	w := getGzipWriter()
	defer putGzipWriter(w)
	w.Reset(buf)
	assignutil.Assign(w.Write([]byte(data)))
	errcheck.CheckAndPanic(w.Flush())
	errcheck.CheckAndPanic(w.Close())
	out = convutil.Bytes2str(buf.Bytes())
	return
}

// Gunzip returns the gunzip of the data.
func Gunzip(data []byte) (out []byte, e error) {
	defer rescueutil.Recover(func(err any) {
		fmt.Printf("Gunzip error: %v \n", err)
		e = err.(error)
	})
	buf := getBytesBuf()
	defer putBytesBuf(buf)
	buf.Write(data)
	r := assignutil.Assign(gzip.NewReader(buf))
	defer func() {
		errcheck.CheckAndPanic(r.Close())
	}()
	out = assignutil.Assign(ioutil.ReadAll(r))
	return
}

// GunzipString returns the gunzip string of the data.
func GunzipString(data string) (out string, e error) {
	defer rescueutil.Recover(func(err any) {
		fmt.Printf("GunzipString error: %v \n", err)
		e = err.(error)
	})
	buf := getBytesBuf()
	defer putBytesBuf(buf)
	buf.Write(convutil.Str2bytes(data))
	r := assignutil.Assign(gzip.NewReader(buf))
	defer func() {
		errcheck.CheckAndPanic(r.Close())
	}()
	out = convutil.Bytes2str(assignutil.Assign(ioutil.ReadAll(r)))
	return
}

// GzipThenBase64 returns the gzip then base64 of the data.
func GzipThenBase64(data string) (out string, e error) {
	defer rescueutil.Recover(func(err any) {
		fmt.Printf("GzipString error: %v \n", err)
		e = err.(error)
	})
	buf := getBytesBuf()
	defer putBytesBuf(buf)
	w := getGzipWriter()
	defer putGzipWriter(w)
	w.Reset(buf)
	assignutil.Assign(w.Write(convutil.Str2bytes(data)))
	errcheck.CheckAndPanic(w.Flush())
	errcheck.CheckAndPanic(w.Close())
	out = base64util.EncodeBytes(buf.Bytes())
	return
}

// GunzipWithBase64 returns the gunzip then base64 of the data.
func GunzipWithBase64(data string) (out string, e error) {
	defer rescueutil.Recover(func(err any) {
		fmt.Printf("GunzipString error: %v \n", err)
		e = err.(error)
	})
	data = regexutil.SpaceReg.ReplaceAllString(data, "")
	buf := getBytesBuf()
	defer putBytesBuf(buf)
	b := assignutil.Assign(base64util.DecodeString(data))
	buf.Write(b)
	r := assignutil.Assign(gzip.NewReader(buf))
	defer func() {
		errcheck.CheckAndPanic(r.Close())
	}()
	out = convutil.Bytes2str(assignutil.Assign(ioutil.ReadAll(r)))
	return
}
