package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	in := &MyData{One: 1, Two: "second"}
	fmt.Println(in)

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(in)
	json.Unmarshal(inrec, &inInterface)

	fmt.Println(inInterface)

	// iterate through inrecs
	for field, val := range inInterface {
		fmt.Println("KV Pair: ", field, val)
	}

	// x := struct {
	// 	Name     string
	// 	Surname  string
	// 	LastName string
	// 	Login    string
	// 	Password string
	// }{"Jahongir", "Niyozbekov", "Dalerovich", "jony@gmail.com", "jony123"}

	// v := reflect.ValueOf(x)

	// values := make([]interface{}, v.NumField())

	// for i := 0; i < v.NumField(); i++ {
	// 	values[i] = v.Field(i).Interface()
	// 	fmt.Println(values[i])
	// }
}
