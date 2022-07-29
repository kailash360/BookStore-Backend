package bookRoutes

import (
	"github.com/gorilla/mux"
	"github.com/kailash360/BookStore-Backend/controllers/bookController"
)

func Route(router *mux.Router) {

	router.HandleFunc("", bookController.GetBooks).Methods("GET")
	router.HandleFunc("/{id}", bookController.GetBook).Methods("GET")
	router.HandleFunc("", bookController.AddBook).Methods("POST")
	router.HandleFunc("/{id}", bookController.UpdateBook).Methods("PUT")
	router.HandleFunc("/{id}", bookController.DeleteBook).Methods("DELETE")

}
