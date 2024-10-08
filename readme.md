# CRUD API For Cars

Basic API for get, post, put and delete cars data

## Package Used - ( Gorilla Mux )

Gorilla Mux is to handle routing and it maps HTTP requests (like GET, POST, PUT, DELETE) to specific handler functions

# Usage

- **POST** - http://localhost:8080/cars
- **GET** - http://localhost:8080/cars
- **GET BY ID** - http://localhost:8080/cars/{ID}
- **PUT** - http://localhost:8080/cars/{ID}
- **DELETE** - http://localhost:8080/cars/{ID}

## How to run

To run this file import the package

go get -u github.com/gorilla/mux

go run main.go
