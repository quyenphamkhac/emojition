package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/quyenphamkhac/emojition/config"
	"github.com/quyenphamkhac/emojition/db"
	"github.com/quyenphamkhac/emojition/migration"
	"github.com/quyenphamkhac/emojition/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config.Init(os.Getenv("ENV"))
	db.Init()
	migration.Init()
	server.Init()
}
