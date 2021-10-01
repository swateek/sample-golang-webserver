package main

import (
	"log"
	"net/http"
	"sample-golang-bookstore/models"
	"time"

	"github.com/gorilla/mux"
)

var (
	AppObj = &models.App{}
)

func main() {

	AppObj.InitDB("sqlite3", "test.db")

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/books", AppObj.ListHandler).Methods("GET")
	r.HandleFunc("/books/{id:.+}", AppObj.ViewHandler).Methods("GET")
	r.HandleFunc("/books", AppObj.CreateHandler).Methods("POST")
	r.HandleFunc("/books/{id:.+}", AppObj.UpdateHandler).Methods("PATCH")
	r.HandleFunc("/books/{id:.+}", AppObj.DeleteHandler).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	defer AppObj.DB.Close()
}
