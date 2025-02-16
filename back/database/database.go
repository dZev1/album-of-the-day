package database

import (
	"album_of_the_day/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

var db *sql.DB

func InitDB(connStr string) error {
	var err error

	db, err = sql.Open("mysql", connStr)
	if err != nil {
		return fmt.Errorf("Could not connect to database: %v.", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Could not verify connection: %v.", err)
	}

	fmt.Println("Connection to database has been established.")
	return nil
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func GetAlbumOfTheDay() (models.Album, error) {
	var album models.Album

	now := time.Now()
	dayOfTheYear := now.YearDay()
	currYear := now.Year()

	query := `
		SELECT artist, album, release_year, cover
		FROM albums
		ORDER BY id
		LIMIT 1 OFFSET ?
	`

	offset := ((dayOfTheYear - 1) + (currYear % GetTotalAlbums())) % GetTotalAlbums()
	row := db.QueryRow(query, offset)

	err := row.Scan(&album.Artist, &album.Album, &album.ReleaseYear, &album.Cover)
	if err != nil {
		return models.Album{}, fmt.Errorf("Could not fetch album: %v", err)
	}

	return album, nil
}

func GetTotalAlbums() int {
	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM albums").Scan(&count)
	if err != nil {
		log.Fatalf("Error al contar los álbumes: %v", err)
	}
	return count
}
