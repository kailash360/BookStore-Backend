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
	expirationTime := time.Now().Add(5 * time.Minute)

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

func VerifyToken(tokenString string) bool {

	//Create the claims
	claims := &models.Claims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_SECRET"), nil
	})

	return err == nil
}
