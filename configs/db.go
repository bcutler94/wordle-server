package configs

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
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

	db := os.Getenv("MONGODB_DB")
	if db == "" {
		log.Fatal("No MONGODB_DB env var found")
	}

	log.Print("Connected to DB: ", db)
	return client.Database(db)
}

var DB *mongo.Database = Connect()
