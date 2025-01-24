package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

type EBook struct {
	Book
	FileSize int
}

type BookInterface interface {
	DisplayDetails()
}

type Library struct {
	Collection []BookInterface
}

func (b *Book) DisplayDetails() {
	availability := "No"
	if b.Available {
		availability = "Yes"
	}
	fmt.Printf("Title: %s, Author: %s, ISBN: %s, Available: %s\n", b.Title, b.Author, b.ISBN, availability)
}

func (e *EBook) DisplayDetails() {
	availability := "No"
	if e.Available {
		availability = "Yes"
	}
	fmt.Printf("Title: %s, Author: %s, ISBN: %s, File Size: %dMB, Available: %s\n", e.Title, e.Author, e.ISBN, e.FileSize, availability)
}

func (l *Library) AddBook(book BookInterface) {
	l.Collection = append(l.Collection, book)
	fmt.Println("Book/EBook added successfully!")
}

func (l *Library) RemoveBook(isbn string) {
	for i, book := range l.Collection {
		switch b := book.(type) {
		case *Book:
			if b.ISBN == isbn {
				l.Collection = append(l.Collection[:i], l.Collection[i+1:]...)
				fmt.Println("Book removed successfully!")
				return
			}
		case *EBook:
			if b.ISBN == isbn {
				l.Collection = append(l.Collection[:i], l.Collection[i+1:]...)
				fmt.Println("EBook removed successfully!")
				return
			}
		}
	}
	fmt.Println("Book/EBook with the given ISBN not found.")
}

func (l *Library) SearchBookByTitle(title string) {
	found := false
	for _, book := range l.Collection {
		switch b := book.(type) {
		case *Book:
			if strings.Contains(strings.ToLower(b.Title), strings.ToLower(title)) {
				b.DisplayDetails()
				found = true
			}
		case *EBook:
			if strings.Contains(strings.ToLower(b.Title), strings.ToLower(title)) {
				b.DisplayDetails()
				found = true
			}
		}
	}
	if !found {
		fmt.Println("No books found with the given title.")
	}
}

func (l *Library) ListBooks() {
	if len(l.Collection) == 0 {
		fmt.Println("No books/eBooks in the library.")
		return
	}
	for _, book := range l.Collection {
		book.DisplayDetails()
	}
}

func main() {
	library := Library{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a Book/EBook")
		fmt.Println("2. Remove a Book/EBook")
		fmt.Println("3. Search for Books by Title")
		fmt.Println("4. List all Books/EBooks")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\n1. Add a Book\n2. Add an EBook")
			fmt.Print("Choose type: ")
			var bookType int
			fmt.Scanln(&bookType)

			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := scanner.Text()

			fmt.Print("Enter ISBN: ")
			scanner.Scan()
			isbn := scanner.Text()

			fmt.Print("Is it available? (yes/no): ")
			scanner.Scan()
			availableInput := scanner.Text()
			available := strings.ToLower(availableInput) == "yes"

			if bookType == 1 {
				book := &Book{
					Title:     title,
					Author:    author,
					ISBN:      isbn,
					Available: available,
				}
				library.AddBook(book)
			} else if bookType == 2 {
				fmt.Print("Enter File Size (in MB): ")
				var fileSize int
				fmt.Scanln(&fileSize)

				ebook := &EBook{
					Book: Book{
						Title:     title,
						Author:    author,
						ISBN:      isbn,
						Available: available,
					},
					FileSize: fileSize,
				}
				library.AddBook(ebook)
			} else {
				fmt.Println("Invalid choice.")
			}

		case 2:
			fmt.Print("Enter ISBN of the Book/EBook to remove: ")
			scanner.Scan()
			isbn := scanner.Text()
			library.RemoveBook(isbn)

		case 3:
			fmt.Print("Enter Title to search: ")
			scanner.Scan()
			title := scanner.Text()
			library.SearchBookByTitle(title)

		case 4:
			library.ListBooks()

		case 5:
			fmt.Println("Exiting the system.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
