package handlers

import (
	"myapp/internal/models"
	"myapp/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDevelopers(c *gin.Context) {
	developers := []models.Developer{}
	for _, dev := range storage.Developers {
		developers = append(developers, *dev)
	}
	c.JSON(http.StatusOK, developers)
}

