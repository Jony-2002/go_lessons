package main

import (
	"fmt"
)

func main() {
	// fmt.Printf("---------- TASK 1 ----------\n")
	// 1. Log the user whose name starts with "M"

	// var usersSlice = []string{"Jahongir", "Masrur"}
	// filteredUsers := []string{}

	// usersSlice = append(usersSlice, "Khushnud", "Mustafo")
	// fmt.Println(usersSlice)

	// for i := 0; i < len(usersSlice); i++ {
	// 	name := usersSlice[i]

	// 	chars := strings.Split(name, "")

	// 	if strings.ToLower(chars[0]) == "m" {
	// 		newName := strings.Join(chars, "")
	// 		filteredUsers = append(filteredUsers, newName)
	// 		fmt.Printf("%v starts with - %v\n", newName, chars[0])
	// 	}
	// }
	// fmt.Println(filteredUsers)

	fmt.Printf("---------- TASK 2 ----------\n")
	// 2. Add new element to slice
	carsSlice := []string{"BMW"}
	userChoice := false
	userTyping := ""

	fmt.Println(carsSlice)

	fmt.Println("Do you want to add new car?")
	fmt.Println("y - yes")
	fmt.Println("n - no")
	fmt.Scanln(&userTyping)

	result := checkUserAnswer(userTyping, userChoice)

	if result == true {
		createNewCar(carsSlice, userChoice, userTyping)
	} else {
		fmt.Println(carsSlice)
	}

	// fmt.Printf("---------- TASK 3 ----------\n")
	// fmt.Printf("---------- TASK 4 ----------\n")
	// fmt.Printf("---------- TASK 5 ----------\n")

}

func createNewCar(carsSlice []string, userChoice bool, userTyping string) {
	for i := 0; i < 1; i++ {
		var newCar string
		fmt.Println("Please, provide new car: ")
		fmt.Scanln(&newCar)
		carsSlice = append(carsSlice, newCar)
		fmt.Println(carsSlice)
		fmt.Println("Successefully added new car!")
	}

	fmt.Println("Do you want to add new car?")
	fmt.Println("y - yes")
	fmt.Println("n - no")
	fmt.Scanln(&userTyping)

	result := checkUserAnswer(userTyping, userChoice)

	if result == true {
		createNewCar(carsSlice, userChoice, userTyping)
	} else {
		fmt.Println(carsSlice)
	}
}

func checkUserAnswer(userTyping string, userChoice bool) bool {
	switch userTyping {
	case "y":
		userChoice = true
		return true
	case "n":
		userChoice = false
		return false
	default:
		userChoice = false
		return false
	}
}
