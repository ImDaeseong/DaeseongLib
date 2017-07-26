// DelDirInfo
package DaeseongLib

import (
	_ "fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

var (
	allfileList = []string{}
)

func FindFiles(sPath string) {

	files, err := ioutil.ReadDir(sPath)
	if err != nil {
		panic(err)
	}

	for _, f := range files {

		sDir := filepath.Join(sPath, f.Name())
		if f.IsDir() {
			FindFiles(sDir)
		} else {
			allfileList = append(allfileList, sDir)
		}
	}
}

func FindFileList(sPath string) []string {
	fileList := []string{}
	filepath.Walk(sPath, func(p string, f os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !f.IsDir() {
			fileList = append(fileList, p)
		}
		return nil
	})
	sort.Strings(fileList)
	return fileList
}

func RemoveDir(sPath string) bool {

	err := os.RemoveAll(sPath)
	if err != nil {
		return false
	}
	return true
}

func FindDirList(sPath string) []string {
	fileList := []string{}

	filepath.Walk(sPath, func(p string, f os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if f.IsDir() {
			fileList = append(fileList, p)
		}
		return nil
	})
	return fileList
}

/*
func f1() {
	FindFiles("C:\\test")
	for _, f := range allfileList {
		fmt.Println(f)
	}
}

func f2() {
	fileList := FindFileList("C:\\test")
	for _, f := range fileList {
		fmt.Println(f)
	}
}

func f3() {
	RemoveDir("C:\\test")
}

func f4() {
	DirList := FindDirList("C:\\test")
	for _, dir := range DirList {
		RemoveDir(dir)
	}
}

func main() {
	//f1()
	//f2()
	//f4()
}
*/
