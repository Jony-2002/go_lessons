package main

import "fmt"

func main() {
	fmt.Println("Taks 1")
	// • Напишите программу, которая добавляет элемент в слайс.
	// • затем  выводит первый элемент из этого слайса.
	// • затем которая удаляет первый элемент из слайса.
	// • затем выводит в принт все элементы слайса.

	// names := []string{"Masrur", "Jahongir", "Khushnud"}
	// fmt.Println(names)

	// // 1. Add new element to slice
	// names = append(names, "Shuhrat", "Muyassar", "Mustafo")
	// fmt.Println(names)

	// var userInput string
	// fmt.Println("Please, write username in order to add to slice...")
	// fmt.Scanln(&userInput)
	// names = append(names, userInput)
	// fmt.Println(names)

	// // 2. Delete element (Masrur) from slice
	// names = append(names[1:])
	// fmt.Println(names)

	// // 2.1 Delete element based on its index (Shuhrat)
	// names = append(names[:2], names[2+1:]...)
	// fmt.Println(names)

	fmt.Println("Taks 2")
	// // [“Dushanbe”,”Moskow”,”Parice”,”NewYourk”,”Kabul”]
	// // 1.1 Удалите 2-й элемент из этого массива.
	// // 1.2 Удалите 4-й элемент из этого массива.
	// // 1.3 Удалите 3-й элемент из этого массива.
	// // 1.4 Удалите 5-й элемент из этого массива.

	// var cities = [5]string{"Dushanbe", "Moskow", "Parice", "NewYourk", "Kabul"}
	// fmt.Println(cities)

	fmt.Println("Taks 3")
	// // Создайте slice c типoм string и заполните его данными о овощах - [apple,potato]
	// // Добавьте несколько 5 новых элементов в массив.
	// // Выведите 2 и 3 элемент из массива.
	// // Удалите 1 и 3 элементы из массива.

	// var products = []string{"apple", "potato"}
	// fmt.Println(products)
	
	// // add new 5 element
	// products = append(products, "grapes", "banana", "mango", "strawberies", "watermaloon")
	// fmt.Println(products)
	
	// // log the 2 and 3 element in console
	// fmt.Println(products[1:3])
	
	// // delete 1st and 3rd element from the slice
	// products = append(products[1:])
	// products = append(products[:1], products[1+1:]...)
	// fmt.Println(products)
}
