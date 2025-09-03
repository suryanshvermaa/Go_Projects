package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suryanshvermaa/restaurant-management/controllers"
)

func NoteRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes", controllers.GetNotes())
	incomingRoutes.GET("/notes/:note_id", controllers.GetNote())
	incomingRoutes.POST("/notes", controllers.CreateNote())
	incomingRoutes.PATCH("/notes/:note_id", controllers.UpdateNote())
}
