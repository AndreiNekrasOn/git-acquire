package handlers

import (
	"myapp/internal/models"
	"myapp/storage"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var mutex = &sync.Mutex{}

func GetFiles(c *gin.Context) {
	c.JSON(http.StatusOK, storage.GetFiles())
}

func AddFile(c *gin.Context) {
	var file models.File
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mutex.Lock()
	storage.AddFile(&file)
	mutex.Unlock()
	c.JSON(http.StatusCreated, file)
}

func AssignFiles(c *gin.Context) {
	var req struct {
		Developer string `json:"developer"`
		FileNames []string `json:"fileNames"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	storage.AssignFiles(req.Developer, req.FileNames)
	c.JSON(http.StatusOK, gin.H{"message": "Files assigned successfully"})
}

func DeleteFile(c *gin.Context) {
	name := c.Param("name")
	mutex.Lock()
	defer mutex.Unlock()
	err := storage.DeleteFile(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File deleted"})
}

