package main

import (
	"log"

	"github.com/Akimpupupuu/ClearCity/db"
	"github.com/Akimpupupuu/ClearCity/internal/config"
	"github.com/Akimpupupuu/ClearCity/internal/http"
)

func main() {
	cfg := config.NewConfig()

	db, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServer(cfg, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
