package path

import (
	"io/ioutil"
	"os"
)

//获取指定目录下的所有目录和文件
func GetDirAndFile(dirPth string) []string {
	patas := []string{}
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return patas
	}
	//PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		patas = append(patas, fi.Name())
	}
	return patas
}

//获取指定目录下的所有文件
func Files(dirPth string) []string {
	patafile := []string{}
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return patafile
	}
	for _, fi := range dir {
		if !fi.IsDir() { // 忽略目录
			patafile = append(patafile, fi.Name())
		}
	}
	return patafile
}

//获取指定目录下的所有文件  完整路径
func FileFullPath(dirPth string) []string {
	retFlie := []string{}
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return retFlie
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if !fi.IsDir() { // 忽略目录
			retFlie = append(retFlie, dirPth+PthSep+fi.Name())
		}
	}
	return retFlie
}

//获取指定目录下的所有目录
func Dirs(dirPth string) []string {
	patas := []string{}
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return patas
	}
	for _, fi := range dir {
		if fi.IsDir() {
			patas = append(patas, fi.Name())
		}
	}
	return patas
}

//获取指定目录下的所有目录 完整路径
func DirFullPath(dirPth string) []string {
	patas := []string{}
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return patas
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			patas = append(patas, dirPth+PthSep+fi.Name())
		}
	}
	return patas
}
