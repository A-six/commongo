package encrypt

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(str string) string {
	sum := sha1.New()
	sum.Write([]byte(str))
	return hex.EncodeToString(sum.Sum(nil))
}

func Sha1Equal(a, sh string) bool {
	sum := sha1.New()
	sum.Write([]byte(a))
	if hex.EncodeToString(sum.Sum(nil)) == sh {
		return true
	}
	return false
}
