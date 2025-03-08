package model

type Cast struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Image string `json:"image,omitempty" bson:"image,omitempty"`
}
