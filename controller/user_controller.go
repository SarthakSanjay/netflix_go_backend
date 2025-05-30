package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

	_, _, err := helper.CreateUser(user)
	if err != nil {
		utils.SendJSONResponse(w, map[string]interface{}{
			"error": "user already exist please login",
		}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, map[string]interface{}{
		"message": "success",
	}, http.StatusOK)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	isLoggedIn, accessToken, refreshToken, userF := helper.LoginUser(user)

	if !isLoggedIn {
		utils.SendJSONResponse(w, map[string]string{"error": "login failed"}, http.StatusInternalServerError)
		return
	}
	services.SetTokenCookies(w, "access_token", accessToken)
	services.SetTokenCookies(w, "refresh_token", refreshToken)

	response := map[string]interface{}{
		"message":      "success",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"user":         userF,
	}
	utils.SendJSONResponse(w, response, http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&user)

	// fmt.Println("user is ", user)
	count, err := helper.UpdateUser(params["id"], user)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "unable to update user"}, http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"message":      "success",
		"updatedCount": count,
	}
	utils.SendJSONResponse(w, res, http.StatusOK)
}

func UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var reqBody struct {
		Role model.Role `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}

	if !reqBody.Role.IsValid() {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid role provided"}, http.StatusBadRequest)
		return
	}

	updateUser, err := helper.UpdateUserRole(params["id"], reqBody.Role)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Unable to update user role"}, http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"message":     "success",
		"newRole":     model.RoleAdmin,
		"updateCount": updateUser,
	}

	utils.SendJSONResponse(w, res, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	deletedUser, err := helper.DeleteUser(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "error deleting user"}, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "user deleted successfully",
		"user":    deletedUser,
	}

	utils.SendJSONResponse(w, response, http.StatusOK)
}

func DeleteAllUser(w http.ResponseWriter, r *http.Request) {
	deletedCount, err := helper.DeleteAllUser()
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Unable to delete"}, http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"message":      "success",
		"deletedCount": deletedCount,
	}

	utils.SendJSONResponse(w, res, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user, err := helper.GetUser(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "error finding user"}, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "success",
		"user":    user,
	}

	utils.SendJSONResponse(w, response, http.StatusOK)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	users := helper.GetAllUser()

	utils.SendJSONResponse(w, map[string]interface{}{
		"message": "success",
		"users":   users,
		"total":   len(users),
	}, http.StatusOK)
}
