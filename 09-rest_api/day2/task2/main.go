package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strings"
)

// 1. Can we use the User struct for both registration and login

type User struct {
	Name     string
	Surname  string
	LastName string
	Login    string
	Password string
}

type UserLogin struct {
	Login    string
	Password string
}

var Users  = []User {
	{
		Name: "Jahongir",
		Surname: "Niyozbekov",
		LastName: "Dalerovich",
		Login: "jony@gmail.com",
		Password: "jony123",
	},
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var newUser User

	json.NewDecoder(r.Body).Decode(&newUser)

	newUser.checkUserCredentials(w)
}

func login(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var user UserLogin

	json.NewDecoder(r.Body).Decode(&user)

	user.checkLoginCredentials(w)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	marshaledSlice, _ := json.Marshal(Users)
	// Fprintf - can only takes string as the second argument
	// fmt.Fprintf(w, string(marshaledSlice))

	w.Write(marshaledSlice)

	// 1. What is the difference between w.Write and fmt.Fprintf?

	fmt.Println("Success")
	fmt.Println(string(marshaledSlice))
}

func main() {
	fmt.Println("server is listening on 127.0.0.1")

	http.HandleFunc("/registration", createNewUser)
	http.HandleFunc("/login", login)
	http.HandleFunc("/getUsers", getAllUsers)

	http.ListenAndServe(":9999", nil)
}

func (obj User) checkUserCredentials(w http.ResponseWriter) {
	var isExistingUser bool = false

	for _, v := range Users {
		if obj.Login == v.Login {
			isExistingUser = true
			fmt.Println("User with this email has already exist. Try another one!")
			fmt.Fprintf(w, "User with this email has already exist. Try another one!")
		}
	}

	if isExistingUser == false {
		if obj.Name == "" {
			fmt.Fprintf(w, "Name cannot be empty")
			fmt.Println("Name cannot be empty")
		} else if obj.LastName == "" {
			fmt.Fprintf(w, "LastName cannot be empty")
			fmt.Println("LastName cannot be empty")
		} else if obj.Surname == "" {
			fmt.Fprintf(w, "Surname cannot be empty")
			fmt.Println("Surname cannot be empty")
		} else if obj.Password == "" {
			fmt.Fprintf(w, "Password cannot be empty")
			fmt.Println("Password cannot be empty")
		} else if len(obj.Password) < 7 {
			fmt.Fprintf(w, "Password must contain more than 6 characters")
			fmt.Println("Password must contain more than 6 characters")
		} else if obj.Login == "" {
			fmt.Fprintf(w, "Login cannot be empty")
			fmt.Println("Login cannot be empty")
		} else {
			fmt.Fprintf(w, "Success")
			fmt.Println("Success")
			fmt.Println(obj)

			Users = append(Users, obj)
			fmt.Println(Users)
			fmt.Println("-------------------------------------------------")

			fmt.Fprintf(w, "%+v\n", obj)
			fmt.Fprintf(w, "%+v", Users)
		}
	}
}

func (user UserLogin) checkLoginCredentials(w http.ResponseWriter) {
	var hasFound bool = false

	if user.Login == "" {
		fmt.Println("Login cannot be empty")
		fmt.Fprintf(w, "Login cannot be empty")
	} else if user.Password == "" {
		fmt.Println("Password cannot be empty")
		fmt.Fprintf(w, "Password cannot be empty")
	} else {
		for i := 0; i < len(Users); i++ {
			if user.Login == Users[i].Login {
				hasFound = true
				if user.Password == Users[i].Password {
					fmt.Println("Success")
					fmt.Fprintf(w, "Success")
					break
				} else {
					fmt.Println("Wrong password")
					fmt.Fprintf(w, "Wrong password")
					break
				}
			}
		}

		if hasFound == false {
			fmt.Println("cannot find user")
			fmt.Fprintf(w, "cannot find user")
		}
	}
}
