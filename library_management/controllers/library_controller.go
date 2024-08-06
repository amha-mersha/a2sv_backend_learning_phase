package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func handleADD(lib models.Library) error {
	var (
		title  string
		id     int64
		status string
		author string
	)
	fmt.Println("Enter the title of the book:")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Println("Enter the id of the book:")
	temp, _ := reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	id, _ = strconv.ParseInt(temp, 10, 64)
	status = "Available"
	fmt.Println("Enter the author of the book:")
	author, _ = reader.ReadString('\n')
	newBook := models.Book{
		ID:     id,
		Title:  title,
		Status: status,
		Author: author,
	}
	err := lib.AddBook(newBook)
	if err != nil {
		return err
	}
	return nil
}

func handleREMOVE(lib models.Library) error {
	fmt.Println("Enter the ID of the book:")
	temp, _ := reader.ReadString('\n')
	bookId, _ := strconv.ParseInt(temp, 10, 64)
	lib.RemoveBoook(bookId)
	return nil
}

func handleBORROW(lib models.Library) error {
	fmt.Printf("Enter the ID of the book:")
	temp, _ := reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	bookId, _ := strconv.ParseInt(temp, 10, 64)

	fmt.Print("Enter your ID:")
	temp, _ = reader.ReadString('\n')
	temp = strings.TrimSpace(temp)
	memberId, _ := strconv.ParseInt(temp, 10, 64)
	err := lib.BorrowBook(bookId, memberId)
	if err != nil {
		return err
	}
	return nil
}

func handleRETURN(lib models.Library) error {
	fmt.Printf("Enter the ID of the book:")
	temp, _ := reader.ReadString('\n')
	bookId, _ := strconv.ParseInt(temp, 10, 64)

	fmt.Print("Enter your ID:")
	temp, _ = reader.ReadString('\n')
	memberId, _ := strconv.ParseInt(temp, 10, 64)
	err := lib.ReturnBook(bookId, memberId)
	if err != nil {
		return err
	}
	return nil
}
func handleAVAILABLE(lib models.Library) error {
	lib.ListAvailableBooks()
	return nil
}

func handleUNAVAILABLE(lib models.Library) error {
	fmt.Print("Enter your ID:")
	temp, _ := reader.ReadString('\n')
	memberId, _ := strconv.ParseInt(temp, 10, 64)
	lib.ListBorrowedBooks(memberId)
	return nil
}

func handleQUIT(lib models.Library) error {
	os.Exit(0)
	return nil
}

var availableCommands = map[string]string{
	"Add a book to library":            "ADD",
	"Remove a book from library":       "REMOVE",
	"Borrow a book from library":       "BORROW",
	"Return a book to library":         "RETURN",
	"See available books at library":   "AVAILABLE",
	"See unavailable books at library": "UNAVAILABLE",
	"Quit out of the system":           "QUIT",
}

var CommandtoHandler = map[string]func(models.Library) error{
	"ADD":         handleADD,
	"REMOVE":      handleREMOVE,
	"BORROW":      handleBORROW,
	"RETURN":      handleRETURN,
	"AVAILABLE":   handleAVAILABLE,
	"UNAVAILABLE": handleUNAVAILABLE,
	"QUIT":        handleQUIT,
}

func GetInput(input string) (cmd string, err error) {
	input = strings.TrimSpace(input)
	words := strings.Split(input, " ")
	if CommandtoHandler[words[0]] == nil {
		return cmd, fmt.Errorf("not a valid command input")
	}
	if len(words) > 2 {
		return cmd, fmt.Errorf("not a valid opitions input")
	}
	return words[0], nil
}

func Run(lib *models.Library) error {
	fmt.Println("Hello there, this is a library managment system.")
	fmt.Println("Here are the available commands in this library:")
	maxDescWidth := 0

	for desc := range availableCommands {
		if len(desc) > maxDescWidth {
			maxDescWidth = len(desc)
		}
	}
	for desc, cmd := range availableCommands {
		fmt.Printf("\tTo %-*v :\t %v\n", maxDescWidth, desc, cmd)
	}
	fmt.Println("Enter your command after >")

	for {
		fmt.Printf("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		cmd, err := GetInput(input)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		err = CommandtoHandler[cmd](*lib)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

	}
}
