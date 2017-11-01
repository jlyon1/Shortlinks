package main

import (
	"docs/api"
	"docs/database"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

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

	r := mux.NewRouter()

	// Public
	r.HandleFunc("/", api.IndexHandler).Methods("GET")
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe("0.0.0.0:8080", r)

}
