package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"khorogtech/types"
)

// func insertData() {

// }

func SignUp(c *gin.Context) {
	// Setup options for connections to db 
	db_options := options.Client().ApplyURI("mongodb://127.0.0.1:27019")
	// connection to db
	client, err := mongo.Connect(context.Background(), db_options)

	if err != nil {
		log.Fatalf("Ocurred error while connecting to DataBase: %v", err)
		c.JSON(500, gin.H{
			"error": err,
		})
	} else {
		// get collection from db
		users_collection := client.Database("kta").Collection("users")

		var NewUser types.User
		c.ShouldBindJSON(&NewUser)

		// insert data to collection in db
		res, err := users_collection.InsertOne(context.Background(), NewUser)

		if err != nil {
			log.Fatalf("%v", err)
		} else {
			log.Printf("%v", res)
		}
	}
}

// func SignIn(c *gin.Context) {
// 	log.Println("Sign In")
// }
