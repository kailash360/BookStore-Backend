package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/kailash360/BookStore-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func Connect() {
	godotenv.Load(".env")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("Error in connecting to database: ", err)
		panic(err)
	}

	database = client.Database(os.Getenv("MONGO_DATABASE_NAME"))
	fmt.Println("Connected to database successfully")

}

func GetClient() *mongo.Database {
	if database == nil {
		Connect()
	}
	return database
}

func GenerateJWT(user models.User) (string, time.Time) {

	//Generate the expiration time
	expirationTime := time.Now().Add(25 * time.Minute)

	//Create the claims of the token
	claims := &models.Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//Declare the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Create the Token
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString, expirationTime
}

func VerifyToken(tokenString string) (bool, primitive.ObjectID) {

	//Create the claims
	claims := &models.Claims{}

	//Populate the claims from the token
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Print(err.Error())
		return false, primitive.NilObjectID
	}

	//Check if object ID is present in the database
	isPresent := database.Collection("users").FindOne(context.TODO(), bson.D{{Key: "_id", Value: claims.ID}})

	//If not present, return an error
	if isPresent == nil {
		return false, primitive.NilObjectID
	}

	return true, claims.ID
}
