package services

import (
	"github.com/google/uuid"
	"github.com/namduong/project-layout/helper"
	"github.com/namduong/project-layout/internal/models"
	"github.com/namduong/project-layout/internal/repositories"
)

type RestaurantService struct {
	RestaurantRepo *repositories.RestaurantRepository
}

func (s *RestaurantService) CreateRestaurant(name, address, code string) helper.Response {
	restaurant, err := s.RestaurantRepo.FindByCode(code)
	if err == nil && restaurant != nil {
		return helper.BuildErrorResponse("Restaurant already exists", "restaurant already exists", nil)
	}
	ID := uuid.New().String()
	newRestaurant := &models.Restaurant{
		ID:      ID,
		Name:    name,
		Address: address,
		Code:    code,
	}
	err = s.RestaurantRepo.Create(newRestaurant)
	if err != nil {
		return helper.BuildErrorResponse("Failed to create restaurant", err.Error(), nil)
	}
	return helper.BuildResponse(true, "Restaurant created successfully", newRestaurant)
}

func (s *RestaurantService) GetAllRestaurants(page, pageSize int) helper.Response {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	restaurants, total, err := s.RestaurantRepo.GetAll(page, pageSize)
	if err != nil {
		return helper.BuildErrorResponse("Failed to get restaurants", err.Error(), nil)
	}

	data := map[string]interface{}{
		"restaurants": restaurants,
		"pagination": map[string]interface{}{
			"page":        page,
			"size":        pageSize,
			"total":       total,
			"total_pages": (int(total) + pageSize - 1) / pageSize,
		},
	}

	return helper.BuildResponse(true, "Restaurants retrieved successfully", data)
}

func (s *RestaurantService) GetRestaurantByID(id string) helper.Response {
	restaurant, err := s.RestaurantRepo.FindByID(id)
	if err != nil {
		return helper.BuildErrorResponse("Restaurant not found", err.Error(), nil)
	}
	return helper.BuildResponse(true, "Restaurant retrieved successfully", restaurant)
}

func (s *RestaurantService) UpdateRestaurant(id string, updates map[string]interface{}) helper.Response {
	// Check if restaurant exists
	_, err := s.RestaurantRepo.FindByID(id)
	if err != nil {
		return helper.BuildErrorResponse("Restaurant not found", err.Error(), nil)
	}

	// Remove empty values and nil
	cleanUpdates := make(map[string]interface{})
	for k, v := range updates {
		if v != nil && v != "" {
			cleanUpdates[k] = v
		}
	}

	if len(cleanUpdates) == 0 {
		return helper.BuildErrorResponse("No valid fields to update", "no valid fields provided", nil)
	}

	err = s.RestaurantRepo.Update(id, cleanUpdates)
	if err != nil {
		return helper.BuildErrorResponse("Failed to update restaurant", err.Error(), nil)
	}

	// Get updated restaurant
	restaurant, _ := s.RestaurantRepo.FindByID(id)
	return helper.BuildResponse(true, "Restaurant updated successfully", restaurant)
}

func (s *RestaurantService) DeleteRestaurant(id string) helper.Response {
	// Check if restaurant exists
	_, err := s.RestaurantRepo.FindByID(id)
	if err != nil {
		return helper.BuildErrorResponse("Restaurant not found", err.Error(), nil)
	}

	err = s.RestaurantRepo.Delete(id)
	if err != nil {
		return helper.BuildErrorResponse("Failed to delete restaurant", err.Error(), nil)
	}

	return helper.BuildResponse(true, "Restaurant deleted successfully", nil)
}

func (s *RestaurantService) SearchRestaurants(query string, page, pageSize int) helper.Response {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	restaurants, total, err := s.RestaurantRepo.Search(query, page, pageSize)
	if err != nil {
		return helper.BuildErrorResponse("Failed to search restaurants", err.Error(), nil)
	}

	data := map[string]interface{}{
		"restaurants": restaurants,
		"pagination": map[string]interface{}{
			"page":        page,
			"size":        pageSize,
			"total":       total,
			"total_pages": (int(total) + pageSize - 1) / pageSize,
		},
		"query": query,
	}

	return helper.BuildResponse(true, "Restaurants found successfully", data)
}

func NewRestaurantService(restaurantRepo *repositories.RestaurantRepository) RestaurantServiceInterface {
	return &RestaurantService{
		RestaurantRepo: restaurantRepo,
	}
}

type RestaurantServiceInterface interface {
	CreateRestaurant(name, address, code string) helper.Response
	GetAllRestaurants(page, pageSize int) helper.Response
	GetRestaurantByID(id string) helper.Response
	UpdateRestaurant(id string, updates map[string]interface{}) helper.Response
	DeleteRestaurant(id string) helper.Response
	SearchRestaurants(query string, page, pageSize int) helper.Response
}
