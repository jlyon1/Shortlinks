package main

import (
	"docs/api"
	"docs/database"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	//"encoding/json"
)

type Article struct{
	Title string `json:"title"`
	Link  string `json:"link"`
	Text  string `json:"text"`
	Image string `json:"image"`
}

func connectDB(db database.DB) {
	fmt.Printf("Connected: %v\n", db.Connect())

}

func main() {
	test := &database.Redis{}
	test.IP = "localhost"
	test.Port = "6379"
	test.DB = 0
	test.Password = ""
	connectDB(test)

	api := api.API{
		Database: test,
	}

	r := mux.NewRouter()

	// Public
	r.HandleFunc("/", api.IndexHandler).Methods("GET")
	r.HandleFunc("/add", api.AddIndexHandler).Methods("GET")
	r.HandleFunc("/add", api.SetHandler).Methods("POST")
	r.HandleFunc("/s/{val}",api.ShortLink).Methods("GET")

	r.HandleFunc("/get/",api.GetHandler).Methods("GET")
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe("0.0.0.0:8081", r)

}
