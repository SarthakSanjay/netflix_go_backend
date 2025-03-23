package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func FavRoutes(protectedRoutes *mux.Router) {
	protectedRoutes.HandleFunc("/favorite", controller.AddMovieToFavorite).Methods("POST")
	protectedRoutes.HandleFunc("/favorite/{profileId}/{contentId}", controller.RemoveMovieFromFavorite).Methods("DELETE")
	protectedRoutes.HandleFunc("/favorites/{id}", controller.GetAllMoviesFromUsersProfileFavorite).Methods("GET")
	protectedRoutes.HandleFunc("/favorite", controller.AddShowToFavorite).Methods("POST")
	protectedRoutes.HandleFunc("/favorite/{profileId}/{contentId}", controller.RemoveShowFromFavorite).Methods("DELETE")
}
