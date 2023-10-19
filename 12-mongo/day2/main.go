package main

import (
	"project/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.Cors())

	router.Run(":9999")
}