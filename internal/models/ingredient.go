package models

import "time"

type Ingredient struct {
	ID        string     `gorm:"column:id;primaryKey"`
	Name      string     `gorm:"column:name"`
	Code      string     `gorm:"column:code"`
	Unit      string     `gorm:"column:unit"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (Ingredient) TableName() string {
	return "ingredient"
}
