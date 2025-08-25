package main

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1cnlhbnNodi51ZzIzLmVlQGRqZG5pdHAuYWMuaW4iLCJ1c2VySWQiOiIxIiwiZXhwIjoxNzU0MTI1MDczfQ.lluAaAGwfylRnFA3TQuerGRybfBH2biQNDNA9hRpsOY"

	mySigningKey := []byte("YW5zaDk1ODAxMDQ3NTMK")
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})

	// 4. Handle errors.
	if err != nil {
		log.Fatalf("Token parsing error: %v", err)
	}

	// 5. Check if the token is valid and extract claims.
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Println("Token is valid! 👍")
		fmt.Println(claims)
	} else {
		fmt.Println("Token is not valid.")
	}
}
