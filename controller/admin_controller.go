package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/dto"
	"github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/utils"
)

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
	utils.SendJSONResponse(w, dto.MovieSuccessResponse{
		Message: "success",
		Movie:   *movie,
	}, http.StatusOK)
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

func CreateShow(w http.ResponseWriter, r *http.Request) {
	var show model.Show

	err := json.NewDecoder(r.Body).Decode(&show)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}

	id, err := helper.InsertShow(show)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "failed to insert show"}, http.StatusInternalServerError)
		return
	}
	respose := map[string]interface{}{
		"message":        "success",
		"insertedShowId": id.Hex(),
		"show":           &show,
	}
	utils.SendJSONResponse(w, respose, http.StatusCreated)
}
