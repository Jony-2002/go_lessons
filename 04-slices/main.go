package main

import "fmt"

func main() {
	fmt.Println("This is Slices topic!!!")

	// 1. Create slice
	// NOTE! We don't mention the size of a slice (which we did with arrays)
	var cars = []string{"BMW", "Mercedes", "Porche"}

	fmt.Printf("cars: %v\n", cars)
	fmt.Printf("cars type is: %T\n", cars)

	// 2. Add new element to slice
	cars = append(cars, "Tangem", "Zil")
	fmt.Printf("cars: %v\n", cars)

	// 3. Slice the slice
	cars = append(cars[1:])
	fmt.Printf("cars: %v\n", cars)

	cars = append(cars[1:3])
	fmt.Printf("cars: %v\n", cars)

	// Practice 1
	names := []string{}
	fmt.Printf("names: %v\n", names)

	names = append(names, "Jahongir", "Masrur", "Khushnud")
	fmt.Printf("names: %v\n", names)
	
	var newUser string
	fmt.Println("Provide new user:")
	fmt.Scanln(&newUser)
	names = append(names, newUser)
	fmt.Printf("new results: %v\n", names)
	fmt.Printf("length: %v\n", len(names))
	fmt.Printf("first three results: %v, and the length is: %v\n", names[:3], len(names[:3]))
	fmt.Printf("last three results: %v\n", names[1:])

	// Questions
	// 1. How to add elements into slice at the beginning?
	// 2. How to remove elements in slice from ending
}
