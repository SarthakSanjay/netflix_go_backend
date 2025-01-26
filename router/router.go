package router

import (
	"github.com/gorilla/mux"
	movieController "github.com/sarthaksanjay/netflix-go/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", movieController.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", movieController.CreateMovie).Methods("POST")

	return router
}
