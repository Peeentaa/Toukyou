package fetcher

import (
	"toukyou/data/db"
)

func getStockData() {
	db.InsertStockData()
}
