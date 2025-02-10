package helper

import (
	"context"
	"log"
	"time"

	"github.com/sarthaksanjay/netflix-go/db"
	"github.com/sarthaksanjay/netflix-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddMovieToWatchlist(movieId string, profileId string) (primitive.ObjectID, error) {
	mId, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Printf("Invalid movieId %v\n", err)
		return primitive.NilObjectID, err
	}
	uId, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		log.Printf("Invalid userId %v\n", err)
		return primitive.NilObjectID, err
	}

	var watchlist model.Watchlist

	watchlist.ProfileId = uId
	watchlist.ContentId = mId
	watchlist.AddedAt = time.Now()

	insertedMovie, err := db.WatchlistCollection.InsertOne(context.Background(), watchlist)
	if err != nil {
		log.Printf("Error inserting movie %v\n", err)
		return primitive.NilObjectID, err
	}

	return insertedMovie.InsertedID.(primitive.ObjectID), nil
}

func GetAllMovieFromUserWatchlist(profileId string) ([]model.Movies, error) {
	pId, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		log.Printf("Invalid profileId: %v\n", err)
		return nil, err
	}

	filter := bson.M{"profileId": pId}
	cursor, err := db.WatchlistCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var watchlistItems []model.Watchlist
	if err := cursor.All(context.Background(), &watchlistItems); err != nil {
		return nil, err
	}

	if len(watchlistItems) == 0 {
		return []model.Movies{}, nil
	}

	var movieIDs []primitive.ObjectID
	for _, item := range watchlistItems {
		movieIDs = append(movieIDs, item.ContentId)
	}

	movieFilter := bson.M{
		"_id": bson.M{"$in": movieIDs},
	}
	movieCursor, err := db.MoviesCollection.Find(context.Background(), movieFilter)
	if err != nil {
		return nil, err
	}

	defer movieCursor.Close(context.Background())

	var movies []model.Movies
	if err := movieCursor.All(context.Background(), &movies); err != nil {
		return nil, nil
	}

	return movies, nil
}

func DeleteMovieFromWatchlist(profileId string, movieId string) (bson.M, error) {
	pID, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	mID, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"profileId": pID,
		"contentId": mID,
	}

	var deletedDoc bson.M
	err = db.WatchlistCollection.FindOneAndDelete(context.Background(), filter).Decode(&deletedDoc)
	if err != nil {
		log.Println("Error deleting movie from watchlist", err)
		return bson.M{}, err
	}

	return deletedDoc, nil
}
