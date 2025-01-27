package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name              string             `json:"name,omitempty"`
	Description       string             `json:"description,omitempty"`
	Image             Images             `json:"image,omitempty"`
	Genre             []string           `json:"genre,omitempty"`
	ReleasedOn        int                `json:"releasedOn,omitempty"`
	Duration          int                `json:"duration,omitempty"`
	Rating            float64            `json:"rating,omitempty"`
	Language          []string           `json:"language,omitempty"`
	Cast              []string           `json:"cast,omitempty"`
	Director          string             `json:"director,omitempty"`
	TrailerUrl        string             `json:"trailer,omitempty"`
	IsFeatured        bool               `json:"isFeatured,omitempty"`
	Tags              []string           `json:"tags,omitempty"`
	Availablity       []string           `json:"availablity,omitempty"`
	AgeRating         string             `json:"ageRating,omitempty"`
	Views             int64              `json:"views,omitempty"`
	AudioLanguages    []string           `json:"audioLanguages,omitempty"`
	SubtitleLanguages []string           `json:"subtitleLanguages,omitempty"`
	AddedDate         time.Time          `json:"addedOn,omitempty"`
}

type Images struct {
	MovieID      primitive.ObjectID `json:"movieId,omitempty"`
	ThumbnailUrl string             `json:"thumbnailUrl,omitempty"`
	Screenshots  []string           `json:"screenshots,omitempty"`
	Poster       string             `json:"poster,omitempty"`
}

type Wishlist struct {
	UserId primitive.ObjectID `json:"_id,omitempty"`
	Movies []Movies           `json:"movies,omitempty"`
}

type History struct {
	UserId primitive.ObjectID `json:"_id,omitempty"`
	Movies []Movies           `json:"movies,omitempty"`
}
