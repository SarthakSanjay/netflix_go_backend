package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
	"github.com/sarthaksanjay/netflix-go/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", controller.CheckHealth).Methods("GET")

	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	AuthRoutes(router)
	UserRoutes(protectedRoutes)
	UserProfileRoutes(protectedRoutes)
	MovieRoutes(protectedRoutes)
	WatchlistRoutes(protectedRoutes)
	FavRoutes(protectedRoutes)

	return router
}
