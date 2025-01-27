package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// user routes

	router.HandleFunc("/api/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controller.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users", controller.DeleteAllUser).Methods("DELETE")
	router.HandleFunc("/api/users", controller.CreateUser).Methods("GET")

	// movies routes
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.GetMovieById).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")

	// watchlist routes

	// router.HandleFunc("/api/watchlist/{id}", .GetUserWatchlist())

	return router
}
