package controllers

import (
	"net/http"

	"github.com/anshvermaa/go-chi-gorm/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, 200, "user fetched(testing route)", map[string]any{
		"name":  "Suryansh Verma",
		"email": "suryanshverma.dev.official@gmail.com",
	})
}
