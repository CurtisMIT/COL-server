package profile

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CurtisMIT/COL-server/controllers/get"
	"github.com/lib/pq"
)

type Profile struct {
	Individual_id int            `json:"id"`
	Title         string         `json:"title"`
	Location      string         `json:"location"`
	Industry      string         `json:"industry"`
	Experience    int            `json:"experience"`
	Earnings      int            `json:"earnings"`
	Expenses      int            `json:"expenses"`
	Quote         string         `json:"quote"`
	Tags          pq.StringArray `json:"tags"`
	Breakdown     []Extra        `json:"breakdownList"`
	ExpenseList   []Extra        `json:"expenseList"`
	PastList      []Past         `json:"pastList"`
	Currency      string         `json:"currency"`
}
type Extra struct {
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
type Past struct {
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Amount int    `json:"amount"`
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	decoder := json.NewDecoder(r.Body)
	var p Profile
	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	// fmt.Println(p)
	insertProfile(p)
	insertTags(p)
	insertEarnings(p)
	insertExpenses(p)
	insertPast(p)
}

func insertProfile(p Profile) {
	db := get.OpenDb()
	_, err := db.Query(`
		INSERT INTO profiles(
		individual_id, title, location, industry, experience, earnings, expenses, quote, currency)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9 )`,
		p.Individual_id,
		p.Title,
		p.Location,
		p.Industry,
		p.Experience,
		p.Earnings,
		p.Expenses,
		p.Quote,
		p.Currency)
	if err != nil {
		panic(err)
	}
}

func insertTags(p Profile) {
	db := get.OpenDb()
	_, err := db.Query(`
		INSERT INTO tags(individual_id, tag)
		SELECT $1 id, x
		FROM UNNEST($2::text[]) x`, p.Individual_id, p.Tags)
	if err != nil {
		panic(err)
	}
}

func insertEarnings(p Profile) {
	db := get.OpenDb()
	samples := p.Breakdown
	query := `insert into earnings(individual_id, category, amount, description) values `
	values := []interface{}{}
	for i, s := range samples {
		values = append(values, p.Individual_id, s.Category, s.Amount, s.Description)
		numFields := 4
		n := i * numFields
		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}
	query = query[:len(query)-1]
	_, err := db.Query(query, values...)
	if err != nil {
		panic(err)
	}
}

func insertExpenses(p Profile) {
	db := get.OpenDb()
	samples := p.ExpenseList
	query := `insert into expenses(individual_id, category, amount, description) values `
	values := []interface{}{}
	for i, s := range samples {
		values = append(values, p.Individual_id, s.Category, s.Amount, s.Description)
		numFields := 4
		n := i * numFields
		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}
	query = query[:len(query)-1]
	_, err := db.Query(query, values...)
	if err != nil {
		panic(err)
	}
}

func insertPast(p Profile) {
	db := get.OpenDb()
	samples := p.PastList
	query := `insert into past(individual_id, title, year, amount) values `
	values := []interface{}{}
	for i, s := range samples {
		values = append(values, p.Individual_id, s.Title, s.Year, s.Amount)
		numFields := 4
		n := i * numFields
		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}
	query = query[:len(query)-1]
	_, err := db.Query(query, values...)
	if err != nil {
		panic(err)
	}
}
