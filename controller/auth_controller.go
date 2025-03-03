package controller

import (
	"net/http"
	"os"

	"github.com/sarthaksanjay/netflix-go/helper"
	"github.com/sarthaksanjay/netflix-go/services"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func RefreshTokens(w http.ResponseWriter, r *http.Request) {
	token, cookieErr := services.GetTokenFromCookie("refresh_token", r)

	tokenFromHeader, headerErr := services.GetTokenFromHeader(r)

	if (cookieErr != nil || token == "") && (tokenFromHeader == "" || headerErr != nil) {
		// fmt.Println("error", err)
		utils.SendJSONResponse(w, map[string]string{"error": "cannot retrive refresh_token"}, http.StatusInternalServerError)
		return
	}

	finalToken := token
	if finalToken == "" {
		finalToken = tokenFromHeader
	}

	_, claims, err := services.VerifyToken(finalToken, []byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "cannot retrive refresh_token"}, http.StatusInternalServerError)
		return
	}

	accessToken, refreshToken, err := helper.RefreshToken(claims.UserId)
	if err != nil {
		utils.SendJSONResponse(w, map[string]error{"error": err}, http.StatusInternalServerError)
		return
	}
	services.SetTokenCookies(w, "access_token", accessToken)
	services.SetTokenCookies(w, "refresh_token", refreshToken)

	utils.SendJSONResponse(w, map[string]string{
		"message":      "success",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}, http.StatusOK)
}
