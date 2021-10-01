package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	//ID     uint
	BookID int    `json:"id" gorm:"unique"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (a *App) ListHandler(w http.ResponseWriter, r *http.Request) {
	var books []Book

	// Select all books and convert to JSON.
	a.DB.Find(&books)
	booksJSON, _ := json.Marshal(books)

	// Write to HTTP response.
	w.WriteHeader(200)
	w.Write([]byte(booksJSON))
}

func (a *App) ViewHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	vars := mux.Vars(r)

	// Select the book with the given name, and convert to JSON.
	a.DB.First(&book, "book_	id = ?", vars["id"])
	bookJSON, _ := json.Marshal(book)

	// Write to HTTP response.
	w.WriteHeader(200)
	w.Write([]byte(bookJSON))
}

func (a *App) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &book)
	a.DB.Create(book)

	// Write to HTTP response.
	w.WriteHeader(201)
}

func (a *App) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	vars := mux.Vars(r)

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &book)

	// Update the star with the given name.
	a.DB.Model(&book).Where("book_id = ?", vars["book_id"]).Updates(&book)

	// Write to HTTP response.
	w.WriteHeader(204)
}

func (a *App) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Delete the star with the given name.
	a.DB.Where("book_id = ?", vars["book_id"]).Delete(Book{})

	// Write to HTTP response.
	w.WriteHeader(204)
}
