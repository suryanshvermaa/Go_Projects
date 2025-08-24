package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseWithJson(w http.ResponseWriter, statusCode int, message string, payload interface{}) {
	responseStruct := Response{Success: true, Message: message, Data: payload}
	if !(statusCode >= 200 && statusCode < 300) {
		responseStruct.Success = false
	}
	data, err := json.Marshal(responseStruct)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
