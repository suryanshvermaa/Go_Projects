package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JsonResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	res := response{Success: true, Message: message, Data: data}
	w.WriteHeader(statusCode)
	if statusCode >= 400 {
		res.Success = false
	}
	resJson, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(resJson)
}
