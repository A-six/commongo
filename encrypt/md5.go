package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	sum := md5.New()
	sum.Write([]byte(str))
	return hex.EncodeToString(sum.Sum(nil))
}

func MD5Equal(a, md string) bool {
	sum := md5.New()
	sum.Write([]byte(a))
	if hex.EncodeToString(sum.Sum(nil)) == md {
		return true
	}
	return false
}
