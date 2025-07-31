package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwt_secret string

type JwtPayLoad struct {
	Email  string `json:"email"`
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func getJwtSecret() string {
	if jwt_secret != "" {
		return jwt_secret
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in loading .env")
	}
	jwt_secret = os.Getenv("JWT_SECRET")
	return jwt_secret
}
func CreateToken(email string, userId string) (string, error) {
	secret := []byte(getJwtSecret())
	expirationTime := time.Now().Add(24 * time.Hour)
	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtPayLoad{
		Email:  email,
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	})
	token, err := payload.SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, err
}
