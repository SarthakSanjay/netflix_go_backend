package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/dto"
	"github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func GetAllShows(w http.ResponseWriter, r *http.Request) {
	shows, err := helper.GetAllShows(r.Context())
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Failed to fetch shows"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.ShowsSuccessResponse{
		Message: "success",
		Shows:   shows,
		Total:   len(shows),
	},
		http.StatusOK)
}

func GetShowById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	show, err := helper.GetShowById(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Failed to fetch show"}, http.StatusInternalServerError)
	}
	utils.SendJSONResponse(w, dto.ShowSuccessResponse{
		Message: "success",
		Show:    *show,
	}, http.StatusOK)
}

func GetShowByGenre(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	limitStr := r.URL.Query().Get("limit")
	limit := 40
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = parsedLimit
		}
	}
	shows, err := helper.GetShowByGenre(params["genre"], limit)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Failed to fetch show"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.ShowsSuccessResponse{
		Message: "success",
		Shows:   shows,
		Total:   len(shows),
	},
		http.StatusOK)
}

func GetTrendingShows(w http.ResponseWriter, r *http.Request) {
	shows, err := helper.TrendingShows()
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "error finding shows"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.ShowsSuccessResponse{
		Message: "success",
		Shows:   shows,
		Total:   len(shows),
	},
		http.StatusOK)
}

func getSimilarShows(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	shows, err := helper.GetShowById(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Failed to fetch shows"}, http.StatusInternalServerError)
		return
	}

	similarShows, err := helper.SimilarShow(shows.Genre)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Failed to get similar movie"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.ShowsSuccessResponse{
		Message: "success",
		Shows:   similarShows,
		Total:   len(similarShows),
	},
		http.StatusOK)
}
