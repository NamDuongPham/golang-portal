package validators

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	RestaurantID    string `json:"restaurant_id" binding:"required"`
}

type UpdateUserRequest struct {
	Username     *string `json:"username"`
	Password     *string `json:"password"`
	RestaurantID *string `json:"restaurant_id"`
}

func ValidateCreateUserRequest(c *gin.Context) (*CreateUserRequest, bool) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return nil, false
	}
	return &req, true
}

func ValidateUpdateUserRequest(c *gin.Context) (*UpdateUserRequest, bool) {
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return nil, false
	}
	return &req, true
}
