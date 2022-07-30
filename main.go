package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kailash360/BookStore-Backend/controllers/basicController"
	"github.com/kailash360/BookStore-Backend/router/api"
	"github.com/kailash360/BookStore-Backend/router/auth"
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

	router.HandleFunc("/", basicController.HandleHome)

	//Handling routes for apis
	apiRouter := router.PathPrefix("/api").Subrouter()
	api.Route(apiRouter)

	//Handling routes for auth
	authRouter := router.PathPrefix("/auth").Subrouter()
	auth.Route(authRouter)

	fmt.Println("Server started successfully!!")
	log.Fatal(http.ListenAndServe(":8080", router))

}
