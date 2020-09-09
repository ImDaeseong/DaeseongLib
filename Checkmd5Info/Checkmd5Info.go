// Checkmd5Info
package DaeseongLib

import (
	//DaeseongLib "DaeseongLib/lib"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	_ "time"
)

func IsFileOpen(sPath string) bool {
	file, err := os.Open(sPath)
	if err != nil {
		return false
	}
	defer file.Close()
	return true
}

func GetFileMd5Info(sFilePath string) string {
	if !IsFileOpen(sFilePath) {
		return ""
	}
	bytes, _ := ioutil.ReadFile(sFilePath)
	valHash := md5.New()
	io.WriteString(valHash, string(bytes))
	return fmt.Sprintf("%x", valHash.Sum(nil))
}

func mergeMap(a map[string]string, b map[string]string) {
	for k, v := range b {
		a[k] = v
	}
}

func findAllMapfile(sPath string) map[string]string {

	m := map[string]string{}

	directory, _ := os.Open(sPath)
	dirlst, _ := directory.Readdir(-1)

	for _, file := range dirlst {

		sPath := sPath + "/" + file.Name()

		if file.IsDir() {

			mergeMap(m, findAllMapfile(sPath))

		} else {
			key := sPath
			//fmt.Printf("key:%s \n", key)
			m[key] = GetFileMd5Info(sPath)
		}
	}
	return m
}

/*
func f1() {

	sPath := "C:\\Go\\src\\DaeseongLib\\lib"
	if DaeseongLib.IsDir(sPath) {
		fmt.Println("Dir")
	} else {
		fmt.Println("Not Dir")
	}
}

func f2() {

	path1 := "C:\\Go\\src\\DaeseongLib\\lib\\BluestackInfo.go"
	path2 := "D:\\Go\\src\\DaeseongLib\\lib\\BluestackInfo.go"

	for {

		if IsFileOpen(path1) && IsFileOpen(path2) {

			md5val1 := GetFileMd5Info(path1)
			md5val2 := GetFileMd5Info(path2)
			if md5val1 == md5val2 {
				fmt.Println("sama")
			} else {
				fmt.Println("tidak sama")
			}
		} else {

			fmt.Println("path1:" + GetFileMd5Info(path1) + " path2:" + GetFileMd5Info(path2))
		}

		time.Sleep(10 * time.Second) //10초마다
	}
}

func f3() {

	sPath := "C:\\Go\\src\\DaeseongLib\\lib"
	x := findAllMapfile(sPath)
	for key, val := range x {
		fmt.Printf("key:%s    value:%s \n", key, val)
	}
}

func main() {
	//f1()
	//f2()
	f3()
}
*/
