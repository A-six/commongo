package encrypt

import (
	"fmt"
	"testing"
)

func TestSha256(t *testing.T) {
	fmt.Println(Sha256("123456"))
	fmt.Println(Sha256("admin"))
	fmt.Println(Sha256("admin@123"))
}

func TestSha256Equal(t *testing.T) {
	md := Sha256("123456")
	fmt.Println(Sha256Equal("123456", md))
}
