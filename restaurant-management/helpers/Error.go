package helpers

import "github.com/gin-gonic/gin"

func NewError(c *gin.Context, errorStatus int, message string) {
	JsonResponse(c, errorStatus, message, nil)
}
