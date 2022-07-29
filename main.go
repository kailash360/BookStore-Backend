package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kailash360/BookStore-Backend/controllers/bookController"
	"github.com/kailash360/BookStore-Backend/utils"
)

func main() {
	fmt.Println("Book Store")

	//Load the environment variables
	godotenv.Load()

	//Connect to the database
	utils.Connect()

	//Add the router
	router := mux.NewRouter()

	router.HandleFunc("/", bookController.GetBooks).Methods("GET")
	router.HandleFunc("/{id}", bookController.GetBook).Methods("GET")
	router.HandleFunc("/", bookController.AddBook).Methods("POST")

	fmt.Println("Server started successfully!!")
	log.Fatal(http.ListenAndServe(":8080", router))

}
