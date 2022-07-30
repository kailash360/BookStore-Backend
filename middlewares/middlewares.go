package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kailash360/BookStore-Backend/models"
	"github.com/kailash360/BookStore-Backend/utils"
)

func IsAuthroized(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("Entering to check authorization")
	return func(response http.ResponseWriter, request *http.Request) {

		//Get the JWT token from the headers
		token := request.Header.Get("token")
		if token == "" {
			_response := models.Response{Success: false, Message: "Token not provided"}
			json.NewEncoder(response).Encode(_response)
		}

		//Verify the token
		verified, userId := utils.VerifyToken(token)

		//If token is not valid, return an error
		if !verified {
			_response := models.Response{Success: false, Message: "User is not authorized"}
			json.NewEncoder(response).Encode(_response)

			return
		}

		//set the userId in request header t pass to callbacks
		request.Header.Set("userId", userId.Hex())

		//If valid move to the callback
		next(response, request)
	}
}
