// mp3Info
package DaeseongLib

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	_ "strings"
)

const (
	titleEnd   = 30
	artistEnd  = 60
	albumEnd   = 90
	yearEnd    = 94
	commentEnd = 124
	tagStart   = 3
	tagSize    = 128
)

type mp3Tag struct {
	title   string
	artist  string
	album   string
	year    int
	comment string
	track   int
	genre   int
}

func FindMp3List(sPath string) []string {

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

	return fileList
}

func getTag(file *os.File, size int) string {

	outByte := make([]byte, size)
	file.Read(outByte)
	sOut := string(outByte)

	return sOut
}

func GetMp3Tag(file *os.File) (mp3Tag, error) {

	tTag := mp3Tag{}
	file.Seek(-128, 2)

	tagBytes := make([]byte, 3)
	file.Read(tagBytes)

	//"TAG" 3
	if tagBytes[0] != 84 || tagBytes[1] != 65 || tagBytes[2] != 71 {
		return tTag, errors.New("error")
	}

	tTag.title = getTag(file, 30)
	tTag.artist = getTag(file, 30)
	tTag.album = getTag(file, 30)
	tTag.year, _ = strconv.Atoi(getTag(file, 4))
	tTag.comment = getTag(file, 30)

	file.Seek(1, 0)
	tTag.track, _ = strconv.Atoi(getTag(file, 1))
	tTag.genre, _ = strconv.Atoi(getTag(file, 28))

	return tTag, nil
}

func byteString(b []byte) string {

	pos := bytes.IndexByte(b, 0)
	if pos == -1 {
		pos = len(b)
	}
	return string(b[0:pos])
}

func GetMp3TagA(sPath string) (map[string]string, error) {

	buff_ := make([]byte, tagSize)

	f, err := os.Open(sPath)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	f.Seek(-tagSize, 2)
	f.Read(buff_)

	//"TAG" 3
	if byteString(buff_[0:tagStart]) != "TAG" {
		return nil, errors.New("error")
	}

	buff := buff_[tagStart:]

	id3tag := map[string]string{}

	id3tag["title"] = byteString(buff[0:titleEnd])
	id3tag["artist"] = byteString(buff[titleEnd:artistEnd])
	id3tag["album"] = byteString(buff[artistEnd:albumEnd])
	id3tag["year"] = byteString(buff[albumEnd:yearEnd])
	id3tag["comment"] = byteString(buff[yearEnd:commentEnd])

	if buff[commentEnd-2] == 0 {
		id3tag["track"] = fmt.Sprintf("%d", buff[commentEnd-1])
	}
	genre_code := buff[commentEnd]
	id3tag["genre"] = fmt.Sprintf("%d", genre_code)

	return id3tag, nil
}

/*
func f1() {

	fileList := FindMp3List("D:\\test")
	for _, fList := range fileList {
		sExt := strings.ToUpper(filepath.Ext(fList))
		if strings.HasSuffix(sExt, ".MP3") {

			tTag, err := GetMp3TagA(fList)
			if err != nil {
				continue
			}

			for skey, sValue := range tTag {
				fmt.Printf("%s:%s\n", skey, sValue)
			}
		}
	}
}

func f2() {

	fileList := FindMp3List("D:\\test")
	for _, fList := range fileList {
		sExt := strings.ToUpper(filepath.Ext(fList))
		if strings.HasSuffix(sExt, ".MP3") {

			file, err := os.Open(fList)
			defer file.Close()

			if err != nil {
				continue
			}

			tTag, err := GetMp3Tag(file)
			if err != nil {
				continue
			}

			fmt.Printf("\n")
			fmt.Println("title:" + tTag.title)
			fmt.Println("artist:" + tTag.artist)
			fmt.Println("album:" + tTag.album)
			fmt.Println(fmt.Sprintf("year:%d", tTag.year))
			fmt.Println("comment:" + tTag.comment)
			fmt.Println(fmt.Sprintf("track:%d", tTag.track))
			fmt.Println(fmt.Sprintf("genre:%d", tTag.genre))
		}
	}
}

func main() {

	f1()
	f2()
}
*/
