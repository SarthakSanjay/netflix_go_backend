package helper

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

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

func UpdateMovie(movieId string, updates model.Movies) (int64, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Printf("Invalid movieId %v\n", err)
		return 0, err
	}

	updateFields := bson.M{}

	if updates.Name != "" {
		updateFields["name"] = updates.Name
	}

	if updates.Description != "" {
		updateFields["description"] = updates.Description
	}

	if len(updates.Genre) > 0 {
		updateFields["genre"] = updates.Genre
	}

	if updates.ReleasedOn != 0 {
		updateFields["releasedOn"] = updates.ReleasedOn
	}

	if updates.Duration != 0 {
		updateFields["duration"] = updates.Duration
	}

	if updates.Rating != 0 {
		updateFields["rating"] = updates.Rating
	}

	// if len(updates.Language) > 0 {
	// 	updateFields["language"] = updates.Language
	// }

	if len(updates.Cast) > 0 {
		updateFields["cast"] = updates.Cast
	}

	if updates.Director != "" {
		updateFields["director"] = updates.Director
	}

	if updates.TrailerUrl != "" {
		updateFields["trailer"] = updates.TrailerUrl
	}

	if len(updates.Tags) > 0 {
		updateFields["tags"] = updates.Tags
	}

	if len(updates.Availablity) > 0 {
		updateFields["availablity"] = updates.Availablity
	}

	if updates.AgeRating != "" {
		updateFields["ageRating"] = updates.AgeRating
	}

	if updates.Views != 0 {
		updateFields["views"] = updates.Views
	}

	if len(updates.AudioLanguages) > 0 {
		updateFields["audioLanguages"] = updates.AudioLanguages
	}

	if len(updates.SubtitleLanguages) > 0 {
		updateFields["subtitleLanguages"] = updates.SubtitleLanguages
	}

	updateFields["addedOn"] = time.Now()

	if len(updateFields) == 0 {
		return 0, errors.New("no fields to update")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateFields}

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

func GetAllMovie(ctx context.Context) ([]model.Movies, error) {
	cursor, err := db.MoviesCollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Printf("Error finding movies %v\n", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var movies []model.Movies

	for cursor.Next(context.Background()) {
		var movie model.Movies
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

func SearchMovie(searchQuery string) ([]model.Movies, error) {
	if searchQuery == "" {
		log.Println("Search query is empty")
		return nil, fmt.Errorf("search query is empty")
	}
	filter := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"description": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"genre": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"language": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"tags": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"director": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"cast": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"audioLanguages": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"subtitleLanguages": bson.M{"$regex": searchQuery, "$options": "i"}},
		},
	}

	var movies []model.Movies

	cursor, err := db.MoviesCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Printf("Error decoding movie %v\n", err)
			continue
		}

		movies = append(movies, movie)

	}

	return movies, nil
}

func PopularMovie() ([]model.Movies, error) {
	var movies []model.Movies
	filter := bson.M{
		"rating": bson.M{"$gt": 3},
	}
	cursor, err := db.MoviesCollection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("No movie found %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Printf("Error decoding movie %v\n", err)
			continue
		}

		movies = append(movies, movie)
	}
	return movies, nil
}

func GetMovieByGenre(genre string) ([]model.Movies, error) {
	var movies []model.Movies
	filter := bson.M{"genre": genre}
	cursor, err := db.MoviesCollection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("No movie found %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Printf("Error decoding movie %v\n", err)
			continue
		}

		movies = append(movies, movie)
	}
	return movies, nil
}

//
// func RecommendedMovie() {
// }

func SimilarMovie(genres []string) ([]model.Movies, error) {
	var movies []model.Movies
	filter := bson.M{
		"genre": bson.M{"$in": genres},
	}

	cursor, err := db.MoviesCollection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Error finding movie %v\n", err)
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatalf("Error decoding movie%v\n", err)
			continue
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
