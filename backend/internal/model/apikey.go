package model

import (
	"time"
)

type APIKey struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	APIKey    string    `gorm:"uniqueIndex;not null" json:"api_key"`
	UserID    string    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
