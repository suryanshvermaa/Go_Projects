package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/signup")
	incomingRoutes.POST("user/login")
}
