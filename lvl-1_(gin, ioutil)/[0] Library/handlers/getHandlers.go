package handlers

import (
	"encoding/json"
	"io/ioutil"
	"library/data"
	"library/interfaces"
	"log"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db, _ := ioutil.ReadFile("data/users.json")
	json.Unmarshal(db, &data.Users)

	cookie, err := c.Request.Cookie("Access_Token")

	if err != nil {
		log.Printf("%v\n", err)
		log.Printf("Invalid cookies\n")

		c.JSON(400, gin.H{
			"message": "Invalid cookie",
		})
	} else {
		var hasCookie bool = false
		var User interfaces.User

		for _, v := range data.Users {
			if cookie.Value == v.Login {
				User = v
				hasCookie = true
				break
			}
		}

		if hasCookie {
			log.Printf("Username: %v, Surname: %v, ID: %v, - got all users\n\n", User.Name, User.Surname, User.ID)

			c.JSON(200, gin.H{
				"message": "Success",
				"data":    data.Users,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "Invalid cookie",
			})
		}
	}
}
