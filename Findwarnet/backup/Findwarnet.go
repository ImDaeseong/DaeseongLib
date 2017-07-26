// Findwarnet
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
Google Places API
https://developers.google.com/places/web-service/search?hl=ko
*/

// Location type
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Viewport type
type Viewport struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

// Geometry type
type Geometry struct {
	Location Location `json:"location"`
	Viewport Viewport `json:"viewport"`
}

// Photo type
type Photo struct {
	Height         int      `json:"height"`
	Width          int      `json:"width"`
	PhotoReference string   `json:"photo_reference"`
	HTMLAttributes []string `json:"html_attributions"`
}

// Result json
type Result struct {
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
	Icon             string   `json:"icon"`
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Photos           []Photo  `json:"photos"`
	PlaceID          string   `json:"place_id"`
	Reference        string   `json:"reference"`
	Types            []string `json:"types"`
}

// Response json
type Response struct {
	Results        []Result `json:"results"`
	Status         string   `json:"status"`
	HTMLAttributes []string `json:"html_attributions"`
}

func readzipFileList(sInPath string) []string {
	fileList := make([]string, 0, 10)

	err := filepath.Walk(sInPath, func(sOutPath string, fi os.FileInfo, err error) error {
		if !fi.IsDir() {
			fileList = append(fileList, sOutPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("[%v]", err)
	}
	return fileList
}

func readZipFile(sPath string) []string {
	linefileList := make([]string, 0, 10)

	file, err := os.Open(sPath)
	if err != nil {
		return nil
	}

	fileScan := bufio.NewScanner(file)
	for fileScan.Scan() {
		linefileList = append(linefileList, fileScan.Text())
	}
	return linefileList
}

func writeGoogleMapString(sPath, sText string) {

	file, err := os.OpenFile(sPath, os.O_RDWR|os.O_APPEND, 0660)
	if os.IsNotExist(err) {
		file, err = os.Create(sPath)
	}
	defer file.Close()

	if err != nil {
		return
	}

	n, err := io.WriteString(file, sText+"\r\n")
	if err != nil {
		fmt.Println(n, err)
		return
	}
}

func getfloat64first(sInput string) string {

	sParam := "float64="

	if strings.Contains(sInput, sParam) {

		Pos1 := strings.Index(sInput, sParam) + len(sParam)
		sValue := sInput[Pos1:]

		Pos2 := strings.Index(sValue, ")")
		sValue = sValue[:Pos2]
		return sValue
	}
	return ""
}

func getfloat64second(sInput string) string {

	sParam := "float64="

	if strings.Contains(sInput, sParam) {

		Pos1 := strings.LastIndex(sInput, sParam) + len(sParam)
		sValue := sInput[Pos1:]

		Pos2 := strings.LastIndex(sValue, ")")
		sValue = sValue[:Pos2]
		return sValue
	}
	return ""
}

func trimspacestrings(sVal string) string {
	return strings.Replace(sVal, " ", "+", -1)
}

func GoogleMapsPlaceSearch(url string) bool {

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()

	var result Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println(result.Status)

	if result.Status == "OVER_QUERY_LIMIT" {
		sleep()
	}

	if result.Status != "OK" {
		writeGoogleMapString("error.txt", url)
	}

	for _, val := range result.Results {
		//fmt.Println(key, val)

		loc := fmt.Sprintf("%s", val.Geometry.Location)
		Lat := getfloat64first(loc)
		Lng := getfloat64second(loc)

		sLocation := fmt.Sprintf("%s|%s|%s", val.Name, Lat, Lng)

		writeGoogleMapString("Mapinfo.txt", sLocation)
	}
	return true
}

func sleep() {
	<-time.After(2 * time.Second)
}

func main() {

	cnl := make(chan string)
	go func() {

		sKey := "API Key"
		sWord := "kind"
		zipList := readzipFileList("C:\\Go\\src\\DaeseongLib\\zipcode")
		for _, file := range zipList {

			for _, line := range readZipFile(file) {

				url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/textsearch/json?query=%s+%s&key=%s", sWord, trimspacestrings(line), sKey)
				cnl <- url
			}
		}
		close(cnl)
	}()

	for zip := range cnl {
		GoogleMapsPlaceSearch(zip)
	}
	fmt.Println("complete")

}
