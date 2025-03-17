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
	movie.AddedDate = time.Now()
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

	if len(updates.Availability) > 0 {
		updateFields["availablity"] = updates.Availability
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
