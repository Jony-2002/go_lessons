package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	// var userName string
	// fmt.Println("Please, provide username")
	// fmt.Scanln(&userName)
	// fmt.Printf("Hello, %v!\n", userName)
	// fmt.Printf("userName type is: %T\n", userName)

	// 1. Create array
	// 2. When we create array, we should mention the size of an araay, (in below array, this is 4) and the type of the data inside the array
	var names [4]string

	// 3. Adding element to an array
	names[0] = "Masrur"
	// 4. When we miss the specific index (in this case [1], there will be more space in the output [check the console])
	names[2] = "Jahongir"
	names[3] = "Khushnud"

	fmt.Println(names)
	// len() - method which shows the array length
	// NOTE: It will show the size of an array (which we specify while creating an array, in this case this is 4), not the actual elements quantity
}
