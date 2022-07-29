package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("Error in connecting to database: ", err)
		panic(err)
	}

	client.Database("Cluster0")
	fmt.Println("Connected to database successfully")

	return client
}
