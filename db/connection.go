package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName = "netflix-go"
)

var (
	MoviesCollection    *mongo.Collection
	WatchlistCollection *mongo.Collection
	UserCollection      *mongo.Collection
	ProfileCollection   *mongo.Collection
	FavoriteCollection  *mongo.Collection
	client              *mongo.Client // Store the client to close it later
)

// ConnectDB initializes the database connection
func ConnectDB() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Get MongoDB connection string
	dbURI := os.Getenv("DB_CONNECTION_STRING")
	if dbURI == "" {
		log.Fatal("DB_CONNECTION_STRING is not set in .env")
	}

	// Create client options
	clientOptions := options.Client().ApplyURI(dbURI)

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Verify connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	fmt.Println("✅ MongoDB connection established successfully!")

	// Assign collections
	WatchlistCollection = client.Database(dbName).Collection("watchlist")
	MoviesCollection = client.Database(dbName).Collection("movies")
	UserCollection = client.Database(dbName).Collection("user")
	ProfileCollection = client.Database(dbName).Collection("profile")
	FavoriteCollection = client.Database(dbName).Collection("favorite")

	fmt.Println("✅ Collection instances are ready!")
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB() {
	if client != nil {
		err := client.Disconnect(context.TODO())
		if err != nil {
			log.Fatal("Error closing MongoDB connection:", err)
		}
		fmt.Println("✅ MongoDB connection closed.")
	}
}
