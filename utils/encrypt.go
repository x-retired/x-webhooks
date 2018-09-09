package utils

import (
	"crypto/md5"
	"fmt"
)

// Md5String return md5 string
func Md5String(raw string) string {
	data := []byte(raw)
	return fmt.Sprintf("%x", md5.Sum(data))
}

// Md5Verify md5 and raw string verify
func Md5Verify(md5str string, raw string) bool {
	if Md5String(raw) == md5str {
		return true
	}
	return false
}
