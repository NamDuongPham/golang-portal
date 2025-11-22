package repositories

import (
	"github.com/namduong/project-layout/internal/models"
	"gorm.io/gorm"
)

type RestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) Create(restaurant *models.Restaurant) error {
	return r.db.Create(restaurant).Error
}
func (r *RestaurantRepository) FindByCode(code string) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := r.db.Where("code = ?", code).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}
