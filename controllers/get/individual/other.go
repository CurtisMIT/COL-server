package individual

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/CurtisMIT/COL-server/controllers/get"
)

type profiles struct {
	Individual_ID int      `json:"individual_id"`
	Title         string   `json:"title"`
	Location      string   `json:"location"`
	Industry      string   `json:"industry"`
	Experience    int      `json:"experience"`
	Earnings      int      `json:"earnings"`
	Expenses      int      `json:"expenses"`
	Quote         string   `json:"quote"`
	Created_at    string   `json:"created_at"`
	Tags          []string `json:"tags"`
}
type Profiles []profiles

func ReturnOthersReq(w http.ResponseWriter, r *http.Request) {
	// can remove in prod, depending on origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// grabbing the parameter for id
	location := strings.TrimPrefix(r.URL.Path, "/individual/others/")
	profiles := returnOthersDB(location)
	// fmt.Println(profiles)
	json.NewEncoder(w).Encode(profiles)
	fmt.Println("#User tried to access db.others. Roger.")
}

func returnOthersDB(location string) Profiles {
	db := get.OpenDb()
	rows, err := db.Query(`
		SELECT
			profiles.*,
			DT.tags
		FROM profiles
		INNER JOIN 
			(SELECT individual_id, string_agg(tag, ', ') AS tags 
			FROM tags GROUP  BY 1) DT 
		ON (profiles.individual_id = DT.individual_id)		
		WHERE location = $1
		limit 3
	`, location)
	if err != nil {
		panic(err)
	}
	var othersData Profiles
	for rows.Next() {
		p := profiles{}
		var Created_at time.Time
		var Tags string
		rows.Scan(
			&p.Individual_ID, &p.Title,
			&p.Location, &p.Industry, &p.Experience, &p.Earnings,
			&p.Expenses, &p.Quote, &Created_at, &Tags)
		// convert string to array for FE
		p.Tags = strings.Split(Tags, ", ")
		p.Created_at = Created_at.Format("January 2, 2006")
		othersData = append(othersData, p)
	}
	db.Close()
	return othersData
}
