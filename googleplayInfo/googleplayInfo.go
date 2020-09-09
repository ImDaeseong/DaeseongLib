// googleplayInfo
package DaeseongLib

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func selectPage() (int, string) {
	sSelectPage := flag.Int("page", 1, "page")
	flag.Parse()

	var sUrl string

	if *sSelectPage == 1 {
		sUrl = "https://play.google.com/store/apps/collection/topselling_free"
	} else if *sSelectPage == 2 {
		sUrl = "https://play.google.com/store/apps/collection/topselling_paid"
	} else if *sSelectPage == 3 {
		sUrl = "https://play.google.com/store/apps/collection/topgrossing"
	} else if *sSelectPage == 4 {
		sUrl = "https://play.google.com/store/apps/category/GAME/collection/topselling_free"
	} else if *sSelectPage == 5 {
		sUrl = "https://play.google.com/store/apps/category/GAME/collection/topselling_paid"
	} else if *sSelectPage == 6 {
		sUrl = "https://play.google.com/store/apps/category/GAME/collection/topgrossing"
	}

	return *sSelectPage, sUrl
}

func WriteGoogleAppString(sPath, sText string) {

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

func GoogleAppParser(index int, sContent string) {

	sFilePath := fmt.Sprintf("D:\\Daeseong\\%d.txt", index)

	var arialabel = regexp.MustCompile(`aria-label="(.*?)"`)
	var targetexp = regexp.MustCompile(`<a class="card-click-target" href=(.*?)>`)
	match := targetexp.FindAllString(string(sContent), -1)
	if match != nil {
		for _, val := range match {
			sName := arialabel.FindAllString(val, -1)
			if len(sName) == 0 {
				continue
			}

			indexF1 := strings.Index(sName[0], "\"")
			indexL1 := strings.LastIndex(sName[0], "\"")
			sValue := strings.TrimSpace(sName[0][indexF1+1 : indexL1])

			sLabel := strings.Split(sValue, ".")[0]
			sDisplay := strings.Split(sValue, ".")[1]
			sDisplay = strings.TrimLeft(sDisplay, " ")
			//fmt.Println(sLabel, sDisplay)

			sContent := fmt.Sprintf("%s|%s\r\n", sLabel, sDisplay)
			WriteGoogleAppString(sFilePath, sContent)
		}
	}

	fmt.Println("complete")
}

func GoogleAppInfo() {

	index, url := selectPage()

	res, err := http.Get(url)

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

	GoogleAppParser(index, string(body))
}

/*
func main() {

	GoogleAppInfo()
}
*/
