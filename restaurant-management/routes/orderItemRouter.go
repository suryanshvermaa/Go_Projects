package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suryanshvermaa/restaurant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
