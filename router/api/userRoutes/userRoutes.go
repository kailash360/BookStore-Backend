package userRoutes

import (
	"github.com/gorilla/mux"
	"github.com/kailash360/BookStore-Backend/controllers/userController"
	"github.com/kailash360/BookStore-Backend/middlewares"
)

func Route(router *mux.Router) {
	router.HandleFunc("/{id}", middlewares.IsAuthroized(userController.GetUserDetails)).Methods("GET")
}
