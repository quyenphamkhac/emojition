package models

import "gorm.io/gorm"

// User model
type User struct {
	ID    string `gorm:"primary_key;auto_increment" json:"id"`
	Name  string `gorm:"size:255;not null;unique" json:"name"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
}

// SaveUser func: Create & save user
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
