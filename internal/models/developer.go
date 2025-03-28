package models

type Developer struct {
	Name  string `json:"name"`
	Files []int  `json:"files"` // Store file IDs
}

