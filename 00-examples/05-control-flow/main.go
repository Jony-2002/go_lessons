package main

import (
	"fmt"
	"math"
)

func main() {
	// Statements
	x := 10
	y := 20
	z := x + y
	fmt.Println("The sum of x and y is", z)

	// Switch statement
	grade := 85
	switch {
	case grade >= 90:
		fmt.Println("Grade is A")
	case grade >= 80:
		fmt.Println("Grade is B")
	case grade >= 70:
		fmt.Println("Grade is C")
	case grade >= 60:
		fmt.Println("Grade is D")
	default:
		fmt.Println("Grade is F")
	}

	// Error handling
	num := -10.0
	if num < 0 {
		err := fmt.Errorf("can't calculate square root of a negative number: %f", num)
		fmt.Println(err)
	} else {
		result := math.Sqrt(num)
		fmt.Printf("The square root of %f is %f\n", num, result)
	}
}
