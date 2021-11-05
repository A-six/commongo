package encrypt

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	fmt.Println(MD5("123456"))
	fmt.Println(MD5("admin"))
	fmt.Println(MD5("admin@123"))
}

func TestMD5Equal(t *testing.T) {
	md := MD5("123456")
	fmt.Println(MD5Equal("123456", md))
}
