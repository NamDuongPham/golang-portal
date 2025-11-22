package repositories

import (
	"time"

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

func (r *RestaurantRepository) FindByID(id string) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (r *RestaurantRepository) GetAll(page, pageSize int) ([]models.Restaurant, int64, error) {
	var restaurants []models.Restaurant
	var total int64

	offset := (page - 1) * pageSize

	query := r.db.Model(&models.Restaurant{}).Where("deleted_at IS NULL")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&restaurants).Error; err != nil {
		return nil, 0, err
	}

	return restaurants, total, nil
}

func (r *RestaurantRepository) Update(id string, updates map[string]interface{}) error {
	return r.db.Model(&models.Restaurant{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(updates).Error
}

func (r *RestaurantRepository) Delete(id string) error {
	return r.db.Model(&models.Restaurant{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", time.Now()).Error
}

func (r *RestaurantRepository) Search(query string, page, pageSize int) ([]models.Restaurant, int64, error) {
	var restaurants []models.Restaurant
	var total int64

	offset := (page - 1) * pageSize

	dbQuery := r.db.Model(&models.Restaurant{}).
		Where("deleted_at IS NULL").
		Where("name ILIKE ? OR address ILIKE ? OR code ILIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%")

	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := dbQuery.Offset(offset).Limit(pageSize).Find(&restaurants).Error; err != nil {
		return nil, 0, err
	}

	return restaurants, total, nil
}
