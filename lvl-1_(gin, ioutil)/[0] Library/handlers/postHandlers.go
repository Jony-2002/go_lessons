package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"library/data"
	"library/interfaces"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"time"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var NewUser interfaces.User
	c.ShouldBindJSON(&NewUser)
	NewUser.ID = fmt.Sprintf("%v", rand.Intn(1000000000)) //? Convert integer to string with Sprintf

	values := reflect.ValueOf(NewUser)
	interfaceTYPES := values.Type()

	var isOK bool = true

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface() == "" {
			isOK = false

			log.Printf("Field %s is required", interfaceTYPES.Field(i).Name)
			var ErrorMessage string = interfaceTYPES.Field(i).Name + " is required"

			c.JSON(400, gin.H{
				"message": ErrorMessage,
			})

			break
		}
	}

	if isOK {
		res, _ := ioutil.ReadFile("data/users.json")
		json.Unmarshal(res, &data.Users)

		var hasFound bool = false

		for _, v := range data.Users {
			if NewUser.Login == v.Login {
				hasFound = true

				c.JSON(400, gin.H{
					"message": "User with this Login has already exist. Try another one",
				})

				break
			}
		}
		if !hasFound {
			// NewUser.ID = fmt.Sprintf("%v", rand.Intn(1000000000)) //? Convert integer to string with Sprintf
			data.Users = append(data.Users, NewUser)
			marshalledData, _ := json.Marshal(data.Users)
			ioutil.WriteFile("data/users.json", marshalledData, 0644)

			log.Printf("Successfully registered user: %+v\n", NewUser)

			c.JSON(200, gin.H{
				"message": "Success",
				"data":    NewUser,
			})
		}
	}
}

func SignIn(c *gin.Context) {
	var User interfaces.Login
	c.ShouldBindJSON(&User)

	values := reflect.ValueOf(User)
	typeOfS := values.Type()

	var isOK bool = true

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface() == "" {
			isOK = false
			log.Printf("Field %v is required! Please try again\n", typeOfS.Field(i).Name)
			var ErrorMessage string = typeOfS.Field(i).Name + " is required"

			c.JSON(400, gin.H{
				"message": ErrorMessage,
			})

			break
		}
	}

	if isOK {
		db, _ := ioutil.ReadFile("data/users.json")
		json.Unmarshal(db, &data.Users)

		var isExisting bool = false

		for _, v := range data.Users {
			if User.Login == v.Login && User.Password == v.Password {
				isExisting = true

				http.SetCookie(c.Writer, &http.Cookie{
					Name:     "Access_Token",
					Value:    User.Login,
					Secure:   false,
					HttpOnly: false, // if false - javascript has no access to cookies
					Expires:  time.Now().Add(3 * time.Hour),
				})

				c.JSON(200, gin.H{
					"message": "Success",
					"data":    v,
				})

				break
			}
		}

		if !isExisting {
			c.JSON(404, gin.H{
				"message": "Invalid login or password",
			})
		}
	}
}

func AddEmployee(c *gin.Context) {
	res, _ := ioutil.ReadFile("data/employees.json")
	json.Unmarshal(res, &data.Employees)

	var NewEmployee interfaces.Employee
	c.ShouldBindJSON(&NewEmployee)
	NewEmployee.ID = fmt.Sprintf("%v", rand.Intn(1000000000))

	// binaryImage, _ := base64.StdEncoding.DecodeString(NewEmployee.ImageUrl)
	// ioutil.WriteFile("static/image-1.png", binaryImage, 0644)

	var isOK bool = true

	values := reflect.ValueOf(NewEmployee)
	interfaceKeys := values.Type()

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface() == "" || values.Field(i).Interface() == 0 {
			isOK = false

			log.Printf("Field %s is required", interfaceKeys.Field(i).Name)
			var ErrorMessage string = interfaceKeys.Field(i).Name + " is required"

			c.JSON(400, gin.H{
				"message": ErrorMessage,
			})

			break
		}
	}

	if isOK {
		var hasFound bool = false

		for _, v := range data.Employees {
			if NewEmployee.Name == v.Name || NewEmployee.Surname == v.Surname {
				hasFound = true
				break
			}
		}

		if hasFound {
			c.JSON(400, gin.H{
				"message": "Employee with these credentials has already exist!",
			})
		} else {
			data.Employees = append(data.Employees, NewEmployee)
			marshalledData, _ := json.Marshal(data.Employees)
			
			//! Store image
			binaryImage, _ := base64.StdEncoding.DecodeString(NewEmployee.ImageUrl)
			ioutil.WriteFile("static/image-1.png", binaryImage, 0644)
			
			//! store employee
			ioutil.WriteFile("data/employees.json", marshalledData, 0644)

			c.JSON(200, gin.H{
				"message": "Success",
				"data":    NewEmployee,
			})
		}
	}
}

func DeleteEmployee(c *gin.Context) {
	// id := c.Request.URL.Query().Get("id")
	id := c.Query("id")
	fmt.Println(id)

	employeesJSON, _ := ioutil.ReadFile("data/employees.json")
	json.Unmarshal(employeesJSON, &data.Employees)

	var hasFound bool = false
	
	for _, v := range data.Employees {
		if id == v.ID {
			hasFound = true
			break
		}
	}
	
	if hasFound {
		log.Println("Found employee with this ID")
	}
}