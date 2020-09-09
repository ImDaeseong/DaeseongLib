// Zipper
package DaeseongLib

import (
	"archive/zip"
	_ "fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Zipfile(sSource, sTarget string) error {
	zipfile, err := os.Create(sTarget)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(sSource)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(sSource)
	}

	filepath.Walk(sSource, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			//fmt.Println(path + "  " + sSource)
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, sSource))
		}
		//fmt.Println(header.Name)

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate //ZIP Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)

		return err

	})

	return err
}

func UnZipfile(sSource, sTarget string) error {
	reader, err := zip.OpenReader(sSource)
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := os.MkdirAll(sTarget, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(sTarget, file.Name)
		//fmt.Println(path)

		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

/*
func f1() {

	Zipfile("C:\\Go\\src\\DaeseongLib\\lib", "C:\\lib.zip")
	Zipfile("C:\\Go\\src\\DaeseongLib\\lib", "lib.zip")
	Zipfile("C:\\test.exe", "C:\\test.zip")
}

func f2() {
	UnZipfile("C:\\lib.zip", "C:\\lib")
	UnZipfile("C:\\test.zip", "C:\\test")
}

func main() {
    f1()
	f2()
}
*/
