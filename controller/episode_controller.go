package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/db"
	"github.com/sarthaksanjay/netflix-go/dto"
	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSeasonEpisodes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	showId, err := primitive.ObjectIDFromHex(params["showId"])
	if err != nil {
		log.Println("Invalid showID")
		return
	}

	seasonId, err := primitive.ObjectIDFromHex(params["seasonId"])
	if err != nil {
		log.Println("Invalid seasonId")
		return
	}

	log.Println("showId and seasonID", showId, seasonId)

	filter := bson.M{
		"showId":   showId,
		"seasonId": seasonId,
	}
	cursor, err := db.EpisodesCollection.Find(context.Background(), filter)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error fetching episodes"}, http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	var episodes []model.Episode
	for cursor.Next(context.Background()) {
		var episode model.Episode
		err := cursor.Decode(&episode)
		if err != nil {
			log.Printf("Error decoding episode %v\n", err)
			continue
		}

		episodes = append(episodes, episode)
	}
	if err := cursor.Err(); err != nil {
		log.Printf("Cursor iteration err: %v\n", err)
	}

	utils.SendJSONResponse(w, dto.SuccessResponse{
		Message: "success",
		Data:    episodes,
	}, http.StatusOK)
}
