package base64util

import "encoding/base64"

// EncodeBytes encode bytes to string
func EncodeBytes(in []byte) string {
	return base64.StdEncoding.EncodeToString(in)
}

// DecodeString decode string to bytes
func DecodeString(in string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(in)
}
