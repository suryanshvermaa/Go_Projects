package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	user.Password = utils.HashPassword(user.Password)
	u, dbError := user.CreateUser()
	if dbError != nil {
		utils.JsonResponse(w, 400, dbError.Error(), nil)
		return
	}
	token, err := utils.CreateToken(user.Email, strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		utils.JsonResponse(w, 400, "Error in login", nil)
		return
	}
	u.Password = "" // sensitive cannot be shared
	utils.JsonResponse(w, 201, "user created", map[string]interface{}{"token": token, "user": u})
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
	if !utils.VerifyPassword(user.Password, data.Password) {
		utils.JsonResponse(w, 400, "Email or password is wrong", nil)
		return
	}
	token, err := utils.CreateToken(user.Email, strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		utils.JsonResponse(w, 400, "Error in login", nil)
		return
	}
	user.Password = "" // sensitive cannot be shared
	utils.JsonResponse(w, 400, "Login successful", map[string]interface{}{"token": token, "user": user})
}
