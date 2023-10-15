package handlers

import (
	"encoding/json"
	"io/ioutil"
	"mail/data"
	"github.com/gin-gonic/gin"
)

func Messages(c *gin.Context) {
	messagesJSON, _ := ioutil.ReadFile("data/files.json")
	json.Unmarshal(messagesJSON, &data.Files)

	// c.JSON(200, gin.H{
	// 	"message": "Success",
	// 	"data": data.Files,
	// })
	c.JSON(200, data.Files)
}
