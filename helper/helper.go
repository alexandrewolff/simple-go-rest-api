package helper

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connects to the database
func ConnectDB() *mongo.Collection {
	// Sets the db address
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// Connects to Mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected")

	// Gets the desired collection
	collection := client.Database("rest-api").Collection("cars")

	return collection
}
