package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupUploadRoutes(router *gin.Engine) {
	routesGroup := router.Group("/api")
	routesGroup.POST("/upload/media", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["files[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload
			ctx.SaveUploadedFile(file, "../../temp")
		}

		ctx.JSON(http.StatusOK, gin.H{
			"filesIds": []string{},
		})
	})
}
