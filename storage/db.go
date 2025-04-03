package storage

import (
	"database/sql"
	"errors"
	"log"
	"myapp/internal/models"
	"sync"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB   *sql.DB
	mutex = &sync.Mutex{}
)

// InitDB initializes SQLite database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./devtracker.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	_, err = DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("Failed to enable foreign keys:", err)
	}
	// Create tables if they don't exist
	createTables()
}

// createTables ensures the tables exist in the database
func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS files (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS devs (
		name TEXT PRIMARY KEY
	);
	CREATE TABLE IF NOT EXISTS file2dev (
		file_id INTEGER NOT NULL,
		dev_name TEXT NOT NULL,
		UNIQUE(file_id, dev_name),
		FOREIGN KEY (file_id) REFERENCES files (id) ON DELETE CASCADE,
		FOREIGN KEY (dev_name) REFERENCES devs (name) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS users (
		name TEXT NOT NULL,
		password TEXT NOT NULL
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}
}

// GetFiles retrieves all files from the database
func GetFiles() []models.File {
	rows, err := DB.Query(`
	SELECT f.id, f.name, d.name FROM files f
	LEFT JOIN file2dev f2d on f2d.file_id = f.id
	LEFT JOIN devs d on d.name = f2d.dev_name;
	`)
	files := []models.File{}
	if err != nil {
		log.Println("Error retrieving files:", err)
		return files
	}
	defer rows.Close()

	for rows.Next() {
		var file models.File
		var dev sql.NullString;
		err := rows.Scan(&file.ID, &file.Name, &dev)
		if dev.Valid {
			file.Developer = dev.String;
		}
		if err != nil {
			log.Println("Error scanning file:", err)
			continue
		}
		files = append(files, file)
	}
	log.Println(files)
	return files
}

// AddFile inserts a new file into the database
func AddFile(file *models.File) {
	_, err := DB.Exec("INSERT INTO files (name) VALUES (?);", file.Name)
	if err != nil {
		log.Println("Error inserting file:", err)
	}
}

func fileHasDev(id int) *models.File {
	var file models.File
	err := DB.QueryRow(`SELECT f2d.file_id FROM file2dev f2d WHERE f2d.file_id = ?;`, id).Scan(&file.ID)
	if err != nil {
		return nil
	}
	return &file
}

func AssignFiles(devName string, fileIds []int) {
	// AssignFiles updates file ownership
	mutex.Lock()
	defer mutex.Unlock()
	log.Println(fileIds)
	_ = GetDeveloperByName(devName)
	for _, id := range fileIds {
		file := fileHasDev(id)
		var query string;
		if file == nil {
			query = "INSERT INTO file2dev (dev_name, file_id) VALUES (?, ?)"
		} else {
			query = "UPDATE file2dev SET dev_name = ? WHERE file_id = ?"
		}
		log.Println("Running query: ", query)
		log.Println("for file: ", file)
		_, err := DB.Exec(query, devName, id)
		if err != nil {
			log.Println("Error assigning file:", err)
		}
	}
}

// DeleteFile removes a file from the database
func DeleteFile(fileId int) error {
	mutex.Lock()
	defer mutex.Unlock()
	result, err := DB.Exec("DELETE FROM files WHERE id = ?;", fileId)
	if err != nil {
		return errors.New("failed to delete file")
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("file not found")
	}

	return nil
}

func GetDevelopers() []models.Developer {
	developers := []models.Developer{}
	rows, err := DB.Query(`
	SELECT d.name, f.id, f.name FROM devs d
	LEFT JOIN file2dev f2d on f2d.dev_name == d.name
	LEFT JOIN files f on f.id == f2d.file_id;
	`)
	if err != nil {
		log.Println("Error retrieving developers:", err)
		return developers
	}
	defer rows.Close()
	for rows.Next() {
		var dev models.Developer
		var file models.File
		err := rows.Scan(&dev.Name, &file.ID, &file.Name)
		if err != nil {
			log.Println("Error scanning developer:", err)
			continue
		}
		contains := false
		for _, d := range developers {
			if d.Name == dev.Name {
				contains = true
				d.Files = append(d.Files, file.ID)
			}
		}
		if !contains {
			dev.Files = append(dev.Files, file.ID)
			developers = append(developers, dev)
		}
		log.Println(developers)
	}
	return developers
}

// GetDeveloperByName fetches a developer or creates one
func GetDeveloperByName(name string) *models.Developer {
	var dev models.Developer
	err := DB.QueryRow("SELECT name FROM devs WHERE name = ?;", name).Scan(&dev.Name)
	if err != nil {
		// Developer not found, create one
		_, err := DB.Exec("INSERT INTO devs (name) VALUES (?);", name)
		if err != nil {
			log.Println("Error inserting developer:", err)
			return nil
		}
		dev.Name = name
	}
	return &dev
}

func ContainsUser(user string, password string) bool {
	var correctPassword string
	err := DB.QueryRow("SELECT password FROM users WHERE name = ?;", user).Scan(&correctPassword)
	if err == nil {
		return false
	}
	return correctPassword == password
}

func GetAccounts() gin.Accounts {
	accounts := gin.Accounts{}
	rows, err := DB.Query("SELECT name, password from users;")
	if err != nil {
		log.Println("Error retrieving developers:", err)
		return accounts
	}
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Name, &user.Password)
		accounts[user.Name] = user.Password
	}
	return accounts
}

