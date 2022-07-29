package api

import (
	"github.com/gorilla/mux"
	"github.com/kailash360/BookStore-Backend/controllers/basicController"
	"github.com/kailash360/BookStore-Backend/router/api/bookRoutes"
)

func Route(router *mux.Router) {

	//Routes for books
	bookRouter := router.PathPrefix("/books").Subrouter()
	bookRoutes.Route(bookRouter)

	router.HandleFunc("", basicController.HandleIncorrectRoute)
}
