package models

import "time"

type User struct {
	ID           string     `gorm:"column:id;primaryKey"`
	UserName     string     `gorm:"column:username"`
	Password     string     `gorm:"column:password"`
	RestaurantID string     `gorm:"column:restaurant_id"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "user"
}
