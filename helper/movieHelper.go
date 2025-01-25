package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertMovie(movie model.Movies) {
	inserted, err := utils.MoviesCollection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in the db with id :", inserted.InsertedID)
}

func UpdateMovie(movieId string, key string, value string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{key: value}}

	result, err := utils.MoviesCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated movie", result.ModifiedCount)
}

func DeleteMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	result, err := utils.MoviesCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted movie ", result.DeletedCount)
}

func GetAllMovie() []bson.M {
	cursor, err := utils.MoviesCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}
