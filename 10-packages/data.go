package main

import "fmt"

type User struct {
	Name string
	Age  int
}

var Users = []User{
	{
		Name: "Jahongir",
		Age:  22,
	},
	{
		Name: "Masrur",
		Age:  21,
	},
	{
		Name: "Khayom",
		Age:  19,
	},
}

func data() {
	fmt.Println("Data function")
}
