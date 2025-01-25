package utils

import "go.mongodb.org/mongo-driver/mongo"

var (
	MoviesCollection    *mongo.Collection
	WatchlistCollection *mongo.Collection
)
