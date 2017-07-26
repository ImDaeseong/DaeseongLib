// NaverMovieRank
package DaeseongLib

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func NavermovieInfo() {

	sUrl := "http://movie.naver.com/movie/sdb/rank/rmovie.nhn"

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

	sSlice, _ := readParserNaverMv()
	sText := fmt.Sprintf("%s", sSlice)

	naverMovieParser(sText)

	err = os.Remove("parser.txt")
	if err != nil {
		log.Println("Remove:", err)
	}
	fmt.Println("complete")
}

func getMovieName(sInput string) string {

	sParam := "title="

	if strings.Contains(sInput, sParam) {

		Pos := strings.Index(sInput, sParam) + len(sParam)
		sValue := betweenMovie(sInput[Pos:], "\"", "\"")
		return sValue
	}
	return ""
}

func betweenMovie(sInput string, sFirst string, sEnd string) string {
	if strings.Contains(sInput, sFirst) && strings.Contains(sInput, sEnd) {
		PosF := strings.Index(sInput, sFirst) + len(sFirst)
		PosE := strings.LastIndex(sInput, sEnd)
		return sInput[PosF:PosE]
	}
	return ""
}

func writeNaverMovieString(sPath, sText string) {

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

func naverMovieParser(sContent string) {

	err := os.Remove("date.txt")
	if err != nil {
		log.Println("Remove:", err)
	}

	var divexp = regexp.MustCompile(`<div class="tit3">(.*?)</div>`)
	var tbodyexp = regexp.MustCompile(`<tbody>(.*?)</tbody>`)
	match := tbodyexp.FindAllString(sContent, -1)
	if match != nil {
		for _, val := range match {

			td1 := divexp.FindAllString(val, -1)

			if len(td1) == 0 {
				continue
			}

			var index int = 0
			for _, link := range td1 {
				index = index + 1
				sContent := fmt.Sprintf("%d|%s\r\n", index, getMovieName(link))
				writeNaverMovieString("date.txt", sContent)
			}
		}
	}
}

func readParserNaverMv() ([]string, error) {

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
	NavermovieInfo()
}
*/
