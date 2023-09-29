package main

import (
	"fmt"
	"strings"
)

func main() {
	// theory()
	// task1()
	// task2()
	// buyIceCreamBaseOnIndex()
	// buyIceCreamBaseOnValue()

	// splitFunction()
	// clearText()

	reverseText()
}

func theory() {
	var elements = []string{"Car", "House", "Dog", "Cap", "Cat"}
	// 1. Log all elements in terminal

	// with for/range
	fmt.Printf("\nfor/range\n\n")
	for index, el := range elements {
		fmt.Printf("%v - %v\n", index, el)
	}

	// with for
	fmt.Printf("\nfor\n\n")
	for i := 0; i < len(elements); i++ {
		fmt.Printf("%v - %v\n", i, elements[i])
	}
}

func task1() {
	var nums = []int{4, 1, 2, 3, 6}
	var filteredNums = []int{}

	for index, value := range nums {
		if index == 0 {
			filteredNums = append(filteredNums, value)
			continue
		}

		filteredNums = append(filteredNums, value)
	}

	fmt.Println(filteredNums)
}

func task2() {
	var prices []int = []int{3, 2, 4, 5}
	var coins int = 4

	for _, v := range prices {
		if coins >= v {
			fmt.Println(v)
		}
	}
}

func buyIceCreamBaseOnIndex() {
	var prices []int = []int{3, 2, 4, 5}
	var coins int = 4

	var input int
	fmt.Println(prices)
	fmt.Printf("\nYou have %v coins\n", coins)
	fmt.Println("Which ice-cream do you want to buy?")
	fmt.Scanln(&input)

	input = input - 1
	for i := 0; i < len(prices); i++ {

		if input < len(prices) {
			fmt.Println("Please, try another number...")

			if coins >= prices[input] {
				fmt.Printf("Yes, you can buy this (%v) ice-cream\n", prices[input])
				break
			} else {
				fmt.Printf("No, you can not buy this (%v) ice-cream\n", prices[input])
			}
		}
	}
}

func buyIceCreamBaseOnValue() {
	var prices []int = []int{3, 2, 4, 5}
	var coins int = 4

	var input int
	fmt.Println(prices)
	fmt.Printf("\nYou have %v coins\n", coins)
	fmt.Println("Which ice-cream do you want to buy?")
	fmt.Scanln(&input)

	hasFound := false

	for _, value := range prices {
		if value == input {
			fmt.Printf("Found the ice-cream - %v\n", input)

			if coins >= value {
				fmt.Printf("Yes, you can buy this - (%v) ice-cream", value)
				break
			} else {
				fmt.Printf("No, you can not buy this - (%v) ice-cream", value)
			}
			hasFound = true
			break
		}
	}

	if hasFound == false {
		fmt.Printf("Can not find the ice-cream - %v\n", input)
	}
}

func splitFunction() {
	email := "masrur@gmail.com"

	splitedEmail := strings.Split(email, "@")

	fmt.Println(email)
	fmt.Println(splitedEmail)
	// fmt.Println(splitedEmail[0])
	// fmt.Println(splitedEmail[1])

	userEmail := "abc@gmail.com"
	emailSplitted := strings.Split(userEmail, "")
	// fmt.Println(emailSplitted)

	var chars string

	for i := len(emailSplitted); i > 0; i-- {
		if len(emailSplitted) >= 13 {
			chars += emailSplitted[i-1]
		}
	}

	fmt.Println(chars)

	password := "mas12"
	splittedPassword := strings.Split(password, "")
	joinedText := strings.Join(splittedPassword, "")

	fmt.Println(splittedPassword)
	fmt.Println(joinedText)

	if len(splittedPassword) > 6 {
		fmt.Println("correct")
	} else {
		fmt.Println("wrong")
	}
}

func clearText() {
	someText := "Hel#lo World!"
	splittedText1 := strings.Split(someText, "#")
	joinedText1 := strings.Join(splittedText1, "")
	fmt.Println(splittedText1)
	fmt.Println(joinedText1)
}

func reverseText() {
	var text string = "masrur"
	var reversedText []string

	splittedText := strings.Split(text, "")
	// fmt.Printf("splitted name - %v\n", splittedText)
	fmt.Printf("%v\n", splittedText)

	for i := len(splittedText); i > 0; i-- {
		reversedText = append(reversedText, splittedText[i-1])
	}

	finalResult := strings.Join(reversedText, "")

	// fmt.Printf("final result - %v\n", finalResult)
	fmt.Printf("%v\n", finalResult)
}
