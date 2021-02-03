package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/therealsandro/restapi/models"
)

// Books is a http.Handler
type Books struct {
	l *log.Logger
}

// NewBooksController return new Book Handler
func NewBooksController(l *log.Logger) *Books {
	return &Books{l}
}

// GetBooks return All Books on memory
func (b Books) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := models.GetBooks()

	err := data.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetBook returning single book with a certain ID
func (b Books) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	data, err := models.GetBook(params["id"])

	if err == models.ErrBookNotFound {
		http.Error(w, "Book not found", http.StatusNotFound)
	}

	err = data.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// CreateBook adding new books to repository
func (b Books) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.Book

	err := data.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	models.CreateBook(&data)

	err = data.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// UpdateBook update book by overwriting it
func (b Books) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var data models.Book

	err := data.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = models.UpdateBook(params["id"], &data)
	if err == models.ErrBookNotFound {
		http.Error(w, "Book not found", http.StatusNotFound)
	}

	err = data.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// // delete Books
// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }
