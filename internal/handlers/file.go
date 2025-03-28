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
	c.JSON(http.StatusOK, storage.Files)
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
	storage.Files = append(storage.Files, file)
	mutex.Unlock()

	c.JSON(http.StatusCreated, file)
}

func UpdateFileDeveloper(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	var request struct {
		Developer string `json:"developer"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Check if file exists
	var fileExists bool
	for _, file := range storage.Files {
		if file.ID == id {
			fileExists = true
			break
		}
	}
	if !fileExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Add developer to storage if they don't exist
	if _, exists := storage.Developers[request.Developer]; !exists {
		storage.Developers[request.Developer] = &models.Developer{
			Name:  request.Developer,
			Files: []int{},
		}
	}

	// Assign file to developer
	developer := storage.Developers[request.Developer]
	for _, assignedFile := range developer.Files {
		if assignedFile == id {
			c.JSON(http.StatusOK, developer) // Already assigned, no need to update
			return
		}
	}
	developer.Files = append(developer.Files, id)
	for i, file := range storage.Files {
		if file.ID == id {
			storage.Files[i].Developer = developer.Name
		}
	}
	c.JSON(http.StatusOK, developer)
}

func DeleteFile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	// TODO: also remove file from developer
	for i, file := range storage.Files {
		if file.ID == id {
			storage.Files = append(storage.Files[:i], storage.Files[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "File deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
}

