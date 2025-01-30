package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
}

var (
	accessSecret  = []byte(os.Getenv("ACCESS_SECRET"))
	refreshSecret = []byte(os.Getenv("REFRESH_SECRET"))
	resetSecret   = []byte(os.Getenv("RESET_SECRET"))
)

type Claims struct {
	UserId primitive.ObjectID `json:"userId,omitempty"`
	Email  string             `json:"email,omitempty"`
	jwt.RegisteredClaims
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
	return GenerateJWT(userId, email, accessSecret, 15*time.Minute)
}

func GenerateRefreshToken(userId primitive.ObjectID, email string) (string, error) {
	return GenerateJWT(userId, email, refreshSecret, 7*24*time.Hour)
}

func GenerateResetToken(userId primitive.ObjectID, email string) (string, error) {
	return GenerateJWT(userId, email, refreshSecret, 15*time.Minute)
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
