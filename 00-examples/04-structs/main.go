package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var people []Person
	p1 := Person{Name: "Alice", Age: 25}
	p2 := Person{Name: "Bob", Age: 30}
	p3 := Person{Name: "Charlie", Age: 35}
	people = append(people, p1, p2, p3)
	for _, person := range people {
		if person.Age < 30 {
			fmt.Printf("%s is under 30 years old\n", person.Name)
		} else {
			fmt.Printf("%s is 30 years old or older\n", person.Name)
		}
	}
}
