package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func MovieRoutes(protectedRoutes *mux.Router) {
	// movies routes
	protectedRoutes.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	protectedRoutes.HandleFunc("/movies", controller.DeleteAllMovie).Methods("DELETE")
	protectedRoutes.HandleFunc("/movie", controller.CreateMovie).Methods("POST")
	protectedRoutes.HandleFunc("/movie/{id}", controller.GetMovieById).Methods("GET")
	protectedRoutes.HandleFunc("/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	protectedRoutes.HandleFunc("/movie/{id}", controller.UpdateMovie).Methods("PUT")
	protectedRoutes.HandleFunc("/movie/search", controller.SearchMovie).Methods("GET")
	protectedRoutes.HandleFunc("/movie/popular", controller.PopularMovie).Methods("GET")
	// protectedRoutes.HandleFunc("/movie/recommended", controller.RecommendedMovie).Methods("GET")
	protectedRoutes.HandleFunc("/movie/{id}/similar", controller.SimilarMovie).Methods("GET")
	protectedRoutes.HandleFunc("/content/{genre}", controller.GetMoviesByGenre).Methods(("GET"))

	protectedRoutes.HandleFunc("/content/trending", controller.GetTrendingMovies).Methods("GET")
}
