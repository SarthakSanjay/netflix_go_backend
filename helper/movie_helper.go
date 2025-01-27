package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/sarthaksanjay/netflix-go/db"
	"github.com/sarthaksanjay/netflix-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertMovie(movie model.Movies) (primitive.ObjectID, error) {
	inserted, err := db.MoviesCollection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Printf("Error inserting movie : %v\n", err)
	}

	fmt.Println("Inserted 1 movie in the db with id :", inserted.InsertedID)
	return inserted.InsertedID.(primitive.ObjectID), nil
}

func UpdateMovie(movieId string, updates map[string]interface{}) (int64, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Printf("Invalid movieId %v\n", err)
		return 0, err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updates}

	result, err := db.MoviesCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Unable to update movie : %v\n ", err)
		return 0, err
	}
	fmt.Println("updated movie", result.ModifiedCount)
	return result.ModifiedCount, nil
}

func DeleteMovie(movieId string) (int64, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Printf("Invalid movieId : %v\n", err)
		return 0, err
	}

	filter := bson.M{"_id": id}

	result, err := db.MoviesCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Printf("Error deleting movie : %v\n", err)
		return 0, err
	}

	fmt.Println("Deleted movie ", result.DeletedCount)

	return result.DeletedCount, nil
}

func DeleteAllMovie() (int64, error) {
	movie, err := db.MoviesCollection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Printf("Error delete all movies : %v\n", err)
		return 0, err
	}
	fmt.Println("Deleted all movies")
	return movie.DeletedCount, nil
}

func GetAllMovie(ctx context.Context) ([]bson.M, error) {
	cursor, err := db.MoviesCollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Printf("Error finding movies %v\n", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var movies []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Printf("Error decoding movie %v\n", err)
			continue
		}

		movies = append(movies, movie)
	}
	if err := cursor.Err(); err != nil {
		log.Printf("Cursor iteration err: %v\n", err)
	}
	return movies, nil
}

func GetMovieById(movieId string) (*model.Movies, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Printf("Invalid movie id %v\n", err)
	}
	filter := bson.M{"_id": id}

	var movie model.Movies
	err = db.MoviesCollection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		log.Printf("Movie not found%v\n", err)
	}

	return &movie, nil
}

// func SearchMovie() {
// }
//
// func PopularMovie() {
// }
//
// func RecommendedMovie() {
// }
//
// func SimilarMovie() {
// }
