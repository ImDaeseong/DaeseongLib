// main
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type JSONData struct {
	Response struct {
		Header struct {
			ResultCode string `json:"resultCode"`
			ResultMsg  string `json:"resultMsg"`
		} `json:"header"`
		Body struct {
			DataType string `json:"dataType"`
			Items    struct {
				Item []struct {
					Tm         string `json:"tm"`
					Rnum       string `json:"rnum"`
					StnID      string `json:"stnId"`
					StnNm      string `json:"stnNm"`
					Ta         string `json:"ta"`
					TaQcflg    string `json:"taQcflg"`
					Rn         string `json:"rn"`
					RnQcflg    string `json:"rnQcflg"`
					Ws         string `json:"ws"`
					WsQcflg    string `json:"wsQcflg"`
					Wd         string `json:"wd"`
					WdQcflg    string `json:"wdQcflg"`
					Hm         string `json:"hm"`
					HmQcflg    string `json:"hmQcflg"`
					Pv         string `json:"pv"`
					Td         string `json:"td"`
					Pa         string `json:"pa"`
					PaQcflg    string `json:"paQcflg"`
					Ps         string `json:"ps"`
					PsQcflg    string `json:"psQcflg"`
					Ss         string `json:"ss"`
					SsQcflg    string `json:"ssQcflg"`
					Icsr       string `json:"icsr"`
					Dsnw       string `json:"dsnw"`
					Hr3Fhsc    string `json:"hr3Fhsc"`
					Dc10Tca    string `json:"dc10Tca"`
					Dc10LmcsCa string `json:"dc10LmcsCa"`
					ClfmAbbrCd string `json:"clfmAbbrCd"`
					LcsCh      string `json:"lcsCh"`
					Vs         string `json:"vs"`
					GndSttCd   string `json:"gndSttCd"`
					DmstMtphNo string `json:"dmstMtphNo"`
					Ts         string `json:"ts"`
					TsQcflg    string `json:"tsQcflg"`
					M005Te     string `json:"m005Te"`
					M01Te      string `json:"m01Te"`
					M02Te      string `json:"m02Te"`
					M03Te      string `json:"m03Te"`
				} `json:"item"`
			} `json:"items"`
			PageNo     int `json:"pageNo"`
			NumOfRows  int `json:"numOfRows"`
			TotalCount int `json:"totalCount"`
		} `json:"body"`
	} `json:"response"`
}

type XMLData struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		ResultCode string `xml:"resultCode"`
		ResultMsg  string `xml:"resultMsg"`
	} `xml:"header"`
	Body struct {
		Text     string `xml:",chardata"`
		DataType string `xml:"dataType"`
		Items    struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text       string `xml:",chardata"`
				Tm         string `xml:"tm"`
				Rnum       string `xml:"rnum"`
				StnId      string `xml:"stnId"`
				StnNm      string `xml:"stnNm"`
				Ta         string `xml:"ta"`
				TaQcflg    string `xml:"taQcflg"`
				Rn         string `xml:"rn"`
				RnQcflg    string `xml:"rnQcflg"`
				Ws         string `xml:"ws"`
				WsQcflg    string `xml:"wsQcflg"`
				Wd         string `xml:"wd"`
				WdQcflg    string `xml:"wdQcflg"`
				Hm         string `xml:"hm"`
				HmQcflg    string `xml:"hmQcflg"`
				Pv         string `xml:"pv"`
				Td         string `xml:"td"`
				Pa         string `xml:"pa"`
				PaQcflg    string `xml:"paQcflg"`
				Ps         string `xml:"ps"`
				PsQcflg    string `xml:"psQcflg"`
				Ss         string `xml:"ss"`
				SsQcflg    string `xml:"ssQcflg"`
				Icsr       string `xml:"icsr"`
				Dsnw       string `xml:"dsnw"`
				Hr3Fhsc    string `xml:"hr3Fhsc"`
				Dc10Tca    string `xml:"dc10Tca"`
				Dc10LmcsCa string `xml:"dc10LmcsCa"`
				ClfmAbbrCd string `xml:"clfmAbbrCd"`
				LcsCh      string `xml:"lcsCh"`
				Vs         string `xml:"vs"`
				GndSttCd   string `xml:"gndSttCd"`
				DmstMtphNo string `xml:"dmstMtphNo"`
				Ts         string `xml:"ts"`
				TsQcflg    string `xml:"tsQcflg"`
				M005Te     string `xml:"m005Te"`
				M01Te      string `xml:"m01Te"`
				M02Te      string `xml:"m02Te"`
				M03Te      string `xml:"m03Te"`
			} `xml:"item"`
		} `xml:"items"`
		NumOfRows  string `xml:"numOfRows"`
		PageNo     string `xml:"pageNo"`
		TotalCount string `xml:"totalCount"`
	} `xml:"body"`
}

