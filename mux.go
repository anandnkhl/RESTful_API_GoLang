package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Product struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Products []Product

func allProducts(w http.ResponseWriter, r *http.Request){
	products := Products{
		Product{ Title:"Test Title", Desc: "Test desc", Content: "Test Content"},
		Product{ Title: "New", Desc: "New Desc", Content:"New Content"},
	}

	fmt.Println("Endpoint hit allProducts Endpoint")
	json.NewEncoder(w).Encode(products)
}

func testPostProducts(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Use GET request instead of POST")
}

func homePage2(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "Homepage2, Endpoint hit")
}

func handleRequests2(){

	// mux is HTTP request multiplexer
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage2)

	//when /products is called with 'GET' method
	myRouter.HandleFunc("/products", allProducts).Methods("GET")
	//when /products is called with 'POST' method
	myRouter.HandleFunc("/products", testPostProducts).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main(){
	handleRequests2()
}
