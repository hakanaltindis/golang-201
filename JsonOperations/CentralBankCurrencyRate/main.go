package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

type CurrencyDay struct {
	Id         string
	Date       time.Time
	DayNo      string
	Currencies []Currency
}

type Currency struct {
	Code            string
	CrossOrder      int
	Unit            int
	CurrencyNameTR  string
	CurrencyName    string
	BanknoteBuying  float64
	BanknoteSelling float64
	ForexBuying     float64
	ForexSelling    float64
	CrossRateUSD    float64
	CrossRateOther  float64
}
type tarih_Date struct {
	XMLName   xml.Name `xml:"Tarih_Date"`
	Tarih     string   `xml:"Tarih,attr"`
	Date      string   `xml:"Date,attr"`
	Bulten_No string   `xml:"Bulten_No,attr"`
	Currency  []xmlCurrency
}

type xmlCurrency struct {
	Kod             string `xml:"Kod,attr"`
	CrossOrder      string `xml:"CrossOrder,attr"`
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	Unit            string `xml:"Unit"`
	Isim            string `xml:"Isim"`
	CurrencyName    string `xml:"CurrencyName"`
	ForexBuying     string `xml:"ForexBuying"`
	ForexSelling    string `xml:"ForexSelling"`
	BanknoteBuying  string `xml:"BanknoteBuying"`
	BanknoteSelling string `xml:"BanknoteSelling"`
	CrossRateUSD    string `xml:"CrossRateUSD"`
	CrossRateOther  string `xml:"CrossRateOther"`
}

func (c *CurrencyDay) GetData(currencyDate time.Time) {
	xDate := currencyDate
	t := new(tarih_Date)
	currDay := t.getData(currencyDate, xDate)

	for {
		if currDay == nil {
			currencyDate = currencyDate.AddDate(0, 0, -1)
			currDay := t.getData(currencyDate, xDate)
			if currDay != nil {
				break
			}
		} else {
			break
		}
	}
}

func (c *tarih_Date) getData(currencyDate time.Time, xDate time.Time) *CurrencyDay {
	currDay := new(CurrencyDay)
	var resp *http.Response
	var err error
	var url string

	currDay = new(CurrencyDay)
	url = "https://www.tcmb.gov.tr/kurlar/" + currencyDate.Format("200601") + "/" + currencyDate.Format("02012006") + ".xml"
	fmt.Println("The url:", url)
	resp, err = http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			tarih := new(tarih_Date)
			d := xml.NewDecoder(resp.Body)
			marshalError := d.Decode(&tarih)

			if marshalError != nil {
				log.Printf("error: %v\n", marshalError)
			}

			c = &tarih_Date{}
			currDay.Id = xDate.Format("20060102")
			currDay.Date = xDate
			currDay.DayNo = tarih.Bulten_No
			currDay.Currencies = make([]Currency, len(tarih.Currency))
			for i, curr := range tarih.Currency {
				currDay.Currencies[i].Code = curr.CurrencyCode
				currDay.Currencies[i].CurrencyName = curr.CurrencyName
				currDay.Currencies[i].CurrencyNameTR = curr.Isim
				currDay.Currencies[i].BanknoteBuying, _ = strconv.ParseFloat(curr.BanknoteBuying, 64)
				currDay.Currencies[i].BanknoteSelling, _ = strconv.ParseFloat(curr.BanknoteSelling, 64)
				currDay.Currencies[i].ForexBuying, _ = strconv.ParseFloat(curr.ForexBuying, 64)
				currDay.Currencies[i].ForexSelling, _ = strconv.ParseFloat(curr.ForexSelling, 64)
				currDay.Currencies[i].CrossOrder, _ = strconv.Atoi(curr.CrossRateOther)
				currDay.Currencies[i].CrossRateUSD, _ = strconv.ParseFloat(curr.CrossRateUSD, 64)
				currDay.Currencies[i].CrossRateOther, _ = strconv.ParseFloat(curr.CrossRateOther, 64)
				currDay.Currencies[i].Unit, _ = strconv.Atoi(curr.Unit)
			}

			fmt.Println(currDay)
			SaveJSON("currencies.json", currDay)
		} else {
			currDay = nil
		}

	}
	return currDay
}

func SaveJSON(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error :", err.Error())
		os.Exit(1)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	startTime := time.Now()
	CurrencyDay := new(CurrencyDay)
	CurrencyDate := time.Now()
	CurrencyDay.GetData(CurrencyDate)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s", elapsedTime)
}
