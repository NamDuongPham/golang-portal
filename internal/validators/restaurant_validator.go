package validators

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRestaurantRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Code    string `json:"code" binding:"required"`
}

type UpdateRestaurantRequest struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
	Code    *string `json:"code"`
}

func ValidateCreateRestaurantRequest(c *gin.Context) (*CreateRestaurantRequest, bool) {
	var req CreateRestaurantRequest
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

func ValidateUpdateRestaurantRequest(c *gin.Context) (*UpdateRestaurantRequest, bool) {
	var req UpdateRestaurantRequest
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
