package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	jwt.RegisteredClaims
	Email  string             `json:"email,omitempty"`
	UserId primitive.ObjectID `json:"userId,omitempty"`
}

func GenerateJWT(userId primitive.ObjectID, email string, secret []byte, expiry time.Duration) (string, error) {
	claims := &Claims{
		UserId: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateAccessToken(userId primitive.ObjectID, email string) (string, error) {
	return GenerateJWT(userId, email, []byte(os.Getenv("ACCESS_SECRET")), 15*time.Minute)
}

func GenerateRefreshToken(userId primitive.ObjectID, email string) (string, error) {
	return GenerateJWT(userId, email, []byte(os.Getenv("REFRESH_SECRET")), 7*24*time.Hour)
}

func GenerateResetToken(userId primitive.ObjectID, email string) (string, error) {
	return GenerateJWT(userId, email, []byte(os.Getenv("RESET_TOKEN")), 15*time.Minute)
}

func VerifyToken(tokenString string, secret []byte) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, claims, nil
}
