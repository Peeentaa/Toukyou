package fetcher

import "time"

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

func transformSymbol(symbol rawStockSymbol) transformedStockSymbol {
	stockSymbol := transformedStockSymbol{
		Currency: symbol.Currency,
		Symbol:   symbol.Symbol,
		Type:     symbol.Type,
	}

	return stockSymbol
}

func transformStockQuote(quote rawStockQuote) transformedStockQuote {
	stockQuote := transformedStockQuote{
		currentPrice:  quote.C,
		change:        quote.H,
		percentChange: quote.L,
		highPrice:     quote.O,
		lowePrice:     quote.Pc,
		previousPrice: quote.T,
		date:          time.Now(),
	}
	return stockQuote
}
