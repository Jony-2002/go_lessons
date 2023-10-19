package helpers

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// Setup options for connections to db 
	db_options := options.Client().ApplyURI("mongodb://127.0.0.1:27018")
	// connection to db
	client, err := mongo.Connect(context.Background(), db_options)

	if err != nil {
		log.Printf("Ocurred error while connecting to DataBase: %v", err)
	}
	return client
}