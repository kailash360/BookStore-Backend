package authController

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/kailash360/BookStore-Backend/models"
	"github.com/kailash360/BookStore-Backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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

	//Hash the pasword
	cost, _ := strconv.Atoi(os.Getenv("PASSWORD_COST"))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)

	if err != nil {
		_response := models.Response{Success: false, Message: err.Error()}
		json.NewEncoder(response).Encode(_response)

		return
	}

	//Set the hashed password as new password
	user.Password = string(hashedPassword)

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

func SignIn(response http.ResponseWriter, request *http.Request) {
	//Set the response headers
	response.Header().Set("Content-Type", "application/json")

	//Extract the credentials from the request body
	var credentials models.Credentials
	json.NewDecoder(request.Body).Decode(&credentials)

	//Check if the email is present in the database
	database := utils.GetClient()
	result := database.Collection("users").FindOne(context.TODO(), bson.D{{Key: "email", Value: credentials.Email}})

	//Decode the result
	var user models.User
	result.Decode(&user)

	//If user is not present, return an error
	if user.ID == primitive.NilObjectID {
		_response := models.Response{Success: false, Message: "Incorrect credentials"}
		json.NewEncoder(response).Encode(_response)

		return
	}

	//If present, compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))

	//If error is true, then hashes do not match
	if err != nil {
		_response := models.Response{Success: false, Message: "Incorrect password"}
		json.NewEncoder(response).Encode(_response)

		return
	}

	//If error is not found, then password is termed correct
	//Move ahead to generate the JWT token
	token, expirationTime := utils.GenerateJWT(user)

	//Save in the cookies
	http.SetCookie(response, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

	//Send the response
	_response := models.Response{Success: true, Data: bson.M{"user": user, "token": token}, Message: "Logged in successfully"}
	json.NewEncoder(response).Encode(_response)

}
