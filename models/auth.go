package models

import (
	"time"

	"github.com/quyenphamkhac/emojition/db"

	"github.com/quyenphamkhac/emojition/config"

	"github.com/dgrijalva/jwt-go"
)

// AccessToken ...
type AccessToken struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	TTL       uint      `json:"ttl"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	ExpireAt  time.Time `json:"-"`
}

// AuthModel ...
type AuthModel struct{}

// GenerateToken ...
func (a *AuthModel) GenerateToken(user User) (*AccessToken, error) {
	c := config.GetConfig()
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["name"] = user.Name
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24 * 14).Unix()

	tokenInstance := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := c.GetString("auth.secret")
	token, err := tokenInstance.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	accessToken := &AccessToken{
		ID:     token,
		TTL:    1209600,
		UserID: user.ID,
	}

	return accessToken, nil
}

// SaveToken ...
func (a *AuthModel) SaveToken(accessToken *AccessToken) error {
	db := db.GetDB()
	accessToken.ExpireAt = time.Now().Add(time.Hour * 24 * 14)
	result := db.Create(accessToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
