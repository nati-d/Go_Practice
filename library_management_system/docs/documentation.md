# Library Management System

## Overview

The Library Management System is a console-based application built with Go. It provides functionalities to manage books and members in a library. Users can add, delete, borrow, return, and search for books, as well as manage member information.



## Features

- Add, delete, borrow, and return books.
- List available and borrowed books.
- Search for books by title or ID.
- Manage library members.


## Usage

Once the application is running, you will see a menu with various options. Enter the corresponding number for the action you want to perform and follow the prompts.



## Functionality

### Add New Book

To add a new book to the library, select the `Add new book` option from the menu and provide the required details:
- Book ID
- Book Title
- Book Author

### Delete Book

To delete a book from the library, select the `Delete book` option from the menu and provide the Book ID.

### Borrow Book

To borrow a book, select the `Borrow book` option from the menu and provide the Member ID and Book ID. The book status will be updated to "Borrowed".

### Return Book

To return a borrowed book, select the `Return book` option from the menu and provide the Member ID and Book ID. The book status will be updated to "Available".

### List Available Books

To list all available books, select the `List available books` option from the menu. A table of available books will be displayed.

### List Borrowed Books

To list all books borrowed by a member, select the `List borrowed books` option from the menu and provide the Member ID. A table of borrowed books will be displayed.

### Search Book by Title

To search for a book by its title, select the `Search book by title` option from the menu and provide the Book Title. A table of matching books will be displayed.

### Search Book by ID

To search for a book by its ID, select the `Search book by id` option from the menu and provide the Book ID. The details of the matching book will be displayed.

### Add New Member

To add a new member to the library, select the `Add new member` option from the menu and provide the required details:
- Member ID
- Member Name

### Delete Member

To delete a member from the library, select the `Delete member` option from the menu and provide the Member ID.

## Validation and Edge Cases

The application includes validations and handles edge cases to ensure data integrity:
- Ensures unique Book IDs and Member IDs.
- Validates that input fields are not empty.
- Validates that numerical inputs are valid integers.
- Ensures books can only be borrowed if they are available.
- Ensures books can only be returned if they are currently borrowed.
- Ensures Member IDs exist before performing member-related operations.
