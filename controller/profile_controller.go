package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/dto"
	"github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddNewProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Invalid req body"}, http.StatusBadRequest)
		return
	}
	var req dto.CreateProfileDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Invalid req body"}, http.StatusBadRequest)
	}

	defer r.Body.Close()

	newProfile := model.Profile{
		UserId:    userId,
		Name:      req.Name,
		Avatar:    req.Avatar,
		History:   []model.History{},
		Watchlist: []model.Watchlist{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	insertedId, err := helper.CreateUserProfile(newProfile)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "error creating user profile"}, http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"message":    "success",
		"insertedId": insertedId,
		"profile":    newProfile,
	}

	utils.SendJSONResponse(w, data, http.StatusOK)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := helper.GetProfileById(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error finding profile with this ID"}, http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"message": "success",
		"profile": result,
	}

	utils.SendJSONResponse(w, data, http.StatusOK)
}

func GetAllUserProfiles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	profiles, err := helper.GetAllUserProfiles(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error finding profiles with given userId"}, http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"message":  "success",
		"profiles": profiles,
	}

	utils.SendJSONResponse(w, data, http.StatusOK)
}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var req model.Profile
	json.NewDecoder(r.Body).Decode(&req)
	updateCount, err := helper.UpdateProfile(params["id"], req)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error updating user profile"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.SuccessResponse{Message: "success", Data: updateCount}, http.StatusOK)
}

func DeleteUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	deleteCount, err := helper.DeleteProfile(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error deleting user profile"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.SuccessResponse{Message: "success", Data: deleteCount}, http.StatusOK)
}
