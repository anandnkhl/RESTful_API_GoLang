package main

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{ Title:"Test Title", Desc: "Test desc", Content: "Test Content"},
		Article{ Title: "New", Desc: "New Desc", Content:"New Content"},
	}

	fmt.Println("Endpoint hit allArticles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "Homepage, Endpoint hit")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func main(){
	handleRequests()
}