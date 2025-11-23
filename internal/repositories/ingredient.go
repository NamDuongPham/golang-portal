package repositories

import (
	"time"

	"github.com/namduong/project-layout/internal/models"
	"gorm.io/gorm"
)

type IngredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) *IngredientRepository {
	return &IngredientRepository{db: db}
}

func (r *IngredientRepository) Create(ingredient *models.Ingredient) error {
	return r.db.Create(ingredient).Error
}

func (r *IngredientRepository) FindByCode(code string) (*models.Ingredient, error) {
	var ingredient models.Ingredient
	if err := r.db.Where("code = ?", code).First(&ingredient).Error; err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (r *IngredientRepository) FindByID(id string) (*models.Ingredient, error) {
	var ingredient models.Ingredient
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&ingredient).Error; err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (r *IngredientRepository) GetAll(page, pageSize int) ([]models.Ingredient, int64, error) {
	var ingredient []models.Ingredient
	var total int64

	offset := (page - 1) * pageSize

	query := r.db.Model(&models.Ingredient{}).Where("deleted_at IS NULL")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&ingredient).Error; err != nil {
		return nil, 0, err
	}

	return ingredient, total, nil
}

func (r *IngredientRepository) Update(id string, updates map[string]interface{}) error {
	return r.db.Model(&models.Ingredient{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(updates).Error
}

func (r *IngredientRepository) Delete(id string) error {
	return r.db.Model(&models.Ingredient{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", time.Now()).Error
}

func (r *IngredientRepository) Search(query string, page, pageSize int) ([]models.Ingredient, int64, error) {
	var ingredient []models.Ingredient
	var total int64

	offset := (page - 1) * pageSize

	dbQuery := r.db.Model(&models.Ingredient{}).
		Where("deleted_at IS NULL").
		Where("name ILIKE ? OR address ILIKE ? OR code ILIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%")

	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := dbQuery.Offset(offset).Limit(pageSize).Find(&ingredient).Error; err != nil {
		return nil, 0, err
	}

	return ingredient, total, nil
}
