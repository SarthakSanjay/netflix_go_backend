package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cast struct {
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Image  string             `json:"image,omitempty" bson:"image,omitempty"`
	Gender string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Age    string             `json:"age,omitempty" bson:"age,omitempty"`
	Bio    string             `json:"bio,omitempty" bson:"bio,omitempty"`
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
