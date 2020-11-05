package models

import (
	"gorm.io/gorm"
)

// Product Model
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// CreateProduct func
func (p *Product) CreateProduct(db *gorm.DB) (*Product, error) {
	result := db.Create(p)
	if result.Error != nil {
		return &Product{}, result.Error
	}
	return p, nil
}
