package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suryanshvermaa/restaurant-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetOrderTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetOrderTable())
	incomingRoutes.POST("/tables", controllers.CreateOrderTable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.UpdateOrderTable())
}
