package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sarthaksanjay/netflix-go/dto"
	"github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/services"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	log.Println("hello")
	id, err := utils.ExtractUserIdFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("actual id after extraction ", id)

	userId := id.Hex()
	fmt.Printf("Type of user id %T", userId)

	updateCount, err := helper.LogoutUser(userId)
	if err != nil {
		utils.SendJSONResponse(w, dto.ErrorResponseDTO{Error: "Error logging out user"}, http.StatusInternalServerError)
		return
	}

	services.ClearTokenCookies(w, "access_token")
	services.ClearTokenCookies(w, "refresh_token")

	utils.SendJSONResponse(w, dto.SuccessResponse{
		Message: "success",
		Data:    updateCount,
	}, http.StatusOK)
}
