// main
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/tealeg/xlsx"
)

func DeleteFile(sPath string) error {
	file, err := os.Stat(sPath)
	if err != nil {
		return err
	}

	if file.IsDir() {
		err := os.RemoveAll(sPath)
		if err != nil {
			return err
		}
	} else {
		err := os.Remove(sPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func readFileData() ([]string, error) {

	file, err := os.Open("a.html")
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

func writeParseString(sPath, sText string) {

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

func getDatatext(sInput string) string {

	index1 := strings.Index(sInput, "<td>")
	index2 := strings.LastIndex(sInput, "</td>")
	sTemp := sInput[index1:index2]

	sValue := strings.TrimSpace(sTemp[strings.LastIndex(sTemp, ">")+1:])
	return sValue
}

func readParser(sText string) {

	var p3 = regexp.MustCompile(`<td>([\w\W]+?)</td>`)
	var p2 = regexp.MustCompile(`<tr>([\w\W]+?)</tr>`)
	var p1 = regexp.MustCompile(`<div class="box" id="user-data">(.*?)</div>`)
	match := p1.FindAllString(sText, -1)
	//fmt.Println(match)

	if match != nil {
		for _, val := range match {

			tr1 := p2.FindAllString(val, -1)

			if len(tr1) == 0 {
				continue
			}

			//var index = 0
			for _, tr := range tr1 {

				td1 := p3.FindAllString(tr, -1)

				if len(td1) == 0 {
					continue
				}

				slice := strings.Split(td1[5], "<td>")
				//sslice1 := strings.Replace(string(slice[0]), "</td>", "", -1)
				sslice2 := strings.Replace(string(slice[1]), "</td>", "", -1)
				sslice3 := strings.Replace(string(slice[2]), "</td>", "", -1)
				sData2 := strings.TrimSpace(sslice2)
				sData3 := strings.TrimSpace(sslice3)
				//index++

				sVal := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s\r\n", getDatatext(td1[0]), getDatatext(td1[1]), getDatatext(td1[2]), getDatatext(td1[3]), getDatatext(td1[4]), sData2, sData3)
				fmt.Println(sVal)
				writeParseString("a.txt", sVal)
			}

		}
	}
}

func writeExecel(sText string) {

	var p3 = regexp.MustCompile(`<td>([\w\W]+?)</td>`)
	var p2 = regexp.MustCompile(`<tr>([\w\W]+?)</tr>`)
	var p1 = regexp.MustCompile(`<div class="box" id="user-data">(.*?)</div>`)
	match := p1.FindAllString(sText, -1)
	//fmt.Println(match)

	//엑셀로 저장하기
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell1, cell2, cell3, cell4, cell5, cell6, cell7 *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	if match != nil {
		for _, val := range match {

			tr1 := p2.FindAllString(val, -1)

			if len(tr1) == 0 {
				continue
			}

			//var index = 0
			for _, tr := range tr1 {

				td1 := p3.FindAllString(tr, -1)

				if len(td1) == 0 {
					continue
				}

				slice := strings.Split(td1[5], "<td>")
				//sslice1 := strings.Replace(string(slice[0]), "</td>", "", -1)
				sslice2 := strings.Replace(string(slice[1]), "</td>", "", -1)
				sslice3 := strings.Replace(string(slice[2]), "</td>", "", -1)
				sData2 := strings.TrimSpace(sslice2)
				sData3 := strings.TrimSpace(sslice3)
				//index++

				row = sheet.AddRow()

				cell1 = row.AddCell()
				cell1.Value = getDatatext(td1[0])

				cell2 = row.AddCell()
				cell2.Value = getDatatext(td1[1])

				cell3 = row.AddCell()
				cell3.Value = getDatatext(td1[2])

				cell4 = row.AddCell()
				cell4.Value = getDatatext(td1[3])

				cell5 = row.AddCell()
				cell5.Value = getDatatext(td1[4])

				cell6 = row.AddCell()
				cell6.Value = sData2

				cell7 = row.AddCell()
				cell7.Value = sData3
			}

			err = file.Save("a.xlsx")
			if err != nil {
				fmt.Printf(err.Error())
			}
		}
	}
}

func main() {

	DeleteFile("a.xlsx")
	DeleteFile("a.txt")

	sSlice, _ := readFileData()
	sText := fmt.Sprintf("%s", sSlice)
	readParser(sText)
	writeExecel(sText)
	fmt.Println("complete")
}
