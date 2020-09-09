// NewRequestTest
package DaeseongLib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func GetFileList(sPath string) (list []string, count int) {

	mList := make([]string, 0)

	files, err := ioutil.ReadDir(sPath)
	if err != nil {
		println(err)
	}

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		filetype := path.Ext(file.Name())
		if filetype == ".mp3" {
			filePath := path.Join(sPath, file.Name())
			mList = append(mList, filePath)
		}
	}
	return mList, len(mList)
}

func GetWebPage(sUrl string) (string, error) {
	req, err := http.NewRequest("GET", sUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", "Daeseonglib")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

func GetDownloadFile(sUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", sUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Daeseonglib")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bByte, nil
}

func WriteFile(sFileName string, bByte []byte) bool {
	err := ioutil.WriteFile(sFileName, bByte, 0644)
	if err != nil {
		return false
	}
	return true
}

/*
func f1() {
	filelist, _ := GetFileList("C:\\lyricsPlayer")
	for _, file := range filelist {
		fmt.Println(file)
	}
}

func f2() {
	contents, _ := GetWebPage("https://github.com/ImDaeseong")
	fmt.Println(contents)
}

func f3() {
	bByte, _ := GetDownloadFile("https://github.com/ImDaeseong/DaeseongLib/blob/master/DaeseongLib/BluestackInfo.go")
	bWrite := WriteFile("BluestackInfo.go", bByte)
	fmt.Println(bWrite)
}

func main() {
	f3()
}
*/