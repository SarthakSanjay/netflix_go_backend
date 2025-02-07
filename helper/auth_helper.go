package helper

import (
	"context"
	"log"
	"time"

	"github.com/sarthaksanjay/netflix-go/db"
	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RefreshToken(userId primitive.ObjectID) (string, string, error) {
	var user model.User
	filter := bson.M{"_id": userId}
	err := db.UserCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Printf("Error finding user %v\n", err)
		return "", "", err
	}

	accessToken, err := services.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		log.Printf("Error generating access token  %v\n", err)
		return "", "", err

	}
	refreshToken, err := services.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		log.Printf("Error generating refresh token  %v\n", err)
		return "", "", err
	}

	update := bson.M{
		"$set": bson.M{
			"refresh_tokens": model.RefreshToken{
				Token:     refreshToken,
				ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
			},
		},
	}

	_, err = db.UserCollection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, update)
	if err != nil {
		log.Printf("Error updating refresh token %v\n", err)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
