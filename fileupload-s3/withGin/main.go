package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"sucess":  true,
			"message": "healty",
		})
	})
	r.Run(":8000")
}
