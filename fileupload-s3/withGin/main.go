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
	r.POST("/file", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(400, gin.H{
				"success": false,
				"message": "Failed to upload file",
			})
			return
		}
		err = ctx.SaveUploadedFile(file, "./uploads/"+file.Filename)
		if err != nil {
			ctx.JSON(500, gin.H{
				"success": false,
				"message": "Failed to save file",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "File uploaded successfully",
			"file":    file.Filename,
		})
	})
	r.Run(":8000")
}
