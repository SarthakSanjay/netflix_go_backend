package helper

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sarthaksanjay/netflix-go/db"
	"github.com/sarthaksanjay/netflix-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddContentToWatchlist(contentId string, profileId string, contentType string) (primitive.ObjectID, error) {
	mId, err := primitive.ObjectIDFromHex(contentId)
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
	watchlist.ContentType = contentType
	watchlist.AddedAt = time.Now()

	insertedMovie, err := db.WatchlistCollection.InsertOne(context.Background(), watchlist)
	if err != nil {
		log.Printf("Error inserting movie %v\n", err)
		return primitive.NilObjectID, err
	}

	return insertedMovie.InsertedID.(primitive.ObjectID), nil
}

func GetAllContentFromUserWatchlist(profileId string, contentType string) (interface{}, error) {
	pId, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		log.Printf("Invalid profileId: %v\n", err)
		return nil, err
	}

	filter := bson.M{
		"profileId":   pId,
		"contentType": contentType,
	}
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
		// Return empty array based on contentType
		if contentType == "movie" {
			return []model.Movies{}, nil
		}
		return []model.Show{}, nil
	}

	var contentIDs []primitive.ObjectID
	for _, item := range watchlistItems {
		contentIDs = append(contentIDs, item.ContentId)
	}

	contentFilter := bson.M{
		"_id": bson.M{"$in": contentIDs},
	}

	switch contentType {
	case "movie":
		movieCursor, err := db.MoviesCollection.Find(context.Background(), contentFilter)
		if err != nil {
			return nil, err
		}
		defer movieCursor.Close(context.Background())

		var movies []model.Movies
		if err := movieCursor.All(context.Background(), &movies); err != nil {
			return nil, err
		}
		return movies, nil

	case "show":
		showCursor, err := db.ShowsCollection.Find(context.Background(), contentFilter)
		if err != nil {
			return nil, err
		}
		defer showCursor.Close(context.Background())

		var shows []model.Show
		if err := showCursor.All(context.Background(), &shows); err != nil {
			return nil, err
		}
		return shows, nil

	default:
		return nil, fmt.Errorf("unsupported contentType: %s", contentType)
	}
}

func DeleteContentFromWatchlist(profileId string, contentId string) (bson.M, error) {
	pID, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	mID, err := primitive.ObjectIDFromHex(contentId)
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

func DeleteAllContentFromWatchlist(profileId string, contentType string) (int, error) {
	pID, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"profileId": pID, "contentType": contentType}
	result, err := db.WatchlistCollection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting content from watchlist")
		return 0, err
	}

	return int(result.DeletedCount), nil
}
