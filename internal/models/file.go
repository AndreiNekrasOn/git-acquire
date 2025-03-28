package models

type File struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Developer  string `json:"developer"`
	Branch     string `json:"branch"`
}
