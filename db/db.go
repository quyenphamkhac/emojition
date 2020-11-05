package db

import (
	"fmt"

	"github.com/quyenphamkhac/emojition/config"

	"github.com/quyenphamkhac/emojition/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init func
func Init() {
	var err error
	c := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", c.GetString("db.host"), c.GetString("db.user"), c.GetString("db.password"), c.GetString("db.database"), c.GetInt("db.port"), c.GetString("db.sslmode"), c.GetString("db.timezone"))
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
