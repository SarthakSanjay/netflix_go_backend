package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func CastRoutes(protectedRoute *mux.Router) {
	protectedRoute.HandleFunc("/cast", controller.AddCast).Methods("POST")
}
