package services

import (
	"fmt"
	"net/http"
	"time"
)

func SetTokenCookies(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
	fmt.Println("Access token set in cookie")
}

func GetTokenFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
