package main

import "fmt"

func main() {
	fmt.Println("Bubble sorting")

	var nums []int = []int{54, 23, 11, 65, 76, 54, 34, 44, 22}
	fmt.Println(nums)

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				var temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}

	fmt.Println(nums)
}
