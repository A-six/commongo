package encrypt

import (
	"fmt"
	"testing"
)

func TestSha1(t *testing.T) {
	fmt.Println(Sha1("123456"))
	fmt.Println(Sha1("admin"))
	fmt.Println(Sha1("admin@123"))
}

func TestSha1Equal(t *testing.T) {
	md := Sha1("123456")
	fmt.Println(Sha1Equal("123456", md))
}
