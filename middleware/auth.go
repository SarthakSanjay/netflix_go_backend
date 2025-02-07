package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/sarthaksanjay/netflix-go/services"
	"github.com/sarthaksanjay/netflix-go/utils"
)

type contextKey string

const UserContextKey contextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := services.GetTokenFromCookie("access_token", r)
		if err != nil || tokenString == "" {
			utils.SendJSONResponse(w, map[string]string{"message": "access token is missing please login"}, http.StatusUnauthorized)
			return
		}

		token, claims, err := services.VerifyToken(tokenString, []byte(os.Getenv("ACCESS_SECRET")))
		if err != nil {
			fmt.Println("error in auth", err)
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
		// fmt.Println("claims", claims)

		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
