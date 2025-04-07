package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sarthaksanjay/netflix-go/db"
	"github.com/sarthaksanjay/netflix-go/dto"
	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func InsertSeason(w http.ResponseWriter, r *http.Request) {
	var season model.Season
	err := json.NewDecoder(r.Body).Decode(&season)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Invalid request body"}, http.StatusBadRequest)
		return
	}

	result, err := db.SeasonsCollection.InsertOne(context.Background(), season)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "error inserting season"}, http.StatusInternalServerError)
	}

	utils.SendJSONResponse(w, dto.SuccessResponse{
		Message: "success",
		Data:    result,
	}, http.StatusOK)
}

func InsertEpisode(w http.ResponseWriter, r *http.Request) {
	var episode model.Episode
	err := json.NewDecoder(r.Body).Decode(&episode)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Invalid request body"}, http.StatusBadRequest)
		log.Println("error", err)
		return
	}

	result, err := db.EpisodesCollection.InsertOne(context.Background(), episode)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "error inserting episode"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.SuccessResponse{
		Message: "success",
		Data:    result,
	}, http.StatusOK)
}
