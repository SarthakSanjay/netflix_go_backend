package controller

import (
	"encoding/json"
	"net/http"

	helper "github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/model"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	allMovies := helper.GetAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "POST")

	var movie model.Movies
	json.NewDecoder(r.Body).Decode(&movie)
	helper.InsertMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
}
