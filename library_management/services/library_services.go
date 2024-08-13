package services

import (
	"errors"
	"fmt"
	"library_management/models"
	"strings"
)

type LibraryManager struct {
	Library *models.Library
}

func (lib *LibraryManager) AddBook(book models.Book) error {
	if len(lib.Library.Books) > 0 && lib.Library.Books[book.ID] != nil {
		return errors.New("book already present in library")
	} else {
		lib.Library.Books[book.ID] = &book
		fmt.Println("Book added to library.")
		fmt.Println(lib.Library.Books)
		return nil
	}
}

func (lib *LibraryManager) RemoveBoook(bookID int64) {
	delete(lib.Library.Books, bookID)
	fmt.Println("Book removed")
}

func (lib *LibraryManager) BorrowBook(bookID int64, memberID int64) error {
	if lib.Library.Books[bookID] == nil {
		return fmt.Errorf("no book exists with the book id: %d", bookID)
	} else if strings.Compare("Available", lib.Library.Books[bookID].Status) != 0 {
		return fmt.Errorf("the book with the book id: %d is already borrowed", bookID)
	} else {
		lib.Library.Books[bookID].Status = "Borrowed"
		if len(lib.Library.Member) == 0 || lib.Library.Member[memberID] == nil {
			newMember := models.Member{
				ID:            memberID,
				Name:          "",
				BorrowedBooks: []*models.Book{},
			}
			lib.Library.Member[memberID] = &newMember
		}
		for _, book := range lib.Library.Member[memberID].BorrowedBooks {
			if bookID == book.ID {
				return fmt.Errorf("you already have borrowed that book")
			}
		}
		lib.Library.Member[memberID].BorrowedBooks = append(lib.Library.Member[memberID].BorrowedBooks, lib.Library.Books[bookID])
		fmt.Println("Here is your book:")

		fmt.Printf("ID : %q\n", lib.Library.Books[bookID].ID)
		fmt.Printf("Title : %q\n", lib.Library.Books[bookID].Title)
		fmt.Printf("Author : %q\n", lib.Library.Books[bookID].Author)
		return nil
	}
}

func (lib *LibraryManager) ReturnBook(bookID int64, memberID int64) error {
	if lib.Library.Books[bookID] == nil {
		return fmt.Errorf("no book exists with the book id: %v", bookID)
	} else if strings.Compare("Borrowed", lib.Library.Books[bookID].Status) != 0 {
		return fmt.Errorf("the book with the book id: %v isn't borrowed", bookID)
	} else {
		lib.Library.Books[bookID].Status = "Available"
		delete(lib.Library.Books, bookID)
		fmt.Println("Book have been returned to library")
		return nil
	}
}

func (lib *LibraryManager) ListAvailableBooks() []models.Book {
	availableBooks := make([]models.Book, 0, len(lib.Library.Books))
	for _, book := range lib.Library.Books {
		if strings.Compare("Available", book.Status) == 0 {
			availableBooks = append(availableBooks, *book)
		}
	}
	fmt.Println(availableBooks)
	return availableBooks
}

func (lib *LibraryManager) ListBorrowedBooks(memberID int64) ([]*models.Book, error) {
	if len(lib.Library.Member) == 0 || lib.Library.Member[memberID] == nil {
		newMember := models.Member{
			ID:            memberID,
			Name:          "",
			BorrowedBooks: []*models.Book{},
		}
		lib.Library.Member[memberID] = &newMember
	}

	return lib.Library.Member[memberID].BorrowedBooks, nil
}
