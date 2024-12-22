package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Yershuaq/Assignment1/library"
)

func main() {
	lib := library.NewLibrary()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Add\n2. Borrow\n3. Return\n4. List\n5. Exit")
		fmt.Print("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter Book ID: ")
			id, _ := reader.ReadString('\n')
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			fmt.Print("Enter Author: ")
			author, _ := reader.ReadString('\n')

			lib.AddBook(library.Book{
				ID:     strings.TrimSpace(id),
				Title:  strings.TrimSpace(title),
				Author: strings.TrimSpace(author),
			})
		case "2":
			fmt.Print("Enter Book ID to Borrow: ")
			id, _ := reader.ReadString('\n')
			lib.BorrowBook(strings.TrimSpace(id))
		case "3":
			fmt.Print("Enter Book ID to Return: ")
			id, _ := reader.ReadString('\n')
			lib.ReturnBook(strings.TrimSpace(id))
		case "4":
			lib.ListBooks()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
