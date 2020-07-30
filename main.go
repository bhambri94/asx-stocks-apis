package main

import (
	// "time"

	config "github.com/bhambri94/asx-stocks-apis/configs"
	"github.com/bhambri94/asx-stocks-apis/sheets"
	"github.com/bhambri94/asx-stocks-apis/stocks"
)

func main() {
	config.SetConfig()
	// GetLatestMarketData()
	// DailyAlerts()
	// HistorySheetData := sheets.BatchGet("Sheet9!B2:D81")
	// values:=stocks.GetHistoryData(HistorySheetData)
	// sheets.BatchWrite("TestMe!E159", values)
}

func GetLatestMarketData() {
	HistorySheetData := sheets.BatchGet(config.Configurations.ReadHistorySheetDetails)
	values := stocks.GetLatestData(HistorySheetData)
	SheetName := "LatestData" + "!B1"
	sheets.BatchWrite(SheetName, values)
}

func DailyAlerts() {
	dailyAlertsData := sheets.BatchGet(config.Configurations.ReadSymbolCodeFrom)
	values := stocks.GenerateFinalDailyAlertsSheet(dailyAlertsData)
	SheetName := "DailyAlerts" + "!A3"
	sheets.BatchWrite(SheetName, values)
}
