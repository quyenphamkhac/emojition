package migration

import (
	"github.com/quyenphamkhac/emojition/db"
	"github.com/quyenphamkhac/emojition/models"
)

// Init ...
func Init() {
	db := db.GetDB()
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})
}
