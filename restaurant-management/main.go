package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/suryanshvermaa/restaurant-management/helpers"
	"github.com/suryanshvermaa/restaurant-management/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in loading .env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	// auth middleware

	// health route
	router.GET("/health", func(ctx *gin.Context) {
		helpers.JsonResponse(ctx, 200, "healty", nil)
	})
	// all routes setup
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.NoteRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
