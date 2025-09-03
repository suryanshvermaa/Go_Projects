package routes

import "github.com/gin-gonic/gin"

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus")
	incomingRoutes.GET("/menus/:menu_id")
	incomingRoutes.POST("/menus")
	incomingRoutes.PATCH("/menus/:menu_id")
}
