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

func GetUserFavoriteFromProfile(profileId string) ([]model.Favorite, error) {
	pId, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"profileId": pId,
	}
	cursor, err := db.FavoriteCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Error fetching user favorite")
		return nil, err
	}

	var favorites []model.Favorite
	if err := cursor.All(context.Background(), &favorites); err != nil {
		log.Println("Error decoding user favorites")
		return nil, err
	}

	if len(favorites) == 0 {
		return []model.Favorite{}, nil
	}

	return favorites, nil
}
