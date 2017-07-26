// Bithumburl_Public
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
https://www.bithumb.com/u1/US127
*/

//bithumb 거래소 마지막 거래 정보
type result_data struct {
	Opening_price float64 `json:"opening_price,string"`
	Closing_price float64 `json:"closing_price,string"`
	Min_price     float64 `json:"min_price,string"`
	Max_price     float64 `json:"max_price,string"`
	Average_price float64 `json:"average_price,string"`
	Units_traded  float64 `json:"units_traded,string"`
	Volume_1day   float64 `json:"volume_1day,string"`
	Volume_7day   float64 `json:"volume_7day,string"`
	Buy_price     float64 `json:"buy_price,string"`
	Sell_price    float64 `json:"sell_price,string"`
	Date          float64 `json:"date,string"`
}

type Response struct {
	Status string      `json:"status"`
	Data   result_data `json:"data"`
}

func getbithumbTicker() bool {

	res, err := http.Get("https://api.bithumb.com/public/ticker/BTC")
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	bytes := []byte(string(body))

	var result Response
	json.Unmarshal(bytes, &result)

	if result.Status != "0000" {
		fmt.Println(result.Status)
		return false
	}

	fmt.Printf("Opening_price:%.0f\n", result.Data.Opening_price)
	fmt.Printf("Closing_price:%.0f\n", result.Data.Closing_price)
	fmt.Printf("Min_price:%.0f\n", result.Data.Min_price)
	fmt.Printf("Max_price:%.0f\n", result.Data.Max_price)
	fmt.Printf("Average_price:%.8f\n", result.Data.Average_price)
	fmt.Printf("Units_traded:%.8f\n", result.Data.Units_traded)
	fmt.Printf("Volume_1day:%.8f\n", result.Data.Volume_1day)
	fmt.Printf("Volume_7day:%.8f\n", result.Data.Volume_7day)
	fmt.Printf("Buy_price:%.0f\n", result.Data.Buy_price)
	fmt.Printf("Sell_price:%.0f\n", result.Data.Sell_price)
	fmt.Printf("Date:%.0f\n", result.Data.Date)

	return true
}

//bithumb 거래소 판/구매 등록 대기 또는 거래 중 내역 정보
type price_data struct {
	Quantity float64 `json:"quantity,string"`
	Price    string  `json:"price"`
}

type result_data1 struct {
	Timestamp        float64      `json:"timestamp,string"`
	Order_currency   string       `json:"order_currency"`
	Payment_currency string       `json:"payment_currency"`
	Bids             []price_data `json:"bids"`
	Asks             []price_data `json:"asks"`
}

type Response1 struct {
	Status string       `json:"status"`
	Data   result_data1 `json:"data"`
}

func getbithumborderbook() bool {
	res, err := http.Get("https://api.bithumb.com/public/orderbook/BTC")
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	bytes := []byte(string(body))

	var result Response1
	json.Unmarshal(bytes, &result)

	if result.Status != "0000" {
		fmt.Println(result.Status)
		return false
	}

	//fmt.Println(string(body))

	fmt.Printf("Timestamp:%.0f\n", result.Data.Timestamp)
	fmt.Printf("Timestamp:%s\n", result.Data.Order_currency)
	fmt.Printf("Timestamp:%s\n", result.Data.Payment_currency)
	//fmt.Println(result.Data.Bids)
	//fmt.Println(result.Data.Asks)

	for _, val := range result.Data.Bids {
		fmt.Printf("bids Quantity:%.0f\n", val.Quantity)
		fmt.Printf("bids Timestamp:%s\n", val.Price)
	}

	for _, val := range result.Data.Asks {
		fmt.Printf("asks Quantity:%.0f\n", val.Quantity)
		fmt.Printf("asks Timestamp:%s\n", val.Price)
	}

	return true
}

//bithumb 거래소 거래 체결 완료 내역
type transaction_data struct {
	Transaction_date string `json:"transaction_date"`
	Type             string `json:"type"`
	Units_traded     string `json:"units_traded"`
	Price            string `json:"price"`
	Total            string `json:"total"`
}

type Response2 struct {
	Status string             `json:"status"`
	Data   []transaction_data `json:"data"`
}

func getbithumbrecent_transactions() bool {
	res, err := http.Get("https://api.bithumb.com/public/recent_transactions/BTC")
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	bytes := []byte(string(body))

	var result Response2
	json.Unmarshal(bytes, &result)

	if result.Status != "0000" {
		fmt.Println(result.Status)
		return false
	}

	//fmt.Println(string(body))

	for _, val := range result.Data {
		fmt.Printf("Transaction_date:%s\n", val.Transaction_date)
		fmt.Printf("Type:%s\n", val.Type)
		fmt.Printf("Units_traded:%s\n", val.Units_traded)
		fmt.Printf("Price:%s\n", val.Price)
		fmt.Printf("Total:%s\n", val.Total)
	}

	return true
}

func main() {

	getbithumbTicker()

	getbithumborderbook()

	getbithumbrecent_transactions()
}
