package main

import (
	//	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the homepage")
	fmt.Println("Endpoint Hit: HomePage")
}

func handleRequest(port string) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello2", Desc: "Desc 2", Content: "Content 2"},
	}
	port := os.Getenv("PORT")
	fmt.Println(os.Getenv("PORT"))
	fmt.Println(os.Getenv("DATABASE_URL"))
	//	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	//	if err != nil {
	//		fmt.Println("error in opening")
	//		panic(err)
	//	}
	//	defer db.Close()
	//
	//	err = db.Ping()
	//	if err != nil {
	//		panic(err)
	//	}

	fmt.Println("CONNECTED")

	handleRequest(port)
}
