package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name      string
	Age       int
	Gender    string
	IsMarried bool
}

var Users = []User{}

func Append(w http.ResponseWriter, r *http.Request) {

	var newUser User

	json.NewDecoder(r.Body).Decode(&newUser)

	if newUser.Age < 18 {
		fmt.Println("User should be older than 18")
	} else {
		Users = append(Users, newUser)

		fmt.Println("")
		fmt.Printf("%+v\n", newUser)
		fmt.Println("")
		fmt.Println(Users)
		fmt.Println("-----------------------------------------------")

		fmt.Fprintf(w, "Added new user\n")
		fmt.Fprintf(w, "%+v\n", newUser)
	}
}

func main() {
	http.HandleFunc("/appendData", Append)

	fmt.Println("server is listening on: 127.0.0.1:5555")

	http.ListenAndServe(":5555", nil)
}
