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

func findFileById(id int) *models.File {
	for i, file := range storage.Files {
		if file.ID == id {
			return &storage.Files[i]
		}
	}
	return nil
}

func assignFileToDev(name string, fileId int) {
	if _, exists := storage.Developers[name]; !exists {
		storage.Developers[name] = &models.Developer{
			Name:  name,
			Files: []int{},
		}
	}
	// Assign file to developer
	developer := storage.Developers[name]
	for _, assignedFile := range developer.Files {
		if assignedFile == fileId {
			return
		}
	}
	developer.Files = append(developer.Files, fileId)
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
	for _, id := range req.FileIds {
		file := findFileById(id)
		if file != nil {
			file.Developer = req.Developer
			assignFileToDev(req.Developer, id)
		}
	}
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

