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
		"result":  result,
	}

	utils.SendJSONResponse(w, data, http.StatusOK)
}
