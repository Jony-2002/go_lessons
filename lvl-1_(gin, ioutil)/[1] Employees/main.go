package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/SignUp")
	router.POST("/SignIn")
	
	router.Run(":9999")
}