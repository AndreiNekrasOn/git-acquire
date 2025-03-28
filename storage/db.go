package storage

import (
	"myapp/internal/models"
	"sync"
)

var (
	Files  = []models.File{}
	mutex  = &sync.Mutex{}
	Developers = make(map[string]*models.Developer) // New storage
)

// InitDB initializes fake data (optional)
func InitDB() {
	mutex.Lock()
	defer mutex.Unlock()
	Files = []models.File{} // Empty DB on start
}

