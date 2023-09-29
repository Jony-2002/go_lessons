package main

import (
	"fmt"
	"net/http"
)

const PORT = 8080

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Good bye!")
}

func main() {
	fmt.Println("Server is running on: 8080")

	// Rest api or HTTP Server
	// In order to communicate front (client) with back (server) it usually use provider or HTTP

	// In the first days, we will use browser and postman

	// creating handlers or routes in order to catch different functions
	http.HandleFunc("/greeting", sayHello)
	http.HandleFunc("/goodBye", sayBye)

	http.ListenAndServe(":8080", nil)
	// ListenAndServe
	// ListenAndServe - is function which takes two parameters
	// 1 - port - (in order to make sure which is front or back)
	// 2 - 

	// ip - for different computers
	// port - for different applications in computer
	// route or functions - which function to call to make different requests
}
