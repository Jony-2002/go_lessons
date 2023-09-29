package main

import "fmt"

func main() {
	// Declare and initialize a map
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Cathy": 35,
	}
	fmt.Println("Ages:", ages)

	// Add a new key-value pair to the map
	ages["Dave"] = 40
	fmt.Println("Ages after adding Dave:", ages)

	// Retrieve a value from the map
	aliceAge := ages["Alice"]
	fmt.Println("Alice's age:", aliceAge)

	// Check if a key exists in the map
	_, isPresent := ages["Bob"]
	fmt.Println("Is Bob present?", isPresent)

	// Delete a key-value pair from the map
	delete(ages, "Cathy")
	fmt.Println("Ages after deleting Cathy:", ages)

	// Iterate over the map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
