package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarthaksanjay/netflix-go/dto"
	"github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func AddContentToFavorite(w http.ResponseWriter, r *http.Request) {
	var req dto.FavoriteRequestDto
	json.NewDecoder(r.Body).Decode(&req)

	profileId := req.ProfileId.Hex()
	contentId := req.ContentId.Hex()

	result, err := helper.AddToFavorite(profileId, contentId)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error adding content to favorites"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.SuccessResponse{Message: "success", Data: result}, http.StatusOK)
}

func RemoveContentFromFavorite(w http.ResponseWriter, r *http.Request) {
	var req dto.FavoriteRequestDto
	json.NewDecoder(r.Body).Decode(&req)

	profileId := req.ProfileId.Hex()
	contentId := req.ContentId.Hex()

	result, err := helper.RemoveFromFavorite(profileId, contentId)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error removing content from favorite"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, dto.SuccessResponse{Message: "success", Data: result}, http.StatusOK)
}

func GetAllContentFromUsersProfileFavorite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	favorites, err := helper.GetUserFavoriteFromProfile(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error finding users favorite"}, http.StatusInternalServerError)
	}

	utils.SendJSONResponse(w, dto.SuccessResponse{Message: "success", Data: favorites}, http.StatusOK)
}
