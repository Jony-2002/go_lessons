package main

import "fmt"

type User struct {
	Username  string
	Password  string
	Age       int
	IsMarried bool
}

func main() {
	// theory()
	task2()
}

func theory() {
	fmt.Println("Maps lessons!")

	languages := make(map[string]string)
	fmt.Println("languages", languages)
	fmt.Printf("languages type is: %T\n", languages)

	// add new element
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["CPP"] = "C++"

	fmt.Println("languages", languages)

	// get value by key
	fmt.Println("JS is stands for:", languages["JS"])

	// delete value by key
	delete(languages, "CPP")
	fmt.Println("languages", languages)

	// loop through the map
	for key, value := range languages {
		fmt.Printf("\nkey is - %v, and value is - %v\n", key, value)
	}
}

func task2() {
	// 1. Create map
	var usersMap = map[int]User{
		0: {
			Username:  "Jahongir",
			Password:  "12345",
			Age:       20,
			IsMarried: false,
		},
	}

	// 2. Append to MAP
	usersMap[1] = User{Username: "Masrur", Password: "54321", Age: 19, IsMarried: true}

	handleProvideChoice(usersMap)
}

func handleProvideChoice(usersMap map[int]User) string {
	var input string

	fmt.Printf("\nPlease choose command in order to start...\n")
	fmt.Println("\npress '1' - add new user")
	fmt.Println("press '2' - delete user")
	fmt.Println("press '3' - find more details about the user")
	fmt.Println("press '4' - show all users")
	fmt.Println("press '5' - exit")

	fmt.Println("")
	fmt.Scanln(&input)

	handleCheckTheChoice(input, usersMap)
	return input
}

func handleCheckTheChoice(choice string, usersMap map[int]User) {
	switch choice {
	case "1":
		handleAddNewUser(usersMap)
	case "2":
		handleDeleteUser(usersMap)
	case "3":
		handleFindUser(usersMap)
	case "4":
		handleShowAllUsers(usersMap)
	case "5":
		handleExit()
	default:
		fmt.Println("default")
	}
}

func handleAddNewUser(usersMap map[int]User) {
	var userNameInput string
	var passwordInput string
	var ageInput int
	var isMarriedInput bool

	fmt.Println("Please provide username:")
	fmt.Scanln(&userNameInput)
	fmt.Println("Please provide your password:")
	fmt.Scanln(&passwordInput)
	fmt.Println("Please provide user age:")
	fmt.Scanln(&ageInput)
	fmt.Println("Please provide if user is married:")
	fmt.Scanln(&isMarriedInput)

	usersMap[2] = User{Username: userNameInput, Password: passwordInput, Age: ageInput, IsMarried: isMarriedInput}
	fmt.Printf("\n%v\n", usersMap)

	handleProvideChoice(usersMap)
}

func handleDeleteUser(usersMap map[int]User) {
	var input int

	fmt.Println(usersMap)
	fmt.Printf("\nPlease, provide the key in order to delete a user...\n")
	fmt.Scanln(&input)

	user, isFound := usersMap[input]

	if isFound == true {
		delete(usersMap, input)
		fmt.Printf("\n%v - successfully deleted...\n", user.Username)
		fmt.Println(usersMap)
	} else {
		fmt.Printf("Can not find user with this key (%v)", input)
	}

	handleProvideChoice(usersMap)
}

func handleFindUser(usersMap map[int]User) {
	var input int

	fmt.Printf("\nProvide key in order to find more details\n")
	fmt.Scanln(&input)

	user, isFound := usersMap[input]

	if isFound == true {
		fmt.Printf("\n%+v\n", user)
	} else {
		fmt.Printf("\nCan not find user with this key (%v)\n", input)
	}

	handleProvideChoice(usersMap)
}

func handleShowAllUsers(usersMap map[int]User) {
	fmt.Printf("\n%v\n", usersMap)

	handleProvideChoice(usersMap)
}

func handleExit() {
	fmt.Println("exiting...")
}
