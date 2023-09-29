package main

import (
	// "bufio"
	"fmt"
	// "os"
)

func main() {
	fmt.Printf("Hello, world! \n")

	// // 1
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter the course rating")

	// // comma ok syntax - try catch alternative
	// // result is try - the second argument will be the error, (if we do not want to catch the error, just write "_" as below)
	// result, _ := reader.ReadString('\n')
	// fmt.Printf("thanks for rating, %v\n", result)
	// fmt.Printf("result type is: %T", result)

	// 2
	var x string
	print("Give x number: ")
	fmt.Scanln(&x)
	fmt.Println("x - value is: ", x)
	fmt.Printf("x type is: %T", x)
}
