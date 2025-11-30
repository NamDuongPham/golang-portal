package services

import (
	"github.com/google/uuid"
	"github.com/namduong/project-layout/helper"
	"github.com/namduong/project-layout/internal/models"
	"github.com/namduong/project-layout/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo       *repositories.UserRepository
	RestaurantRepo *repositories.RestaurantRepository
}

func (s *UserService) CreateUser(username, password, confirmPassword, restaurantID string) helper.Response {
	if password != confirmPassword {
		return helper.BuildErrorResponse("Passwords do not match", "passwords do not match", nil)
	}

	restaurant, err := s.RestaurantRepo.FindByID(restaurantID)
	if err != nil || restaurant == nil {
		return helper.BuildErrorResponse("Restaurant not found", "restaurant not found", nil)
	}

	existingUser, err := s.UserRepo.FindByUsername(username)
	if err == nil && existingUser != nil {
		return helper.BuildErrorResponse("User already exists", "username already exists", nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return helper.BuildErrorResponse("Failed to hash password", err.Error(), nil)
	}

	ID := uuid.New().String()
	newUser := &models.User{
		ID:           ID,
		UserName:     username,
		Password:     string(hashedPassword),
		RestaurantID: restaurantID,
	}

	err = s.UserRepo.Create(newUser)
	if err != nil {
		return helper.BuildErrorResponse("Failed to create user", err.Error(), nil)
	}

	// Don't return password in response
	newUser.Password = ""
	return helper.BuildResponse(true, "User created successfully", newUser)
}

func (s *UserService) GetAllUsers(page, pageSize int) helper.Response {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	users, total, err := s.UserRepo.GetAll(page, pageSize)
	if err != nil {
		return helper.BuildErrorResponse("Failed to get users", err.Error(), nil)
	}

	// Remove passwords from response
	for i := range users {
		users[i].Password = ""
	}

	data := map[string]interface{}{
		"users": users,
		"pagination": map[string]interface{}{
			"page":        page,
			"size":        pageSize,
			"total":       total,
			"total_pages": (int(total) + pageSize - 1) / pageSize,
		},
	}

	return helper.BuildResponse(true, "Users retrieved successfully", data)
}

func (s *UserService) GetUsersByRestaurantID(restaurantID string, page, pageSize int) helper.Response {
	restaurant, err := s.RestaurantRepo.FindByID(restaurantID)
	if err != nil || restaurant == nil {
		return helper.BuildErrorResponse("Restaurant not found", "restaurant not found", nil)
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	users, total, err := s.UserRepo.GetByRestaurantID(restaurantID, page, pageSize)
	if err != nil {
		return helper.BuildErrorResponse("Failed to get users", err.Error(), nil)
	}

	// Remove passwords from response
	for i := range users {
		users[i].Password = ""
	}

	data := map[string]interface{}{
		"users": users,
		"pagination": map[string]interface{}{
			"page":        page,
			"size":        pageSize,
			"total":       total,
			"total_pages": (int(total) + pageSize - 1) / pageSize,
		},
	}

	return helper.BuildResponse(true, "Users retrieved successfully", data)
}

func (s *UserService) GetUserByID(id string) helper.Response {
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		return helper.BuildErrorResponse("User not found", err.Error(), nil)
	}

	// Don't return password
	user.Password = ""
	return helper.BuildResponse(true, "User retrieved successfully", user)
}

func (s *UserService) UpdateUser(id string, updates map[string]interface{}) helper.Response {
	_, err := s.UserRepo.FindByID(id)
	if err != nil {
		return helper.BuildErrorResponse("User not found", err.Error(), nil)
	}

	if password, ok := updates["password"].(string); ok && password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return helper.BuildErrorResponse("Failed to hash password", err.Error(), nil)
		}
		updates["password"] = string(hashedPassword)
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

	err = s.UserRepo.Update(id, cleanUpdates)
	if err != nil {
		return helper.BuildErrorResponse("Failed to update user", err.Error(), nil)
	}

	user, _ := s.UserRepo.FindByID(id)
	user.Password = ""
	return helper.BuildResponse(true, "User updated successfully", user)
}

func (s *UserService) DeleteUser(id string) helper.Response {
	_, err := s.UserRepo.FindByID(id)
	if err != nil {
		return helper.BuildErrorResponse("User not found", err.Error(), nil)
	}

	err = s.UserRepo.Delete(id)
	if err != nil {
		return helper.BuildErrorResponse("Failed to delete user", err.Error(), nil)
	}

	return helper.BuildResponse(true, "User deleted successfully", nil)
}

func (s *UserService) SearchUsers(query string, page, pageSize int) helper.Response {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	users, total, err := s.UserRepo.Search(query, page, pageSize)
	if err != nil {
		return helper.BuildErrorResponse("Failed to search users", err.Error(), nil)
	}

	// Remove passwords from response
	for i := range users {
		users[i].Password = ""
	}

	data := map[string]interface{}{
		"users": users,
		"pagination": map[string]interface{}{
			"page":        page,
			"size":        pageSize,
			"total":       total,
			"total_pages": (int(total) + pageSize - 1) / pageSize,
		},
		"query": query,
	}

	return helper.BuildResponse(true, "Users found successfully", data)
}

func NewUserService(userRepo *repositories.UserRepository, restaurantRepo *repositories.RestaurantRepository) UserServiceInterface {
	return &UserService{
		UserRepo:       userRepo,
		RestaurantRepo: restaurantRepo,
	}
}

type UserServiceInterface interface {
	CreateUser(username, password, confirmPassword, restaurantID string) helper.Response
	GetAllUsers(page, pageSize int) helper.Response
	GetUsersByRestaurantID(restaurantID string, page, pageSize int) helper.Response
	GetUserByID(id string) helper.Response
	UpdateUser(id string, updates map[string]interface{}) helper.Response
	DeleteUser(id string) helper.Response
	SearchUsers(query string, page, pageSize int) helper.Response
}
