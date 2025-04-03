package api

import (
	"myapp/internal/handlers"
	"myapp/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS Middleware (this fixes your issue)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	accounts := storage.GetAccounts()

	// Routes
	r.GET("/files", handlers.GetFiles)
	r.POST("/files", handlers.AddFile)
	r.DELETE("/files/:id", handlers.DeleteFile)
	r.GET("/developers", handlers.GetDevelopers)
	r.POST("/assign", gin.BasicAuth(accounts), handlers.AssignFiles)
	r.POST("/login", handlers.Login)
	return r
}

