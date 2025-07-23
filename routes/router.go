package routes

import (
	"marine-ar-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Server static files (models & thumbnails)
	r.Static("/uploads", "./uploads")

	// API route
	r.GET("/api/models", controllers.GetModels)

	return r
}
