package repositories

import (
	"github.com/namduong/project-layout/internal/models"
	"gorm.io/gorm"
)

type PortalRepository struct {
	db *gorm.DB
}

func NewPortalRepository(db *gorm.DB) *PortalRepository {
	return &PortalRepository{db: db}
}

func (r *PortalRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
