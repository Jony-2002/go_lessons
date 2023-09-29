package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	fmt.Println("JSON is coming...")
	// marshalling - process of sending data from back to front

	var testSlice = []string{"Hello", "World"}
	fmt.Println(testSlice)

	marshaledSlice, _ := json.Marshal(testSlice)
	fmt.Println(marshaledSlice)
	fmt.Println(string(marshaledSlice))

	fmt.Println(reflect.TypeOf(testSlice))
	fmt.Println(reflect.TypeOf(marshaledSlice))
	fmt.Println(reflect.TypeOf(string(marshaledSlice)))
	
}
