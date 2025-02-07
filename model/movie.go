package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status int

const (
	Active Status = iota
	Inactive
	Cancelled
	Suspended
)

func (s Status) String() string {
	return [...]string{"Active", "Inactive", "Cancelled", "Suspended"}[s]
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

type Watchlist struct {
	UserId    primitive.ObjectID `json:"userId,omitempty"`
	ContentId primitive.ObjectID `json:"contentId,omitempty"`
	AddedAt   time.Time          `json:"addedAt,omitempty"`
}

type History struct {
	UserId    primitive.ObjectID `json:"userId,omitempty"`
	ContentId primitive.ObjectID `json:"contentId,omitempty"`
	AddedAt   time.Time          `json:"addedAt,omitempty"`
}

type Subscription struct {
	Plan      string    `json:"plan,omitempty"`
	Status    Status    `json:"status,omitempty"`
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty"`
	Payments  []Payment `json:"payments,omitempty"`
}

type Payment struct {
	PaymentDate   time.Time `json:"paymentDate,omitempty"`
	Amount        float64   `json:"amount,omitempty"`
	PaymentMethod string    `json:"paymentMethod,omitempty"`
}

type Playback struct {
	ContentId   primitive.ObjectID `json:"contentId,omitempty"`
	Progress    float64            `json:"progress,omitempty"`
	LastWatched time.Time          `json:"lastWatched,omitempty"`
}

type Review struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	ContentId primitive.ObjectID `json:"contentId,omitempty"`
	UserId    primitive.ObjectID `json:"userId,omitempty"`
	Text      string             `json:"text,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty"`
}

type Rating struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	ContentId primitive.ObjectID `json:"contentId,omitempty"`
	UserId    primitive.ObjectID `json:"userId,omitempty"`
	Rating    int                `json:"rating,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty"`
}
