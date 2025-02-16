package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	helper "github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/model"
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

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.Movies

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}

	id, err := helper.InsertMovie(movie)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Failed to insert movie"}, http.StatusInternalServerError)
		return
	}
	respose := map[string]interface{}{
		"message":         "success",
		"insertedMovieId": id.Hex(),
		"movie":           &movie,
	}
	utils.SendJSONResponse(w, respose, http.StatusCreated)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if len(params["id"]) != 24 {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid movie id"}, http.StatusBadRequest)
		return
	}
	count, err := helper.DeleteMovie(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Failed to delete movie"}, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "successfully deleted movie",
		"id":      params["id"],
		"count":   count,
	}

	utils.SendJSONResponse(w, response, http.StatusOK)
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	count, err := helper.DeleteAllMovie()
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Error deleting All Movies"}, http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"message": "successfull deleted all movies",
		"count":   count,
	}

	utils.SendJSONResponse(w, response, http.StatusOK)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "GET")

	params := mux.Vars(r)
	movie, err := helper.GetMovieById(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if len(params["id"]) != 24 {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid movie Id"}, http.StatusBadRequest)
		return
	}

	var updates model.Movies
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	count, err := helper.UpdateMovie(params["id"], updates)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Failed to update movie"}, http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"message": "movie updated successfully",
		"id":      params["id"],
		"count":   count,
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
