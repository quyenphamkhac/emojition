package models

import (
	"time"

	"gorm.io/gorm"
)

// Product Model
type Product struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Code      string     `json:"code"`
	Price     uint       `json:"price"`
}

// CreateProduct func
func (p *Product) CreateProduct(db *gorm.DB) (*Product, error) {
	result := db.Create(p)
	if result.Error != nil {
		return &Product{}, result.Error
	}
	return p, nil
}
