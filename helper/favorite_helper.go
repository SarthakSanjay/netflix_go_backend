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

func AddToFavorite(profileId string, contentId string) (primitive.ObjectID, error) {
	pId, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return primitive.NilObjectID, err
	}
	cId, err := primitive.ObjectIDFromHex(contentId)
	if err != nil {
		return primitive.NilObjectID, err
	}

	doc := model.Favorite{
		ProfileId: pId,
		ContentId: cId,
		AddedOn:   time.Now(),
	}
	result, err := db.FavoriteCollection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Println("Error adding to favorite", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func RemoveFromFavorite(profileId string, contentId string) (int64, error) {
	pId, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return 0, err
	}
	cId, err := primitive.ObjectIDFromHex(contentId)
	if err != nil {
		return 0, err
	}

	filter := bson.M{
		"profileId": pId,
		"contentId": cId,
	}
	result, err := db.FavoriteCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error adding to favorite")
		return 0, err
	}

	return result.DeletedCount, nil
}

func GetUserFavoriteFromProfile(profileId string) ([]model.Movies, error) {
	pId, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		log.Printf("Invalid profileId: %v\n", err)
		return nil, err
	}

	filter := bson.M{"profileId": pId}
	cursor, err := db.FavoriteCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var favoriteItems []model.Favorite
	if err := cursor.All(context.Background(), &favoriteItems); err != nil {
		return nil, err
	}

	if len(favoriteItems) == 0 {
		return []model.Movies{}, nil
	}

	var movieIDs []primitive.ObjectID
	for _, item := range favoriteItems {
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
