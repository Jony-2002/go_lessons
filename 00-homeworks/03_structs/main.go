package main

import (
	"fmt"
	"strconv"
	"strings"
)

type DriverInterface struct {
	Car_Number     int64
	Driver_Name    string
	Package_Name   string
	Package_Weight int
}

type Employee struct {
	Name     string
	Salary   int
	Position string
}

// How to convert string to int (not int64)

func main() {
	// classWork()
	homework()
}

func classWork() {
	var DriversInfo []DriverInterface

	element1 := DriverInterface{
		Car_Number:     4004,
		Driver_Name:    "Jahongir",
		Package_Name:   "Oneplus 10 PRO",
		Package_Weight: 1000,
	}
	element2 := DriverInterface{
		Car_Number:     5274,
		Driver_Name:    "Masrur",
		Package_Name:   "Redmi 9",
		Package_Weight: 3212,
	}
	element3 := DriverInterface{
		Car_Number:     3234,
		Driver_Name:    "Khushnud",
		Package_Name:   "Samsung",
		Package_Weight: 2313,
	}

	DriversInfo = append(DriversInfo, element1, element2, element3)
	fmt.Printf("%+v\n", DriversInfo)

	var userInput string
	fmt.Println("Please provide car number:")
	fmt.Scanln(&userInput)

	intInput, err := strconv.ParseInt(userInput, 0, 0)
	notFound := false

	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(DriversInfo); i++ {
			if intInput == DriversInfo[i].Car_Number {
				fmt.Printf("%+v\n", DriversInfo[i])
				notFound = false
				break
			} else {
				notFound = true
				// i++
				// continue
			}
		}
		if notFound == true {
			fmt.Println("Водитель не найден")
		}
	}
}

func homework() {
	var employees []Employee = []Employee{}

	emp1 := Employee{Name: "Jahongir", Salary: 1500, Position: "Front-End"}
	emp2 := Employee{Name: "Masrur", Salary: 3321, Position: "Back-End"}
	emp3 := Employee{Name: "Hisom", Salary: 5500, Position: "UX/UI"}

	employees = append(employees, emp1, emp2, emp3)
	fmt.Printf("%v\n", employees)

	provideChoice(employees)
}

func checkUserChoice() string {
	fmt.Println("\npress '1' - Add new Employee")
	fmt.Println("press '2' - Find Employee")
	fmt.Println("press '3' - Exit")

	var userChoice string
	fmt.Scanln(&userChoice)

	switch userChoice {
	case "1":
		return "add"
	case "2":
		return "find"
	case "3":
		return "exit"
	default:
		return "not found"
	}
}

func addNewEmployee(Employees []Employee) {
	var name string
	var salary int
	var position string

	fmt.Println("Employee Name:")
	fmt.Scanln(&name)
	fmt.Println("Employee Salary:")
	fmt.Scanln(&salary)
	fmt.Println("Employee Position:")
	// type - "Project Manager" and check this !!!
	fmt.Scanln(&position)

	var newEmployee = Employee{Name: name, Salary: salary, Position: position}

	Employees = append(Employees, newEmployee)
	fmt.Println(Employees)

	provideChoice(Employees)
}

func findEmployee(Employees []Employee) {
	var input string
	fmt.Println("\nProvide Employee Name:")
	fmt.Scanln(&input)

	isFound := false

	for _, employee := range Employees {
		if strings.ToLower(input) == strings.ToLower(employee.Name) {
			fmt.Printf("\nEmployee Details:\n")
			fmt.Printf("%+v\n", employee)
			isFound = true
			break
		} 
	}

	if isFound == false {
		fmt.Printf("\nCan not find employee with name (%v)\n", input)
	}

	provideChoice(Employees)
}

func provideChoice(Employees []Employee) {
	res := checkUserChoice()

	switch res {
	case "add":
		addNewEmployee(Employees)
	case "find":
		findEmployee(Employees)
	case "exit":
		exit()
	case "not found":
		notFound(Employees)
	}
}

func notFound(Employees []Employee) {
	fmt.Println("\nInvalid command! Try again...")
	provideChoice(Employees)
}

func exit() {
	fmt.Println("\nexiting...")
}
