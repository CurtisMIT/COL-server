package individual

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/CurtisMIT/COL-server/controllers/get"
)

type market struct {
	Title      string `json:"title"`
	Earnings   int    `json:"earnings"`
	Expenses   int    `json:"expenses"`
	Experience int    `json:"experience"`
}
type Market []market

func ReturnMarketReq(w http.ResponseWriter, r *http.Request) {
	// can remove in prod, depending on origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// grabbing the parameter for id
	location := strings.TrimPrefix(r.URL.Path, "/individual/market/")
	individual := returnMarketDB(location)
	json.NewEncoder(w).Encode(individual)
	fmt.Println("#User tried to access db.earnigns. Roger.")
}

func returnMarketDB(location string) Market {
	db := get.OpenDb()
	rows, err := db.Query(`
		SELECT 
			title,
			earnings,
			expenses,			
			experience
		FROM profiles
		WHERE location = $1
	`, location)
	if err != nil {
		panic(err)
	}
	var marketData Market
	for rows.Next() {
		m := market{}
		rows.Scan(&m.Title, &m.Earnings, &m.Expenses, &m.Experience)
		marketData = append(marketData, m)
	}
	db.Close()
	return marketData
}
