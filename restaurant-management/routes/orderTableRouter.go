package routes

import "github.com/gin-gonic/gin"

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables")
	incomingRoutes.GET("/tables/:table_id")
	incomingRoutes.POST("/tables")
	incomingRoutes.PATCH("/tables/:table_id")
}
