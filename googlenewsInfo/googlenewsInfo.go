// googlenewsInfo
package DaeseongLib

import (
	"flag"
	"fmt"
	_ "html"
	_ "io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func getSrc(sInput string) string {

	sParam := "src="

	if strings.Contains(sInput, sParam) {

		Pos1 := strings.Index(sInput, sParam) + len(sParam)
		sValue := sInput[Pos1+1:]

		Pos2 := strings.Index(sValue, "\"")
		sValue = sValue[:Pos2]
		return sValue
	}
	return ""
}

func getSpanTitletext(sInput string) string {
	indexF1 := strings.Index(sInput, ">")
	indexL1 := strings.LastIndex(sInput, "<")
	sValue := strings.TrimSpace(sInput[indexF1+1 : indexL1])
	return sValue
}

func selectNews() (int, string) {
	sSelectPage := flag.Int("page", 1, "page")
	flag.Parse()

	var sUrl string

	if *sSelectPage == 1 {
		sUrl = "https://news.google.co.kr/news/section?cf=all&ned=id_id&topic=tc&siidp=05e410791582a70bcd4a9d0a651fd02762df&ict=ln"
	} else if *sSelectPage == 2 {
		sUrl = "https://news.google.co.kr/news/section?cf=all&pz=1&ned=kr&topic=t&siidp=5977dcc9cbd0afde8a75a5d16cbeb3ab9554&ict=ln"
	} else if *sSelectPage == 3 {
		sUrl = "https://news.google.co.kr/news/section?cf=all&ned=us&topic=tc&siidp=99c746aa2f624f4eb8b8ad5c3bfd279d76a2&ict=ln"
	}

	return *sSelectPage, sUrl
}

func GoogleNewsInfo() {

	_, sUrl := selectNews()

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

	GoogleNewsParser()

	err = os.Remove("parser.txt")
	if err != nil {
		log.Println("Remove:", err)
	}
	fmt.Println("complete")
}

func GoogleNewsParser() {

	err := os.Remove("date.txt")
	if err != nil {
		log.Println("Remove:", err)
	}

	var data []string
	bytes, err := ioutil.ReadFile("parser.txt")
	if err != nil {
		log.Println("ReadFile:", err)
	}
	data = strings.Split(string(bytes), "\n")

	var sText string
	for _, val := range data {
		if len(val) == 0 {
			continue
		}
		sText += val
	}
	//fmt.Println(sText)

	imgexp := regexp.MustCompile(`<img class="esc-thumbnail-image" src="(.*?)"`)
	spanexp := regexp.MustCompile(`<span class="titletext">(.*?)</span>`)
	divexp := regexp.MustCompile(`<div class="esc-lead-snippet-wrapper">(.*?)</div>`)

	id1 := imgexp.FindAllString(sText, -1)
	id2 := spanexp.FindAllString(sText, -1)
	id3 := divexp.FindAllString(sText, -1)
	//fmt.Println(id1)
	//fmt.Println(id2)
	//fmt.Println(id3)

	fmt.Println("//--------image list--------//")
	if len(id1) > 0 {
		for i, val := range id1 {
			fmt.Println(i, getSrc(val))
		}
	}
	fmt.Println()

	fmt.Println("//--------headline list--------//")
	if len(id2) > 0 {
		for i, val := range id2 {
			fmt.Println(i, getSpanTitletext(val))
		}
	}
	fmt.Println()

	fmt.Println("//--------Contents list--------//")
	if len(id3) > 0 {
		for i, val := range id3 {
			fmt.Println(i, getSpanTitletext(val))
		}
	}
}

/*
func main() {

	GoogleNewsInfo()
}
*/
