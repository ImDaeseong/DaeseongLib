// baseballInfo
package DaeseongLib

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type tagHead struct {
	sVal1  string
	sVal2  string
	sVal3  string
	sVal4  string
	sVal5  string
	sVal6  string
	sVal7  string
	sVal8  string
	sVal9  string
	sVal10 string
	sVal11 string
}

var mTeamRank = make(map[string]tagHead)

func leftA(sString string, nCount int) string {

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

func getThead(sInput string) string {

	sParam1 := "<thead>"
	sParam2 := "</thead>"

	if strings.Contains(sInput, sParam1) && strings.Contains(sInput, sParam2) {

		Pos1 := strings.Index(sInput, sParam1) + len(sParam1)
		Pos2 := strings.Index(sInput[Pos1:], sParam2)
		sValue := leftA(sInput[Pos1:], Pos2)
		return sValue
	}
	return ""
}

func getTr(sInput string) string {

	sParam1 := "<tr>"
	sParam2 := "</tr>"

	if strings.Contains(sInput, sParam1) && strings.Contains(sInput, sParam2) {

		Pos1 := strings.Index(sInput, sParam1) + len(sParam1)
		Pos2 := strings.Index(sInput[Pos1:], sParam2)
		sValue := leftA(sInput[Pos1:], Pos2)
		return sValue
	}
	return ""
}

func getTh(sInput string) string {

	sParam1 := "<th>"
	sParam2 := "</th>"

	if strings.Contains(sInput, sParam1) && strings.Contains(sInput, sParam2) {

		Pos1 := strings.Index(sInput, sParam1) + len(sParam1)
		Pos2 := strings.Index(sInput[Pos1:], sParam2)
		sValue := leftA(sInput[Pos1:], Pos2)
		return sValue
	}
	return ""
}

func getTd(sInput string) string {

	sParam1 := "<td>"
	sParam2 := "</td>"

	if strings.Contains(sInput, sParam1) && strings.Contains(sInput, sParam2) {

		Pos1 := strings.Index(sInput, sParam1) + len(sParam1)
		Pos2 := strings.Index(sInput[Pos1:], sParam2)
		sValue := leftA(sInput[Pos1:], Pos2)
		return sValue
	}
	return ""
}

func getTdfirst(sInput string) string {

	sParam1 := "<td>"

	if strings.Contains(sInput, sParam1) {
		Pos := strings.Index(sInput, sParam1) + len(sParam1)
		sValue := sInput[Pos:]
		return sValue
	}
	return ""
}

func getSpanTitle(sInput string) string {

	sParam1 := "span class=\"title\""
	sParam2 := "</span>"

	if strings.Contains(sInput, sParam1) && strings.Contains(sInput, sParam2) {

		Pos1 := strings.Index(sInput, sParam1) + len(sParam1) + 1
		Pos2 := strings.Index(sInput[Pos1:], sParam2)
		sValue := leftA(sInput[Pos1:], Pos2)
		return sValue
	}
	return ""
}

func betweenA(sInput string, sFirst string, sEnd string) string {
	if strings.Contains(sInput, sFirst) && strings.Contains(sInput, sEnd) {
		PosF := strings.Index(sInput, sFirst) + len(sFirst)
		PosE := strings.LastIndex(sInput, sEnd)
		return sInput[PosF:PosE]
	}
	return ""
}

func getRankhref(sInput string) string {

	sParam1 := "<a href="
	sParam2 := "</a>"
	sParam3 := ">"

	if strings.Contains(sInput, sParam1) && strings.Contains(sInput, sParam2) {

		Pos1 := strings.Index(sInput, sParam1) + len(sParam1) + 1
		Pos2 := strings.Index(sInput[Pos1:], sParam2)
		sValue := leftA(sInput[Pos1:], Pos2)

		if strings.Contains(sValue, sParam3) {
			Pos3 := strings.Index(sValue, sParam3) + len(sParam3)
			sValue = sValue[Pos3:]
		}

		return sValue
	}
	return ""
}

func getRankList(sInput string, nType int) string {

	var sParam1, sParam2 string

	if nType == 0 {
		sParam1 = "span class='rank1 name'"
		sParam2 = "</span>"
	} else if nType == 1 {
		sParam1 = "span class='rank2 name'"
		sParam2 = "</span>"
	} else if nType == 2 {
		sParam1 = "span class='rank3 name'"
		sParam2 = "</span>"
	} else if nType == 3 {
		sParam1 = "span class='rank4 name'"
		sParam2 = "</span>"
	} else if nType == 4 {
		sParam1 = "span class='rank5 name'"
		sParam2 = "</span>"
	} else if nType == 5 {
		sParam1 = "span class=\"team\""
		sParam2 = "</span>"
	} else if nType == 6 {
		sParam1 = "span class=\"rr\""
		sParam2 = "</span>"
	}

	if strings.Contains(sInput, sParam1) && strings.Contains(sInput, sParam2) {

		Pos1 := strings.Index(sInput, sParam1) + len(sParam1) + 1
		Pos2 := strings.Index(sInput[Pos1:], sParam2)
		sValue := leftA(sInput[Pos1:], Pos2)
		return sValue
	}
	return ""
}

func TeamRankParser(sContent string) {

	var trexp = regexp.MustCompile(`<tr>([\w\W]+?)</tr>`)
	var thexp = regexp.MustCompile(`<th>([\w\W]+?)</th>`)

	var tbodyexp = regexp.MustCompile(`<tbody>([\w\W]+?)</tbody>`)
	var theadexp = regexp.MustCompile(`<thead>([\w\W]+?)</thead>`)

	var tableexp = regexp.MustCompile(`summary="(.*?)" class="tData">([\w\W]+?)</table>`)
	match := tableexp.FindAllString(sContent, -1)
	if match != nil {
		for _, val := range match {

			td1 := theadexp.FindAllString(val, -1)
			td2 := tbodyexp.FindAllString(val, -1)

			if len(td1) == 0 || len(td2) == 0 {
				continue
			}

			link1 := getThead(td1[0])
			if len(link1) == 0 {
				continue
			}

			sTh := thexp.FindAllString(link1, -1)
			sTr := trexp.FindAllString(td2[0], -1)

			if len(sTh) == 0 || len(sTr) == 0 {
				continue
			}

			//check data
			/*
				var tagHeda []string = make([]string, 0)
				for _, th := range sTh {
					tagHeda = append(tagHeda, getTh(th))
				}
				fmt.Println(tagHeda)
			*/
			mTeamRank[getTh(sTh[0])] = tagHead{getTh(sTh[0]), getTh(sTh[1]), getTh(sTh[2]),
				getTh(sTh[3]), getTh(sTh[4]), getTh(sTh[5]), getTh(sTh[6]),
				getTh(sTh[7]), getTh(sTh[8]), getTh(sTh[9]), getTh(sTh[10])}

			//check data
			/*
				for _, tr := range sTr {
					td := getTr(tr)

					var tagTd []string = make([]string, 0)
					for _, val := range strings.Split(td, "</td>") {
						tagTd = append(tagTd, getTdfirst(val))
					}
					fmt.Println(tagTd)
				}
			*/
			for _, tr := range sTr {
				td := getTr(tr)

				tagTd := strings.Split(td, "</td>")

				mTeamRank[getTdfirst(tagTd[0])] = tagHead{getTdfirst(tagTd[0]), getTdfirst(tagTd[1]),
					getTdfirst(tagTd[2]), getTdfirst(tagTd[3]), getTdfirst(tagTd[4]),
					getTdfirst(tagTd[5]), getTdfirst(tagTd[6]), getTdfirst(tagTd[7]),
					getTdfirst(tagTd[8]), getTdfirst(tagTd[9]), getTdfirst(tagTd[10])}
			}

		}
	}

	//data
	for key, val := range mTeamRank {
		fmt.Println(key, val)
	}

	fmt.Println("complete")
}

func MainRankParser(sContent string) {

	var link1 []string = make([]string, 0)
	var link2 map[int][]string = make(map[int][]string, 0)
	//var link3 map[string][]string = make(map[string][]string, 0)

	var liexp = regexp.MustCompile(`<li>([\w\W]+?)</li>`)
	var spanexp = regexp.MustCompile(`<span class="title">(.*?)</span>`)

	var titleexp = regexp.MustCompile(`<div class="title_bar">([\w\W]+?)</div>`)
	var rankListexp = regexp.MustCompile(`<ol class="rankList">([\w\W]+?)</ol>`)

	var divexp = regexp.MustCompile(`<div class="(.*?)">([\w\W]+?)</div>`)
	match := divexp.FindAllString(sContent, -1)
	if match != nil {
		for _, val := range match {

			td1 := titleexp.FindAllString(val, -1)

			if len(td1) == 0 {
				continue
			}

			title := spanexp.FindAllString(td1[0], -1)
			if len(title) == 0 {
				continue
			}

			link1 = append(link1, getSpanTitle(title[0]))
		}
	}

	var index int = 0
	match = divexp.FindAllString(sContent, -1)
	if match != nil {
		for _, val := range match {

			td1 := rankListexp.FindAllString(val, -1)

			if len(td1) == 0 {
				continue
			}

			rank := liexp.FindAllString(td1[0], -1)
			if len(rank) == 0 {
				continue
			}

			link2[index] = rank
			index++
		}
	}

	for i, val := range link1 {

		if len(link2[i]) == 0 {
			continue
		}

		fmt.Println(val)
		for i, val := range link2[i] {
			sVal := fmt.Sprintf("%s|%s|%s", getRankhref(val), getRankList(val, 5), getRankList(val, 6))
			fmt.Println(i+1, sVal)
		}
		fmt.Println()
	}

	fmt.Println("complete")
}

func selectbaseballPage() (int, string) {

	sSelectPage := flag.Int("page", 2, "page")
	flag.Parse()

	var sUrl string

	if *sSelectPage == 1 {
		sUrl = "http://www.koreabaseball.com/TeamRank/TeamRank.aspx"
	} else if *sSelectPage == 2 {
		sUrl = "http://www.koreabaseball.com/Record/Main.aspx"
	}

	return *sSelectPage, sUrl
}

func BaseballInfo() {

	index, sUrl := selectbaseballPage()

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

	if index == 1 {
		TeamRankParser(string(bytes))
	} else if index == 2 {
		MainRankParser(string(bytes))
	}
}

/*
func main() {
	BaseballInfo()
}
*/
