package config

import (
	"context"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB() *mongo.Client {
	MONGO_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	MongoClient = client
	return client
}

func DisconnectMongoDB(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Printf("Error disconnecting MongoDB: %v", err)
	}
	log.Println("Disconnected from MongoDB.")
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database("golangDb").Collection(collectionName)
}
