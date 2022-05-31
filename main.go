package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint hit! Homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit! Articles")
	json.NewEncoder(w).Encode(Articles)
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key)
	for _, a := range Articles {
		if a.Id == key {
			json.NewEncoder(w).Encode(a)
			return
		}
	}
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage)
	r.HandleFunc("/articles", returnAllArticles)
	r.HandleFunc("/articles/{id}", returnArticle)
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", r))
}

func main() {
	fmt.Println("Rest API v2. - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Description", Content: "Content"},
		Article{Id: "2", Title: "Hello1", Desc: "Description1", Content: "Content1"},
	}
	handleRequests()
}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:Content`
}

var Articles []Article
