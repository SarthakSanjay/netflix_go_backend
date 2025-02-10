package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddMovieDTO struct {
	ProfileId primitive.ObjectID `json:"profileId"`
	ContentId primitive.ObjectID `json:"contentId"`
}
