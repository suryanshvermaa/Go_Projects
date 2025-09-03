package routes

import "github.com/gin-gonic/gin"

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders")
	incomingRoutes.GET("/orders/:order_id")
	incomingRoutes.POST("/orders")
	incomingRoutes.PATCH("/orders/:order_id")
}
