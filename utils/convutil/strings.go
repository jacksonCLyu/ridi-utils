package convutil

import "unsafe"

// Str2bytes (efficient)
func Str2bytes(oriStr string) (target []byte) {
	if len(oriStr) == 0 {
		return
	}
	x := (*[2]uintptr)(unsafe.Pointer(&oriStr))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2str (efficient)
func Bytes2str(b []byte) (target string) {
	if len(b) == 0 {
		return
	}
	return *(*string)(unsafe.Pointer(&b))
}
