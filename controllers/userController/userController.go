package userController

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kailash360/BookStore-Backend/models"
	"github.com/kailash360/BookStore-Backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserDetails(response http.ResponseWriter, request *http.Request) {
	//set the response headers
	response.Header().Set("Content-Type", "application/json")

	//Extract the id and convert into objectID
	id := mux.Vars(request)["id"]
	objectId, _ := primitive.ObjectIDFromHex(id)

	//Find the user in the database
	database := utils.GetClient()
	result := database.Collection("users").FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}})

	//Decode the result
	var user models.User
	result.Decode(&user)

	//If user is not found in the database, send an error
	if user.ID == primitive.NilObjectID {
		_response := models.Response{Success: false, Message: "User does not exist"}
		json.NewEncoder(response).Encode(_response)

		return
	}

	//Send the user data
	_response := models.Response{Success: true, Data: user}
	json.NewEncoder(response).Encode(_response)
}
