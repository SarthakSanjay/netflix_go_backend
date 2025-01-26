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

// first make a connection string give by mongodb
//

const (
	dbName = "netflix-go"
)

var (
	MoviesCollection    *mongo.Collection
	WatchlistCollection *mongo.Collection
)

func init() {
	// create client options
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	clientOption := options.Client().ApplyURI(os.Getenv("DB_CONNECTION_STRING"))

	// connect to mongodb

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB conneciton established successfully")

	// collection instances

	WatchlistCollection = client.Database(dbName).Collection("watchlist")
	MoviesCollection = client.Database(dbName).Collection("movies")

	fmt.Println("Collection instance/ref are now ready!!")
}
