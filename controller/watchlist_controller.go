package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/dto"
	"github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func AddMovieToWatchlist(w http.ResponseWriter, r *http.Request) {
	var req dto.AddMovieDTO
	json.NewDecoder(r.Body).Decode(&req)

	profileId := req.ProfileId.Hex()
	contentId := req.ContentId.Hex()

	fmt.Println(profileId, contentId)
	// fmt.Println("body", r.Body)
	insertedContentId, err := helper.AddMovieToWatchlist(contentId, profileId)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "error adding movie to watchlist"},
			http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"message":   "success",
		"contentId": insertedContentId,
	}
	utils.SendJSONResponse(w, data, http.StatusOK)
}

func GetUserWatchlist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	watchlist, err := helper.GetAllMovieFromUserWatchlist(params["id"])
	if err != nil {
		log.Println(err)
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "error finding movies in watchlist"}, http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"message":   "success",
		"watchlist": watchlist,
	}

	utils.SendJSONResponse(w, data, http.StatusOK)
}

func DeleteMovieFromWatchlist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	profileId := params["profileId"]
	movieId := params["contentId"]

	deletedDoc, err := helper.DeleteMovieFromWatchlist(profileId, movieId)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error deleting movie from watchlist"}, http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"message":    "success",
		"deletedDoc": deletedDoc,
	}
	utils.SendJSONResponse(w, data, http.StatusOK)
}

func DeleteAllMovieFromWatchlist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	deletedCount, err := helper.DeleteAllMovieFromWatchlist(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error deleting movie from watchlist"}, http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"message":      "success",
		"deletedCount": deletedCount,
	}
	utils.SendJSONResponse(w, data, http.StatusOK)
}
