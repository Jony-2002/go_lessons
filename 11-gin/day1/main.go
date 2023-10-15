package main

import (
	"gin/module/functions"
	"github.com/gin-gonic/gin"
)

func main() {
	// route will be on the main.go file
	Router := gin.Default()

	Router.GET("/getHello", functions.GetHello) // 127.0.0.1:9999/getHello
	Router.POST("/login", functions.Login) // 127.0.0.1:9999/getHello

	Router.Run(":9999")
}
