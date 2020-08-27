package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	OpenPriceMap    = make(map[string]float64)
	ClosePriceMap   = make(map[string]float64)
	LatestTradeData string
	finalvalues     [][]interface{}
)

type ASXData struct {
	Code                        string  `json:"code"`
	IsinCode                    string  `json:"isin_code"`
	DescFull                    string  `json:"desc_full"`
	LastPrice                   float64 `json:"last_price"`
	OpenPrice                   float64 `json:"open_price"`
	DayHighPrice                float64 `json:"day_high_price"`
	DayLowPrice                 float64 `json:"day_low_price"`
	ChangePrice                 float64 `json:"change_price"`
	ChangeInPercent             string  `json:"change_in_percent"`
	Volume                      int     `json:"volume"`
	BidPrice                    float64 `json:"bid_price"`
	OfferPrice                  float64 `json:"offer_price"`
	PreviousClosePrice          float64 `json:"previous_close_price"`
	PreviousDayPercentageChange string  `json:"previous_day_percentage_change"`
	YearHighPrice               float64 `json:"year_high_price"`
	LastTradeDate               string  `json:"last_trade_date"`
	YearHighDate                string  `json:"year_high_date"`
	YearLowPrice                float64 `json:"year_low_price"`
	YearLowDate                 string  `json:"year_low_date"`
	Pe                          int     `json:"pe"`
	Eps                         int     `json:"eps"`
	AverageDailyVolume          int     `json:"average_daily_volume"`
	AnnualDividendYield         int     `json:"annual_dividend_yield"`
	MarketCap                   int     `json:"market_cap"`
	NumberOfShares              int     `json:"number_of_shares"`
	DeprecatedMarketCap         int     `json:"deprecated_market_cap"`
	DeprecatedNumberOfShares    int     `json:"deprecated_number_of_shares"`
	Suspended                   bool    `json:"suspended"`
}

func GetLatestData(symbols [][]string) [][]interface{} {
	for i := range symbols {
		req, err := http.NewRequest("GET", "https://www.asx.com.au/asx/1/share/"+symbols[i][2], nil)
		if err != nil {
			// handle err
		}
		req.Header.Set("Cookie", "JSESSIONID=.node205; TS0122d459=01a4bc132057a0c4bd1a60c9baf03882902de0d7c9b210ff7595f6427e3dca77446473becff3acb3e879e2d65e5953d1316e90e748eb5e49336ac73fa2495f38fd1682d6b9")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			// handle err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		var ASXLatestData ASXData
		json.Unmarshal(body, &ASXLatestData)

		OpenPriceMap[ASXLatestData.Code] = ASXLatestData.OpenPrice
		ClosePriceMap[ASXLatestData.Code] = ASXLatestData.LastPrice
		LatestTradeData = ASXLatestData.LastTradeDate

	}

	var values [][]interface{}
	firstHeader := false
	secondHeader := true

	for i := range symbols {
		var row []interface{}
		j := 0
		for j = range symbols[i] {
			if i == 0 {
				row = append(row, symbols[i][j])
			} else if i == 1 {
				row = append(row, symbols[i][j])
			} else {
				row = append(row, symbols[i][j])
			}
		}
		if firstHeader {
			newDate := LatestTradeData[5:7] + "-" + LatestTradeData[8:10] + "-" + LatestTradeData[0:4]
			t, err := time.Parse("01-02-2006", newDate)
			if err != nil {
				panic(err)
			}
			week := t.Weekday().String()
			row = append(row, week[0:3]+LatestTradeData[5:7]+"-"+LatestTradeData[8:10]+"-"+LatestTradeData[0:4], LatestTradeData[5:7]+"-"+LatestTradeData[8:10]+"-"+LatestTradeData[0:4])
			firstHeader = false
			secondHeader = true
		} else if secondHeader {
			row = append(row, "O", "C")
			secondHeader = false
		} else {
			row = append(row, fmt.Sprintf("%.3f", OpenPriceMap[symbols[i][2]]), fmt.Sprintf("%.3f", ClosePriceMap[symbols[i][2]]))

		}
		values = append(values, row)
	}
	return values
}

func GenerateFinalDailyAlertsSheet(symbols [][]string) ([][]interface{}, [][]interface{}) {
	var symbolValues [][]interface{}
	var closePriceValues [][]interface{}

	for i := range symbols {
		var row1 []interface{}
		var row2 []interface{}
		for j := range symbols[i] {
			row1 = append(row1, symbols[i][j])
		}
		row2 = append(row2, ClosePriceMap[symbols[i][2]])
		closePriceValues = append(closePriceValues, row2)
		symbolValues = append(symbolValues, row1)
	}
	return symbolValues, closePriceValues
}
