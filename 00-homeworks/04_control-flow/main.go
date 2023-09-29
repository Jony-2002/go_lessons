package main

import "fmt"

type TUser struct {
	Age int
}

func main() {
	fmt.Println("Control flow task")

	var users map[string]TUser = map[string]TUser{
		"Jon":   {Age: 13},
		"Sam":   {Age: 33},
		"Sara":  {Age: 28},
		"Conor": {Age: 12},
	}

	handleProvideChoice(users)

	// user, isFound := users[input]

	// if isFound == true {
	// 	if user.Age < 18 {
	// 		fmt.Printf("\n You (%v) are not able to vote\n", input)
	// 	} else {
	// 		fmt.Printf("\n Okay, you (%v) can easily vote\n", input)
	// 	}
	// }
}

func handleProvideChoice(usersMap map[string]TUser) {
	var input string
	println("\nProvide username...\n")
	fmt.Scanln(&input)

	handleCheckTheChoice(input, usersMap)
	// return input
}

func handleCheckTheChoice(input string, usersMap map[string]TUser) {
	user, isFound := usersMap[input]

	if isFound == true {
		if user.Age < 18 {
			fmt.Printf("\n You (%v) are not able to vote\n", input)
		} else {
			fmt.Printf("\n Okay, you (%v) can easily vote\n", input)
		}
	}

	handleProvideChoice(usersMap)
}
