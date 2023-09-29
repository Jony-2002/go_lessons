package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name   string
	Gender string
}

var userSlice []User

func addUser(w http.ResponseWriter, r *http.Request) {
	var user = User{
		Name: "Sam",
		Gender: "male",
	}

	userSlice = append(userSlice, user)

	fmt.Fprintf(w, "Success")

	fmt.Println(userSlice)
}

func main() {
	fmt.Println("Hello world!")

	http.HandleFunc("/addSam", addUser)

	http.ListenAndServe(":8080", nil)
}
