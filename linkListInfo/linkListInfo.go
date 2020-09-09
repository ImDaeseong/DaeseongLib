// linkListInfo
package DaeseongLib

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	PAGES []string
)

func linkList(sContent string) {

	var href = regexp.MustCompile(`<link(.*?)href="((http|https)://)(.*?)"`)
	match := href.FindAllString(sContent, -1)
	if match != nil {
		for i, val := range match {

			link := fmt.Sprintf("[%d]%s \r\n", i, val)
			WriteString("D:\\linkInfo.txt", link)
		}
	}
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

func DownloadLinkTest() {
	PAGES = append(PAGES, "http://www.naver.com")
	PAGES = append(PAGES, "http://www.daum.net")

	for _, url := range PAGES {

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

		linkList(string(body))
	}

	fmt.Println("complete")
}

/*
func main() {

	DownloadLinkTest()
}
*/
