package utils

import "github.com/gin-gonic/gin"

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(c *gin.Context, status int, message string, data interface{}) {
	success := false
	if status >= 200 && status < 300 {
		success = true
	}
	c.JSON(status, response{success, message, data})
}
