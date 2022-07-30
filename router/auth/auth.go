package auth

import (
	"github.com/gorilla/mux"
	"github.com/kailash360/BookStore-Backend/controllers/authController"
)

func Route(router *mux.Router) {
	router.HandleFunc("/signup", authController.SignUp).Methods("POST")
}
