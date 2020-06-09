package main

import (
	"log"
	"os"

	"encoding/json"
	"fmt"

	// "log"

	"net/http"

	"github.com/CurtisMIT/COL-server/controllers/get"
	"github.com/CurtisMIT/COL-server/controllers/get/individual"
	"github.com/CurtisMIT/COL-server/controllers/post/profile"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	fmt.Println(Articles)
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the best go app ever.")
	fmt.Println("Endpoint Hit: HomePage")
}

func handleRequest(port string) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/profiles", get.ReturnProfilesReq)
	http.HandleFunc("/currency", get.ReturnCurrencyReq)
	http.HandleFunc("/individual/header/", individual.ReturnHeaderReq)
	http.HandleFunc("/individual/earnings/", individual.ReturnEarningsReq)
	http.HandleFunc("/individual/growth/", individual.ReturnGrowthReq)
	http.HandleFunc("/individual/expenses/", individual.ReturnExpensesReq)
	http.HandleFunc("/individual/market/", individual.ReturnMarketReq)
	http.HandleFunc("/individual/others/", individual.ReturnOthersReq)
	// post
	http.HandleFunc("/post/profile", profile.CreateProfile)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello2", Desc: "Desc 2", Content: "Content 2"},
	}
	port := os.Getenv("PORT")
	fmt.Println(os.Getenv("PORT"))
	handleRequest(port)

}
