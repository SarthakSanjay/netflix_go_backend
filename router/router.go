package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")

	return router
}
