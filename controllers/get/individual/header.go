package individual

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/CurtisMIT/COL-server/database"
	_ "github.com/lib/pq"
)

type individual struct {
	Title      string   `json:"title"`
	Location   string   `json:"location"`
	Industry   string   `json:"industry"`
	Experience int      `json:"experience"`
	Quote      string   `json:"quote"`
	Created_at string   `json:"created_at"`
	Currency   string   `json:"currency"`
	Tags       []string `json:"tags"`
}

type Individual []individual

func ReturnHeaderReq(w http.ResponseWriter, r *http.Request) {
	// can remove in prod, depending on origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// grabbing the parameter for id
	id := strings.TrimPrefix(r.URL.Path, "/individual/header/")
	individual := returnHeaderDB(id)
	json.NewEncoder(w).Encode(individual)
	fmt.Println("#User tried to access db.individual. Roger.")
}

func returnHeaderDB(id string) Individual {
	db := database.DBCON
	rows, err := db.Query(`
	SELECT 
		profiles.title,
		profiles.location,
		profiles.industry,
		profiles.experience,
		profiles.quote,
		profiles.created_at,
		profiles.currency,
		DT.tags 
	FROM profiles 
	INNER JOIN 
		(SELECT individual_id, string_agg(tag, ', ') AS tags 
		FROM tags GROUP  BY 1) DT 
	ON (profiles.individual_id = DT.individual_id)
	WHERE profiles.individual_id = $1`, id)

	if err != nil {
		panic(err)
	}
	var headerData Individual
	var Created_at time.Time
	var Tags string
	for rows.Next() {
		i := individual{}
		rows.Scan(
			&i.Title, &i.Location, &i.Industry,
			&i.Experience, &i.Quote, &Created_at, &i.Currency, &Tags)
		// conversion for FE
		i.Tags = strings.Split(Tags, ", ")
		i.Created_at = Created_at.Format("January 2, 2006")
		headerData = append(headerData, i)
	}
	return headerData
}
