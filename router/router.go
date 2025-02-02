package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
	"github.com/sarthaksanjay/netflix-go/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", controller.LoginUser).Methods("POST")
	// user routes

	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// protectedRoutes.HandleFunc("/api/user/{id}", controller.GetUser).Methods("GET")
	// protectedRoutes.HandleFunc("/api/user/{id}", controller.UpdateUser).Methods("PUT")
	// protectedRoutes.HandleFunc("/api/user/{id}", controller.DeleteUser).Methods("DELETE")
	// protectedRoutes.HandleFunc("/api/users", controller.DeleteAllUser).Methods("DELETE")
	protectedRoutes.HandleFunc("/users", controller.GetAllUser).Methods("GET")

	// movies routes
	protectedRoutes.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	protectedRoutes.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")
	protectedRoutes.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	protectedRoutes.HandleFunc("/api/movie/{id}", controller.GetMovieById).Methods("GET")
	protectedRoutes.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	protectedRoutes.HandleFunc("/api/movie/{id}", controller.UpdateMovie).Methods("PUT")
	protectedRoutes.HandleFunc("/api/movie/search", controller.SearchMovie).Methods("GET")
	protectedRoutes.HandleFunc("/api/movie/popular", controller.PopularMovie).Methods("GET")
	// protectedRoutes.HandleFunc("/api/movie/recommended", controller.RecommendedMovie).Methods("GET")
	protectedRoutes.HandleFunc("/api/movie/{id}/similar", controller.SimilarMovie).Methods("GET")

	//
	// watchlist routes

	// router.HandleFunc("/api/watchlist/{id}", .GetUserWatchlist())

	return router
}
