package routes

import (
	"khorogtech/controllers"
	"net/http"
)

func AuthRouter() {
	http.HandleFunc("/signUp", controllers.SignUp)
}
