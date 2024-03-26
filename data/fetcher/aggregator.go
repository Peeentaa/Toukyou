package fetcher

import (
	"time"
	"toukyou/data/model"
)

type rawStockSymbol struct {
	Currency      string
	Description   string
	DisplaySymbol string
	Figi          string
	Mic           string
	Symbol        string
	Type          string
}

type transformedStockSymbol struct {
	Currency string
	Symbol   string
	Type     string
}

type rawStockQuote struct {
	C  float64 // current price
	H  float64 // change
	L  float64 // percent change
	O  float64 // high price of the day
	Pc float64 // low price of the day
	T  int     // previous close price
}

type transformedStockQuote struct {
	currentPrice  float64 // current price
	change        float64 // change
	percentChange float64 // percent change
	highPrice     float64 // high price of the day
	lowePrice     float64 // low price of the day
	previousPrice int     // previous close price
	date          time.Time
}

func transformSymbol(symbol rawStockSymbol) model.StockSymbol {
	stockSymbol := model.StockSymbol{
		Currency: symbol.Currency,
		Symbol:   symbol.Symbol,
		Type:     symbol.Type,
	}

	return stockSymbol
}

func transformStockQuote(quote rawStockQuote, stocksymbol string) model.StockQuote {
	stockQuote := model.StockQuote{
		CurrentPrice:  quote.C,
		Change:        quote.H,
		PercentChange: quote.L,
		HighPrice:     quote.O,
		LowPrice:      quote.Pc,
		PreviousPrice: quote.T,
		Date:          time.Now(),
		StockSymbol:   stocksymbol,
	}
	return stockQuote
}
