# Library Management System

## Overview

This is a simple console-based library management system implemented in Go. It provides functionalities for managing books and library members, including adding, removing, borrowing, and returning books. The system tracks the books borrowed by each member and offers a console interface for user interaction.

## Structs

- **Book**: Represents a book with the following fields:

  - `ID` (int64): Unique identifier for the book.
  - `Title` (string): Title of the book.
  - `Author` (string): Author of the book.
  - `Status` (string): Current status of the book (e.g., "Available", "Borrowed").

- **Member**: Represents a library member with the following fields:
  - `ID` (int64): Unique identifier for the member.
  - `Name` (string): Name of the member.
  - `BorrowedBooks` ([]\*Book): List of books currently borrowed by the member.

## Interfaces

- **LibraryManager**: Defines the methods for library management, including:
  - `AddBook(book Book)`: Adds a new book to the library.
  - `RemoveBook(bookID int64)`: Removes a book from the library.
  - `BorrowBook(bookID int64, memberID int64)`: Allows a member to borrow a book.
  - `ReturnBook(bookID int64, memberID int64)`: Allows a member to return a borrowed book.
  - `ListAvailableBooks()`: Lists all available books in the library.
  - `ListBorrowedBooks(memberID int64)`: Lists all books borrowed by a specific member.
  - `AddMember(member Member)`: Adds a new member to the library.
  - `RemoveMember(memberID int64)`: Removes a member from the library.
  - `ListMembers()`: Lists all members of the library.

## Implementation

- **Library**: Implements the `LibraryManager` interface and contains logic for managing books and members. The `Library` struct includes methods to add, remove, borrow, return books, and manage members.

## Console Interface

The application provides a simple console interface to interact with the library management system. Users can:

- Add a new book.
- Remove an existing book.
- Borrow a book.
- Return a borrowed book.
- List available books.
- List borrowed books.
- Add a new member.
- Remove a member.
- List all members and their borrowed books.

### Example Commands

- **Add Book**: `Enter the ID, title, author, and status of the book.`
- **Remove Book**: `Enter the ID of the book to remove.`
- **Borrow Book**: `Enter the ID of the book and the ID of the member borrowing the book.`
- **Return Book**: `Enter the ID of the book and the ID of the member returning the book.`
- **List Available Books**: `Displays all books with status "Available".`
- **List Borrowed Books**: `Displays all books borrowed by a specified member.`
- **Add Member**: `Enter the ID and name of the new member.`
- **Remove Member**: `Enter the ID of the member to remove.`
- **List Members**: `Displays all library members and the books they have borrowed.`

## Running the Application

To run the application, use the following command:

```sh
go run main.go
```
