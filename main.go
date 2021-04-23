package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the homepage!")
	fmt.Println("Endpoint hit! Homepage")
}

func returnAllArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint hit! Articles")
	json.NewEncoder(writer).Encode(Articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Description", Content:"Content"},
		Article{Title: "Hello1", Desc: "Description1", Content:"Content1"},
	}
	handleRequests()
}

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:Content`
}

var Articles []Article