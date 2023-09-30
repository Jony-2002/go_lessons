package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type User struct {
	UserName string
	Email    string
	Password string
	Role     string
	Id       int
}

type Food struct {
	Name  string
	Price int
}

type Login struct {
	Login    string
	Password string
}

var Foods = []Food{
	{
		Name:  "Burger",
		Price: 29,
	},
	{
		Name:  "Pizza",
		Price: 69,
	},
	{
		Name:  "Hot-Dog",
		Price: 18,
	},
}

var Users = []User{
	{
		UserName: "Jahongir",
		Email:    "thisisjahongir@gmail.com",
		Password: "jony123",
		Role:     "User",
		Id:       rand.Intn(1000000000),
	},
	{
		UserName: "admin",
		Email:    "admin@gmail.com",
		Password: "admin",
		Role:     "Admin",
		Id:       rand.Intn(1000000000),
	},
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var newUser User

	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.Id = rand.Intn(1000000000)
	newUser.Role = "User"
	Users = append(Users, newUser)

	marshalledUser, _ := json.Marshal(newUser)

	w.Write(marshalledUser)
	fmt.Println(newUser)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	var user Login

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)

	if user.Login == "" || user.Password == "" {
		fmt.Fprintf(w, "Fill all the inputs")
	} else {
		hasFound := false

		for _, v := range Users {
			if user.Login == v.UserName && user.Password == v.Password {
				http.SetCookie(w, &http.Cookie{
					Name:  "test",
					Value: "accept",
				})

				fmt.Fprintf(w, "Success")
				fmt.Printf("%v Just logged in who is - %v\n", v.UserName, v.Role)
				hasFound = true
				break
			}
		}

		if hasFound == false {
			fmt.Fprintf(w, "Invalid user! Try again...")
		}
	}
}

func getMenu(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	fmt.Println("/getMenu")

	marshalledData, _ := json.Marshal(Foods)

	cookie, err := r.Cookie("test")

	if err != nil {
		fmt.Println("Invalid cookie name")
	} else if cookie.Value == "accept" {
		fmt.Println(Foods)
		w.Write(marshalledData)
	} else {
		fmt.Fprintf(w, "You don't have access to menu")
	}

}

func main() {
	fmt.Println("Server is running on: 127.0.0.1:4455")

	http.HandleFunc("/signUp", signUp)
	http.HandleFunc("/signIn", signIn)
	http.HandleFunc("/getMenu", getMenu)

	http.ListenAndServe(":4455", nil)
}
