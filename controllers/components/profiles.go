package components

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"COL-server/account"

	"github.com/lib/pq"
)

type profile struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	// experience int
	// location   string
	// industry   string
	// earnings   int
	// expenses   int
	// quote      string
}

type Profiles []profile

func OpenDb() *sql.DB {
	url := account.Url
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("err")
		log.Println(err)
	}

	fmt.Println("init db")
	return db
}

func ProfileList(w http.ResponseWriter, r *http.Request) {
	profiles := GetProfiles()
	fmt.Println(profiles)
	json.NewEncoder(w).Encode(profiles)
}

func GetProfiles() Profiles {
	db := OpenDb()

	rows, err := db.Query("SELECT id, title FROM profiles")
	if err != nil {
		panic(err)
	}
	var profiles Profiles
	for rows.Next() {
		var id int
		var title string
		rows.Scan(&id, &title)
		fmt.Printf("%3v | %8v\n", id, title)
		profiles = append(profiles, profile{Id: id, Title: title})
	}
	db.Close()
	return profiles
}
