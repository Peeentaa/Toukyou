package model

import "time"

type StockSymbol struct {
	Currency string
	Symbol   string
	Type     string
}

type StockQuote struct {
	CurrentPrice  float64 // current price
	Change        float64 // change
	PercentChange float64 // percent change
	HighPrice     float64 // high price of the day
	LowPrice      float64 // low price of the day
	PreviousPrice int     // previous close price
	Date          time.Time
	StockSymbol   string
}

type CompanyNews struct {
}
