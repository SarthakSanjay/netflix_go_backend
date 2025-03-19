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

func GetAllShows(ctx context.Context) ([]model.Show, error) {
	cursor, err := db.ShowsCollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Printf("Error finding shows %v\n", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var shows []model.Show

	for cursor.Next(context.Background()) {
		var show model.Show
		err := cursor.Decode(&show)
		if err != nil {
			log.Printf("Error decoding show %v\n", err)
			continue
		}

		shows = append(shows, show)
	}
	if err := cursor.Err(); err != nil {
		log.Printf("Cursor iteration err: %v\n", err)
	}
	return shows, nil
}

func GetShowById(showId string) (*model.Show, error) {
	id, err := primitive.ObjectIDFromHex(showId)
	if err != nil {
		log.Printf("Invalid show id %v\n", err)
	}
	filter := bson.M{"_id": id}

	var show model.Show
	err = db.ShowsCollection.FindOne(context.Background(), filter).Decode(&show)
	if err != nil {
		log.Printf("show not found%v\n", err)
	}

	return &show, nil
}

func SearchShow(searchQuery string) ([]model.Show, error) {
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

	var shows []model.Show

	cursor, err := db.ShowsCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var show model.Show
		err := cursor.Decode(&show)
		if err != nil {
			log.Printf("Error decoding show %v\n", err)
			continue
		}

		shows = append(shows, show)

	}

	return shows, nil
}

func PopularShow() ([]model.Show, error) {
	var shows []model.Show
	filter := bson.M{
		"rating": bson.M{"$gt": 8},
	}
	cursor, err := db.ShowsCollection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("No show found %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var show model.Show
		err := cursor.Decode(&show)
		if err != nil {
			log.Printf("Error decoding show %v\n", err)
			continue
		}

		shows = append(shows, show)
	}
	return shows, nil
}

func GetShowByGenre(genre string, limit int) ([]model.Show, error) {
	var shows []model.Show
	filter := bson.M{"genre": genre}
	cursor, err := db.ShowsCollection.Find(context.Background(), filter, options.Find().SetLimit(int64(limit)))
	if err != nil {
		log.Printf("No show found %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var show model.Show
		err := cursor.Decode(&show)
		if err != nil {
			log.Printf("Error decoding show %v\n", err)
			continue
		}

		shows = append(shows, show)
	}
	return shows, nil
}

func SimilarShow(genres []string) ([]model.Show, error) {
	var shows []model.Show
	filter := bson.M{
		"genre": bson.M{"$in": genres},
	}

	cursor, err := db.ShowsCollection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Error finding show %v\n", err)
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var show model.Show
		err := cursor.Decode(&show)
		if err != nil {
			log.Fatalf("Error decoding show%v\n", err)
			continue
		}
		shows = append(shows, show)
	}
	return shows, nil
}

func TrendingShows() ([]model.Show, error) {
	var shows []model.Show
	filter := bson.M{
		"isFeatured": true,
	}

	cursor, err := db.ShowsCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var show model.Show
		err := cursor.Decode(&show)
		if err != nil {
			log.Println("Error decoding show")
			continue
		}
		shows = append(shows, show)
	}
	log.Println("trending shows", shows)
	return shows, nil
}
