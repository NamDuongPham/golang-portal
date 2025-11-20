package repositories

import (
	"github.com/namduong/project-layout/internal/models"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func (r *RefreshTokenRepository) Create(refreshToken *models.RefreshToken) error {
	return r.db.Create(refreshToken).Error
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}
