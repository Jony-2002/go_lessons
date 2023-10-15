package main

import (
	"mail/handlers"
	"mail/middlewares"

	"github.com/gin-gonic/gin"
)

// func Cors(c *gin.Context) {
// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://192.168.43.28:5500") // http://192.168.43.28:5500/addFile.html?surname=admin
// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")

// 	if c.Request.Method == "OPTIONS" {
// 		c.AbortWithStatus(200)
// 		return
// 	}

// 	c.Next()
// }

func main() {
	r := gin.Default()

	r.Use(middlewares.Cors())
	r.StaticFS("/static", gin.Dir("./static", true)) // if false, we cannot access the files in front just typing ip/static

	r.GET("/messages", handlers.Messages)
	r.POST("/file", handlers.File)

	r.Run(":5657")
}
