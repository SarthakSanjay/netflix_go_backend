package helper

import (
	"context"
	"log"

	"github.com/sarthaksanjay/netflix-go/db"
	"github.com/sarthaksanjay/netflix-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUserProfile(newProfile model.Profile) (primitive.ObjectID, error) {
	result, err := db.ProfileCollection.InsertOne(context.Background(), newProfile)
	if err != nil {
		log.Println("Error creating profile", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func GetProfileById(profileId string) (model.Profile, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		log.Println("Invalid profileId")
		return model.Profile{}, err
	}
	var profile model.Profile
	filter := bson.M{"_id": id}
	err = db.ProfileCollection.FindOne(context.Background(), filter).Decode(&profile)
	if err != nil {
		log.Println("Unable to find user profile")
		return model.Profile{}, err
	}

	return profile, nil
}
