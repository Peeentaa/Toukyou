package fetcher

func GetStockData() {
	fetchStockSymbols()
	fetchStockQuotes()
}

func fetchStockQuotes() []transformedStockQuote {
	stockQuote := rawStockQuote{
		C:  123,
		H:  123,
		L:  13123,
		O:  12313,
		Pc: 13123,
		T:  131232,
	}
	rawStockQuotes := []rawStockQuote{stockQuote}

	var transformedStockQuotes []transformedStockQuote

	for _, q := range rawStockQuotes {
		transformedStockQuotes = append(transformedStockQuotes, transformStockQuote(q))
	}

	return transformedStockQuotes
}

func fetchStockSymbols() []transformedStockSymbol {
	symbol := rawStockSymbol{
		Currency:      "USD",
		Description:   "dasd",
		DisplaySymbol: "AAPL",
		Figi:          "asdasd",
		Mic:           "asdasd",
		Symbol:        "AAPL",
		Type:          "Common",
	}
	rawStockSymbols := []rawStockSymbol{symbol}

	var transformedStockSymbols []transformedStockSymbol

	for _, s := range rawStockSymbols {
		transformedStockSymbols = append(transformedStockSymbols, transformSymbol(s))
	}

	return transformedStockSymbols
}
