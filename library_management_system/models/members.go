package models

type Member struct {
    ID            int
    Name          string
    BorrowedBooks map[int]Book
}
