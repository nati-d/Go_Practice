package main

type Member struct{
	MemberID int `json:"member_id"`
	MemberName string `json:"member_name"`
	BorrowedBooks []Books `json:"borrowed_books"`
}