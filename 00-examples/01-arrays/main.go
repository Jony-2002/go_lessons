package main

import "fmt"

func main() {
	// Array manipulation
	arr1 := []int{1, 2, 3, 4, 5}
	for i, num := range arr1 {
		if num%2 == 0 {
			arr1[i] = num * num
		} else {
			arr1[i] = num * num * num
		}
	}
	fmt.Println("Array after manipulation:", arr1)

	// Array rotation
	arr2 := []int{1, 2, 3, 4, 5}
	rotation := 2
	arr2 = append(arr2[rotation:], arr2[:rotation]...)
	fmt.Println("Array after rotation:", arr2)

	// Array concatenation
	arr3 := []int{1, 2, 3}
	arr4 := []int{4, 5, 6}
	arr3 = append(arr3, arr4...)
	fmt.Println("Concatenated array:", arr3)

	// Array intersection
	arr5 := []int{1, 2, 3, 4, 5}
	arr6 := []int{3, 4, 5, 6, 7}
	var arr7 []int
	for _, num := range arr5 {
		for _, num2 := range arr6 {
			if num == num2 {
				arr7 = append(arr7, num)
			}
		}
	}
	fmt.Println("Intersection of two arrays:", arr7)

	// Array reversal
	arr8 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr8)/2; i++ {
		j := len(arr8) - i - 1
		arr8[i], arr8[j] = arr8[j], arr8[i]
	}
	fmt.Println("Reversed array:", arr8)
}
