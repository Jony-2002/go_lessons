package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name     string
	Surname  string
	Phone    string
	Login    string
	Password string
}

type Login struct {
	Login    string
	Password string
}

type Book struct {
	Name   string
	Author string
}

type GetBook struct {
	Name string
}

type BringBack struct {
	Login string
	Name  string
}

var Users = []User{}

var Books = []Book{
	{
		Name:   "backend",
		Author: "Masrur",
	},
	{
		Name:   "front-end",
		Author: "Jahongir",
	},
	{
		Name:   "fullstack",
		Author: "Shuhrat",
	},
}

var Cart = []BringBack{}

func signUp(w http.ResponseWriter, r *http.Request) {
	var NewUser User

	json.NewDecoder(r.Body).Decode(&NewUser)

	if NewUser.Name == "" || NewUser.Surname == "" || NewUser.Phone == "" || NewUser.Login == "" || NewUser.Password == "" {
		fmt.Fprintf(w, "All fields are required!")
	}

	Users = append(Users, NewUser)
	fmt.Println(Users)

	marshalledUser, _ := json.Marshal(NewUser)

	w.Write(marshalledUser)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	var User Login

	json.NewDecoder(r.Body).Decode(&User)

	hasFound := false

	for _, v := range Users {
		if User.Login == v.Login && User.Password == v.Password {
			hasFound = true

			http.SetCookie(w, &http.Cookie{
				Name:  "Access_Token",
				Value: v.Login,
			})

			marshalledData, _ := json.Marshal(v)

			w.WriteHeader(200)
			w.Write(marshalledData)
			// fmt.Fprintf(w, "Success")
		}
	}

	if hasFound == false {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Not found")
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {

	marshalledData, _ := json.Marshal(Books)

	cookie, err := r.Cookie("Access_Token")

	if err != nil {
		w.WriteHeader(400)
		fmt.Println(err)
		fmt.Fprintf(w, "You don't have access to the book list!")
	} else {
		hasCookie := false

		for _, v := range Users {
			if cookie.Value == v.Login {
				hasCookie = true
				break
			}
		}

		if hasCookie == true {
			w.WriteHeader(200)
			w.Write(marshalledData)
		} else {
			w.WriteHeader(400)
			fmt.Fprintf(w, "You don't have cookie")
		}

	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var NewBook GetBook

	json.NewDecoder(r.Body).Decode(&NewBook)

	cookie, err := r.Cookie("Access_Token")

	if err != nil {
		w.WriteHeader(400)
		fmt.Println(err)
		fmt.Fprintf(w, "You don't have access to get book")
	} else {
		hasCookie := false

		for _, v := range Users {
			// Find Login which matches cookie
			if cookie.Value == v.Login {
				hasCookie = true
				fmt.Println("hasCookie")
				break
			}
		}

		if hasCookie == true {
			hasFound := false

			// Find User
			for _, v := range Users {
				if cookie.Value == v.Login {
					hasFound = true
					fmt.Println("Found user")
				}
			}

			if hasFound == false {
				fmt.Fprintf(w, "Invalid User")
			} else {
				// Find Book
				for _, v := range Books {
					if NewBook.Name == v.Name {
						fmt.Println("Found Book")

						Cart = append(Cart, BringBack{
							Login: cookie.Value,
							Name: v.Name,
						})

						marshalledCart, _ := json.Marshal(Cart)

						w.WriteHeader(200)
						w.Write(marshalledCart)
					}
				}
			}
		} else {
			fmt.Fprintf(w, "You don't have cookie")
		}
	}
}

func bringBack(w http.ResponseWriter, r *http.Request) {
	var Book BringBack

	json.NewDecoder(r.Body).Decode(&Book)

	cookie, err := r.Cookie("Access_Token")

	if err != nil {
		fmt.Fprintf(w, "You don't have cookie to bring back book")
	} else {
		hasCookie := false

		// Check cookie
		for _, v := range Users {
			if cookie.Value == v.Login {
				fmt.Println("Yast Cookie")
				hasCookie = true
				break
			}
		}

		if hasCookie == true {
			// Find book
			for i, v := range Cart {
				if cookie.Value == v.Login && Book.Name == v.Name {
					fmt.Println("Virudum")

					Cart = append(Cart[:i], Cart[i+1:]...)
					// Cart = append(Cart[:i])

					marshalledCart, _ := json.Marshal(Cart)
					w.Write(marshalledCart)
				}
			}
		} else {
			fmt.Fprintf(w, "You don't have cookie to bring back book")
		}
	}

}

func main() {
	fmt.Println("Server is running on: 127.0.0.1")

	http.HandleFunc("/signUp", signUp)
	http.HandleFunc("/signIn", signIn)
	http.HandleFunc("/getBooks", getBooks)
	http.HandleFunc("/getBook", getBook)
	http.HandleFunc("/bringBack", bringBack)

	http.ListenAndServe(":9999", nil)
}
