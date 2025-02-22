package book

// Book represents a book in the library
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Image  string `json:"image"` // URL or path to the book cover image
}

// GetBooks returns a list of books
func GetBooks() []Book {
	return []Book{
		{
			ID:     "1",
			Title:  "The Go Programming Language",
			Author: "Alan A. A. Donovan",
			Year:   2015,
			Image:  "https://images-na.ssl-images-amazon.com/images/I/41a6U2U+1bL._SX331_BO1,204,203,200_.jpg",
		},
		{
			ID:     "2",
			Title:  "Clean Code",
			Author: "Robert C. Martin",
			Year:   2008,
			Image:  "https://images-na.ssl-images-amazon.com/images/I/41xShlnTZTL._SX376_BO1,204,203,200_.jpg",
		},
		{
			ID:     "3",
			Title:  "Designing Data-Intensive Applications",
			Author: "Martin Kleppmann",
			Year:   2017,
			Image:  "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg",
		},
		{
			ID:     "4",
			Title:  "The Power of Cloud Native",
			Author: "Martin Kleppmann",
			Year:   2017,
			Image:  "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg",
		},
		{
			ID:     "5",
			Title:  "Kubeveral Basics Guide",
			Author: "Martin Kleppmann",
			Year:   2017,
			Image:  "https://images-na.ssl-images-amazon.com/images/I/51ZSpMl1-LL._SX379_BO1,204,203,200_.jpg",
		},
	}
}
