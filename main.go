package main

import (
	"log"
	"net/http"
	"simple-go-rest-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the mux router
	r := mux.NewRouter()

	// Creates endpoints
	r.HandleFunc("/api/cars", routes.GetCars).Methods("GET")
	r.HandleFunc("/api/cars/{id}", routes.GetCar).Methods("GET")
	r.HandleFunc("/api/cars", routes.CreateCar).Methods("POST")
	r.HandleFunc("/api/cars/{id}", routes.UpdateCar).Methods("PUT")
	r.HandleFunc("/api/cars/{id}", routes.DeleteCar).Methods("DELETE")

	// Runs the server
	log.Fatal(http.ListenAndServe(":8000", r))

}
