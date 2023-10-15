package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mail/data"
	customTypes "mail/interfaces"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func File(c *gin.Context) {
	var NewFile customTypes.File

	c.ShouldBindJSON(&NewFile)

	if NewFile.Login == "" || NewFile.ImageUrl == "" {
		log.Printf("All fields are required. Please try again!\n")

		c.JSON(409, gin.H{
			"message": "All fields are required. Please try again!",
		})
	} else {
		// Generate random name for each file
		imageID := rand.Intn(10000) * time.Now().Second()
		fileName := fmt.Sprintf("upload-%v.png", imageID)

		// Get the messages from json
		messagesJSON, _ := ioutil.ReadFile("data/files.json")
		json.Unmarshal(messagesJSON, &data.Files)

		// Create image with ImageUrl (which is byte format)
		binaryImage, _ := base64.StdEncoding.DecodeString(NewFile.ImageUrl)
		// save the file
		// fileName - is the generated name
		// binaryImage - is the image that created - base64.StdEncoding.DecodeString //! Check line 38
		// 0644 - system permission (use this while creating files)
		ioutil.WriteFile("./static/"+fileName, binaryImage, 0644)
		//! Change the filename which is byte format by default
		NewFile.ImageUrl = fileName

		data.Files = append(data.Files, NewFile)
		marshalledMessages, _ := json.Marshal(data.Files)
		ioutil.WriteFile("data/files.json", marshalledMessages, 0644)

		log.Printf("Success\n")

		c.JSON(200, gin.H{
			"message": "Success",
			"data":    marshalledMessages,
		})
	}
}
