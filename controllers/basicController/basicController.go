package basicController

import (
	"encoding/json"
	"net/http"

	"github.com/kailash360/BookStore-Backend/models"
)

func HandleHome(response http.ResponseWriter, request *http.Request) {
	_response := models.Response{Success: true, Message: "Server is running successfully"}
	json.NewEncoder(response).Encode(_response)
}

func HandleIncorrectRoute(response http.ResponseWriter, request *http.Request) {
	_response := models.Response{Success: false, Message: "Incorrect API endpoint"}
	json.NewEncoder(response).Encode(_response)
}
