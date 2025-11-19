package models

import "time"

type Admin struct {
	ID        string    `gorm:"column:id;primaryKey"`
	UserName  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

func (Admin) TableName() string {
	return "admin"
}
