package services

import (
	"github.com/google/uuid"
	"github.com/namduong/project-layout/helper"
	"github.com/namduong/project-layout/internal/models"
	"github.com/namduong/project-layout/internal/repositories"
)

type IngredientService struct {
	IngredientRepo *repositories.IngredientRepository
}

func (s *IngredientService) CreateIngredient(name, unit, code string) helper.Response {
	ingredient, err := s.IngredientRepo.FindByCode(code)
	if err == nil && ingredient != nil {
		return helper.BuildErrorResponse("Ingredient already exists", "Ingredient already exists", nil)
	}
	ID := uuid.New().String()
	newIngredient := &models.Ingredient{
		ID:   ID,
		Name: name,
		Unit: unit,
		Code: code,
	}
	err = s.IngredientRepo.Create(newIngredient)
	if err != nil {
		return helper.BuildErrorResponse("Failed to create Ingredient", err.Error(), nil)
	}
	return helper.BuildResponse(true, "Ingredient created successfully", newIngredient)
}

func (s *IngredientService) GetAllIngredient(page, pageSize int) helper.Response {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	ingredients, total, err := s.IngredientRepo.GetAll(page, pageSize)
	if err != nil {
		return helper.BuildErrorResponse("Failed to get ingredients", err.Error(), nil)
	}

	data := map[string]interface{}{
		"ingredients": ingredients,
		"pagination": map[string]interface{}{
			"page":        page,
			"size":        pageSize,
			"total":       total,
			"total_pages": (int(total) + pageSize - 1) / pageSize,
		},
	}

	return helper.BuildResponse(true, "Ingredients retrieved successfully", data)
}

func (s *IngredientService) GetIngredientByID(id string) helper.Response {
	ingredient, err := s.IngredientRepo.FindByID(id)
	if err != nil {
		return helper.BuildErrorResponse("Restaurant not found", err.Error(), nil)
	}
	return helper.BuildResponse(true, "Restaurant retrieved successfully", ingredient)
}

func (s *IngredientService) UpdateIngredient(id string, updates map[string]interface{}) helper.Response {
	_, err := s.IngredientRepo.FindByID(id)
	if err != nil {
		return helper.BuildErrorResponse("Ingredient not found", err.Error(), nil)
	}

	cleanUpdates := make(map[string]interface{})
	for k, v := range updates {
		if v != nil && v != "" {
			cleanUpdates[k] = v
		}
	}

	if len(cleanUpdates) == 0 {
		return helper.BuildErrorResponse("No valid fields to update", "no valid fields provided", nil)
	}

	err = s.IngredientRepo.Update(id, cleanUpdates)
	if err != nil {
		return helper.BuildErrorResponse("Failed to update ingredient", err.Error(), nil)
	}

	ingredient, _ := s.IngredientRepo.FindByID(id)
	return helper.BuildResponse(true, "Ingredient updated successfully", ingredient)
}

func (s *IngredientService) DeleteIngredient(id string) helper.Response {
	_, err := s.IngredientRepo.FindByID(id)
	if err != nil {
		return helper.BuildErrorResponse("Ingredient not found", err.Error(), nil)
	}

	err = s.IngredientRepo.Delete(id)
	if err != nil {
		return helper.BuildErrorResponse("Failed to delete ingredient", err.Error(), nil)
	}

	return helper.BuildResponse(true, "Ingredient deleted successfully", nil)
}

func (s *IngredientService) SearchIngredient(query string, page, pageSize int) helper.Response {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	ingredients, total, err := s.IngredientRepo.Search(query, page, pageSize)
	if err != nil {
		return helper.BuildErrorResponse("Failed to search ingredients", err.Error(), nil)
	}

	data := map[string]interface{}{
		"ingredients": ingredients,
		"pagination": map[string]interface{}{
			"page":        page,
			"size":        pageSize,
			"total":       total,
			"total_pages": (int(total) + pageSize - 1) / pageSize,
		},
		"query": query,
	}

	return helper.BuildResponse(true, "Ingredient found successfully", data)
}

func NewIngredientService(IngredientRepo *repositories.IngredientRepository) IngredientServiceInterface {
	return &IngredientService{
		IngredientRepo: IngredientRepo,
	}
}

type IngredientServiceInterface interface {
	CreateIngredient(name, unit, code string) helper.Response
	GetAllIngredient(page, pageSize int) helper.Response
	GetIngredientByID(id string) helper.Response
	UpdateIngredient(id string, updates map[string]interface{}) helper.Response
	DeleteIngredient(id string) helper.Response
	SearchIngredient(query string, page, pageSize int) helper.Response
}
