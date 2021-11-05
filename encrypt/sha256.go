package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(str string) string {
	sum := sha256.New()
	sum.Write([]byte(str))
	return hex.EncodeToString(sum.Sum(nil))
}

func Sha256Equal(a, sh string) bool {
	sum := sha256.New()
	sum.Write([]byte(a))
	if hex.EncodeToString(sum.Sum(nil)) == sh {
		return true
	}
	return false
}
