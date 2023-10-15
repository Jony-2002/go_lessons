package main

import (
	"library/handlers"
	"library/middlewares"

	"github.com/gin-gonic/gin"
)

// "Access-Control-Allow-Origin", evt.AccessIP
// "Access-Control-Allow-Credentials", "true"
// "Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
// "Access-Control-Allow-Methods", "POST, GET, PUT, DELETE"

// c.Next()

func main() {
	router := gin.Default()

	router.Use(middlewares.Cors())
	
	router.POST("/signUp", handlers.SignUp)
	router.POST("/signIn", handlers.SignIn)
	router.GET("/getUsers", handlers.GetUsers)
	router.POST("/addEmployee", handlers.AddEmployee)
	router.POST("/deleteEmployee", handlers.DeleteEmployee)

	router.Run(":9999")
}
