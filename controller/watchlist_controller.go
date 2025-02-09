package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func AddMovieToWatchlist(w http.ResponseWriter, r *http.Request) {
	var req model.AddMovieDTO
	json.NewDecoder(r.Body).Decode(&req)

	profileId := req.ProfileId.Hex()
	movieId := req.MovieId.Hex()

	insertedCount, err := helper.AddMovieToWatchlist(movieId, profileId)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "error adding movie to watchlist"},
			http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"message": "success",
		"movieId": insertedCount,
	}
	utils.SendJSONResponse(w, data, http.StatusOK)
}

func GetUserWatchlist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	watchlist, err := helper.GetAllMovieFromUserWatchlist(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, model.ErrorResponseDTO{Error: "error finding movies in watchlist"}, http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"message":   "success",
		"watchlist": watchlist,
	}

	utils.SendJSONResponse(w, data, http.StatusOK)
}

func DeleteMovieFromWatchlist() {
}

func DeleteAllMovieFromWatchlist() {
}
