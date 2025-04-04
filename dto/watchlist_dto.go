package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddContentDTO struct {
	ProfileId primitive.ObjectID `json:"profileId"`
	ContentId primitive.ObjectID `json:"contentId"`
}
