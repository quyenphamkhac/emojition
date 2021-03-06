package models

import (
	"errors"
	"time"

	"github.com/quyenphamkhac/emojition/db"
	"golang.org/x/crypto/bcrypt"

	"github.com/quyenphamkhac/emojition/forms"
)

// User ...
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserModel ...
type UserModel struct{}

var authModel = new(AuthModel)

// Login ...
func (u *UserModel) Login(input forms.LoginInput) (AccessToken, error) {
	var err error
	var user User
	db := db.GetDB()
	userCheck := db.Where("username = ?", input.Username).First(&user)
	if userCheck.Error != nil {
		return AccessToken{}, userCheck.Error
	}
	bytePassword := []byte(input.Password)
	byteHashedPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		return AccessToken{}, err
	}

	token, err := authModel.GenerateToken(user)
	if err != nil {
		return AccessToken{}, err
	}

	err = authModel.SaveToken(token)
	if err != nil {
		return AccessToken{}, err
	}
	return *token, nil
}

// SignUp ...
func (u *UserModel) SignUp(input forms.SignUpInput) (user User, err error) {
	db := db.GetDB()

	userCheck := db.Where("username = ?", input.Username).Find(&user)
	if userCheck.Error != nil {
		return user, userCheck.Error
	}

	if userCheck.RowsAffected > 0 {
		return user, errors.New("user already exists")
	}

	bytePassword := []byte(input.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user = User{
		Username: input.Username,
		Name:     input.Name,
		Password: string(hashedPassword),
	}

	result := db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, err
}
