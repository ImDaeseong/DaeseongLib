// imgListDownLoad
package DaeseongLib

import (
	"fmt"
	_ "io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	IMGS []string
)

func downloadbytes(sUrl string) ([]byte, error) {
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

func writeImgFile(sFileName string, bByte []byte) bool {
	err := ioutil.WriteFile(sFileName, bByte, 0644)
	if err != nil {
		return false
	}
	return true
}

func DownloadImgFile(sUrl string, sFilePath string) (err error) {

	Replacer := strings.NewReplacer(
		"[", "",
		"]", "",
		"\"", "",
	)
	sUrl = Replacer.Replace(sUrl)

	bytes, err := downloadbytes(sUrl)
	if err != nil {
		return err
	}
	writeImgFile(sFilePath, bytes)

	return nil
}

func createFolder(sfolder string) {
	if _, err := os.Stat(sfolder); os.IsNotExist(err) {
		os.Mkdir(sfolder, 0777)
	}
}

func getFilename(linkUrl string) (sName string) {

	//i := strings.Split(linkUrl, "/")
	//sName = i[len(i)-1]

	sName = linkUrl[strings.LastIndex(linkUrl, "/")+1:]
	sName = strings.Replace(sName, "\"", "", -1)

	return sName
}

func imgList(sContent string) {

	createFolder("D:\\Daeseong")

	var img = regexp.MustCompile(`"(.*?)"`)
	var src = regexp.MustCompile(`src="(.*?)"`)
	var href = regexp.MustCompile(`<img(.*?)src="((http|https)://)(.*?)"`)
	match := href.FindAllString(sContent, -1)
	if match != nil {
		for _, val := range match {

			html := src.FindAllString(val, -1)
			linkUrl := img.FindAllString(html[0], -1)
			filename := getFilename(linkUrl[0])
			saveFile := fmt.Sprintf("D:\\Daeseong\\%s", filename)

			go DownloadImgFile(linkUrl[0], saveFile)
		}
	}
}

func DownloadImgTest() {
	IMGS = append(IMGS, "http://www.naver.com")
	IMGS = append(IMGS, "http://www.daum.net")

	for _, url := range IMGS {

		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		res, err := http.Get(url)
		if err != nil {
			continue
		}

		if res.StatusCode != http.StatusOK {
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			continue
		}

		imgList(string(body))
	}

	fmt.Println("complete")
}

/*
func main() {
	DownloadImgTest()
}
*/
