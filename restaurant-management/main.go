package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	// all routes setup

	router.Run(":" + port)
}
