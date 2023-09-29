package main

import "fmt"

// How to get all the keys in MAP

func dqd() {
	task1()
}

func task1() {
	var users = make(map[int]string)
	fmt.Println("names\n", users)

	// add some default values
	users[0] = "Jahongir"
	users[1] = "Masrur"
	users[2] = "Khushnud"
	users[3] = "Mustafo"
	fmt.Println("names\n", users)

	var newKey int
	var newUser string

	fmt.Println("\nPlease, provide key: ")
	fmt.Scanln(&newKey)
	fmt.Println("\nPlease, provide user: ")
	fmt.Scanln(&newUser)

	users[newKey] = newUser
	fmt.Println("names\n", users)

	// delete user
	fmt.Println("\nPlease, provide key for deleting: ")
	fmt.Scanln(&newKey)
	delete(users, newKey)
	fmt.Println("names\n", users)
}
