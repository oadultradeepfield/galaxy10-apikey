package model

import (
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	APIKeys   []APIKey  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"api_keys"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}