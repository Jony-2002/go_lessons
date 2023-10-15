package handlers

import (
	"gin/Library/data"
	"log"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var AllBooks = data.Books
	
	cookie, err := c.Request.Cookie("Access_Token")

	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Invalid cookie",
		})
	} else {
		log.Printf("Current user is trying to get data with cookie: %v\n\n", cookie)
		c.ShouldBindJSON(AllBooks)
		c.JSON(200, gin.H{
			"message": "Success",
			"data": AllBooks,
		})
	}
	

	log.Println(AllBooks)
}