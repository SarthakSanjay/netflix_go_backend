package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddMovieDTO struct {
	ProfileId primitive.ObjectID `json:"profileId"`
	MovieId   primitive.ObjectID `json:"movieId"`
}
