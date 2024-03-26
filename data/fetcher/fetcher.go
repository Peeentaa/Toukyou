package fetcher

import (
	"fmt"
	"toukyou/data/db"
)

func FetchStockData() {
	fetchStockSymbols()
	fetchStockQuotes()
}

func fetchStockQuotes() {

	stockSymbols := db.GetStockSymbols()

	fmt.Println(stockSymbols)
	for _, s := range stockSymbols {
		stockQuote := rawStockQuote{
			C:  123,
			H:  123,
			L:  13123,
			O:  12313,
			Pc: 13123,
			T:  131232,
		}
		rawStockQuotes := []rawStockQuote{stockQuote}

		for _, q := range rawStockQuotes {
			db.InsertStockQuote(transformStockQuote(q, s.Symbol))
		}
	}
}

func fetchStockSymbols() {
	symbol := rawStockSymbol{
		Currency:      "USD",
		Description:   "dasd",
		DisplaySymbol: "MSFT",
		Figi:          "asdasd",
		Mic:           "asdasd",
		Symbol:        "MSFT",
		Type:          "Common",
	}
	symbol1 := rawStockSymbol{
		Currency:      "USD",
		Description:   "dasd",
		DisplaySymbol: "AAPL",
		Figi:          "asdasd",
		Mic:           "asdasd",
		Symbol:        "AAPL",
		Type:          "Common",
	}
	rawStockSymbols := []rawStockSymbol{symbol, symbol1}

	for _, s := range rawStockSymbols {
		db.InsertStockSymbol(transformSymbol(s))
	}

}
