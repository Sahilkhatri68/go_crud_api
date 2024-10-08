package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// code to get car by ID

// function to get car by ID
func getCarByID(id int) (*Car, int) {
	for i, car := range cars {
		if car.ID == id {
			return &car, i
		}
	}
	return nil, -1
}

func getCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}
	car, _ := getCarByID(id)
	if car == nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func main() {
	router := mux.NewRouter()

	// Routes for CRUD operations
	router.HandleFunc("/cars", createCar).Methods("POST")
	router.HandleFunc("/cars", getCars).Methods("GET")
	router.HandleFunc("/cars/{id}", getCar).Methods("GET")

	// Start the server
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
