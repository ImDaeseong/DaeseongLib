// navermusicInfo
package DaeseongLib

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type tagMusic struct {
	Level string
	Name  string
	Title string
}

var mMusic = make(map[string]tagMusic)

func selectMusicPage() string {
	sSelectPage := flag.Int("page", 1, "page")
	flag.Parse()

	var sUrl string

	if *sSelectPage == 1 {
		sUrl = "http://music.naver.com/listen/top100.nhn?domain=TOTAL"
	} else if *sSelectPage == 2 {
		sUrl = "http://music.naver.com/listen/top100.nhn?domain=DOMESTIC"
	} else if *sSelectPage == 3 {
		sUrl = "http://music.naver.com/listen/top100.nhn?domain=OVERSEA"
	} else if *sSelectPage == 4 {
		sUrl = "http://music.naver.com/listen/history/index.nhn"
	}

	return sUrl
}

func readParser() ([]string, error) {

	file, err := os.Open("parser.txt")
	if err != nil {
		log.Println("Open:", err)
	}
	defer file.Close()

	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	return lines, err
}

func writeNaverMusicString(sPath, sText string) {

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

func getTitle(sInput string) string {

	sValue := between(sInput, ">", "<")
	return sValue
}

func left(sString string, nCount int) string {

	if nCount < 0 {
		return sString
	}

	nLength := len(sString)
	if nLength <= nCount {
		return sString
	}
	sString = string(sString[:nCount])
	return sString
}

func getLevel(sInput string) string {

	sParam1 := "r:"
	sParam2 := ","

	if strings.Contains(sInput, sParam1) && strings.Contains(sInput, sParam2) {

		Pos1 := strings.Index(sInput, sParam1) + len(sParam1)
		Pos2 := strings.Index(sInput[Pos1:], sParam2)
		sValue := left(sInput[Pos1:], Pos2)
		return sValue
	}
	return ""
}

func getName(sInput string) string {

	sParam := "title="

	if strings.Contains(sInput, sParam) {

		Pos := strings.Index(sInput, sParam) + len(sParam)
		sValue := between(sInput[Pos:], "\"", "\"")
		return sValue
	}
	return ""
}

func between(sInput string, sFirst string, sEnd string) string {
	if strings.Contains(sInput, sFirst) && strings.Contains(sInput, sEnd) {
		PosF := strings.Index(sInput, sFirst) + len(sFirst)
		PosE := strings.LastIndex(sInput, sEnd)
		return sInput[PosF:PosE]
	}
	return ""
}

func naverMusicParserA(sContent string) {

	var spanexp = regexp.MustCompile(`<span class="ellipsis">(.*?)</span>`)
	var titleexp = regexp.MustCompile(`href="(.*?)" title="(.*?)"`)

	var tdname = regexp.MustCompile(`<td class="name">([\w\W]+?)</td>`)
	var tdartist = regexp.MustCompile(`<td class="_artist artist">([\w\W]+?)</td>`)

	var trexp = regexp.MustCompile(`<tr class="_tracklist_move data(.*?)"(.*?)</tr>`)
	match := trexp.FindAllString(sContent, -1)
	if match != nil {
		for _, val := range match {

			td1 := tdname.FindAllString(val, -1)
			td2 := tdartist.FindAllString(val, -1)

			if len(td1) == 0 || len(td2) == 0 {
				continue
			}

			link1 := spanexp.FindAllString(td1[0], -1)
			link2 := titleexp.FindAllString(td2[0], -1)

			if len(link1) == 0 || len(link2) == 0 {
				continue
			}

			sTitle := getTitle(link1[0])
			sLevel := getLevel(link2[0])
			sName := getName(link2[0])

			//sContent := fmt.Sprintf("%s | %s - %s\r\n", sLevel, sName, sTitle)
			//writeNaverMusicString("date.txt", sContent)

			mMusic[sLevel] = tagMusic{sLevel, sName, sTitle}
		}
	}

	//data
	for key, val := range mMusic {
		fmt.Println(key, val)
	}

	fmt.Println("complete")
}

func naverMusicParser(sContent string) {

	err := os.Remove("date.txt")
	if err != nil {
		log.Println("Remove:", err)
	}

	var classexp = regexp.MustCompile(`class="(.*?)"`)
	var nameexp = regexp.MustCompile(`title="(.*?)"`)

	var spanexp = regexp.MustCompile(`<span class="ellipsis">(.*?)</span>`)
	var titleexp = regexp.MustCompile(`href="(.*?)" title="(.*?)"`)

	var tdname = regexp.MustCompile(`<td class="name">([\w\W]+?)</td>`)
	var tdartist = regexp.MustCompile(`<td class="_artist artist">([\w\W]+?)</td>`)

	var trexp = regexp.MustCompile(`<tr class="_tracklist_move data(.*?)"(.*?)</tr>`)
	match := trexp.FindAllString(sContent, -1)
	if match != nil {
		for _, val := range match {

			td1 := tdname.FindAllString(val, -1)
			td2 := tdartist.FindAllString(val, -1)

			if len(td1) == 0 || len(td2) == 0 {
				continue
			}

			link1 := spanexp.FindAllString(td1[0], -1)
			link2 := titleexp.FindAllString(td2[0], -1)

			if len(link1) == 0 || len(link2) == 0 {
				continue
			}

			//title
			indexF1 := strings.Index(link1[0], ">")
			indexL1 := strings.LastIndex(link1[0], "<")
			sTitle := strings.TrimSpace(link1[0][indexF1+1 : indexL1])

			data1 := classexp.FindAllString(link2[0], -1)
			data2 := nameexp.FindAllString(link2[0], -1)

			//Level
			indexF2 := strings.Index(data1[0], "r:")
			indexL2 := strings.LastIndex(data1[0], ",")
			sLevel := strings.TrimSpace(data1[0][indexF2+2 : indexL2])

			//name
			indexF3 := strings.Index(data2[0], "\"")
			indexL3 := strings.LastIndex(data2[0], "\"")
			sName := strings.TrimSpace(data2[0][indexF3+1 : indexL3])

			sContent := fmt.Sprintf("%s | %s - %s\r\n", sLevel, sName, sTitle)
			writeNaverMusicString("date1.txt", sContent)
		}
	}

	fmt.Println("complete")
}

func NavermusicInfo() {

	sUrl := selectMusicPage()

	var res *http.Response
	var err error
	var bytes []byte

	for {
		res, err = http.Get(sUrl)
		if err != nil {
			log.Println(err)
			continue
		}
		defer res.Body.Close()

		bytes, err = ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("ReadAll:", err)
			continue
		}
		break
	}

	err = ioutil.WriteFile("parser.txt", bytes, 0644)
	if err != nil {
		log.Println("WriteFile:", err)
	}

	sSlice, _ := readParser()
	sText := fmt.Sprintf("%s", sSlice)

	//first method
	naverMusicParser(sText)

	//second method
	naverMusicParserA(sText)

	err = os.Remove("parser.txt")
	if err != nil {
		log.Println("Remove:", err)
	}
}

/*
func main() {
	NavermusicInfo()
}
*/
