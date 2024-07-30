package main


type Books struct {
	BookID int `json:"book_id"`
	BookTitle string `json:"book_name"`
	BookAuthor string `json:"book_author"`
	BookStatus string `json:"book_status"`
}