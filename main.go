package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Car struct {
	ID    int     `json:"id"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

var cars []Car // In-memory slice to store cars
var nextID = 1 // Variable to generate unique IDs for each car

// Create a new car (POST /cars)
func createCar(w http.ResponseWriter, r *http.Request) {
	var newCar Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	newCar.ID = nextID
	nextID++
	cars = append(cars, newCar)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
}

// code to get all cars
// Get all cars (GET /cars)
func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func main() {
	router := mux.NewRouter()

	// Routes for CRUD operations
	router.HandleFunc("/cars", createCar).Methods("POST")
	router.HandleFunc("/cars", getCars).Methods("GET")

	// Start the server
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
