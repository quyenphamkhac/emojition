package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init func
func Init() {
	var err error
	dsn := "host=localhost user=emojition password=emojition dbname=emojition port=5432 sslmode=disable TimeZone=Asia/SaiGon"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection database error!")
	}
}

// GetDB func
func GetDB() *gorm.DB {
	return db
}
