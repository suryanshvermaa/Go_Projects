package helpers

import "github.com/gin-gonic/gin"

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JsonResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	success := true
	if statusCode >= 300 {
		success = false
	}
	res := response{
		Success: success,
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, res)
}
