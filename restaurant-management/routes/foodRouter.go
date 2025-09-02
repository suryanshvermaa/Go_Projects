package routes

import (
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods")
	incomingRoutes.GET("/foods/:food_id")
	incomingRoutes.POST("/foods")
	incomingRoutes.PATCH("/foods")
}
