package handlers

import (
	"myapp/internal/models"
	"myapp/storage"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var idCounter = 1
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
	file.ID = idCounter
	idCounter++
	storage.AddFile(&file)
	mutex.Unlock()
	c.JSON(http.StatusCreated, file)
}

func AssignFiles(c *gin.Context) {
	var req struct {
		Developer string `json:"developer"`
		FileIds   []int  `json:"fileIds"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	storage.AssignFiles(req.Developer, req.FileIds)
	c.JSON(http.StatusOK, gin.H{"message": "Files assigned successfully"})
}

func DeleteFile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	err = storage.DeleteFile(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File deleted"})
}

