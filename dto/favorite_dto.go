package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type FavoriteRequestDto struct {
	ProfileId primitive.ObjectID `json:"profileId" bson:"profileId"`
	ContentId primitive.ObjectID `json:"contentId" bson:"contentId"`
}
