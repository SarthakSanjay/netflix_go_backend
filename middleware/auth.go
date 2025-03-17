package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/sarthaksanjay/netflix-go/model"
	"github.com/sarthaksanjay/netflix-go/services"
	"github.com/sarthaksanjay/netflix-go/types"
	"github.com/sarthaksanjay/netflix-go/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, cookieErr := services.GetTokenFromCookie("access_token", r)

		tokenStringFromHeader, headerErr := services.GetTokenFromHeader(r)

		if (cookieErr != nil || tokenString == "") && (headerErr != nil || tokenStringFromHeader == "") {
			utils.SendJSONResponse(w, map[string]string{"message": "access token is missing, please login"}, http.StatusUnauthorized)
			return
		}

		finalToken := tokenString
		if finalToken == "" {
			finalToken = tokenStringFromHeader
		}

		token, claims, err := services.VerifyToken(finalToken, []byte(os.Getenv("ACCESS_SECRET")))
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

		ctx := context.WithValue(r.Context(), types.UserContextKey, claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func RequiredRole(requiredRole model.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(types.UserContextKey).(*services.Claims)
			if !ok {
				utils.SendJSONResponse(w, map[string]string{
					"error": "invalid context",
				}, http.StatusUnauthorized)
				return
			}

			if claims.Role != requiredRole {
				utils.SendJSONResponse(w, map[string]string{
					"error": "insufficient permissions",
				}, http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
