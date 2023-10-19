package main

import (
	"github.com/gin-gonic/gin"
	"khorogtech/handlers"
	"khorogtech/middlewares"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.Cors())
	router.StaticFS("/static", gin.Dir("./static", true)) // if false, we cannot access the files in front just typing ip/static

	router.POST("/SignUp", handlers.SignUp)
	// router.POST("/SignIn", handlers.SignIn)
	router.GET("/GetUsers", handlers.GetUsers)

	router.Run(":5555")
}
