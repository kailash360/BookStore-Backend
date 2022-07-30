package authController

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kailash360/BookStore-Backend/models"
	"github.com/kailash360/BookStore-Backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(response http.ResponseWriter, request *http.Request) {
	//Set the response headers
	response.Header().Set("Content-Type", "application/json")

	//Extract the json body
	var user models.User
	json.NewDecoder(request.Body).Decode(&user)

	//Check if the email is already registered
	database := utils.GetClient()
	result := database.Collection("users").FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}})

	var existingUser models.User
	result.Decode(&existingUser)

	if existingUser.ID != primitive.NilObjectID {
		_response := models.Response{Success: false, Message: "Email is already registered"}
		json.NewEncoder(response).Encode(_response)

		return
	}

	//Add the user to the database if not already registered
	inserted, err := database.Collection("users").InsertOne(context.TODO(), user)

	//Check if any error was encountered
	if err != nil {
		_response := models.Response{Success: false, Message: err.Error()}
		json.NewEncoder(response).Encode(_response)

		return
	}

	_response := models.Response{Success: true, Data: inserted}
	json.NewEncoder(response).Encode(_response)
}