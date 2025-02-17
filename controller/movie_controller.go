package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	helper "github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := helper.GetAllMovie(r.Context())
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Failed to fetch movies"}, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "success",
		"total":   len(movies),
		"movies":  movies,
	}

	utils.SendJSONResponse(w, response, http.StatusOK)
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

	response := map[string]interface{}{
		"message": "movie found",
		"total":   len(movies),
		"movies":  movies,
	}

	utils.SendJSONResponse(w, response, http.StatusOK)
}

func PopularMovie(w http.ResponseWriter, r *http.Request) {
	movies, err := helper.PopularMovie()
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Failed to fetch movies"}, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "found popular movies",
		"total":   len(movies),
		"movies":  movies,
	}

	utils.SendJSONResponse(w, response, http.StatusOK)
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

	res := map[string]interface{}{
		"message":      "success",
		"total":        len(similarMovie),
		"similarMovie": similarMovie,
	}

	utils.SendJSONResponse(w, res, http.StatusOK)
}

func GetMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movies, err := helper.GetMovieByGenre(params["genre"])
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Movie not found"}, http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"message": "success",
		"total":   len(movies),
		"movies":  movies,
	}

	utils.SendJSONResponse(w, res, http.StatusOK)
}
