// RemoveFolder
package DaeseongLib

import (
	"bufio"
	"fmt"
	"os"
	_ "path"
	"path/filepath"
	"time"
)

type File struct {
	FileName string
	FileSize int64
	FileMode os.FileMode
	FileTime time.Time
	IsDir    bool
}

func GetFileInfoList(sInPath string) ([]File, error) {
	fileInfoList := []File{}

	err := filepath.Walk(sInPath, func(sOutPath string, fi os.FileInfo, err error) error {
		file := File{}
		file.FileName = fi.Name()
		file.FileSize = fi.Size()
		file.FileMode = fi.Mode()
		file.FileTime = fi.ModTime()
		file.IsDir = fi.IsDir()
		fileInfoList = append(fileInfoList, file)
		return nil
	})

	if err != nil {
		fmt.Printf("[%v]", err)
	}
	return fileInfoList, nil
}

func GetFileList(sInPath string) []string {
	fileList := make([]string, 0, 10)

	err := filepath.Walk(sInPath, func(sOutPath string, fi os.FileInfo, err error) error {
		if !fi.IsDir() {
			fileList = append(fileList, sOutPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("[%v]", err)
	}
	return fileList
}

func GetDirList(sInPath string, sFolderName string) []string {
	fileList := make([]string, 0, 10)

	err := filepath.Walk(sInPath, func(sOutPath string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			if fi.Name() == sFolderName {
				fileList = append(fileList, sOutPath)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("[%v]", err)
	}
	return fileList
}

func DeleteFile(sPath string) error {
	file, err := os.Stat(sPath)
	if err != nil {
		return err
	}

	if file.IsDir() {
		err := os.RemoveAll(sPath)
		if err != nil {
			return err
		}
	} else {
		err := os.Remove(sPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadFile(sPath string) {
	var line int = 0

	file, err := os.Open(sPath)
	if err != nil {
		return
	}

	fileScan := bufio.NewScanner(file)
	for fileScan.Scan() {
		line++
		fmt.Printf("[%d:%s]\n", line, fileScan.Text())
	}
}

func IsDir(sPath string) bool {
	file, err := os.Stat(sPath)
	if err != nil {
		return false
	}
	return file.IsDir()
}

/*
func f1() {

	var sPath string = "C:\\test"
	if !IsDir(sPath) {
		return
	}

	fileList := GetFileList(sPath)
	for i, p := range fileList {
		fmt.Printf("[%d:%s]\n", i, path.Base(p))
	}
}

func f2() {

	var sPath string = "C:\\test"
	var sFolder string = ".svn"
	if !IsDir(sPath) {
		return
	}

	DirList := GetDirList(sPath, sFolder)
	for i, p := range DirList {
		fmt.Printf("[%d:%s]\n", i, p)
	}
}

func f3() {

	var sPath string = "C:\\test"
	if !IsDir(sPath) {
		return
	}

	fileList := GetFileList(sPath)
	for _, p := range fileList {
		ReadFile(p)
	}
}

func f4() {

	var sPath string = "C:\\test"
	if !IsDir(sPath) {
		return
	}

	fileList, _ := GetFileInfoList(sPath)
	for i, fi := range fileList {
		fmt.Printf("[%d:%s  %v  %v  %v %v]\n", i, fi.FileName, fi.FileSize, fi.FileMode, fi.FileTime, fi.IsDir)
	}
}

func f5() {

	var sPath string = "C:\\test"
	var sFolder string = ".svn"
	if !IsDir(sPath) {
		return
	}

	DirList := GetDirList(sPath, sFolder)
	for _, p := range DirList {
		DeleteFile(p)
	}
	fmt.Println("End")
}

func main() {
	//f5()
}
*/
