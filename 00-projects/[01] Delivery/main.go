package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type User struct {
	UserName string
	Password string
	Address  string
	Phone    string
	Id       int
}

type Food struct {
	Name        string
	Price       int
	IsAvailable bool
	Quantity    int
	Place       string
	Tags        []string
}

type Place struct {
	Name    string
	Address string
	Foods   []Food
}

type Cart struct {
	Name        string
	Price       int
	IsAvailable bool
	Quantity    int
	Place       string
	ItemId      int
}

var Users = []User{
	{
		UserName: "Jahongir Niyozbekov",
		Password: "jony2002",
		Address:  "Gulaken",
		Phone:    "+992935224004",
		Id:       rand.Intn(1000000000),
	},
	{
		UserName: "Masrur Qurbonbekov",
		Password: "mas2003",
		Address:  "Bar",
		Phone:    "+992934236835",
		Id:       rand.Intn(1000000000),
	},
}

var Foods = []Food{
	{
		Name:        "Pizza",
		Price:       65,
		IsAvailable: true,
		Quantity:    10,
		Place:       "World Cousin",
		Tags:        []string{"Fast Food"},
	},
	{
		Name:        "Burger",
		Price:       35,
		IsAvailable: true,
		Quantity:    23,
		Place:       "Surayo",
		Tags:        []string{"Fast Food"},
	},
}

var Carts = []Cart{
	{
		Name:        "Burger",
		Price:       35,
		IsAvailable: true,
		Quantity:    23,
		Place:       "Surayo",
		ItemId:      123,
	},
}

func signUp(w http.ResponseWriter, r *http.Request) {
	// 127.0.0.1:1111/signUp
	fmt.Println("SignUp Logic will be here")

	var newUser User

	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.Id = rand.Intn(1000000000)

	fmt.Println(newUser)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	// 127.0.0.1:1111/signIn
	fmt.Println("SignIn Logic will be here")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// 127.0.0.1:1111/getUsers
	fmt.Println("getUsers Logic will be here")
}

func addToCard(w http.ResponseWriter, r *http.Request) {
	// 127.0.0.1:1111/addToCard
	// id := r.URL.Query().Get("id")

	fmt.Println("Add To Card Logic will be here")

	var newItem Cart

	json.NewDecoder(r.Body).Decode(&newItem)
	Carts = append(Carts, newItem)

	fmt.Println(Carts)
	fmt.Println(newItem)

	marshalledSlice, _ := json.Marshal(Carts)

	w.Write(marshalledSlice)
}

func removeFromCart(w http.ResponseWriter, r *http.Request) {
	// 127.0.0.1:1111/addToCard
	// id := r.URL.Query().Get("id")
}

func addFood(w http.ResponseWriter, r *http.Request) {
	// 127.0.0.1:1111/addFood
	fmt.Println("addFood Logic will be here")

	var newFood Food

	json.NewDecoder(r.Body).Decode(&newFood)

	Foods = append(Foods, newFood)

	fmt.Println(newFood)
	fmt.Println(Foods)

	marshalledFood, _ := json.Marshal(newFood)
	marshalledFoods, _ := json.Marshal(Foods)

	w.Write(marshalledFood)
	w.Write(marshalledFoods)
}

func getFoods(w http.ResponseWriter, r *http.Request) {
	// 127.0.0.1:1111/getFoods

	fmt.Println(Foods)

	marshalledFoods, _ := json.Marshal(Foods)
	w.Write(marshalledFoods)
}

func main() {
	fmt.Println("Server is running on: 127.0.0.1:1111")

	// fmt.Printf("%+v\n", Users)
	// fmt.Printf("\n%+v\n", Carts)

	http.HandleFunc("/signUp", signUp)
	http.HandleFunc("/signIn", signIn)
	http.HandleFunc("/addToCard", addToCard)
	http.HandleFunc("/addFood", addFood)
	http.HandleFunc("/getFoods", getFoods)

	http.ListenAndServe(":1111", nil)
}
