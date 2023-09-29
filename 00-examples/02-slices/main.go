package main

import "fmt"

func main() {

	var numberss [3]int = [3]int{1, 2, 3}           // array
	var names [2]string = [2]string{"Alice", "Bob"} // array

	var numbersSlice []int = []int{1, 2, 3}            // slice
	var namesSlice []string = []string{"Alice", "Bob"} // slice

	fmt.Println(numberss, names)
	fmt.Println(numbersSlice, namesSlice)

	var numbers [5]int = [5]int{10, 20, 30, 40, 50}

	// Sum of array elements
	var sum int = 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Printf("Sum of array elements: %d\n", sum)

	// Average of array elements
	var average float64 = float64(sum) / float64(len(numbers))
	fmt.Printf("Average of array elements: %.2f\n", average)

	// Slice of first three elements of the array
	var firstThree []int = numbers[:3]
	fmt.Printf("First three elements of the array: %v\n", firstThree)

	// Append an element to the slice
	firstThree = append(firstThree, 60)
	fmt.Printf("After appending an element to the slice: %v\n", firstThree)

	// Length and capacity of the slice
	fmt.Printf("Length of the slice: %d\n", len(firstThree))
	fmt.Printf("Capacity of the slice: %d\n", cap(firstThree))

	// Iterate over the slice
	for i, value := range firstThree {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}

}
