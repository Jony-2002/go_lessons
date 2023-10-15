package routes

import (
	"khorogtech/controllers"
	"net/http"
)

func UserRouters() {
	http.HandleFunc("/getUsers", controllers.GetUsers)
}