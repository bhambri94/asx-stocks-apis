package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

type YahooData struct {
	Prices []struct {
		Num0 struct {
			Date     int     `json:"date"`
			Open     float64 `json:"open"`
			High     float64 `json:"high"`
			Low      float64 `json:"low"`
			Close    float64 `json:"close"`
			Volume   int     `json:"volume"`
			Adjclose float64 `json:"adjclose"`
		} `json:"0,omitempty"`
		Num1 struct {
			Date     int     `json:"date"`
			Open     float64 `json:"open"`
			High     float64 `json:"high"`
			Low      float64 `json:"low"`
			Close    float64 `json:"close"`
			Volume   int     `json:"volume"`
			Adjclose float64 `json:"adjclose"`
		} `json:"1,omitempty"`
	} `json:"prices"`
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
	firstHeader := true
	secondHeader := false

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
			row = append(row, LatestTradeData[6:10], LatestTradeData[6:10])
			firstHeader = false
			secondHeader = true
		} else if secondHeader {
			row = append(row, "O", "C")
			secondHeader = false
		} else {
			row = append(row, OpenPriceMap[symbols[i][2]], ClosePriceMap[symbols[i][2]])

		}
		values = append(values, row)
	}
	return values
}

func GenerateFinalDailyAlertsSheet(symbols [][]string) [][]interface{} {
	var values [][]interface{}

	for i := range symbols {
		var row []interface{}
		for j := range symbols[i] {
			row = append(row, symbols[i][j])
		}
		row = append(row, ClosePriceMap[symbols[i][2]])
		values = append(values, row)
	}
	return values
}

func GetHistoryData(symbols [][]string) [][]interface{} {
	// i := 0
	// for i <= 1 {
	// 	url := "https://apidojo-yahoo-finance-v1.p.rapidapi.com/stock/v2/get-historical-data?frequency=1d&filter=history&period1=1586205431&period2=1596141431&symbol=" + symbols[i][2] + ".AX"

	// 	req, _ := http.NewRequest("GET", url, nil)

	// 	req.Header.Add("x-rapidapi-host", "apidojo-yahoo-finance-v1.p.rapidapi.com")
	// 	req.Header.Add("x-rapidapi-key", "c39f558a71msh3e3e3309afb685dp16544djsnd0d85557e9dc")

	// 	resp, _ := http.DefaultClient.Do(req)

	// 	defer resp.Body.Close()

	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	var ASXLatestData YahooData
	// 	json.Unmarshal(body, &ASXLatestData)
	// 	fmt.Println(ASXLatestData)

	// 	// ineedamap := make(map[string]interface{})
	// 	// ineedamap[symbols[i][2]] = ASXLatestData.Prices
	// 	// ineedamap[]
	// 	// for j := range ASXLatestData.Prices {
	// 	// 	i, err := strconv.ParseInt(strconv.Itoa(ASXLatestData.Prices[j].Num0.Date), 10, 64)
	// 	// 	if err != nil {
	// 	// 		panic(err)
	// 	// 	}
	// 	// 	tm := time.Unix(i, 0)
	// 	// 	OpenPriceMap[ASXLatestData.Data[j].Symbol+tm.Day()] = ASXLatestData.Data[j].Open
	// 	// 	ClosePriceMap[ASXLatestData.Data[j].Symbol+DaymonthDate] = ASXLatestData.Data[j].Close
	// 	// 	fmt.Println(ASXLatestData.Data[j].Symbol + DaymonthDate)
	// 	// 	j++
	// 	// }
	// 	break
	// }

	var values [][]interface{}
	var row []interface{}
	for i := range symbols {
		row = append(row, symbols[i][0], symbols[i][2])
	}
	values = append(values, row)

	fmt.Println(values)
	return values
}
