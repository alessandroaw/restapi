package models

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

// Book Struct (Model)
type Book struct {
	ID     int     `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Books is collection of books
type Books []*Book

// ToJSON serialize book object to JSON
func (b *Book) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}

// ToJSON serialize book object to JSON
func (b *Books) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}

// FromJSON deserialize JSON to book object
func (b *Book) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(b)
}

// GetBooks get all the book saved in memory
func GetBooks() Books {
	return bookList
}

// GetBook returns book with specific id
func GetBook(id string) (*Book, error) {
	book, _, err := findBook(id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// CreateBook add new book to the book list
func CreateBook(b *Book) {
	b.ID = getNextID()
	bookList = append(bookList, b)
}

// UpdateBook update book information by overwriting it
func UpdateBook(id string, b *Book) error {
	book, index, err := findBook(id)

	if err == ErrBookNotFound {
		return err
	}

	b.ID = book.ID
	bookList[index] = b

	return nil
}

// ======================
// HELPER FUNCTION BELOW
// ======================

// ErrBookNotFound is thrown when no book with given signature is found
var ErrBookNotFound = fmt.Errorf("Book not found")

func findBook(id string) (*Book, int, error) {
	idInt, _ := strconv.Atoi(id)

	for index, item := range bookList {
		if item.ID == idInt {
			return item, index, nil
		}
	}
	return nil, -1, ErrBookNotFound
}

func getNextID() int {
	lastBook := bookList[len(bookList)-1]
	return lastBook.ID + 1
}

// Example data
var bookList = []*Book{
	{
		ID:    1,
		Isbn:  "44324434",
		Title: "Book One",
		Author: &Author{
			Firstname: "john",
			Lastname:  "doe",
		},
	},

	{
		ID:    2,
		Isbn:  "44324435",
		Title: "Book Two",
		Author: &Author{
			Firstname: "john",
			Lastname:  "doer",
		},
	},
}
