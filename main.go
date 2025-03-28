package main

import (
	"log"
	"myapp/api"
	"myapp/storage"
)

func main() {
	// Initialize storage (DB or in-memory)
	storage.InitDB()

	// Start server
	r := api.SetupRouter()
	log.Println("Server is running on :8080")
	r.Run(":8080")
}

