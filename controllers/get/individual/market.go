package individual

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/CurtisMIT/COL-server/database"
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
	id := strings.TrimPrefix(r.URL.Path, "/individual/market/")
	individual := returnMarketDB(id)
	json.NewEncoder(w).Encode(individual)
	fmt.Println("#User tried to access db.earnigns. Roger.")
}

func returnMarketDB(id string) Market {
	db := database.DBCON
	rows, err := db.Query(`
		SELECT 
			profiles.title, 
			profiles.earnings, 
			profiles.expenses, 
			profiles.experience
		FROM profiles
		INNER JOIN
			(SELECT * FROM profiles where individual_id = $1) DT
		on profiles.location = DT.location
	`, id)
	if err != nil {
		panic(err)
	}
	var marketData Market
	for rows.Next() {
		m := market{}
		rows.Scan(&m.Title, &m.Earnings, &m.Expenses, &m.Experience)
		marketData = append(marketData, m)
	}
	return marketData
}
