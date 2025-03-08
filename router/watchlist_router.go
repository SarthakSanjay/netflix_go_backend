package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func WatchlistRoutes(protectedRoutes *mux.Router) {
	// watchlist routes

	// protectedRoutes.HandleFunc("/api/watchlist/{id}", .GetUserWatchlist())
	protectedRoutes.HandleFunc("/watchlist", controller.AddMovieToWatchlist).Methods("POST")
	protectedRoutes.HandleFunc("/watchlist/{id}", controller.GetUserWatchlist).Methods("GET")
	protectedRoutes.HandleFunc("/watchlist", controller.DeleteMovieFromWatchlist).Methods("DELETE")
	protectedRoutes.HandleFunc("/watchlist/all/{id}", controller.DeleteAllMovieFromWatchlist).Methods("DELETE")
}
