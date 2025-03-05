package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/dto"
	helper "github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := helper.GetAllMovie(r.Context())
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Failed to fetch movies"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.MovieSuccessResponse{
		Message: "success",
		Movies:  movies,
		Total:   len(movies),
	},
		http.StatusOK)
}

func SearchMovie(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		utils.SendJSONResponse(w, map[string]string{"error": "Missing 'q' query params"}, http.StatusBadRequest)
		return
	}

	movies, err := helper.SearchMovie(query)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Movies not found"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.MovieSuccessResponse{
		Message: "success",
		Movies:  movies,
		Total:   len(movies),
	},
		http.StatusOK)
}

func PopularMovie(w http.ResponseWriter, r *http.Request) {
	movies, err := helper.PopularMovie()
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Failed to fetch movies"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.MovieSuccessResponse{
		Message: "success",
		Movies:  movies,
		Total:   len(movies),
	},
		http.StatusOK)
}

//
// func RecommendedMovie(w http.ResponseWriter , r *http.Request) {
// }

func SimilarMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movie, err := helper.GetMovieById(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Movie not found"}, http.StatusInternalServerError)
		return
	}

	similarMovie, err := helper.SimilarMovie(movie.Genre)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Failed to get similar movie"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.MovieSuccessResponse{
		Message: "success",
		Movies:  similarMovie,
		Total:   len(similarMovie),
	},
		http.StatusOK)
}

func GetMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	limitStr := r.URL.Query().Get("limit")
	limit := 40
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = parsedLimit
		}
	}
	movies, err := helper.GetMovieByGenre(params["genre"], limit)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Movie not found"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.MovieSuccessResponse{
		Message: "success",
		Movies:  movies,
		Total:   len(movies),
	},
		http.StatusOK)
}

func GetTrendingMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TrendingMovie")
	movies, err := helper.TrendingMovie()
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "error finding movies"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.MovieSuccessResponse{
		Message: "success",
		Movies:  movies,
		Total:   len(movies),
	},
		http.StatusOK)
}
