package api

import (
	"myapp/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS Middleware (this fixes your issue)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// Routes
	r.GET("/files", handlers.GetFiles)
	r.POST("/files", handlers.AddFile)
	r.PUT("/files/:id", handlers.UpdateFileDeveloper)
	r.PUT("/developers/:id", handlers.AssignFileToDeveloper)
	r.GET("/developers", handlers.GetDevelopers)


	return r
}

