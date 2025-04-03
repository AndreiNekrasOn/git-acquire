package models

type Developer struct {
	Name  string `json:"name"`
	Files []string `json:"files"`
}

