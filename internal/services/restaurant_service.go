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
	return helper.BuildResponse(true, "Restaurant created successfully", restaurant)
}
func NewRestaurantService(restaurantRepo *repositories.RestaurantRepository) RestaurantServiceInterface {
	return &RestaurantService{
		RestaurantRepo: restaurantRepo,
	}
}

type RestaurantServiceInterface interface {
	CreateRestaurant(name, address, code string) helper.Response
}
