// Findwarnet_placeidDetail
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
https://developers.google.com/places/web-service/details?hl=ko
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

type Reviews struct {
	Authorname              string `json:"author_name"`
	Authorurl               string `json:"author_url"`
	Profilephotourl         string `json:"profile_photo_url"`
	Relativetimedescription string `json:"relative_time_description"`
	Text                    string `json:"text"`
}

// Result json
type Result struct {
	FormattedAddress         string    `json:"formatted_address"`
	Formattedphonenumber     string    `json:"formatted_phone_number"`
	Geometry                 Geometry  `json:"geometry"`
	Icon                     string    `json:"icon"`
	ID                       string    `json:"id"`
	Internationalphonenumber string    `json:"international_phone_number"`
	Name                     string    `json:"name"`
	PlaceID                  string    `json:"place_id"`
	Url                      string    `json:"url"`
	Reviews                  []Reviews `json:"reviews"`
}

// Response json
type Response struct {
	Results        Result   `json:"result"`
	Status         string   `json:"status"`
	HTMLAttributes []string `json:"html_attributions"`
}

func writeGoogleMapDetailString(sPath, sText string) {

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

func GoogleMapsPlaceDetailSearch(url string) bool {

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
		writeGoogleMapDetailString("error.txt", url)
	}

	var Authorname string
	var Authorurl string
	var Profilephotourl string
	var Relativetimedescription string
	var TextDesc string
	for _, val := range result.Results.Reviews {
		Authorname = val.Authorname
		Authorurl = val.Authorurl
		Profilephotourl = val.Profilephotourl
		Relativetimedescription = val.Relativetimedescription
		TextDesc = val.Text
	}

	loc := fmt.Sprintf("%s", result.Results.Geometry.Location)
	Lat := getfloat64first(loc)
	Lng := getfloat64second(loc)

	sLocation := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s", result.Results.Name, Lat, Lng,
		result.Results.PlaceID, result.Results.FormattedAddress,
		result.Results.Formattedphonenumber,
		result.Results.Internationalphonenumber, result.Results.Url,
		Authorname, Authorurl, Profilephotourl, Relativetimedescription, TextDesc)

	writeGoogleMapDetailString("GooglePlacesDetailInfo.txt", sLocation)

	return true
}

func main() {
	file, err := os.OpenFile("GooglePlacesInfo.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		line = strings.TrimSpace(line)
		GoogleMapsPlaceDetailSearch(line)
	}
	fmt.Println("complete")
}
