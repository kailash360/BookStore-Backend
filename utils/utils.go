package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
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
