package main

import (
	// "encoding/json"
	"fmt"
	"reflect"
)

// type User struct {
// 	Name     string
// 	Surname  string
// 	LastName string
// 	Login    string
// 	Password string
// }

type MyData struct {
	One   int
	Two   string
	Three int
}

type Student struct {
	Fname  string
	Lname  string
	City   string
	Mobile int64
}

func main() {
	// in := &MyData{One: 1, Two: "second"}
	// fmt.Println(in)

	// var inInterface map[string]interface{}
	// inrec, _ := json.Marshal(in)
	// json.Unmarshal(inrec, &inInterface)

	// fmt.Println(inInterface)

	// // iterate through inrecs
	// for field, val := range inInterface {
	// 	fmt.Println("KV Pair: ", field, val)
	// }

	// x := struct {
	// 	Name     string
	// 	Surname  string
	// 	LastName string
	// 	Login    string
	// 	Password string
	// }{"Jahongir", "Niyozbekov", "Dalerovich", "jony@gmail.com", "jony123"}

	// v := reflect.ValueOf(x)
	// fmt.Println("v = ", v)

	// values := make([]interface{}, v.NumField())
	// fmt.Println("values = ", values)

	// for i := 0; i < v.NumField(); i++ {
	// 	values[i] = v.Field(i).Interface()
	// 	fmt.Println(values[i])
	// }

	s := Student{"Chetan", "Kumar", "Bangalore", 7777777777}
	v := reflect.ValueOf(s) // returns
	fmt.Println("reflect.ValueOf(s)", v) // returns the struct length
	fmt.Println("v.NumField()", v.NumField())
	typeOfS := v.Type()

	// 					0 < 4 (because there are 4 Field in the Student struct)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}

	// typeOfS.Field(i).Name - returns the struct key
	// v.Field(i).Interface() - returns the struct value
}
