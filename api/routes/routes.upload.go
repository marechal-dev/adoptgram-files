package routes

import "github.com/gin-gonic/gin"

func SetupUploadRoutes(router *gin.Engine) {
	routesGroup := router.Group("/api")
	routesGroup.POST("/upload/media", func(ctx *gin.Context) {})
}
