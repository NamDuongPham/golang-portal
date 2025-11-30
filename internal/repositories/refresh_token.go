package repositories

import (
	"time"

	"github.com/namduong/project-layout/internal/models"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func (r *RefreshTokenRepository) Create(refreshToken *models.RefreshToken) error {
	return r.db.Create(refreshToken).Error
}
func (r *RefreshTokenRepository) FindByUserID(userID string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	if err := r.db.Model(&models.RefreshToken{}).
		Where("user_id = ? AND deleted_at IS NULL", userID).
		First(&refreshToken).Error; err != nil {
		return nil, err
	}
	return &refreshToken, nil
}
func (r *RefreshTokenRepository) DeleteByUserID(userID string) error {
	return r.db.Model(&models.RefreshToken{}).
		Where("user_id = ?", userID).
		Update("deleted_at", time.Now()).Error
}
func (r *RefreshTokenRepository) FindByID(id string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	if err := r.db.Model(&models.RefreshToken{}).
		Where("id = ? AND deleted_at IS NULL", id).
		First(&refreshToken).Error; err != nil {
		return nil, err
	}
	return &refreshToken, nil
}
func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}
