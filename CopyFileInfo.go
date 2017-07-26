// CopyFileInfo
package DaeseongLib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	_ "path/filepath"
	"strings"
)

func IsDirExist(sPath string) bool {
	_, err := os.Stat(sPath)
	return err == nil || os.IsExist(err)
}

func FindImageFiles(sPath string) (sFileLst []string) {
	imgLst, err := ioutil.ReadDir(sPath)
	if err != nil {
		return
	}

	for _, file := range imgLst {
		onlyFilename := strings.ToUpper(file.Name())

		if strings.HasSuffix(onlyFilename, ".GIF") || strings.HasSuffix(onlyFilename, ".PNG") {
			sFileLst = append(sFileLst, sPath+"/"+file.Name())
			//fmt.Println(sPath + "/" + file.Name())
		}
	}
	return
}

func CopyFile(sSource, sDest string) (err error) {
	sSourcefile, err := os.Open(sSource)
	if err != nil {
		return err
	}
	defer sSourcefile.Close()

	sDestfile, err := os.Create(sDest)
	if err != nil {
		return err
	}
	defer sDestfile.Close()

	_, err = io.Copy(sDestfile, sSourcefile)
	if err != nil {
		return err
	}
	return
}

func CopyFolder(sSource, sDest string) (err error) {
	sSourceinfo, err := os.Stat(sSource)
	if err != nil {
		return err
	}

	err = os.MkdirAll(sDest, sSourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(sSource)
	//fmt.Printf(directory.Name() + "\n")

	dirlst, err := directory.Readdir(-1)
	for _, file := range dirlst {

		sPath := sSource + "/" + file.Name()
		sCopyPath := sDest + "/" + file.Name()

		if file.IsDir() {

			err = CopyFolder(sPath, sCopyPath)
			if err != nil {
				fmt.Println(err)
			}

		} else {

			err = CopyFile(sPath, sCopyPath)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
	return
}

/*
func f1() {
	var sSourceFolder string = "C:\\Go\\src\\DaeseongLib\\image"
	var sDestFolder string = "D:\\image"

	fList := FindImageFiles(sSourceFolder)
	for _, file := range fList {

		filename := file[strings.LastIndex(file, "/")+1:]

		//fileExt := filename[strings.LastIndex(filename, ".")+1:]
		//fileExt := filepath.Ext(filename)
		//fmt.Printf(fileExt + "\n")

		if !IsDirExist(sDestFolder) {
			err := os.Mkdir(sDestFolder, os.ModePerm)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		sNewfilename := fmt.Sprintf("%s\\%s", sDestFolder, filename)
		CopyFile(file, sNewfilename)
	}
}

func f2() {
	CopyFolder("C:\\Go\\src\\DaeseongLib\\lib", "D:\\Go\\src\\DaeseongLib\\lib")
}

func main() {
	//f1()
	f2()
}
*/
