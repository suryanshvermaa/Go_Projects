package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Auth struct {
	secretKey string
}

type JwtPayLoad struct {
	Email  string `json:"email"`
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func (a *Auth) setJwtSecret() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in loading .env")
	}
	a.secretKey = os.Getenv("AUTH_SECRET_KEY")
}

func (a *Auth) CreateToken(email string, userId string) (string, error) {
	if a.secretKey == "" {
		a.setJwtSecret()
	}
	secret := []byte(a.secretKey)
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

func (a *Auth) VerifyToken(tokenString string) (*JwtPayLoad, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token is not valid or expired")
		}

		return []byte(a.secretKey), nil
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
