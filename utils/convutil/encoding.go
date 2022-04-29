package convutil

import (
	"fmt"
	"sync"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/jacksonCLyu/ridi-utils/utils/assignutil"
	"github.com/jacksonCLyu/ridi-utils/utils/rescueutil"
)

var _gbkDecoderPool = sync.Pool{
	New: func() interface{} {
		return simplifiedchinese.GBK.NewDecoder()
	},
}

var _gbkEncoderPool = sync.Pool{
	New: func() interface{} {
		return simplifiedchinese.GBK.NewEncoder()
	},
}

func getGBKEncoder() *encoding.Encoder {
	return _gbkEncoderPool.Get().(*encoding.Encoder)
}

func returnGBKEncoder(encoder *encoding.Encoder) {
	encoder.Reset()
	_gbkEncoderPool.Put(encoder)
}

func getGBKDecoder() *encoding.Decoder {
	return _gbkDecoderPool.Get().(*encoding.Decoder)
}

func returnGBKDecoder(decoder *encoding.Decoder) {
	decoder.Reset()
	_gbkDecoderPool.Put(decoder)
}

// IsGBK 判断是否为 GBK 编码
func IsGBK(data []byte) bool {
	defer rescueutil.Recover(func(err any) {
		fmt.Printf("IsGBK err: %v \n", err)
	})
	decoder := getGBKDecoder()
	defer returnGBKDecoder(decoder)
	_ = assignutil.Assign(decoder.Bytes(data))
	return true
}

func preNUm(data byte) int {
	var mask byte = 0x80
	var num int = 0
	//8bit中首个0bit前有多少个1bits
	for i := 0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}

// IsUtf8 判断是否为 utf-8 编码
func IsUtf8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

// Utf82GBK 将 utf-8 编码转换为 GBK 编码
func Utf82GBK(in []byte) ([]byte, error) {
	encoder := getGBKEncoder()
	defer returnGBKEncoder(encoder)
	return encoder.Bytes(in)
}

// GBK2Utf8 将 GBK 编码转换为 utf-8 编码
func GBK2Utf8(in []byte) ([]byte, error) {
	decoder := getGBKDecoder()
	defer returnGBKDecoder(decoder)
	return decoder.Bytes(in)
}
