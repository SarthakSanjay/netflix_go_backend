package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func WatchlistRoutes(protectedRoutes *mux.Router) {
	protectedRoutes.HandleFunc("/watchlist", controller.AddMovieToWatchlist).Methods("POST")
	protectedRoutes.HandleFunc("/watchlist/movie/{id}", controller.GetMoviesFromUserWatchlist).Methods("GET")
	protectedRoutes.HandleFunc("/watchlist/{profileId}/{contentId}", controller.DeleteContentFromWatchlist).Methods("DELETE")
	protectedRoutes.HandleFunc("/watchlist/all/{contentType}/{id}", controller.DeleteAllContentFromWatchlist).Methods("DELETE")
	protectedRoutes.HandleFunc("/watchlist/show/{id}", controller.GetShowsFromUserWatchlist).Methods("GET")
}
