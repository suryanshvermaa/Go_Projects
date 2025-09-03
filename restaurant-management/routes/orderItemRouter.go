package routes

import "github.com/gin-gonic/gin"

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems")
	incomingRoutes.GET("/orderItems/:orderItem_id")
	incomingRoutes.GET("/orderItems-order/:order_id")
	incomingRoutes.POST("/orderItems")
	incomingRoutes.PATCH("/orderItems/:orderItem_id")
}
