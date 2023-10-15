package handlers

import (
	"gin/Library/data"
	"gin/Library/interfaces"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"time"
)

func SignUp(c *gin.Context) {
	var newUser interfaces.IUser

	c.ShouldBindJSON(&newUser)

	if newUser.FirstName == "" {
		c.JSON(400, gin.H{
			"message": "FirstName is required",
		})
	} else if newUser.Surname == "" {
		c.JSON(400, gin.H{
			"message": "Surname is required",
		})
	} else if newUser.Phone == "" {
		c.JSON(400, gin.H{
			"message": "Phone is required",
		})
	} else if newUser.Login == "" {
		c.JSON(400, gin.H{
			"message": "Login is required",
		})
	} else if newUser.Password == "" {
		c.JSON(400, gin.H{
			"message": "Password is required",
		})
	} else {
		hasFound := false

		if len(data.Users) > 0 {
			for _, v := range data.Users {
				if newUser.Login == v.Login {
					hasFound = true

					c.JSON(400, gin.H{
						"message": "User with this Login has already exist. Try another one",
					})

					break
				}
			}
			// hasFound == false
			if !hasFound {
				data.Users = append(data.Users, interfaces.IUser{
					FirstName: newUser.FirstName,
					Surname:   newUser.Surname,
					Phone:     newUser.Phone,
					Login:     newUser.Login,
					Password:  newUser.Password,
				})

				log.Println(data.Users)

				c.JSON(200, gin.H{
					"message": "You are successfully registered to our service!",
					"data":    newUser,
				})
			}
		} else {
			data.Users = append(data.Users, interfaces.IUser{
				FirstName: newUser.FirstName,
				Surname:   newUser.Surname,
				Phone:     newUser.Phone,
				Login:     newUser.Login,
				Password:  newUser.Password,
			})

			log.Println(data.Users)

			c.JSON(200, gin.H{
				"message": "You are successfully registered to our service!",
				"data":    newUser,
			})
		}
	}
}

func Login(c *gin.Context) {
	var User interfaces.ILogin

	c.ShouldBindJSON(&User)

	values := reflect.ValueOf(User)
	typeOfS := values.Type()

	var isOK bool = true

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface() == "" {
			isOK = false
			log.Printf("Feld %s is required", typeOfS.Field(i).Name)
			var ErrorMessage string = typeOfS.Field(i).Name + " is required"

			// How can i pass dynamic value to JSON
			c.JSON(400, gin.H{
				"message": ErrorMessage,
			})
			break
		}
	}

	if isOK {
		hasFound := false

		for _, v := range data.Users {
			if User.Login == v.Login && User.Password == v.Password {
				hasFound = true

				http.SetCookie(c.Writer, &http.Cookie{
					Name:    "Access_Token",
					Value:   User.Login,
					Expires: time.Now().Add(3 * time.Hour),
				})

				c.JSON(200, gin.H{
					"message": "Successfully logged in",
					"data":    User,
				})

				break
			}
		}
		// hasFound == false
		if !hasFound {
			c.JSON(404, gin.H{
				"message": "Invalid Login or Password",
			})
		}
	}
}
