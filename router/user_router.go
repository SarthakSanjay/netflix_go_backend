package router

import (
	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/controller"
)

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/api/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", controller.LoginUser).Methods("POST")
	router.HandleFunc("/api/refresh", controller.RefreshTokens).Methods("POST")
}

func UserRoutes(protectedRoutes *mux.Router) {
	protectedRoutes.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")
	protectedRoutes.HandleFunc("/user/{id}", controller.UpdateUser).Methods("PUT")
	protectedRoutes.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")
	protectedRoutes.HandleFunc("/users", controller.DeleteAllUser).Methods("DELETE")
	protectedRoutes.HandleFunc("/users", controller.GetAllUser).Methods("GET")
	protectedRoutes.HandleFunc("/user/{id}/role", controller.UpdateUserRole).Methods("PUT")
	protectedRoutes.HandleFunc("/user/logout", controller.LogoutUser).Methods("POST")
}

func UserProfileRoutes(protectedRoutes *mux.Router) {
	protectedRoutes.HandleFunc("/profile/{id}", controller.AddNewProfile).Methods("POST")
	protectedRoutes.HandleFunc("/profile/{id}", controller.GetUserProfile).Methods("GET")
	protectedRoutes.HandleFunc("/user/profiles/{id}", controller.GetAllUserProfiles).Methods("GET")
	protectedRoutes.HandleFunc("/profile/{id}", controller.UpdateUserProfile).Methods("PUT")
	protectedRoutes.HandleFunc("/profile/{id}", controller.DeleteUserProfile).Methods("DELETE")
}
