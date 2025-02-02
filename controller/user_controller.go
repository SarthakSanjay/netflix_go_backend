package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/services"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Decode request body and handle errors
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	message, accessToken, err := helper.CreateUser(user)
	fmt.Println("messssssssage", message)
	if err != nil {
		statusCode := http.StatusInternalServerError

		// Handle "user already exists" separately
		if message == "user already exists" {
			statusCode = http.StatusConflict
		}

		utils.SendJSONResponse(w, map[string]interface{}{
			"error": message,
		}, statusCode)
		return
	}

	// Set auth cookies
	services.SetTokenCookies(w, accessToken)

	// Send success response
	utils.SendJSONResponse(w, map[string]interface{}{
		"message":     "success",
		"accessToken": accessToken,
	}, http.StatusOK)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	isLoggedIn, _ := helper.LoginUser(user)

	if !isLoggedIn {
		utils.SendJSONResponse(w, map[string]string{"error": "login failed"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, map[string]string{"message": "login success"}, http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteAllUser(w http.ResponseWriter, r *http.Request) {
}

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Content-Allow-Methods", "POST")

	users := helper.GetAllUser()
	json.NewEncoder(w).Encode(users)
}
