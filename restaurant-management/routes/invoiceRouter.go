package routes

import "github.com/gin-gonic/gin"

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices")
	incomingRoutes.GET("/invoices/:invoice_id")
	incomingRoutes.POST("/invoices")
	incomingRoutes.PATCH("/invoices/:invoice_id")
}
