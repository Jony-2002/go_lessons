// cannot find with space in name (Element 1)

package main

import (
	"fmt"
	"strings"
)

type RandomStruct struct {
	Name     string
	Price    int
	Quantity int
}

var randomSlice []RandomStruct

var defaultElement1 = RandomStruct{Name: "Element1", Price: 200, Quantity: 5}
var defaultElement2 = RandomStruct{Name: "Element2", Price: 400, Quantity: 10}
var defaultElement3 = RandomStruct{Name: "Element3", Price: 600, Quantity: 3}

func main() {
	// homeWorkLevel1()
	homeWorkLevel2()
}

func homeWorkLevel1() {
	printYourName()
	addTwoNumbers()
	subtractTwoNumbers()
	printName("Masrur")
	printMinus(2)
}

func printYourName() {
	fmt.Println("Jahongir")
}

func addTwoNumbers() {
	fmt.Println(1 + 1)
}

func subtractTwoNumbers() {
	fmt.Println(1 - 1)
}

func printName(name string) {
	fmt.Println(name)
}

func printMinus(number int) {
	fmt.Println(number - 1)
}

func homeWorkLevel2() {
	randomSlice = append(randomSlice, defaultElement1, defaultElement2, defaultElement3)
	fmt.Println(randomSlice)
	
	provideChoice()
}

func provideChoice() {
	fmt.Printf("\nPlease, choose command in order to continue...\n")

	var userInput string
	fmt.Println("\n1 - Add")
	fmt.Println("2 - Find")
	fmt.Println("3 - Delete")
	fmt.Println("4 - Edit")

	fmt.Scanln(&userInput)
	checkProvidedChoice(userInput)
}

func checkProvidedChoice(input string) {
	switch input {
	case "1":
		addNewElement()
	case "2":
		findElement()
	case "3":
		deleteElement()
	case "4":
		editElement()
	default:
		fmt.Println("unknown command")
		provideChoice()
	}
}

func addNewElement() {
	var name string
	var price int
	var quantity int

	fmt.Println("Please provide Name...")
	fmt.Scanln(&name)
	fmt.Println("Please provide Price...")
	fmt.Scanln(&price)
	fmt.Println("Please provide Quantity...")
	fmt.Scanln(&quantity)

	var newElement = RandomStruct{Name: name, Price: price, Quantity: quantity}
	fmt.Printf("\n%+v\n", newElement)

	randomSlice = append(randomSlice, newElement)
	fmt.Println(randomSlice)

	fmt.Println("success")
	provideChoice()
}

func findElement() {
	fmt.Println("Enter the element name you want to search for")

	var name string
	fmt.Scanln(&name)

	hasFound := false

	for _, v := range randomSlice {
		if strings.ToLower(name) == strings.ToLower(v.Name) {
			fmt.Println("found")
			fmt.Println(v)
			hasFound = true
			provideChoice()
			break
		} else {
			hasFound = false
		}
	}

	if hasFound != true {
		fmt.Println("not found")
		provideChoice()
	}
}

func deleteElement() {
	fmt.Println("Please, provide name for deleting...")

	var input string
	fmt.Scanln(&input)

	hasFound := false

	for i := 0; i < len(randomSlice); i++ {
		if strings.ToLower(input) == strings.ToLower(randomSlice[i].Name) {
			fmt.Println("found")
			randomSlice = append(randomSlice[:i], randomSlice[i+1:]...)
			fmt.Println(randomSlice)
			hasFound = true
			provideChoice()
			break
		} else {
			hasFound = false
		}
	}

	if hasFound == false {
		fmt.Println("not found")
		provideChoice()
	}
}

func editElement() {
	fmt.Println("Please, provide Name for editing...")

	var input string
	fmt.Scanln(&input)

	hasFound := false

	for index, v := range randomSlice {
		if strings.ToLower(input) == strings.ToLower(v.Name) {
			fmt.Println("found")
			var newName string
			fmt.Scanln(&newName)

			v.Name = newName
			fmt.Println(v)

			randomSlice = append(randomSlice[:index], randomSlice[index+1:]...)
			randomSlice = append(randomSlice, v)

			fmt.Println(randomSlice)
			provideChoice()

			hasFound = true
			break
		} else {
			hasFound = false
		}
	}

	if hasFound == false {
		fmt.Println("not found")
		provideChoice()
	}
}
