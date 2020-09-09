// youtubeviewcountInfo
package DaeseongLib

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	lowerRe, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	brRe, _    = regexp.Compile("<br.*?>")
	tagRe, _   = regexp.Compile("<.*?>")
)

func StripTags(html string) string {
	html = lowerRe.ReplaceAllStringFunc(html, strings.ToLower)
	html = strings.Replace(html, "\n", " ", -1)
	html = strings.Replace(html, "\r", "", -1)
	html = strings.Replace(html, "&nbsp;", " ", -1)
	html = strings.Replace(html, "&#39", "", -1)
	html = strings.Replace(html, ";", "", -1)
	html = strings.Replace(html, ":", "", -1)
	html = brRe.ReplaceAllString(html, "")
	html = tagRe.ReplaceAllString(html, "")
	return html
}

type tagYtlockup struct {
	lockup      string
	content     string
	meta        string
	description string
}

var mYt = make(map[int]tagYtlockup)

func leftYt(sString string, nCount int) string {

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

func getTitleYt(sInput string) string {

	sParam := "title="

	if strings.Contains(sInput, sParam) {

		Pos1 := strings.Index(sInput, sParam) + len(sParam)
		sValue := sInput[Pos1+1:]

		Pos2 := strings.Index(sValue, "\"")
		sValue = sValue[:Pos2]
		//fmt.Println(sValue)

		return sValue
	}
	return ""
}

func getCoundYt(sInput string) string {

	sParam1 := "<li>"
	sParam2 := "</li>"

	if strings.Contains(sInput, sParam1) && strings.Contains(sInput, sParam2) {

		Pos1 := strings.Index(sInput, sParam1) + len(sParam1)
		Pos2 := strings.Index(sInput[Pos1:], sParam2)
		sValue1 := leftYt(sInput[Pos1:], Pos2)

		Pos3 := strings.LastIndex(sInput, sParam1) + len(sParam1)
		Pos4 := strings.LastIndex(sInput[Pos3:], sParam2)
		sValue2 := leftYt(sInput[Pos3:], Pos4)

		sResult := fmt.Sprintf("%s|%s", sValue1, sValue2)
		//fmt.Println(sResult)

		return sResult
	}
	return ""
}

func getDescription(sInput string) string {

	sParam1 := ">"
	sParam2 := "<a href="
	sParam3 := "</div>"

	if strings.Contains(sInput, sParam1) {

		Pos1 := strings.Index(sInput, sParam1) + len(sParam1)
		sValue := sInput[Pos1:]

		Pos2 := strings.Index(sValue, sParam2)
		Pos3 := strings.Index(sValue, sParam3)

		if Pos2 > 0 {
			sValue = leftYt(sValue, Pos2)
		} else if Pos3 > 0 {
			sValue = leftYt(sValue, Pos3)
		}

		sValue = StripTags(sValue)
		//fmt.Println(sValue)

		return sValue
	}
	return ""
}

func youtubeviewParser(sContent string) {

	var RankIndex int = 0
	var sVal1, sVal2, sVal3, sVal4 string

	var lockupexp = regexp.MustCompile(`<div class="yt-lockup(.*?)"(.*?)</div>`)
	match := lockupexp.FindAllString(sContent, -1)
	if match != nil {
		for _, val := range match {

			if strings.Contains(val, "<div class=\"yt-lockup-badges\">") {
				continue
			}

			if strings.Contains(val, "<div class=\"yt-lockup yt-lockup-tile") {
				sVal1 = val
			}

			if strings.Contains(val, "<div class=\"yt-lockup-content\">") {
				sVal2 = getTitleYt(val)
			}

			if strings.Contains(val, "<div class=\"yt-lockup-meta") {
				sVal3 = getCoundYt(val)
			}

			if strings.Contains(val, "<div class=\"yt-lockup-description") {
				sVal4 = getDescription(val)

				mYt[RankIndex] = tagYtlockup{sVal1, sVal2, sVal3, sVal4}
				RankIndex++
			}

			/*
				if i%4 == 0 {
					sVal1 = val
				} else if i%4 == 1 {
					sVal2 = val
				} else if i%4 == 2 {
					sVal3 = val
				} else if i%4 == 3 {
					sVal4 = val
					mYt[i] = tagYtlockup{sVal1, sVal2, sVal3, sVal4}
				}
			*/
		}
	}

	for i := 0; i < len(mYt); i++ {
		rank := i + 1
		sResult := fmt.Sprintf("%d : %s | %s", rank, mYt[i].content, mYt[i].meta)
		fmt.Println(sResult)
	}

	/*
		for key, val := range mYt {
			fmt.Println(key, val.content, val.meta, val.description)
		}
	*/

	fmt.Println("complete")
}

func searchYt() string {

	sSearch := flag.String("search", "", "search")
	flag.Parse()
	return *sSearch
}

func YoutubeviewCountInfo() {

	sQuery := searchYt()
	sQuery = strings.Replace(sQuery, " ", "+", -1)

	sUrl := fmt.Sprintf("http://www.youtube.com/results?search_query=%s&search_sort=video_view_count", sQuery)
	//fmt.Println(sUrl)

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

	sSlice, _ := readYtParser()
	sText := fmt.Sprintf("%s", sSlice)

	youtubeviewParser(sText)

	err = os.Remove("parser.txt")
	if err != nil {
		log.Println("Remove:", err)
	}
}

func readYtParser() ([]string, error) {

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

/*
func main() {
	YoutubeviewCountInfo()
}
*/
