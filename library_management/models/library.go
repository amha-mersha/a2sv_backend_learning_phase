package models

type Library struct {
	Books  map[int64]*Book
	Member map[int64]*Member
}

type LibraryManager interface {
	AddBook(book Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book
}
