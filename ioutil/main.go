package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var names = []string{"yakum", "duyum", "seyum"}

func main() {
	res, _ := ioutil.ReadFile("test.json")
	fmt.Println(res)
	
	json.Unmarshal(res, &names)
	fmt.Println(names)

	var newValue string
	
	fmt.Scanln(&newValue)
	names = append(names, newValue)	

	// all access - 0644 - code 
	
	marshalledData, _ := json.Marshal(names)
	
	ioutil.WriteFile("test.json", marshalledData, 0644)
}
