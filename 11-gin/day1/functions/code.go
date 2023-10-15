package functions

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginInterface struct {
	Email    string
	Password string
}

// var names = []string{"1", "2", "3", "4"}

func GetHello(c *gin.Context) {
	// c.Request.Body()
	// fmt.Fprintf(c.Writer, "dqwdqw")

	// c.JSON(200, gin.H{
	// 	"message": "Success",
	// 	"data": names,
	// 	"statusCode": "OK",
	// })
}

func Login(c *gin.Context) {
	var User LoginInterface

	c.ShouldBindJSON(&User)
	c.JSON(200, gin.H{
		"message": "Success",
		"statusCode": "OK",
		"data": User,
	})
	fmt.Printf("User is requiting: %+v\n", User)
}
