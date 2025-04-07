package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func ShowRoutes(protectedRoutes *mux.Router) {
	// shows routes
	protectedRoutes.HandleFunc("/shows", controller.GetAllShows).Methods("GET")
	// protectedRoutes.HandleFunc("/shows", controller.DeleteAllShows).Methods("DELETE")
	protectedRoutes.HandleFunc("/show", controller.CreateShow).Methods("POST")
	protectedRoutes.HandleFunc("/show/{id}", controller.GetShowById).Methods("GET")
	protectedRoutes.HandleFunc("/{contentType}/{id}", controller.DeleteShow).Methods("DELETE")
	protectedRoutes.HandleFunc("/shows", controller.DeleteAllShow).Methods("DELETE")

	// protectedRoutes.HandleFunc("/show/{id}", controller.UpdateShow).Methods("PUT")
	// protectedRoutes.HandleFunc("/show/search", controller.SearchShow).Methods("GET")
	// protectedRoutes.HandleFunc("/show/popular", controller.PopularShows).Methods("GET")
	// protectedRoutes.HandleFunc("/movie/recommended", controller.RecommendedMovie).Methods("GET")
	// protectedRoutes.HandleFunc("/show/{id}/similar", controller.SimilarShows).Methods("GET")
	protectedRoutes.HandleFunc("/shows/{genre}", controller.GetShowByGenre).Methods(("GET"))

	protectedRoutes.HandleFunc("/shows/trending", controller.GetTrendingShows).Methods("GET")

	protectedRoutes.HandleFunc("/show/season", controller.InsertSeason).Methods("POST")
	protectedRoutes.HandleFunc("/show/episode", controller.InsertEpisode).Methods("POST")
}
