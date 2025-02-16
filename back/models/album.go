package models

type Album struct {
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	ReleaseYear int    `json:"release_year"`
	Cover       string `json:"cover"`
}
