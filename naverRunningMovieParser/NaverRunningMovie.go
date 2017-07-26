// NaverRunningMovie
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type movieTag struct {
	title  string
	imgUrl string
}

var mMoive = make(map[int]movieTag)

func getMovieSrc(sInput string) string {

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

func getMovieAlt(sInput string) string {

	sParam := "alt="

	if strings.Contains(sInput, sParam) {

		Pos1 := strings.Index(sInput, sParam) + len(sParam)
		sValue := sInput[Pos1+1:]

		Pos2 := strings.Index(sValue, "\"")
		sValue = sValue[:Pos2]
		return sValue
	}
	return ""
}

func readMoviefile() ([]string, error) {

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

func readMovieParser(sText string) {

	var RankIndex int = 0

	var thumbexp = regexp.MustCompile(`<div class="thumb">(.+?)</div>`)

	var divexp = regexp.MustCompile(`<ul class="lst_detail_t1">([\w\W]+?)</ul>`)
	match := divexp.FindAllString(sText, -1)
	if match != nil {

		for _, val := range match {

			thumb := thumbexp.FindAllString(val, -1)
			if len(thumb) == 0 {
				continue
			}

			for _, val := range thumb {

				mMoive[RankIndex] = movieTag{getMovieSrc(val), getMovieAlt(val)}
				RankIndex++
			}
		}
	}

	err := os.Remove("parser.txt")
	if err != nil {
		log.Println("Remove:", err)
	}
}

func main() {

	sUrl := "http://movie.naver.com/movie/running/current.nhn#"

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

	sSlice, _ := readMoviefile()
	sText := fmt.Sprintf("%s", sSlice)
	readMovieParser(sText)

	for i := 0; i < len(mMoive); i++ {
		rank := i + 1
		sResult := fmt.Sprintf("%d : %s | %s", rank, mMoive[i].title, mMoive[i].imgUrl)
		fmt.Println(sResult)
	}

	fmt.Println("complete")
}
