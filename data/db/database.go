package db

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
	"toukyou/data/model"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() (*sql.DB, error) {
	wd, _ := os.Getwd()                    // .../Toukyou
	wd += "/data/db/"                      // .../Toukyou/data/db/
	dbPath := filepath.Join(wd, "data.db") // .../Toukyou/data/db/data.db

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error pinging database: %v", err)
	}
	return db, nil
}

func InsertStockSymbol(symbol model.StockSymbol) {
	db, err := OpenDatabase()
	if err != nil {
		log.Error("Error opening database in InsertStockSymbols method:", err)
		return
	}
	var count int

	query := "select count(*) from StockSymbols where symbol=? and type=?"

	err = db.QueryRow(query, symbol.Symbol, symbol.Type).Scan(&count)
	if err != nil {
		log.Error("Error checking if symbol exists:", err)
		return
	}
	if count > 0 {
		log.Infof("Symbol (%s) already exists in the database. Skipping insertion.", symbol.Symbol)
		return
	}

	query = `insert into StockSymbols (currency, symbol, type) values (?, ?, ?);`

	_, err = db.Exec(query, symbol.Currency, symbol.Symbol, symbol.Type)
	if err != nil {
		log.Error("Error inserting stock symbol:", err)
		return
	}
	log.Info("Stock symbols inserted successfully")
}

func GetStockSymbols() []model.StockSymbol {
	db, err := OpenDatabase()

	if err != nil {
		log.Error("Error opening database in GetStockSymbols method:", err)
		return nil
	}

	query := "select currency, symbol, type from StockSymbols"
	rows, err := db.Query(query)

	if err != nil {
		log.Error("Error executing query:", err)
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var stockSymbols []model.StockSymbol
	for rows.Next() {
		var symbol model.StockSymbol

		err := rows.Scan(&symbol.Currency, &symbol.Symbol, &symbol.Type)
		if err != nil {
			log.Error("Error scanning row:", err)
			return nil
		}

		stockSymbols = append(stockSymbols, symbol)
	}

	if err := rows.Err(); err != nil {
		log.Error("Error iterating over rows:", err)
		return nil
	}

	return stockSymbols
}

func InsertStockQuote(quote model.StockQuote) {
	db, err := OpenDatabase()
	if err != nil {
		log.Error("Error opening database in InsertStockQuotes method:", err)
		return
	}

	query := "insert into StockQuotes (currentPrice, change, percentChange, highPrice, lowPrice, previousPrice, date, stocksymbol) values (?,?,?,?,?,?,?,?);"

	_, err = db.Exec(query, quote.CurrentPrice, quote.Change, quote.PercentChange, quote.HighPrice, quote.LowPrice, quote.PreviousPrice, time.Now().String(), quote.StockSymbol)
	if err != nil {
		log.Error("Error inserting stock quote:", err)
		return
	}

	log.Info("Inserted Stock Quote: " + quote.StockSymbol)

}
