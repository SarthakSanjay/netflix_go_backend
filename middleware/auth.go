package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sarthaksanjay/netflix-go/services"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessSecret := []byte(os.Getenv("ACCESS_SECRET"))

		tokenString, err := services.GetTokenFromCookie(r)

		if err != nil || tokenString == "" {
			utils.SendJSONResponse(w, map[string]string{"message": "access token is missing please login"}, http.StatusUnauthorized)
			return
		}

		token, claims, err := services.VerifyToken(tokenString, accessSecret)
		if err != nil {
			fmt.Println("error", err)
			utils.SendJSONResponse(w, map[string]interface{}{
				"error": "unauthorized",
				"err":   err,
			}, http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			utils.SendJSONResponse(w, map[string]string{"error": "unauthorized token invalid"}, http.StatusUnauthorized)
			return
		}
		fmt.Println("claims", claims)

		next.ServeHTTP(w, r)
	})
}
