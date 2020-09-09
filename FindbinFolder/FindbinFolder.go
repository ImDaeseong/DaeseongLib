// FindbinFolder
package DaeseongLib

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	directories = make(map[string]bool)
)

func IsBinDir(path string) bool {

	fileStat, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileStat.IsDir()
}

func getRootDrives() (drives []string) {
	for _, val := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		_, err := os.Open(string(val) + ":\\")
		if err == nil {
			drives = append(drives, string(val))
		}
	}
	return
}

func getFoldername(sPath string) string {
	filename := strings.Split(sPath, "\\")
	return filename[len(filename)-1]
}

func WriteFolderString(sPath, sText string) {

	file, err := os.OpenFile(sPath, os.O_RDWR|os.O_APPEND, 0660)
	if os.IsNotExist(err) {
		file, err = os.Create(sPath)
	}
	defer file.Close()

	if err != nil {
		return
	}

	n, err := io.WriteString(file, sText)
	if err != nil {
		fmt.Println(n, err)
		return
	}
}

func findDirectory(sDirectory string) {

	err := filepath.Walk(sDirectory, func(filePath string, f os.FileInfo, err error) error {

		if IsBinDir(filePath) {
			if getFoldername(filePath) == "bin" || getFoldername(filePath) == "obj" || getFoldername(filePath) == "svn" {
				directories[filePath] = false
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func FindAllFolder() {

	fmt.Println(time.Now().Format("2006:01:02 15:04:05"))

	drives := getRootDrives()

	var wait sync.WaitGroup

	wait.Add(len(drives))

	for _, dr := range drives {
		go func(dr string) {
			sVal := fmt.Sprintf("%s:\\", dr)
			findDirectory(sVal)
			wait.Done()
		}(dr)
	}
	wait.Wait()

	for folder := range directories {
		WriteFolderString("folderlist.txt", folder+"\n")
	}

	fmt.Println(time.Now().Format("2006:01:02 15:04:05"))
}

/*
func main() {

	FindAllFolder()
}
*/
