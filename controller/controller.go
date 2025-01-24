package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/sarthaksanjay/netflix-go/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// first make a connection string give by mongodb
//

const connectionString = ""

const (
	dbName = "netflix-go"
)

// create collection
var moviesCollection *mongo.Collection

func inti() {
	// create client options

	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB conneciton established successfully")

	// collection instances

	watchlistCollection := client.Database(dbName).Collection("watchlist")
	moviesCollection := client.Database(dbName).Collection("movies")

	fmt.Println("Collection instance/ref are now ready!!")
}

func insertMovie(movie model.Movies) {
	inserted, err := moviesCollection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in the db with id :", inserted.InsertedID)
}
