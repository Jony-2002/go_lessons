package main

import (
	"github.com/gin-gonic/gin"
	"mongodb/handlers"
	"mongodb/middlewares"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.Cors())
	router.StaticFS("/static", gin.Dir("./static", true)) // if false, we cannot access the files in front just typing ip/static

	router.POST("/SignUp", handlers.SignUp)
	router.POST("/SignIn", handlers.SignIn)
	router.GET("/GetUsers", handlers.GetUsers)
	router.POST("/UpdateUser", handlers.UpdateUser)
	router.POST("/DeleteUser", handlers.DeleteUser)
	router.POST("/InsertManyUsers", handlers.InsertManyUsers)
	router.POST("/Find", handlers.Find)

	router.Run(":5555")
}
