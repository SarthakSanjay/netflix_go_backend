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
)

func GetSeasons(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	filter := bson.M{
		"showId": params["showId"],
	}
	cursor, err := db.SeasonsCollection.Find(context.Background(), filter)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error fetching seasons"}, http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	var seasons []model.Season

	for cursor.Next(context.Background()) {
		var season model.Season
		err := cursor.Decode(&season)
		if err != nil {
			log.Printf("Error decoding season %v\n", err)
			continue
		}
		seasons = append(seasons, season)
	}
	if err := cursor.Err(); err != nil {
		log.Printf("Cursor iteration err: %v\n", err)
	}

	utils.SendJSONResponse(w, dto.SuccessResponse{
		Message: "success",
		Data:    seasons,
	}, http.StatusOK)
}
