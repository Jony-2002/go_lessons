package main

import "fmt"

func main() {
	// 1. strings
	var userName string = "Jahongir"
	fmt.Printf("Hello, %v\n", userName)

	// 2. integers
	var num1 int = 45
	var num2 int = 26
	var result int = num1 + num2

	fmt.Printf("The result of num1(%v) and num2(%v) is: %v\n", num1, num2, result)

	// 3. Booleans
	var isStudent bool = false
	fmt.Printf("isStudent type is: %T\n", isStudent)

	// %T - check the variable type
	// %v - shows the varible value
	// NOTE: we only access these syntax inside (Printf)

	// --------------------------------------------------------
	// 4. no var style
	isLoggedIn := false;
	fmt.Printf("isLoggedIn type is: %T", isLoggedIn)
	// we can "no var style" only inside functions

	// 5. Constants
	// Constants cannot change their value...)
	const PORT int = 8800
}
