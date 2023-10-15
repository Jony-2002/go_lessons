package controllers

import (
	"gin/KhorogTECH/types"
	"gin/KhorogTECH/utils"
	"log"

	"github.com/gin-gonic/gin"
)

// registration
func SignUp(c *gin.Context) {
	var NewUser types.Registration

	c.ShouldBindJSON(&NewUser)
	log.Printf("New user added to the DataBase: %+v", NewUser)

	var res types.IsOKType = utils.CheckData(NewUser)

	if !res.Result {
		c.JSON(400, gin.H{
			"message": res.Message,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Success",
			"data": NewUser,
		})
	}
}

func Login(c *gin.Context) {

}

// login
