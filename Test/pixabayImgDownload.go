// pixabayImgDownload
package DaeseongLib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

var waitGroup sync.WaitGroup

func createPngFolder(sfolder string) {
	if _, err := os.Stat(sfolder); os.IsNotExist(err) {
		os.Mkdir(sfolder, 0777)
	}
}

func downloadpngbytes(sUrl string) ([]byte, error) {
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

func writePngFile(sFileName string, bByte []byte) bool {
	err := ioutil.WriteFile(sFileName, bByte, 0644)
	if err != nil {
		return false
	}
	return true
}

func getPngfilename(linkUrl string) (sName string) {
	sName = linkUrl[strings.LastIndex(linkUrl, "/")+1:]
	return sName
}

func pixPngList(sContent string) {

	createPngFolder("D:\\Daeseong")

	var Imgsrcset = regexp.MustCompile(`<img srcset="(.*?)"`)
	match := Imgsrcset.FindAllString(sContent, -1)
	if match != nil {
		for _, val := range match {

			indexF1 := strings.Index(val, "\"")
			indexL1 := strings.LastIndex(val, "\"")

			pngUrl := strings.Split(val[indexF1+1:indexL1], ",")[0]
			pngUrl = strings.Replace(pngUrl, "1x", "", -1)
			pngUrl = strings.Replace(pngUrl, " ", "", -1)
			//fmt.Println(pngUrl)

			filename := getPngfilename(pngUrl)
			sFilePath := fmt.Sprintf("D:\\Daeseong\\%s", filename)

			bytes, err := downloadpngbytes(pngUrl)
			if err != nil {
				continue
			}
			writePngFile(sFilePath, bytes)
		}
	}

	waitGroup.Done()
}

func pixPngDownload() {

	res, err := http.Get("https://pixabay.com/ko/photos/png/")
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return
	}

	waitGroup.Add(1)

	go pixPngList(string(body))

	waitGroup.Wait()

	fmt.Println("complete")
}

/*
func main() {

	pixPngDownload()
}
*/
