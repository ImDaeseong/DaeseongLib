// mp3rwInfo
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type mp3Tag struct {
	title   string
	artist  string
	album   string
	year    string
	comment string
}

func flags() (string, string, string, string, string, string) {

	sPath := flag.String("Path", "", "Path")
	sTitle := flag.String("title", "", "title")
	sArtist := flag.String("artist", "", "artist")
	sAlbum := flag.String("album", "", "album")
	sYear := flag.String("year", "", "year")
	sComment := flag.String("comment", "", "comment")
	flag.Parse()

	return *sPath, *sTitle, *sArtist, *sAlbum, *sYear, *sComment
}

func getTagBytes(sPath string) ([]byte, error) {

	file, err := os.Open(sPath)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	_, err = file.Seek(-int64(128), os.SEEK_END)
	if err != nil {
		return nil, err
	}

	tagBytes := make([]byte, 128)
	_, err = file.Read(tagBytes)
	if err != nil {
		return nil, err
	}

	return tagBytes, nil
}

func setTagBytes(sPath string, tagBytes []byte) error {

	file, err := os.OpenFile(sPath, os.O_RDWR, 0655)
	defer file.Close()

	if err != nil {
		return err
	}
	_, err = file.Seek(-int64(128), os.SEEK_END)
	if err != nil {
		return err
	}
	_, err = file.Write(tagBytes)
	if err != nil {
		return err
	}

	return nil
}

func SetMp3Info(m mp3Tag, sPath string) error {

	tagBytes := make([]byte, 128)
	copy(tagBytes[:], "TAG")
	copy(tagBytes[3:33], m.title)
	copy(tagBytes[33:63], m.artist)
	copy(tagBytes[63:93], m.album)
	copy(tagBytes[93:97], m.year)
	copy(tagBytes[97:127], m.comment)

	err := setTagBytes(sPath, tagBytes)
	if err != nil {
		return err
	}
	return nil
}

func GetMp3Info(sPath string) (mp3Tag, error) {

	tTag := mp3Tag{}

	tagBytes, err := getTagBytes(sPath)
	if err != nil {
		return tTag, err
	}

	if string(tagBytes[:3]) != "TAG" {
		return tTag, errors.New("error")
	}

	tTag.title = string(tagBytes[3:33])
	tTag.artist = string(tagBytes[33:63])
	tTag.album = string(tagBytes[63:93])
	tTag.year = string(tagBytes[93:97])
	tTag.comment = string(tagBytes[97:127])

	return tTag, nil
}

func changeStrings(sVal string) string {
	return strings.Replace(sVal, "<br>", " ", -1)
}

func WriteString(sPath, sText string) {

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

func f1() {

	file := "D:\\test.mp3"
	tTag := mp3Tag{}
	tTag.title = string("test")
	tTag.artist = string("test")
	tTag.album = string("test")
	tTag.year = string("2017")
	tTag.comment = string("test")
	SetMp3Info(tTag, file)
}

func f2() {

	file := "D:\\test.mp3"
	tTag, err := GetMp3Info(file)
	if err == nil {

		fmt.Println("title:" + tTag.title)
		fmt.Println("artist:" + tTag.artist)
		fmt.Println("album:" + tTag.album)
		fmt.Println("year:" + tTag.year)
		fmt.Println("comment:" + tTag.comment)
	}
}

func f3() {

	if len(os.Args) < 2 {
		return
	}

	Arg1 := fmt.Sprintf("%s", os.Args[0:1])
	Arg2 := fmt.Sprintf("%s", os.Args[1:2])
	Arg3 := fmt.Sprintf("%s", os.Args[2:3])
	Arg4 := fmt.Sprintf("%s", os.Args[3:4])
	Arg5 := fmt.Sprintf("%s", os.Args[4:5])
	Arg6 := fmt.Sprintf("%s", os.Args[5:6])

	file := "D:\\Args.txt"
	WriteString(file, Arg1)
	WriteString(file, Arg2)
	WriteString(file, Arg3)
	WriteString(file, Arg4)
	WriteString(file, Arg5)
	WriteString(file, Arg6)
}

func f4() {

	sPathFlag, sTitleFlag, sArtistFlag, sAlbumFlag, sYearFlag, sCommentFlag := flags()

	replacer := strings.NewReplacer("<br>", " ")
	sPath := replacer.Replace(sPathFlag)
	sTitle := replacer.Replace(sTitleFlag)
	sArtist := replacer.Replace(sArtistFlag)
	sAlbum := replacer.Replace(sAlbumFlag)
	sYear := replacer.Replace(sYearFlag)
	sComment := replacer.Replace(sCommentFlag)

	//fmt.Println("file Path:" + sPath)
	//fmt.Println("title:" + sTitle)
	//fmt.Println("artist:" + sArtist)
	//fmt.Println("album:" + sAlbum)
	//fmt.Println("year:" + sYear)
	//fmt.Println("comment:" + sComment)

	tTag := mp3Tag{}
	tTag.title = sTitle
	tTag.artist = sArtist
	tTag.album = sAlbum
	tTag.year = sYear
	tTag.comment = sComment
	SetMp3Info(tTag, sPath)

	tReadTag, err := GetMp3Info(sPath)
	if err == nil {

		fmt.Println("title:" + tReadTag.title)
		fmt.Println("artist:" + tReadTag.artist)
		fmt.Println("album:" + tReadTag.album)
		fmt.Println("year:" + tReadTag.year)
		fmt.Println("comment:" + tReadTag.comment)
	}
}

func main() {

	f4()
}