func SetDay(nDay int) string {

	now := time.Now()

	var day time.Time

	day = now.AddDate(0, 0, nDay)
	sDay := fmt.Sprintf("%04d%02d%02d", day.Year(), day.Month(), day.Day())
	return sDay
}

func SetMonth(nMonth int) string {

	now := time.Now()

	var month time.Time

	month = now.AddDate(0, nMonth, 0)
	sMonth := fmt.Sprintf("%04d%02d%02d", month.Year(), month.Month(), month.Day())
	return sMonth
}

func GetUrlString(nType int) string {

	spage := "https://apis.data.go.kr/1360000/AsosHourlyInfoService/getWthrDataList?"

	serviceKey := "%2FSWbuoncrZtSM3DaBUA4PJVxqJMFKs0Eu%2F%2FzgFQf8dvVjzIi8ESOjmRaQtAkLKoQUS3S%2BZy%2FwLwR08%2BCT9BWuA%3D%3D"
	pageNo := 1
	numOfRows := 10

	dataType := ""
	if nType == 1 {
		dataType = "XML"
	} else {
		dataType = "JSON"
	}

	dataCd := "ASOS"
	dateCd := "HR"
	startDt := SetDay(-1) //하루전   // SetMonth(-1) 한달전
	startHh := "01"
	endDt := SetDay(-1) // 하루전
	endHh := "01"
	stnIds := "108" // 서울

	sUrl := fmt.Sprintf("%sserviceKey=%s&pageNo=%d&numOfRows=%d&dataType=%s&dataCd=%s&dateCd=%s&startDt=%s&startHh=%s&endDt=%s&endHh=%s&stnIds=%s",
		spage, serviceKey, pageNo, numOfRows, dataType, dataCd, dateCd, startDt, startHh, endDt, endHh, stnIds)

	return sUrl
}

func loadUrl_Xml() {

	nType := 1
	sUrl := GetUrlString(nType)

	res, err := http.Get(sUrl)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return
	}
	//fmt.Println(string(body))

	var result XMLData
	xmlerr := xml.Unmarshal(body, &result)
	if xmlerr != nil {
		panic(xmlerr)
	}
	//fmt.Println(result)

	for _, item := range result.Body.Items.Item {
		fmt.Println("시간: " + item.Tm)
		fmt.Println("지점번호: " + item.StnId)
		fmt.Println("지점명: " + item.StnNm)
		fmt.Println("기온: " + item.Ta)
		fmt.Println("풍속: " + item.Ws)
		fmt.Println("강수량: " + item.Rn)
		fmt.Println("풍향: " + item.Wd)
		fmt.Println("습도: " + item.Hm)
	}
}

func loadUrl_JSON() {

	nType := 2
	sUrl := GetUrlString(nType)

	res, err := http.Get(sUrl)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return
	}
	//fmt.Println(string(body))

	var result JSONData
	jserr := json.Unmarshal(body, &result)
	if jserr != nil {
		panic(jserr)
	}
	//fmt.Println(result)

	for _, item := range result.Response.Body.Items.Item {
		fmt.Println("시간: " + item.Tm)
		fmt.Println("지점번호: " + item.StnID)
		fmt.Println("지점명: " + item.StnNm)
		fmt.Println("기온: " + item.Ta)
		fmt.Println("풍속: " + item.Ws)
		fmt.Println("강수량: " + item.Rn)
		fmt.Println("풍향: " + item.Wd)
		fmt.Println("습도: " + item.Hm)
	}

}

func main() {

	loadUrl_Xml()
	loadUrl_JSON()
}
