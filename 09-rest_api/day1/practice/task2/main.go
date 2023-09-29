package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8888"

type Book struct {
	Id       string
	BookName string
	Author   string
}

var Books = []Book{
	{
		Id:       "1",
		BookName: "MasrurLife",
		Author:   "Masrur",
	},
	{
		Id:       "2",
		BookName: "The Go Programming Language",
		Author:   "Jahongir",
	},
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	for _, book := range Books {
		// Fprintf - only prints string values
		fmt.Fprintf(w, book.BookName)
	}
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	fmt.Println(id)

	for _, book := range Books {
		if book.Id == id {
			fmt.Fprintf(w, book.BookName)
		}
	}
}

func getAuthors(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(Books); i++ {
		fmt.Fprintf(w, Books[i].Author)
	}
}

func getAuthorById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	for i := 0; i < len(Books); i++ {
		if id == Books[i].Id {
			fmt.Fprintf(w, Books[i].Author)
		}
	}
}

func main() {
	fmt.Println("server is running on port:", PORT)

	// routes
	http.HandleFunc("/getBooks", getAllBooks)
	http.HandleFunc("/getSpecificBook", getBookById)
	http.HandleFunc("/getAuthors", getAuthors)
	http.HandleFunc("/getAuthorById", getAuthorById)

	http.ListenAndServe(PORT, nil)
}
