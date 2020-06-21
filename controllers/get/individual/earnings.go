package individual

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/CurtisMIT/COL-server/database"
)

type earnings struct {
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
type Earnings []earnings

func ReturnEarningsReq(w http.ResponseWriter, r *http.Request) {
	// can remove in prod, depending on origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := strings.TrimPrefix(r.URL.Path, "/individual/earnings/")
	individual := returnEarningsDB(id)
	json.NewEncoder(w).Encode(individual)
	fmt.Println("#User tried to access db.earnings. Roger.")
}

func returnEarningsDB(id string) Earnings {
	db := database.DBCON
	rows, err := db.Query(`
		SELECT 			
			category,
			amount, 
			description
		FROM earnings 
		WHERE individual_id = $1
	`, id)
	if err != nil {
		panic(err)
	}
	var earningsData Earnings
	for rows.Next() {
		e := earnings{}
		rows.Scan(&e.Category, &e.Amount, &e.Description)
		earningsData = append(earningsData, e)
	}
	return earningsData
}

type growth struct {
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Amount int    `json:"amount"`
}
type Growth []growth

func ReturnGrowthReq(w http.ResponseWriter, r *http.Request) {
	// can remove in prod, depending on origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := strings.TrimPrefix(r.URL.Path, "/individual/growth/")
	individual := returnGrowthDB(id)
	json.NewEncoder(w).Encode(individual)
	fmt.Println("#User tried to access db.growth. Roger.")
}

func returnGrowthDB(id string) Growth {
	db := database.DBCON
	rows, err := db.Query(`
		SELECT 
			title, 
			year, 
			amount
		FROM PAST
		WHERE individual_id = $1
	`, id)
	if err != nil {
		panic(err)
	}
	var growthData Growth
	for rows.Next() {
		g := growth{}
		rows.Scan(&g.Title, &g.Year, &g.Amount)
		growthData = append(growthData, g)
	}
	return growthData
}
