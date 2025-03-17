package controller

import (
	"net/http"

	"github.com/sarthaksanjay/netflix-go/utils"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONResponse(w, map[string]string{"message": "server healthy"}, http.StatusOK)
}
