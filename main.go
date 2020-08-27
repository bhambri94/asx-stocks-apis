package main

import (
	config "github.com/bhambri94/asx-stocks-apis/configs"
	"github.com/bhambri94/asx-stocks-apis/sheets"
	"github.com/bhambri94/asx-stocks-apis/stocks"
)

func main() {
	config.SetConfig()
	GetLatestMarketData()
	DailyAlerts()
	// GetHistoryData()
}

func GetLatestMarketData() {
	HistorySheetData := sheets.BatchGet(config.Configurations.ReadHistorySheetDetails)
	values := stocks.GetLatestData(HistorySheetData)
	SheetName := "LatestData" + "!B2"
	sheets.BatchWrite(SheetName, values)
}

func DailyAlerts() {
	dailyAlertsData := sheets.BatchGet(config.Configurations.ReadSymbolCodeFrom)
	symbolValues, closePriceValues := stocks.GenerateFinalDailyAlertsSheet(dailyAlertsData)
	SheetName := "DailyAlerts" + "!A3"
	sheets.BatchWrite(SheetName, symbolValues)
	SheetName = "DailyAlerts" + "!G3"
	sheets.BatchWrite(SheetName, closePriceValues)
}

func GetHistoryData() {
	dailyAlertsData := sheets.BatchGet("SPX!B2:D1000")
	WriteToHistorySheet(dailyAlertsData)
}

func WriteToHistorySheet(dailyAlertsData [][]string) {
	var finalValues [][]interface{}
	var row []interface{}
	i := 0
	for i < len(dailyAlertsData) {
		row = append(row, dailyAlertsData[i][0], dailyAlertsData[i][2])
		i++
	}
	finalValues = append(finalValues, row)
	sheets.BatchWrite("LatestData!ES98", finalValues)
}
