package handlers

import (
	"context"
	"log"
	"mongodb/types"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// go get "go.mongodb.org/mongo-driver/mongo"
// go get "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context) {
	// Setup options for connections to db
	db_options := options.Client().ApplyURI("mongodb://192.168.43.28:27018")
	// connection to db
	client, err := mongo.Connect(context.Background(), db_options)

	if err != nil {
		log.Printf("Ocurred error while connecting to DataBase: %v", err)
		c.JSON(500, gin.H{
			"error": err,
		})
	} else {
		var user types.User
		
		// get collection from db
		users_collection := client.Database("FoodsDB").Collection("Users")

		// get single data from collection
		res := users_collection.FindOne(context.Background(), bson.M{
			"firstname": "Jahongir",
		})

		res.Decode(&user)
		log.Println(user)
		c.JSON(200, gin.H{
			"data": user,
			"message": "Success",
		})
	}
}
