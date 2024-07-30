package controllers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/rodaine/table"
	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	service services.LibraryManagement
}

func NewLibraryController() *LibraryController {
	return &LibraryController{
		service: services.NewLibraryService(),
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func acceptInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func validateInput(input string) bool {
	return input != ""
}

func validateIntegerInput(input string) (bool, int) {
	convertedInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input, please enter a valid integer value.")
		return false, 0
	}
	return true, convertedInput
}

func centeredText(text string, width int) string {
	padding := (width - len(text)) / 2
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", padding)
}

func styledHeaders(length int) {
	fmt.Println(strings.Repeat("=", length))
}

func buildTable(books []models.Book) {
	tbl := table.New("ID", "Title", "Author", "Status")
	tbl.WithPadding(10).WithHeaderSeparatorRow(1)
	for _, book := range books {
		tbl.AddRow(book.ID, book.Title, book.Author, book.Status)
	}
	tbl.Print()
}

func (lc *LibraryController) Run() {
	for {
		clearConsole()
		styledHeaders(60)
		fmt.Println(centeredText("Welcome to the Library Management System!", 60))
		styledHeaders(60)
		fmt.Println("")
		fmt.Println("1. Add new book")
		fmt.Println("2. Delete book")
		fmt.Println("3. Borrow book")
		fmt.Println("4. Return book")
		fmt.Println("5. List available books")
		fmt.Println("6. List borrowed books")
		fmt.Println("7. Search book by title")
		fmt.Println("8. Search book by id")
		fmt.Println("9. Add new member")
		fmt.Println("10. Delete member")
		fmt.Println("11. Exit")
		fmt.Print("Enter your choice: ")

		choice := acceptInput()
		switch choice {
		case "1":
			lc.AddNewBook()
		case "2":
			lc.DeleteBook()
		case "3":
			lc.BorrowBook()
		case "4":
			lc.ReturnBook()
		case "5":
			lc.ListAvailableBooks()
		case "6":
			lc.ListBorrowedBooks()
		case "7":
			lc.SearchBookByTitle()
		case "8":
			lc.SearchBookById()
		case "9":
			lc.AddNewMember()
		case "10":
			lc.DeleteMember()
		case "11":
			fmt.Println("Thank you for using the Library Management System! Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
		fmt.Println("\nPress Enter to continue...")
		acceptInput()
	}
}

func (lc *LibraryController) AddNewBook() {
	for {
		fmt.Print("Enter book id: ")
		bookIdInput := acceptInput()
		valid, bookId := validateIntegerInput(bookIdInput)
		if valid {
			for {
				fmt.Print("Enter book title: ")
				bookTitle := acceptInput()
				if validateInput(bookTitle) {
					for {
						fmt.Print("Enter book author: ")
						bookAuthor := acceptInput()
						if validateInput(bookAuthor) {
							book := models.Book{ID: bookId, Title: bookTitle, Author: bookAuthor, Status: "Available"}
							err := lc.service.AddNewBook(book)
							if err != nil {
								fmt.Println(err.Error())
							} else {
								fmt.Println("Book added successfully!")
							}
							return
						} else {
							fmt.Println("Book author cannot be empty, please try again.")
						}
					}
				} else {
					fmt.Println("Book title cannot be empty, please try again.")
				}
			}
		}
	}
}

func (lc *LibraryController) DeleteBook() {
	for {
		fmt.Print("Enter book id: ")
		bookIdInput := acceptInput()
		valid, bookId := validateIntegerInput(bookIdInput)
		if valid {
			err := lc.service.DeleteBook(bookId)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Book deleted successfully!")
			}
			return
		}
	}
}

func (lc *LibraryController) BorrowBook() {
	for {
		fmt.Print("Enter member id: ")
		memberIdInput := acceptInput()
		valid, memberId := validateIntegerInput(memberIdInput)
		if valid {
			for {
				fmt.Print("Enter book id: ")
				bookIdInput := acceptInput()
				valid, bookId := validateIntegerInput(bookIdInput)
				if valid {
					err := lc.service.BorrowBook(bookId, memberId)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("Book borrowed successfully!")
					}
					return
				}
			}
		}
	}
}

func (lc *LibraryController) ReturnBook() {
	for {
		fmt.Print("Enter member id: ")
		memberIdInput := acceptInput()
		valid, memberId := validateIntegerInput(memberIdInput)
		if valid {
			for {
				fmt.Print("Enter book id: ")
				bookIdInput := acceptInput()
				valid, bookId := validateIntegerInput(bookIdInput)
				if valid {
					err := lc.service.ReturnBook(bookId, memberId)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("Book returned successfully!")
					}
					return
				}
			}
		}
	}
}

func (lc *LibraryController) ListAvailableBooks() {
	books, err := lc.service.ListAvailableBooks()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if len(books) == 0 {
			fmt.Println("No available books found.")
		} else {
			styledHeaders(60)
			fmt.Println(centeredText("Available books:", 60))
			styledHeaders(60)
			buildTable(books)
		}
	}
}

func (lc *LibraryController) ListBorrowedBooks() {
	for {
		fmt.Print("Enter member id: ")
		memberIdInput := acceptInput()
		valid, memberId := validateIntegerInput(memberIdInput)
		if valid {
			books, err := lc.service.ListBorrowedBooks(memberId)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				if len(books) == 0 {
					fmt.Println("No borrowed books found.")
				} else {
					styledHeaders(60)
					fmt.Println(centeredText("Borrowed books:", 60))
					styledHeaders(60)
					buildTable(books)
				}
			}
			return
		}
	}
}

func (lc *LibraryController) SearchBookByTitle() {
	for {
		fmt.Print("Enter book title: ")
		bookTitle := acceptInput()
		if validateInput(bookTitle) {
			books, err := lc.service.SearchBookByTitle(bookTitle)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				styledHeaders(60)
				fmt.Println(centeredText("Search results:", 60))
				styledHeaders(60)
				buildTable(books)
			}
			return
		} else {
			fmt.Println("Book title cannot be empty, please try again.")
		}
	}
}

func (lc *LibraryController) SearchBookById() {
	for {
		fmt.Print("Enter book id: ")
		bookIdInput := acceptInput()
		valid, bookId := validateIntegerInput(bookIdInput)
		if valid {
			book, err := lc.service.SearchBookById(bookId)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				styledHeaders(60)
				fmt.Println(centeredText("Book details:", 60))
				styledHeaders(60)
				buildTable([]models.Book{book})
			}
			return
		}
	}
}

func (lc *LibraryController) AddNewMember() {
	for {
		fmt.Print("Enter member id: ")
		memberIdInput := acceptInput()
		valid, memberId := validateIntegerInput(memberIdInput)
		if valid {
			for {
				fmt.Print("Enter member name: ")
				memberName := acceptInput()
				if validateInput(memberName) {
					member := models.Member{ID: memberId, Name: memberName, BorrowedBooks: make(map[int]models.Book)}
					err := lc.service.AddNewMember(member)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("Member added successfully!")
					}
					return
				} else {
					fmt.Println("Member name cannot be empty, please try again.")
				}
			}
		}
	}
}

func (lc *LibraryController) DeleteMember() {
	for {
		fmt.Print("Enter member id: ")
		memberIdInput := acceptInput()
		valid, memberId := validateIntegerInput(memberIdInput)
		if valid {
			err := lc.service.DeleteMember(memberId)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Member deleted successfully!")
			}
			return
		}
	}
}
