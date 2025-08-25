package middlewares

import (
	"net/http"
	"strings"

	"github.com/Go_Projects/postgres-go-crud/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			utils.JsonResponse(w, 401, "Unauthorised", nil)
			return
		}
		token = strings.Split(token, "Bearer ")[1]
		if token == "" {
			utils.JsonResponse(w, 401, "Unauthorised", nil)
			return
		}
		decodedToken, err := utils.VerifyToken(token)
		if err != nil {
			utils.JsonResponse(w, 401, "Unauthorised", nil)
			return
		}
		if decodedToken == nil {
			utils.JsonResponse(w, 401, "Unauthorised", nil)
			return
		}
		r.Header.Set("email", decodedToken.Email)
		r.Header.Set("userId", decodedToken.UserId)
		next.ServeHTTP(w, r)
	})
}
