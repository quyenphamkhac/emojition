package models

import "time"

// AccessToken ...
type AccessToken struct {
	ID        string    `gorm:"primaryKey;unique" json:"id"`
	TTL       int       `json:"ttl"`
	CreatedAt time.Time `json:"created_at"`
	ExpireAt  time.Time `json:"expire_at"`
}
