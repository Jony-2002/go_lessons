package main

import (
	"gin/Library/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/SignUp", handlers.SignUp)
	router.POST("/Login", handlers.Login)
	// router.GET("/GetBooks", handlers.GetgBooks)

	router.Run((":9999"))
}
