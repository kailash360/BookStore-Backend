package bookController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kailash360/BookStore-Backend/models"
	"github.com/kailash360/BookStore-Backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(response http.ResponseWriter, request *http.Request) {
	//set response header
	response.Header().Set("Content-Type", "application/json")

	//Get the database pointer
	database := utils.GetClient()

	//Get all the books from the collection
	cursor, err := database.Collection("books").Find(context.TODO(), bson.D{})
	if err != nil {
		//Send error message
		_response := models.Response{Success: false, Message: err.Error()}
		json.NewEncoder(response).Encode(_response)

		log.Fatal(err)
		return
	}
	defer cursor.Close(context.TODO())

	//Iterate over the MongoDB-based result
	var books []models.Book
	for cursor.Next(context.TODO()) {

		//Decode the Mongo-based result
		var _book models.Book
		cursor.Decode(&_book)

		//Add the book to the result
		books = append(books, _book)
	}

	//create the response to be sent
	_response := models.Response{Success: true, Data: books}

	//Send the response
	json.NewEncoder(response).Encode(_response)
}

func GetBook(response http.ResponseWriter, request *http.Request) {

	//Set the response header
	response.Header().Set("Content-Type", "application/json")

	//Get the id and convert ionto valid object id
	id := mux.Vars(request)["id"]
	objectId, _ := primitive.ObjectIDFromHex(id)

	//Fetch the book from the database
	database := utils.GetClient()
	result := database.Collection("books").FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: objectId}})

	//decode the result
	var book models.Book
	result.Decode(&book)

	//If book is not found, send an error
	if book.ID == primitive.NilObjectID {
		log.Print("Invalid ID used: " + id)
		_response := models.Response{Success: false, Message: "Invalid ID provided"}

		json.NewEncoder(response).Encode(_response)
		return
	}

	//Ceate response to be sent
	_response := models.Response{Success: true, Data: book}

	//Send the response
	json.NewEncoder(response).Encode(_response)
}

func AddBook(response http.ResponseWriter, request *http.Request) {

	//set the response header
	response.Header().Set("Content-Type", "application/json")

	//Decode the request body
	var book models.Book
	json.NewDecoder(request.Body).Decode(&book)

	//Insert the book into the database
	database := utils.GetClient()
	inserted, err := database.Collection("books").InsertOne(context.TODO(), book)

	fmt.Println("book ", book)
	//Return response if any error is encountered
	if err != nil {
		_response := models.Response{Success: false, Message: err.Error()}
		json.NewEncoder(response).Encode(_response)

		log.Fatal(err)
		return
	}

	//create the response to be sent
	_response := models.Response{Success: true, Data: inserted}

	//Send the response
	json.NewEncoder(response).Encode(_response)
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {

	//Set the response header
	response.Header().Set("Content-Type", "application/json")

	//Get the id and convert to object ID
	id := mux.Vars(request)["id"]
	objectId, _ := primitive.ObjectIDFromHex(id)

	//Decode the response and get updated values
	var updatedBook models.Book
	json.NewDecoder(request.Body).Decode(&updatedBook)

	//Fetch the book from the database
	database := utils.GetClient()
	result := database.Collection("books").FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}})

	//Decode the result
	var book models.Book
	result.Decode(&book)

	//If book is not found, return an error
	if book.ID == primitive.NilObjectID {
		_response := models.Response{Success: false, Message: "Invalid book ID"}

		json.NewEncoder(response).Encode(_response)
		return
	}

	//Update the document
	updated, err := database.Collection("books").UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: updatedBook}})

	//Handle if error is encountered
	if err != nil {
		_response := models.Response{Success: false, Message: err.Error()}
		json.NewEncoder(response).Encode(_response)

		log.Print("Error in updating book ID: ", id, err)
		return
	}

	//Send the response
	_response := models.Response{Success: true, Data: updated}
	json.NewEncoder(response).Encode(_response)
}
