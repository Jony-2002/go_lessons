package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	UserName string
	Address  string
	Email    string
	Password string
	Phone    string
}

type Login struct {
	UserName string
	Password string
}

type Response struct {
	Code    int
	Data    any
	Message string
}

var Users = []User{
	{
		UserName: "Jahongir",
		Address:  "Sh Shohtemurov",
		Email:    "thisisjahongeer@gmail.com",
		Password: "jony123",
		Phone:    "+992935224004",
	},
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var newUser User

	json.NewDecoder(r.Body).Decode(&newUser)

	//! Check if any field is empty

	Users = append(Users, newUser)

	
	var res Response = Response{
		Code: 200,
		Message: "Success",
		Data: newUser,
	}
	
	marshalledRes, _ := json.Marshal(res)

	fmt.Printf("\n%+v\n", newUser)
	fmt.Println(Users)
	w.Write(marshalledRes)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	var User Login

	json.NewDecoder(r.Body).Decode(&User)

	if User.UserName == "" || User.Password == "" {
		fmt.Println("All fields are required")
	} else {
		hasFound := false

		for _, v := range Users {
			if User.UserName == v.UserName && User.Password == v.Password {
				hasFound = true

				http.SetCookie(w, &http.Cookie{
					Name:  "Access",
					Value: "True",
				})

				fmt.Fprintf(w, "Found")
				fmt.Println("Found")
				break
			}
		}

		if hasFound == false {
			fmt.Println("Invalid user")
			fmt.Fprintf(w, "Invalid user")
		}
	}
}

func main() {
	fmt.Println("Server is running on: 127.0.0.1:9999")

	http.HandleFunc("/signUp", signUp)
	http.HandleFunc("/signIn", signIn)

	http.ListenAndServe(":9999", nil)
}
