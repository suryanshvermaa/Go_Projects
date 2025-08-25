package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": true, "message": "healty server"})
	})
	r.Run(":3000")
}
