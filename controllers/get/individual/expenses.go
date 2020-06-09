package individual

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/CurtisMIT/COL-server/database"
)

type expenses struct {
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type Expenses []expenses

func ReturnExpensesReq(w http.ResponseWriter, r *http.Request) {
	// can remove in prod, depending on origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := strings.TrimPrefix(r.URL.Path, "/individual/expenses/")
	individual := returnExpensesDB(id)
	json.NewEncoder(w).Encode(individual)
}

func returnExpensesDB(id string) Expenses {
	db := database.DBCON
	rows, err := db.Query(`
		SELECT
			category,
			amount,
			description
		FROM expenses
		WHERE individual_id = $1
	`, id)
	if err != nil {
		panic(err)
	}
	var expensesData Expenses
	for rows.Next() {
		e := expenses{}
		rows.Scan(&e.Category, &e.Amount, &e.Description)
		expensesData = append(expensesData, e)
	}
	return expensesData
}
