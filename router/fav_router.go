package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func FavRoutes(protectedRoutes *mux.Router) {
	protectedRoutes.HandleFunc("/favorite", controller.AddContentToFavorite).Methods("POST")
	protectedRoutes.HandleFunc("/favorite/{profileId}/{contentId}", controller.RemoveContentFromFavorite).Methods("DELETE")
	protectedRoutes.HandleFunc("/favorites/{id}", controller.GetAllContentFromUsersProfileFavorite).Methods("GET")
}
