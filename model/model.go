package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movies struct {
	ID          primitive.ObjectID `json:"_id",omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"name,omitempty"`
	Image       Images             `json:"image,omitempty"`
}

type Images struct {
	Main   string `json:"main,omitempty"`
	Poster string `json:"poster,omitempty"`
}
