package get

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	//"github.com/CurtisMIT/COL-server/account"
	_ "github.com/lib/pq"
)

type profile struct {
	Individual_ID int      `json:"individual_id"`
	Title         string   `json:"title"`
	Location      string   `json:"location"`
	Industry      string   `json:"industry"`
	Experience    int      `json:"experience"`
	Earnings      int      `json:"earnings"`
	Expenses      int      `json:"expenses"`
	Quote         string   `json:"quote"`
	Created_at    string   `json:"created_at"`
	Currency      string   `json:"currency"`
	Tags          []string `json:"tags"`
}
type Profiles []profile

func OpenDb() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	//	connection, _ := pq.ParseURL(url)
	//connection += " sslmode=require"
	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("err")
		log.Println(err)
	}
	fmt.Println("#Successfully connected to db. Roger.")
	return db
}

func ReturnProfilesReq(w http.ResponseWriter, r *http.Request) {
	// can remove in prod, depending on origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	profiles := returnProfilesDB()
	json.NewEncoder(w).Encode(profiles)
	fmt.Println("#User tried to access db.profiles. Roger.")
}

func returnProfilesDB() Profiles {
	db := OpenDb()

	rows, err := db.Query(`
		SELECT 
			profiles.*, 
			DT.tags 
		FROM profiles 
		INNER JOIN 
			(SELECT individual_id, string_agg(tag, ', ') AS tags 
			FROM tags GROUP  BY 1) DT 
		ON (profiles.individual_id = DT.individual_id)`)

	if err != nil {
		panic(err)
	}

	var profilesData Profiles
	for rows.Next() {
		// use profile struct
		p := profile{}
		// type for created_at, tags => helps with FE formatting
		var Created_at time.Time
		var Tags string
		rows.Scan(
			&p.Individual_ID, &p.Title,
			&p.Location, &p.Industry, &p.Experience, &p.Earnings,
			&p.Expenses, &p.Quote, &Created_at, &p.Currency, &Tags)
		// convert string to array for FE
		p.Tags = strings.Split(Tags, ", ")
		p.Created_at = Created_at.Format("January 2, 2006")
		profilesData = append(profilesData, p)
	}

	return profilesData
}
