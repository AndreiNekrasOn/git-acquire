package storage

import (
	"errors"
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

func GetFiles() []models.File {
	return Files
}

func AddFile(file *models.File) {
	Files = append(Files, *file)
}

func FindFileById(id int) *models.File {
	for i, file := range Files {
		if file.ID == id {
			return &Files[i]
		}
	}
	return nil
}

func RemoveFileFromDev(developer *models.Developer, fileId int) {
	for id, assignedFile := range developer.Files {
		if assignedFile == fileId {
			developer.Files = append(developer.Files[:id], developer.Files[id+1:]...)
		}
	}
}

func RemoveFileFromAll(fileId int) {
	for _, dev := range Developers {
		RemoveFileFromDev(dev, fileId)
	}
}

func GetDeveloperByName(name string) *models.Developer {
	if _, exists := Developers[name]; !exists {
		Developers[name] = &models.Developer{ Name:  name, Files: []int{}, }
	}
	return Developers[name]
}

func assignFile(devName string, fileId int) {
	RemoveFileFromAll(fileId)
	if (len(devName) == 0) {
		return
	}
	developer := GetDeveloperByName(devName)
	developer.Files = append(developer.Files, fileId)
}

func AssignFiles(devName string, fileIds []int) {
	for _, id := range fileIds {
		file := FindFileById(id)
		if file != nil {
			file.Developer = devName
			assignFile(devName, id)
		}
	}
}

func DeleteFile(fileId int) error {
	RemoveFileFromAll(fileId)
	for i, file := range Files {
		if file.ID == fileId {
			Files = append(Files[:i], Files[i+1:]...)
			return nil
		}
	}
	return errors.New("File not found")
}

func GetDevelopers() []models.Developer {
	developers := []models.Developer{}
	for _, dev := range Developers {
		developers = append(developers, *dev)
	}
	return developers
}
