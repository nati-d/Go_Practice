package services

import (
	"errors"
	"library_management/models"
)


type LibraryManagement interface {
	AddNewBook(book models.Book) error
	DeleteBook(bookId int) error
	BorrowBook(bookId int, memberId int) error
	ReturnBook(bookId int, memberId int) error
	ListAvailableBooks() ([]models.Book, error)
	ListBorrowedBooks(memberId int) ([]models.Book, error)
	SearchBookByTitle(title string) ([]models.Book, error)
	SearchBookById(bookId int) (models.Book, error)
	AddNewMember(member models.Member) error
	DeleteMember(memberId int) error	
}

type LibraryService struct {
	books map[int]models.Book
	members map[int]models.Member
}

func NewLibraryService() *LibraryService {
	return &LibraryService{
		books: make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}


func (ls *LibraryService) AddNewBook(book models.Book) error {
	_,exists := ls.books[book.ID]
	if exists {
		return errors.New("book already exists")
	}
	ls.books[book.ID] = book
	return nil
}

func(ls *LibraryService) DeleteBook(bookId int) error {
	_,exists := ls.books[bookId]
	if !exists {
		return errors.New("book does not exist")
	}

	delete(ls.books, bookId)
	return nil
}

func (ls *LibraryService) BorrowBook(bookId, memberId int ) error {
	member,exists := ls.members[memberId]
	if !exists {
		return errors.New("member does not exist")
	}
	book,exists := ls.books[bookId]
	if !exists {
		return errors.New("book does not exist")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}
	book.Status = "Borrowed"
	member.BorrowedBooks[bookId] = book
	ls.books[bookId] = book

	return nil
}

func (ls *LibraryService) ReturnBook(bookId, memberId int) error {
	member,exists := ls.members[memberId]
	if !exists {
		return errors.New("member does not exist")
	}
	book,exists := ls.books[bookId]
	if !exists {
		return errors.New("book does not exist")
	}
	if book.Status == "Available" {
		return errors.New("book is not borrowed")
	}

	delete(member.BorrowedBooks, bookId)
	book.Status = "Available"
	ls.books[bookId] = book
	return nil
}

func (ls *LibraryService) ListAvailableBooks() ([]models.Book, error) {
	var books []models.Book
	for _,book := range ls.books {
		if book.Status == "Available" {
			books = append(books, book)
		}
	}
	return books, nil
}

func (ls *LibraryService) ListBorrowedBooks(memberId int) ([]models.Book, error) {
	member,exists := ls.members[memberId]
	if !exists {
		return nil, errors.New("member does not exist")
	}
	var books []models.Book
	for _,book := range member.BorrowedBooks {
		books = append(books, book)
	}
	return books, nil
}

func (ls *LibraryService) SearchBookByTitle(title string) ([]models.Book, error) {
	var books []models.Book
	for _,book := range ls.books {
		if book.Title == title {
			books = append(books, book)
		}
	}
	return books, nil
}

func (ls *LibraryService) SearchBookById(bookId int) (models.Book, error) {
	book,exists := ls.books[bookId]
	if !exists {
		return models.Book{}, errors.New("book does not exist")
	}
	return book, nil
}

func (ls *LibraryService) AddNewMember(member models.Member) error {
	_,exists := ls.members[member.ID]
	if exists {
		return errors.New("member already exists")
	}
	ls.members[member.ID] = member
	return nil
}


func (ls *LibraryService) DeleteMember(memberId int) error {
	_,exist := ls.members[memberId]

	if !exist {
		return errors.New("member does not exist")
	}

	delete(ls.members, memberId)
	return nil
}


