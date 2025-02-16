package main

import (
	"fmt"
	"log"
	"net/http"

	"album_of_the_day/config"
	"album_of_the_day/database"
	"album_of_the_day/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	err = database.InitDB(cfg.Database.ConnStr)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer database.CloseDB()

	http.Handle("/", http.FileServer(http.Dir("../front")))

	http.HandleFunc("/album-of-the-day", handlers.AlbumOfTheDayHandler)
	fmt.Println("Server found at http://localhost:8080/album-of-the-day")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Cound not start server: %v\n", err)
	}
}
