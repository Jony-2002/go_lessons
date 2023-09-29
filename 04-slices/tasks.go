package main

import (
	"fmt"
	"sort"
)

// FIX this error!
func main() {
	fmt.Println("Tasks file is runnig...")

	// 1
	names := []string{"Masrur", "Jahongir", "Shuhrat", "Khushnud", "Mustafo", "Farida"}
	fmt.Println("names", names)

	// delete the element based on index
	// in this case, we delete "Khushnud" from the slice
	slicedNames := append(names[:3], names[3+1:]...)
	fmt.Println("slicedNames", slicedNames)

	// 2
	nums := []int{1, 5, 2, 34, 3, 9, 5, 10, 25}
	filteredNums := nums
	sort.Ints(filteredNums)
	fmt.Printf("---------------BEFORE-------------------\n")
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("filteredNums: %v\n", filteredNums)
	fmt.Printf("---------------AFTER-------------------\n")

	filteredNums = append(filteredNums, 44)

	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("filteredNums: %v\n", filteredNums)

	nums = append(nums, 111)

	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("filteredNums: %v\n", filteredNums)


	
	// FAQs
	// 1. How can we find the element index?
	// 2. How to copy an slice without mutating it?
	// 3. Why I filter the copy slice and it affects the main one
	// 4. Why it does not affect the main slice when I add new element to the copy one
}
