package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/sarthaksanjay/netflix-go/db"
	"github.com/sarthaksanjay/netflix-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllMovie(ctx context.Context) ([]model.Movies, error) {
	cursor, err := db.MoviesCollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Printf("Error finding movies %v\n", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var movies []model.Movies

	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Printf("Error decoding movie %v\n", err)
			continue
		}

		movies = append(movies, movie)
	}
	if err := cursor.Err(); err != nil {
		log.Printf("Cursor iteration err: %v\n", err)
	}
	return movies, nil
}

func GetMovieById(movieId string) (*model.Movies, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Printf("Invalid movie id %v\n", err)
	}
	filter := bson.M{"_id": id}

	var movie model.Movies
	err = db.MoviesCollection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		log.Printf("Movie not found%v\n", err)
	}

	return &movie, nil
}

func SearchMovie(searchQuery string) ([]model.Movies, error) {
	if searchQuery == "" {
		log.Println("Search query is empty")
		return nil, fmt.Errorf("search query is empty")
	}
	filter := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"description": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"genre": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"language": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"tags": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"director": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"cast": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"audioLanguages": bson.M{"$regex": searchQuery, "$options": "i"}},
			{"subtitleLanguages": bson.M{"$regex": searchQuery, "$options": "i"}},
		},
	}

	var movies []model.Movies

	cursor, err := db.MoviesCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Printf("Error decoding movie %v\n", err)
			continue
		}

		movies = append(movies, movie)

	}

	return movies, nil
}

func PopularMovie() ([]model.Movies, error) {
	var movies []model.Movies
	filter := bson.M{
		"rating": bson.M{"$gt": 8},
	}
	cursor, err := db.MoviesCollection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("No movie found %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Printf("Error decoding movie %v\n", err)
			continue
		}

		movies = append(movies, movie)
	}
	return movies, nil
}

func GetMovieByGenre(genre string, limit int) ([]model.Movies, error) {
	var movies []model.Movies
	filter := bson.M{"genre": genre}
	cursor, err := db.MoviesCollection.Find(context.Background(), filter, options.Find().SetLimit(int64(limit)))
	if err != nil {
		log.Printf("No movie found %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Printf("Error decoding movie %v\n", err)
			continue
		}

		movies = append(movies, movie)
	}
	return movies, nil
}

//
// func RecommendedMovie() {
// }

func SimilarMovie(genres []string) ([]model.Movies, error) {
	var movies []model.Movies
	filter := bson.M{
		"genre": bson.M{"$in": genres},
	}

	cursor, err := db.MoviesCollection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Error finding movie %v\n", err)
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatalf("Error decoding movie%v\n", err)
			continue
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func TrendingMovie() ([]model.Movies, error) {
	var movies []model.Movies
	filter := bson.M{
		"isFeatured": true,
	}

	cursor, err := db.MoviesCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie model.Movies
		err := cursor.Decode(&movie)
		if err != nil {
			log.Println("Error decoding movie")
			continue
		}
		movies = append(movies, movie)
	}
	log.Println("trending movies", movies)
	return movies, nil
}
