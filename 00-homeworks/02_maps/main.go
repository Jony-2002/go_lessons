package main

import "fmt"

func main() {
	// Домашнее задание:
	// 1. создайте Мэп с именем froots и с типом [string]int добавьте 3 элемента в неё ([apple:23] это будет продукт и его цена)
	// 2. затем спросите у пользователя имя продукта (с помощью scanln) и используйте поиск (if stetment), чтобы найти его.
	// 3. если он существует (если isExist == true), напечатайте «стоимость вашего продукта составляет 19 $» 19 это его цена из Мэпа
	// 4. затем спросите пользователя, хочет ли он удалить какой-либо продукт (через scanln), и удалите его.

	// 1. Create map and add 3 fruits with name and price
	fruits := make(map[string]int)

	fruits["apple"] = 23
	fruits["banana"] = 45
	fruits["mango"] = 68
	fmt.Println(fruits)

	// 2
	var input string
	fmt.Println("Provide fruit name in order to show the price:")
	fmt.Scanln(&input)

	// 3

	// Check if a key exists in the map
	_, isExist := fruits[input]

	if isExist == true {
		fruitPrice := fruits[input]
		fmt.Printf("«стоимость вашего продукта (%v) составляет - %v $»\n", input, fruitPrice)
	} else {
		fmt.Printf("«не найдено»\n")
	}

	var deleteMessage string
	var deleteInput string

	fmt.Println("Do you want to delete any product?")
	fmt.Println("y - yes")
	fmt.Println("n - no")
	fmt.Scanln(&deleteMessage)

	if deleteMessage == "y" {
		fmt.Println("Provide the fruit name which you want to delete...")
		fmt.Scanln(&deleteInput)

		_, isExist := fruits[deleteInput]

		if isExist == true {
			delete(fruits, deleteInput)
			fmt.Printf("%v was successfully deleted!\n", deleteInput)
			fmt.Println(fruits)
		}
	}
}
