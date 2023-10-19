package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongodb/types"
)

func SignUp(c *gin.Context) {
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
		// get collection from db
		users_collection := client.Database("FoodsDB").Collection("Users")

		var NewUser types.User
		c.ShouldBindJSON(&NewUser)

		// insert data to collection in db
		res, err := users_collection.InsertOne(context.Background(), NewUser)

		if err != nil {
			log.Printf("%v", err)
		} else {
			log.Printf("%v", res)
		}
	}
}

func SignIn(c *gin.Context) {
	log.Println("Sign In")
}

func UpdateUser(c *gin.Context) {
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
		// get collection from db
		users_collection := client.Database("FoodsDB").Collection("Users")

		var user types.User
		c.ShouldBindJSON(&user)

		//! Change specific field
		// users_collection.UpdateOne(context.Background(),
		// 	bson.M{
		// 		"firstname": "Jahongir",
		// 	},
		// 	bson.D{
		// 		{
		// 			"$set", bson.D{
		// 				{"password", "jony123"},
		// 			},
		// 		},
		// 	},
		// )

		//! Change all fields
		users_collection.UpdateOne(context.Background(),
			bson.M{
				"firstname": "idi nakhren",
			},
			bson.D{
				{
					"$set", user,
				},
			},
		)
	}
}

func DeleteUser(c *gin.Context) {
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
		users_collection := client.Database("FoodsDB").Collection("Users")

		_, err := users_collection.DeleteOne(context.Background(), bson.M{
			"firstname": "Masrur",
		})

		if err != nil {
			log.Fatalf("%v\n", err)
		} else {
			log.Println("Success")
		}
	}
}

func InsertManyUsers(c *gin.Context) {
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
		//! Get array from BODY object

		var someUsers= []interface{}{}
		c.ShouldBindJSON(&someUsers)

		// newRes := make([]interface{}, len(someUsers))
		log.Println(someUsers)

		// var someUsers = []interface{}{
		// 	bson.M{
		// 		"FirstName": "dqwdqw",
		// 		"LastName":  "dqwdqwdqwsdqwd",
		// 		"Email":     "dqwodqwoid@gmailcom",
		// 		"Password":  "qsqws12",
		// 	},
		// 	bson.M{
		// 		"FirstName": "111111",
		// 		"LastName":  "1111111111111",
		// 		"Email":     "1111111111@gmailcom",
		// 		"Password":  "1111111111111",
		// 	},
		// 	bson.M{
		// 		"FirstName": "2222222222",
		// 		"LastName":  "222222222222222222",
		// 		"Email":     "2222222222222@gmailcom",
		// 		"Password":  "22222222222222",
		// 	},
		// }

		users_collection := client.Database("FoodsDB").Collection("Users")

		users_collection.InsertMany(context.Background(), someUsers)
	}
}

func Find(c *gin.Context) {
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
		var decodedUser types.User

		users_collection := client.Database("FoodsDB").Collection("Users")

		res, _ := users_collection.Find(context.Background(), bson.M{})

		var copy = []types.User{}

		//! What is Next for?
		for res.Next(context.Background()) {
			res.Decode(&decodedUser)
			log.Println(decodedUser)
			copy = append(copy, decodedUser)
		}

		//! What is Close for?
		res.Close(context.Background())

		c.JSON(200, gin.H{
			"data": copy,
		})
	}
}
