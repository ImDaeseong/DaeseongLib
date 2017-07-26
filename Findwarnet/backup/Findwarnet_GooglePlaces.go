// Findwarnet_GooglePlaces
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func sleep() {
	<-time.After(10 * time.Second)
}

func GoogleMapsPlaceSearch(url string) bool {

	surl := fmt.Sprintf("%s&language=ko", url)

	res, err := http.Get(surl)
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

		sLocation := fmt.Sprintf("%s|%s|%s|%s|%s", val.Name, Lat, Lng, val.PlaceID, val.FormattedAddress)
		writeGoogleMapString("Mapinfo.txt", sLocation)
	}
	return true
}

func main() {

	file, err := os.OpenFile("urlInfo.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		line = strings.TrimSpace(line)
		GoogleMapsPlaceSearch(line)
	}
	fmt.Println("complete")
}
