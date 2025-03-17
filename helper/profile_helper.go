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

func GetAllUserProfiles(userId string) ([]model.Profile, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Println("Invalid UserId")
		return nil, err
	}

	filter := bson.M{"userId": id}
	cursor, err := db.ProfileCollection.Find(context.Background(), filter)
	if err != nil {
		log.Println("Unable to find profiles with given userId")
		return nil, err
	}

	var profiles []model.Profile
	if err := cursor.All(context.Background(), &profiles); err != nil {
		return nil, err
	}

	return profiles, nil
}

func UpdateProfile(profileId string, updates model.Profile) (int, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		log.Println("Invalid ProfileId")
		return 0, err
	}
	log.Println("✅updates✅", updates)
	newProfile := bson.M{}
	if updates.Name != "" {
		newProfile["name"] = updates.Name
	}
	if updates.Avatar != "" {
		newProfile["avatar"] = updates.Avatar
	}

	newProfile["updatedAt"] = time.Now()

	update := bson.M{"$set": newProfile}
	result, err := db.ProfileCollection.UpdateByID(context.Background(), id, update)
	if err != nil {
		log.Println("Error updating user profile", err)
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

func DeleteProfile(profileId string) (int, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		log.Println("Invalid ProfileId")
		return 0, err
	}

	filter := bson.M{"_id": id}
	result, err := db.ProfileCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Unable to delete user profile")
		return 0, err
	}

	return int(result.DeletedCount), nil
}
