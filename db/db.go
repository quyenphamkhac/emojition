package db

import (
	"github.com/quyenphamkhac/emojition/models"
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
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Product{})
}

// GetDB func
func GetDB() *gorm.DB {
	return db
}
