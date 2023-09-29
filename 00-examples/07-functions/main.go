package main

import (
	"fmt"
	"math"
)

// Define a function that takes two integers and returns their sum
func add(x int, y int) int {
	return x + y
}

// Define a struct for a circle and a method to calculate its area
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	// Call the add function
	sum := add(1, 2)
	fmt.Println("Sum:", sum)

	// Create a new circle with radius 2
	c := Circle{radius: 2}

	// Call the area method on the circle
	area := c.area()
	fmt.Println("Area:", area)
}

// In this program, we first define a function add that takes two integers and returns their sum. We then call this function and print out the result.

// Next, we define a struct Circle with a single field radius of type float64. We also define a method area on the Circle struct that calculates the area of the circle using the math.Pi constant and the circle's radius.

// In the main function, we create a new Circle object with a radius of 2. We then call the area method on this object and print out the result.

// Note that the area method is defined on the Circle struct using the syntax func (c Circle) area() float64. This is known as a method receiver, which allows the method to be called on an instance of the Circle struct. The Circle instance is automatically passed to the method as the c parameter.
