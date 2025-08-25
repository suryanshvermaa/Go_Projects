package utils

import (
	"errors"
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

func VerifyToken(tokenString string) (*JwtPayLoad, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token is not valid or expired")
		}

		return []byte(jwt_secret), nil
	})

	// 4. Handle errors.
	if err != nil {
		return nil, err
	}

	// 5. Check if the token is valid and extract claims.
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return &JwtPayLoad{
			Email:  claims["email"].(string),
			UserId: claims["userId"].(string),
		}, nil
	} else {
		return nil, errors.New("token is not valid or expired")
	}
}
