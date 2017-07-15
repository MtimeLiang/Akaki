package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	if len(s) == 0 {
		return ""
	}
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}
