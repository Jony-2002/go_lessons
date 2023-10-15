package main

import (
	"fmt"
	"khorogtech/routes"
	"net/http"
)

func main() {
	fmt.Println("Server is running on: 127.0.0.1:8888")

	routes.UserRouters()
	routes.AuthRouter()
	
	http.ListenAndServe(":8888", nil)
}