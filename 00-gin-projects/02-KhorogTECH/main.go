package main

import (
	"gin/KhorogTECH/controllers"
	"gin/KhorogTECH/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())
	
	router.POST("/signUp", controllers.SignUp)
	router.POST("/login", controllers.Login)
	
	// router.GET("/getUsers")
	// router.GET("/getUserById")
	// router.POST("/updateUser")
	// router.POST("/deleteUser")
	// router.POST("/forgotPassword")

	// router.GET("/getEvents")
	// router.GET("/getEventById")
	// router.POST("/createEvent")
	// router.POST("/updateEvent")
	// router.POST("/deleteEvent")

	// router.GET("getCourses")
	// router.GET("getCourseById")
	// router.POST("createCourse")
	// router.POST("updateCourse")
	// router.POST("deleteCourse")
	// router.POST("buyCourse")

	router.Run(":5555")
}