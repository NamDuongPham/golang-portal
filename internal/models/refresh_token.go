package models

import "time"

type RefreshToken struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Token     string    `gorm:"column:token"`
	UserID    string    `gorm:"column:user_id"`
	ExpiresAt time.Time `gorm:"column:expires_at"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

func (RefreshToken) TableName() string {
	return "refresh_token"
}
