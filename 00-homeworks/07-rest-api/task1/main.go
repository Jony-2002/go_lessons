package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Phone    string
	Password string
}

var usersSlice = []User{
	{
		Phone:    "+992935224004",
		Password: "jony123",
	},
	{
		Phone:    "+992111143040",
		Password: "123mus",
	},
	{
		Phone:    "+992502002002",
		Password: "user999",
	},
}

func getPhones(w http.ResponseWriter, r *http.Request) {
	for _, v := range usersSlice {
		fmt.Fprintf(w, v.Phone, "\n")
		fmt.Println(v.Phone)
	}

	fmt.Println("---------- SUCCESS ----------")
}

func getPasswords(w http.ResponseWriter, r *http.Request) {
	for _, v := range usersSlice {
		fmt.Fprintf(w, v.Password, "\n")
		fmt.Println(v.Password)
	}

	fmt.Println("---------- SUCCESS ----------")
}

func main() {
	fmt.Println("Homework 1")

	http.HandleFunc("/allUsersPhone", getPhones)
	http.HandleFunc("/allUsersPassword", getPasswords)

	http.ListenAndServe(":9999", nil)
}
