package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Go_Projects/postgres-go-crud/models"
	"github.com/Go_Projects/postgres-go-crud/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.PhoneNo) == "" || strings.TrimSpace(user.Password) == "" || strings.TrimSpace(user.Name) == "" {
		utils.JsonResponse(w, 400, "all fields are required", nil)
		return
	}
	if err != nil {
		utils.JsonResponse(w, 400, "error in decoding body", nil)
		return
	}
	u, dbError := user.CreateUser()
	if dbError != nil {
		utils.JsonResponse(w, 400, dbError.Error(), nil)
		return
	}
	utils.JsonResponse(w, 201, "user created", u)
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data LoginPayload
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		utils.JsonResponse(w, 400, "json data not provided", nil)
		return
	}
	if strings.TrimSpace(data.Email) == "" || strings.TrimSpace(data.Password) == "" {
		utils.JsonResponse(w, 400, "all fields are required", nil)
		return
	}
	user, DbErr := models.GetUserByEmail(data.Email)
	if DbErr != nil {
		utils.JsonResponse(w, 400, DbErr.Error(), nil)
		return
	}
	if user.Password != data.Password {
		utils.JsonResponse(w, 400, "Email or password is wrong", nil)
		return
	}
	utils.JsonResponse(w, 400, "Login successful", user)
}
