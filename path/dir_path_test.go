package path

import (
	"fmt"
	"testing"
)

func TestGetDirAndFile(t *testing.T) {
	ds := GetDirAndFile("C:\\Users")
	for _, v := range ds {
		fmt.Println(v)
	}
}

func TestFiles(t *testing.T) {
	ds := Files("C:\\Users")
	for _, v := range ds {
		fmt.Println(v)
	}
}

func TestFileFullPath(t *testing.T) {
	ds := FileFullPath("C:\\Users")
	for _, v := range ds {
		fmt.Println(v)
	}
}

func TestDirs(t *testing.T) {
	ds := Dirs("C:\\Users")
	for _, v := range ds {
		fmt.Println(v)
	}
}

func TestDirFullPath(t *testing.T) {
	ds := DirFullPath("C:\\Users")
	for _, v := range ds {
		fmt.Println(v)
	}
}
