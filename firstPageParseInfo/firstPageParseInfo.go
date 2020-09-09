// firstPageParseInfo
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

var (
	PageIndex int
)

func maxPage() int {
	sMaxPage := flag.Int("page", 16, "page")
	flag.Parse()
	return *sMaxPage
}

func writeString(sPath, sText string) {

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

func GetParserInfo(sContent string) {

	var tplCo = regexp.MustCompile(`<td class="tplCo">([\w\W]+?)</td>`)
	var tplTit = regexp.MustCompile(`<td class="tplTit">([\w\W]+?)</td>`)
	var href = regexp.MustCompile(`<a href="(.+?)">(.+?)</a>`)
	var trexp = regexp.MustCompile(`<tr>([\w\W]+?)</tr>`)

	trTag := trexp.FindAllString(sContent, -1)
	if trTag != nil {
		for _, val := range trTag {

			td1 := tplCo.FindAllString(val, -1)
			td2 := tplTit.FindAllString(val, -1)

			if len(td1) == 0 || len(td2) == 0 {
				continue
			}

			link1 := href.FindAllString(td1[0], -1)
			link2 := href.FindAllString(td2[0], -1)

			if len(link1) == 0 || len(link2) == 0 {
				continue
			}

			indexF1 := strings.Index(link1[0], ">")
			indexL1 := strings.LastIndex(link1[0], "<")

			indexF2 := strings.Index(link2[0], ">")
			indexL2 := strings.LastIndex(link2[0], "<")

			PageIndex++

			link := fmt.Sprintf("[%d] %s - %s\r\n", PageIndex, string(link1[0][indexF1+1:indexL1]), string(link2[0][indexF2+1:indexL2]))
			writeString("D:\\DataInfo.txt", link)
		}

		/*
			var divexp = regexp.MustCompile(`<div class="tplPagination devTplPgn">([\w\W]+?)</div>`)
			var lihref = regexp.MustCompile(`<li><a href="(.+?)">(.+?)</a></li>`)
			var reexp = regexp.MustCompile(`href="(.*?)"`)
			var pexp = regexp.MustCompile(`<p><a href="(.+?)">(.+?)</a></p>`)

			var LIST []string
			trTag = divexp.FindAllString(sContent, -1)
			for _, val := range trTag {
				li := lihref.FindAllString(val, -1)

				for _, val := range li {
					link := reexp.FindAllString(val, -1)
					//LIST = append(LIST, link[0])

					page := strings.Replace(link[0], "\"", "", -1)
					index := strings.LastIndex(page, "=")
					LIST = append(LIST, string(page[index+1:]))
					//fmt.Println(string(page[index+1:]))
				}
			}

			p := pexp.FindAllString(trTag[0], -1)
			if len(p) != 0 {
				link := reexp.FindAllString(p[0], -1)
				//LIST = append(LIST, link[0])

				page := strings.Replace(link[0], "\"", "", -1)
				index := strings.LastIndex(page, "=")
				LIST = append(LIST, string(page[index+1:]))
				//fmt.Println(string(page[index+1:]))
			}
		*/
	}
}

func ParserPages() {

	sUrl := "GI_Area_List.asp?AreaNo=R000&AllStat=0&MapStat=&page="
	nPage := maxPage()
	for i := 1; i <= nPage; i++ {

		searchurl := fmt.Sprintf("%s%d", sUrl, i)
		res, err := http.Get(searchurl)
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

		GetParserInfo(string(body))
	}

	fmt.Println("complete")
}

/*
func main() {
	ParserPages()
}
*/
