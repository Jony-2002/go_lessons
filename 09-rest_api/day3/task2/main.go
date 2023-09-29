package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name     string
	Surname  string
	UserName string
	Email    string
	Password string
}

type Order struct {
	Name  string
	Order string
}

type Food struct {
	Name  string
	Price int
}

var Orders = []Order{
	{
		Name:  "Jahongir",
		Order: "Pizza",
	},
}

var Foods = []Food{
	{
		Name:  "Pizza",
		Price: 45,
	},
	{
		Name:  "Burger",
		Price: 28,
	},
	{
		Name:  "HotDog",
		Price: 16,
	},
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder Order

	json.NewDecoder(r.Body).Decode(&newOrder)

	newOrder.checkOrderFields(w)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	marshalledData, _ := json.Marshal(Orders)
	w.Write(marshalledData)
	fmt.Println(string(marshalledData))
}

func main() {
	fmt.Println("server is listening on: 127.0.0.1:7777")

	http.HandleFunc("/createOrder", createOrder)
	http.HandleFunc("/getOrders", getOrders)

	http.ListenAndServe(":7777", nil)
}

func (item Order) checkOrderFields(response http.ResponseWriter) {
	mapItem := &item
	fmt.Println(mapItem)

	var inInterface map[string]interface{}
	marshalledMap, _ := json.Marshal(mapItem)
	json.Unmarshal(marshalledMap, &inInterface)

	fmt.Println(inInterface)

	var isNotEmpty bool = true

	// iterate through
	for field, val := range inInterface {
		fmt.Println("Data: ", field, val)

		if val == "" {
			fmt.Printf("%v cannot be empty", field)
			isNotEmpty = false
		}
	}

	hasFound := true

	if isNotEmpty == true {
		for k := 0; k < len(Foods); k++ {
			if item.Order == Foods[k].Name {
				Orders = append(Orders, item)

				marshalledData, _ := json.Marshal(item)
				fmt.Fprintf(response, "Your order is: %v", Foods[k].Price)
				// response.Write(marshalledData)
				fmt.Println(string(marshalledData))
				fmt.Println("Success")
				break
			} else {
				hasFound = false
			}
		}
	}

	if hasFound == false {
		fmt.Fprintf(response, "Cannot find food with name: %v", item.Order)
	}
}
