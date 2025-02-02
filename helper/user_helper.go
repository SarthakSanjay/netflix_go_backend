package helper

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sarthaksanjay/netflix-go/db"
	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/services"
	"github.com/sarthaksanjay/netflix-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user model.User) (interface{}, string, error) { // register

	var existingUser model.User

	filter := bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}

	err := db.UserCollection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Printf("error finding user: %v\n", err)
		return "error finding user", "", err
	}

	fmt.Println("existingUser", existingUser)
	if err == nil { // User exists
		log.Printf("user already exists: %s\n", user.Email)
		return "user already exists", "", err
	}

	password, err := utils.HashedPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing the password %v", err)
		return "error hashing password", "", err
	}

	user.Password = password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := db.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Error creating user %v\n", err)
		return "error creating user", "", err
	}

	accessToken, err := services.GenerateAccessToken(result.InsertedID.(primitive.ObjectID), user.Email)
	if err != nil {
		log.Printf("Error generation access token %v\n", err)
		return "error generating access token", "", err
	}

	refreshToken, err := services.GenerateRefreshToken(result.InsertedID.(primitive.ObjectID), user.Email)
	if err != nil {
		log.Printf("Error generation access token %v\n", err)
		return "error generating refresh token", "", err

	}

	update := bson.M{
		"$set": bson.M{
			"refresh_tokens": []model.RefreshToken{
				{Token: refreshToken, ExpiresAt: time.Now().Add(7 * 24 * time.Hour)},
			},
		},
	}

	_, err = db.UserCollection.UpdateOne(context.Background(), bson.M{"_id": result.InsertedID}, update)
	if err != nil {
		log.Printf("Error updating refresh token %v\n", err)
		return "Error updating refresh token", "", err
	}

	return result.InsertedID, accessToken, nil
}

// login
func LoginUser(user model.User) map[string]interface{} {
	filter := bson.M{"email": user.Email}
	var userF model.User
	err := db.UserCollection.FindOne(context.Background(), filter).Decode(&userF)
	if err != nil {
		log.Printf("Invalid email or password%v\n", err)
		return map[string]interface{}{"error": "Invalid email or password"}
	}
	hashedPassword := userF.Password
	if !utils.ComparePassword(hashedPassword, user.Password) {
		log.Printf("Incorrect password ")
		return map[string]interface{}{"error": "password is incorrect"}

	}

	accessToken, err := services.GenerateAccessToken(userF.ID, user.Email)
	if err != nil {
		log.Printf("Error generation access token %v\n", err)
		return map[string]interface{}{"error": "error generating accessToken"}
	}

	refreshToken, err := services.GenerateRefreshToken(userF.ID, user.Email)
	if err != nil {
		log.Printf("Error generation refresh token %v\n", err)
		return map[string]interface{}{"error": "error generating refreshToken"}

	}

	update := bson.M{
		"$push": bson.M{
			"refresh_tokens": []model.RefreshToken{
				{Token: refreshToken, ExpiresAt: time.Now().Add(7 * 24 * time.Hour)},
			},
		},
	}

	_, err = db.UserCollection.UpdateOne(context.Background(), bson.M{"_id": userF.ID}, update)
	if err != nil {
		log.Printf("Error updating refresh token %v\n", err)
		return map[string]interface{}{"error": "error updating refresh token"}
	}
	return map[string]interface{}{
		"message":      "success",
		"isLoggedIn":   true,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
}

func UpdateUser() {
}

func DeleteUser() {
}

func DeleteAllUser() {
}

func GetUser() {
}

func GetAllUser() []model.User {
	cursor, err := db.UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []model.User

	for cursor.Next(context.Background()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}
