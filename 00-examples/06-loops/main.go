package main

import "fmt"

func main() {
	// Array example
	nums := [5]int{1, 2, 3, 4, 5}

	// Looping over an array using a for loop
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// Looping over an array using the range keyword
	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// Modifying an array using a for loop
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * 2
	}
	fmt.Println(nums)

	// Modifying an array using the range keyword
	for i, num := range nums {
		nums[i] = num * 3
	}
	fmt.Println(nums)
}
