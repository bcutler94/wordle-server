package configs

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("No MONGODB_URI env var found")
	}

	connectOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), connectOptions)
	if err != nil {
		panic(err)
	}

	log.Print("Connected to DB")
	return client
}

var DBClient *mongo.Client = Connect()
