package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	helpers "github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// first make a connection string give by mongodb
//

const (
	dbName = "netflix-go"
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

	utils.WatchlistCollection = client.Database(dbName).Collection("watchlist")
	utils.MoviesCollection = client.Database(dbName).Collection("movies")

	fmt.Println("Collection instance/ref are now ready!!")
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	allMovies := helpers.GetAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "POST")

	var movie model.Movies
	json.NewDecoder(r.Body).Decode(&movie)
	helpers.InsertMovie(movie)
	json.NewEncoder(w).Encode(movie)
}
