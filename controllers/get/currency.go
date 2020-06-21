package get

import (
	"encoding/json"
	"net/http"

	"github.com/CurtisMIT/COL-server/database"
)

type currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Currency []currency

func ReturnCurrencyReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	currencies := returnCurrencyDB()
	json.NewEncoder(w).Encode(currencies)
}

func returnCurrencyDB() Currency {
	db := database.DBCON
	rows, err := db.Query(`
		SELECT
			name,
			symbol
		FROM currency
	`)
	if err != nil {
		panic(err)
	}
	var currencyData Currency
	for rows.Next() {
		c := currency{}
		rows.Scan(&c.Name, &c.Symbol)
		currencyData = append(currencyData, c)
	}
	return currencyData
}
