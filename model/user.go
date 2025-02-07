package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

var validRoles = map[Role]bool{
	RoleAdmin: true,
	RoleUser:  true,
}

type User struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username         string             `json:"username,omitempty"`
	Email            string             `json:"email,omitempty"`
	Password         string             `json:"-"`
	PhoneNo          string             `json:"phoneNo,omitempty" bson:"phoneNo,omitempty"`
	Role             Role               `json:"role,omitempty"`
	CreatedAt        time.Time          `json:"createdAt,omitempty"`
	UpdatedAt        time.Time          `json:"updatedAt,omitempty"`
	Profiles         []Profile          `json:"profiles,omitempty"`
	Subscription     Subscription       `json:"-"`
	RefreshTokens    RefreshToken       `json:"refresh_tokens,omitempty" bson:"refresh_tokens,omitempty"`
	ResetToken       string             `json:"resetToken,omitempty"`
	ResetTokenExpiry time.Time          `json:"resetTokenExp,omitempty"`
}

type RefreshToken struct {
	Token     string    `bson:"token"`
	ExpiresAt time.Time `bson:"expires_at"`
}

type Profile struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Avatar    string             `json:"avtar,omitempty"`
	Watchlist []Watchlist        `json:"watchlist,omitempty"`
	History   []History          `json:"history,omitempty"`
}

func (r Role) IsValid() bool {
	return validRoles[r]
}
