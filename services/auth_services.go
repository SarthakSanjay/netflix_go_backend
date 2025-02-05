package services

import (
	"fmt"
	"net/http"
	"time"
)

func SetTokenCookies(w http.ResponseWriter, name string, token string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   false, // true for https , for local development like http use false
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
	fmt.Println("Access token set in cookie")
}

func GetTokenFromCookie(name string, r *http.Request) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
