package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func FavRoutes(protectedRoutes *mux.Router) {
	protectedRoutes.HandleFunc("/favorite/movie", controller.AddMovieToFavorite).Methods("POST")
	protectedRoutes.HandleFunc("/favorite/{profileId}/{contentId}", controller.RemoveMovieFromFavorite).Methods("DELETE")
	protectedRoutes.HandleFunc("/favorites/movie/{id}", controller.GetAllMoviesFromUsersProfileFavorite).Methods("GET")
	protectedRoutes.HandleFunc("/favorite/show", controller.AddShowToFavorite).Methods("POST")
	protectedRoutes.HandleFunc("/favorite/{profileId}/{contentId}", controller.RemoveShowFromFavorite).Methods("DELETE")
	protectedRoutes.HandleFunc("/favorites/show/{id}", controller.GetAllShowsFromUsersProfileFavorite).Methods("GET")
}
