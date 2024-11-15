package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}

	MONGO_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatalf("Error connecting to MongoDB", err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatalf("Error pinging MongoDB", err)
	}

	fmt.Println("Connected to MongoDB!")



}
