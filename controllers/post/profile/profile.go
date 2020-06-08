package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CurtisMIT/COL-server/controllers/get"
)

type big struct {
	// json field needs to be the same
	Test []Test `json:"test"`
}
type Test struct {
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Amount int    `json:"amount"`
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	decoder := json.NewDecoder(r.Body)
	var c big
	// var t test_struct
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
	insertProfile(c)

}

func insertProfile(tag big) {
	db := get.OpenDb()
	// figure out how to enter multiple rows, array of objects
	db.Query(`
		insert into past(individual_id, title, year, amount) 
		SELECT 45 id, x 
		FROM UNNEST($1::text[], $2::integer[], $3::integer[]) x`, tag)

	fmt.Printf("inserted %v\n", tag)
	db.Close()

}

// profile table
// earnings table
// expenses table

// past table
// stored as array of objects

// queries for tags table
// stored as array of strings in FE
// Test pq.StringArray `json:"test"`
// decoder := json.NewDecoder(r.Body)
// var t test_struct
// err := decoder.Decode(&t)
// db.Query(`
// insert into tags(individual_id, tag)
// SELECT 45 id, x
// FROM UNNEST($1::text[]) x`, tag.Test)
