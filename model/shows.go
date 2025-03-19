package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Show struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name              string             `json:"name,omitempty" bson:"name,omitempty"`
	Description       string             `json:"description,omitempty" bson:"description,omitempty"`
	Image             Images             `json:"image,omitempty" bson:"image,omitempty"`
	Genre             []string           `json:"genre,omitempty" bson:"genre,omitempty"`
	ReleasedOn        int                `json:"releasedOn,omitempty" bson:"releasedOn,omitempty"`
	Rating            float64            `json:"rating,omitempty" bson:"rating,omitempty"`
	Cast              []Cast             `json:"cast,omitempty" bson:"cast,omitempty"`
	Seasons           []Season           `json:"seasons,omitempty" bson:"seasons,omitempty"`
	Director          string             `json:"director,omitempty" bson:"director,omitempty"`
	TrailerUrl        string             `json:"trailerUrl,omitempty" bson:"trailerUrl,omitempty"`
	IsFeatured        bool               `json:"isFeatured,omitempty" bson:"isFeatured,omitempty"`
	Tags              []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Availability      []string           `json:"availability,omitempty" bson:"availability,omitempty"`
	AgeRating         string             `json:"ageRating,omitempty" bson:"ageRating,omitempty"`
	Views             int64              `json:"views,omitempty" bson:"views,omitempty"`
	AudioLanguages    []string           `json:"audioLanguages,omitempty" bson:"audioLanguages,omitempty"`
	SubtitleLanguages []string           `json:"subtitleLanguages,omitempty" bson:"subtitleLanguages,omitempty"`
	AddedDate         time.Time          `json:"addedDate,omitempty" bson:"addedDate,omitempty"`
}

type Season struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	number   int8               `json:"number,omitempty" bson:"number,omitempty"`
	episodes []Episode          `json:"episodes,omitempty" bson:"episodes,omitempty"`
}

type Episode struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	thumbnail   string             `json:"thumbnail,omitempty" bson:"thumbnail,omitempty"`
	title       string             `json:"title,omitempty" bson:"title,omitempty"`
	description string             `json:"description,omitempty" bson:"description,omitempty"`
	Duration    int                `json:"duration,omitempty" bson:"duration,omitempty"`
}
