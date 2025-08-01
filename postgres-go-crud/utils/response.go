package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Sucess  bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JsonResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	success := false
	if status >= 200 && status <= 300 {
		success = true
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{success, message, data})
}
