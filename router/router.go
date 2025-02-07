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

	router.HandleFunc("/api/refresh", controller.RefreshTokens).Methods("POST")
	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	protectedRoutes.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")
	protectedRoutes.HandleFunc("/user/{id}", controller.UpdateUser).Methods("PUT")
	protectedRoutes.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")
	protectedRoutes.HandleFunc("/users", controller.DeleteAllUser).Methods("DELETE")
	protectedRoutes.HandleFunc("/users", controller.GetAllUser).Methods("GET")
	protectedRoutes.HandleFunc("/user/{id}/role", controller.UpdateUserRole).Methods("PUT")

	// movies routes
	protectedRoutes.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	protectedRoutes.HandleFunc("/movies", controller.DeleteAllMovie).Methods("DELETE")
	protectedRoutes.HandleFunc("/movie", controller.CreateMovie).Methods("POST")
	protectedRoutes.HandleFunc("/movie/{id}", controller.GetMovieById).Methods("GET")
	protectedRoutes.HandleFunc("/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	protectedRoutes.HandleFunc("/movie/{id}", controller.UpdateMovie).Methods("PUT")
	protectedRoutes.HandleFunc("/movie/search", controller.SearchMovie).Methods("GET")
	protectedRoutes.HandleFunc("/movie/popular", controller.PopularMovie).Methods("GET")
	// protectedRoutes.HandleFunc("/movie/recommended", controller.RecommendedMovie).Methods("GET")
	protectedRoutes.HandleFunc("/movie/{id}/similar", controller.SimilarMovie).Methods("GET")

	//
	// watchlist routes

	// router.HandleFunc("/api/watchlist/{id}", .GetUserWatchlist())
	protectedRoutes.HandleFunc("/watchlist/{id}", controller.AddMovieToWatchlist).Methods("POST")
	router.HandleFunc("/watchlist/{id}", controller.GetUserWatchlist).Methods("GET")

	return router
}
