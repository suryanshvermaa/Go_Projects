package routes

import "github.com/gin-gonic/gin"

func NoteRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes")
	incomingRoutes.GET("/notes/:note_id")
	incomingRoutes.POST("/notes")
	incomingRoutes.PATCH("/notes/:note_id")
}
