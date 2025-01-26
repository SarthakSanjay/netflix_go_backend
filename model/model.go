package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username     string             `json:"username,omitempty"`
	Email        string             `json:"email,omitempty"`
	PhoneNo      int                `json:"_"`
	ProfileImage string             `json:"image,omitempty"`
	Wishlist     Wishlist
	History      History
}

type Movies struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Image       Images             `json:"image,omitempty"`
}

type Images struct {
	Main   string `json:"main,omitempty"`
	Poster string `json:"poster,omitempty"`
}

type Wishlist struct {
	UserId primitive.ObjectID `json:"_id,omitempty"`
	Movies []Movies           `json:"movies,omitempty"`
}

type History struct {
	UserId primitive.ObjectID `json:"_id,omitempty"`
	Movies []Movies           `json:"movies,omitempty"`
}
