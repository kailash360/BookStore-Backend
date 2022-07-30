package api

import (
	"github.com/gorilla/mux"
	"github.com/kailash360/BookStore-Backend/controllers/basicController"
	"github.com/kailash360/BookStore-Backend/router/api/bookRoutes"
	"github.com/kailash360/BookStore-Backend/router/api/userRoutes"
)

func Route(router *mux.Router) {

	//routes for users
	userRouter := router.PathPrefix("/users").Subrouter()
	userRoutes.Route(userRouter)

	//Routes for books
	bookRouter := router.PathPrefix("/books").Subrouter()
	bookRoutes.Route(bookRouter)

	router.HandleFunc("", basicController.HandleIncorrectRoute)
}
