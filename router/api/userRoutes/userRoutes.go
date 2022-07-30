package userRoutes

import (
	"github.com/gorilla/mux"
	"github.com/kailash360/BookStore-Backend/controllers/userController"
)

func Route(router *mux.Router) {
	router.HandleFunc("/{id}", userController.GetUserDetails).Methods("GET")
}
