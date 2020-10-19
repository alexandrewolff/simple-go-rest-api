package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"simple-go-rest-api/helper"
	"simple-go-rest-api/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Runs database connection helper
var collection = helper.ConnectDB()

func GetCars(w http.ResponseWriter, r *http.Request) {
	// Sets header
	w.Header().Set("Content-Type", "application/json")

	var cars []models.Car

	// Gets all data in the database
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
		return
	}

	// Will close the cursor once the function is finished
	defer cur.Close(context.TODO())

	// Puts the found documents inside a slice
	for cur.Next(context.TODO()) {
		var car models.Car

		// Gets the current value and saves it into the car variable
		err := cur.Decode(&car)
		if err != nil {
			log.Fatal(err)
		}

		cars = append(cars, car)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Responds to the HTTP requests with the slice
	json.NewEncoder(w).Encode(cars)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	// Sets header
	w.Header().Set("Content-Type", "application/json")

	var car models.Car

	// Gets the HTTP request params
	var params = mux.Vars(r)

	// Converts the id into Mongodb's format
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// Creates a filter for the database fetch
	filter := bson.M{"_id": id}
	// Fetches the database with the filter and saves the result into the car variable
	err := collection.FindOne(context.TODO(), filter).Decode(&car)

	if err != nil {
		log.Fatal(err)
	}

	// Responds to the HTTP requests with the result
	json.NewEncoder(w).Encode(car)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	// Sets header
	w.Header().Set("Content-Type", "application/json")

	var car models.Car

	// Saves the HTTP request parameters into the car variable
	_ = json.NewDecoder(r.Body).Decode(&car)

	// Inserts the new model into the database
	result, err := collection.InsertOne(context.TODO(), car)

	if err != nil {
		log.Fatal(err)
	}

	// Responds to the HTTP requests with the result
	json.NewEncoder(w).Encode(result)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	// Sets header
	w.Header().Set("Content-Type", "application/json")

	// Gets the HTTP request params
	var params = mux.Vars(r)

	// Converts the id into Mongodb's format
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var car models.Car

	// Creates a filter for the database fetch
	filter := bson.M{"_id": id}

	// Saves the HTTP request parameters into the car variable
	_ = json.NewDecoder(r.Body).Decode(&car)

	// Sets the update model
	update := bson.D{
		{"$set", bson.D{
			{"brand", car.Brand},
			{"model", car.Model},
			{"year", car.Year},
		}},
	}

	// Updates the targeted document and saves previous informations into the car variable
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&car)

	if err != nil {
		log.Fatal(err)
	}

	car.ID = id

	// Responds to the HTTP requests with the old setting
	json.NewEncoder(w).Encode(car)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	// Sets header
	w.Header().Set("Content-Type", "application/json")

	// Gets the HTTP request params
	var params = mux.Vars(r)

	// Converts the id into Mongodb's format
	id, err := primitive.ObjectIDFromHex(params["id"])

	// Creates a filter for the database fetch
	filter := bson.M{"_id": id}

	// Deletes the targeted document
	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	// Responds to the HTTP requests with the number of deleted document
	json.NewEncoder(w).Encode(result)
}
