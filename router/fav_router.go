package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func FavRoutes(protectedRoutes *mux.Router) {
	protectedRoutes.HandleFunc("/api/favorite", controller.AddContentToFavorite).Methods("POST")
	protectedRoutes.HandleFunc("/api/favorite", controller.RemoveContentFromFavorite).Methods("DELETE")
	protectedRoutes.HandleFunc("/api/favorites/{id}", controller.GetAllContentFromUsersProfileFavorite).Methods("GET")
}
