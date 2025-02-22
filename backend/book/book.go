package book

import (
	"errors"
	"strconv"
)

// Book represents a book in the library
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Image  string `json:"image"`
}

// books is a slice of Book structs (in-memory storage)
var books = []Book{
	{ID: "1", Title: "Clean Code", Author: "Robert C. Martin", Year: 2008, Image: "https://images-na.ssl-images-amazon.com/images/I/41xShlnTZTL._SX376_BO1,204,203,200_.jpg"},
	{ID: "2", Title: "Designing Data-Intensive Applications", Author: "Martin Kleppmann", Year: 2017, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "3", Title: "Clean Code", Author: "Robert C. Martin", Year: 2008, Image: "https://images-na.ssl-images-amazon.com/images/I/41xShlnTZTL._SX376_BO1,204,203,200_.jpg"},
	{ID: "4", Title: "Kubernetes: Up and Running", Author: "Kelsey Hightower", Year: 2019, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "5", Title: "Cloud Native Go", Author: "Kevin Hoffman", Year: 2021, Image: "https://images-na.ssl-images-amazon.com/images/I/41a6U2U+1bL._SX331_BO1,204,203,200_.jpg"},
	{ID: "6", Title: "Cloud Native Patterns", Author: "Cornelia Davis", Year: 2019, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "7", Title: "Kubernetes Patterns", Author: "Bilgin Ibryam", Year: 2020, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "8", Title: "Cloud Native DevOps with Kubernetes", Author: "John Arundel", Year: 2020, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "9", Title: "Kubernetes in Action", Author: "Marko Luksa", Year: 2018, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "10", Title: "Cloud Native Java", Author: "Josh Long", Year: 2017, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
}

// GetBooks returns the list of books
func GetBooks() []Book {
	return books
}

// AddBook adds a new book to the list
func AddBook(newBook Book) {
	// Generate a new ID for the book
	newBook.ID = strconv.Itoa(len(books) + 1)
	books = append(books, newBook)
}

// UpdateBook updates an existing book
func UpdateBook(bookID string, updatedBook Book) error {
	for i, book := range books {
		if book.ID == bookID {
			books[i] = updatedBook
			books[i].ID = bookID // Ensure the ID remains the same
			return nil
		}
	}
	return errors.New("book not found")
}

// DeleteBook deletes a book from the list
func DeleteBook(bookID string) error {
	for i, book := range books {
		if book.ID == bookID {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
