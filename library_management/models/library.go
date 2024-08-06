package models

import (
	"errors"
	"fmt"
	"strings"
)

type Library struct {
	Books  map[int64]*Book
	Member map[int64]*Member
}

func (lib *Library) AddBook(book Book) error {
	if len(lib.Books) > 0 && lib.Books[book.ID] != nil {
		return errors.New("book already present in library")
	} else {
		lib.Books[book.ID] = &book
		fmt.Println("Book added to library.")
		fmt.Println(lib.Books)
		return nil
	}
}

func (lib *Library) RemoveBoook(bookID int64) {
	delete(lib.Books, bookID)
	fmt.Println("Book removed")
}

func (lib *Library) BorrowBook(bookID int64, memberID int64) error {
	if lib.Books[bookID] == nil {
		return fmt.Errorf("no book exists with the book id: %d", bookID)
	} else if strings.Compare("Available", lib.Books[bookID].Status) != 0 {
		return fmt.Errorf("the book with the book id: %d is already borrowed", bookID)
	} else {
		lib.Books[bookID].Status = "Borrowed"
		if len(lib.Member) == 0 || lib.Member[memberID] == nil {
			newMember := Member{
				ID:            memberID,
				Name:          "",
				BorrowedBooks: []*Book{},
			}
			lib.Member[memberID] = &newMember
		}
		for _, book := range lib.Member[memberID].BorrowedBooks {
			if bookID == book.ID {
				return fmt.Errorf("you already have borrowed that book")
			}
		}
		lib.Member[memberID].BorrowedBooks = append(lib.Member[memberID].BorrowedBooks, lib.Books[bookID])
		fmt.Println("Here is your book:")

		fmt.Printf("ID : %q\n", lib.Books[bookID].ID)
		fmt.Printf("Title : %q\n", lib.Books[bookID].Title)
		fmt.Printf("Author : %q\n", lib.Books[bookID].Author)
		return nil
	}
}

func (lib *Library) ReturnBook(bookID int64, memberID int64) error {
	if lib.Books[bookID] == nil {
		return fmt.Errorf("no book exists with the book id: %v", bookID)
	} else if strings.Compare("Borrowed", lib.Books[bookID].Status) != 0 {
		return fmt.Errorf("the book with the book id: %v isn't borrowed", bookID)
	} else {
		lib.Books[bookID].Status = "Available"
		delete(lib.Books, bookID)
		fmt.Println("Book have been returned to library")
		return nil
	}
}

func (lib *Library) ListAvailableBooks() []Book {
	availableBooks := make([]Book, 0, len(lib.Books))
	for _, book := range lib.Books {
		if strings.Compare("Available", book.Status) == 0 {
			availableBooks = append(availableBooks, *book)
		}
	}
	fmt.Println(availableBooks)
	return availableBooks
}

func (lib *Library) ListBorrowedBooks(memberID int64) ([]*Book, error) {
	if len(lib.Member) == 0 || lib.Member[memberID] == nil {
		newMember := Member{
			ID:            memberID,
			Name:          "",
			BorrowedBooks: []*Book{},
		}
		lib.Member[memberID] = &newMember
	}

	return lib.Member[memberID].BorrowedBooks, nil
}
