package main

import (
	"fmt"

	"example.com/packages/algorithms"
	"example.com/packages/data"
	// "example.com/packages/util"
)

func main() {
	// fmt.Println("Main function")
	// fmt.Println(utils.StringLength("Hello"))
	// fmt.Println(utils.AddTwoNumbers(4, 2))
	// fmt.Println(data.GetData())
	copiedData := data.GetData()
	fmt.Println(copiedData)
	sortedData := algorithms.SortArray(copiedData)
	fmt.Println(sortedData)
}
