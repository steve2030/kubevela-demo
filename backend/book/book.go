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
	{ID: "1", Title: "Designing Data-Intensive Applications", Author: "Martin Kleppmann", Year: 2017, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "2", Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Year: 2015, Image: "https://images-na.ssl-images-amazon.com/images/I/41a6U2U+1bL._SX331_BO1,204,203,200_.jpg"},
	{ID: "3", Title: "Clean Code", Author: "Robert C. Martin", Year: 2008, Image: "https://images-na.ssl-images-amazon.com/images/I/41xShlnTZTL._SX376_BO1,204,203,200_.jpg"},
	{ID: "4", Title: "Kubernetes: Up and Running", Author: "Kelsey Hightower", Year: 2019, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "5", Title: "Cloud Native Go", Author: "Kevin Hoffman", Year: 2021, Image: "https://images-na.ssl-images-amazon.com/images/I/41a6U2U+1bL._SX331_BO1,204,203,200_.jpg"},
	{ID: "6", Title: "Cloud Native Patterns", Author: "Cornelia Davis", Year: 2019, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "7", Title: "Kubernetes Patterns", Author: "Bilgin Ibryam", Year: 2020, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "8", Title: "Cloud Native DevOps with Kubernetes", Author: "John Arundel", Year: 2020, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "9", Title: "Kubernetes in Action", Author: "Marko Luksa", Year: 2018, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "10", Title: "Cloud Native Java", Author: "Josh Long", Year: 2017, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "11", Title: "Kubernetes Best Practices", Author: "Brendan Burns", Year: 2020, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "12", Title: "Cloud Native Infrastructure", Author: "Justin Garrison", Year: 2019, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "13", Title: "Kubernetes Cookbook", Author: "Sébastien Goasguen", Year: 2018, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "14", Title: "Cloud Native Transformation", Author: "Pini Reznik", Year: 2020, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "15", Title: "Kubernetes Security", Author: "Liz Rice", Year: 2021, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "16", Title: "Cloud Native Observability", Author: "Alexandre Couëdelo", Year: 2021, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "17", Title: "Kubernetes Networking", Author: "James Strong", Year: 2021, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "18", Title: "Cloud Native Microservices", Author: "Kasun Indrasiri", Year: 2021, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "19", Title: "Kubernetes for Developers", Author: "William Denniss", Year: 2021, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
	{ID: "20", Title: "Cloud Native Data Management", Author: "Rohit Bhardwaj", Year: 2021, Image: "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg"},
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
