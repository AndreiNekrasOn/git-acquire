package handlers

import (
	"myapp/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDevelopers(c *gin.Context) {
	c.JSON(http.StatusOK, storage.GetDevelopers())
}

