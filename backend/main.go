package main

import (
	"encoding/json"
	"library/backend/book" // Import the book module
	"net/http"
)

// getBooks returns the list of books as JSON
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book.GetBooks()) // Use the GetBooks function from the book module
}

func main() {
	// Define routes
	http.HandleFunc("/api/books", getBooks)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
