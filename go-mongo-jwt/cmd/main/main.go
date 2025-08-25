package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/suryanshvermaa/Go_Projects/go-mongo-jwt/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(ctx *gin.Context) {
		utils.Response(ctx, 200, "server is healty✅", nil)
	})

	router.Run(":" + port)
}
