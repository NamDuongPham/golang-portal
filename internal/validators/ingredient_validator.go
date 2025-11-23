package validators

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateIngredientRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
	Unit string `json:"unit" binding:"required"`
}

type UpdateIngredientRequest struct {
	Name *string `json:"name"`
	Unit *string `json:"unit" binding:"required"`
	Code *string `json:"code"`
}

func ValidateCreateIngredientRequest(c *gin.Context) (*CreateIngredientRequest, bool) {
	var req CreateIngredientRequest
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

func ValidateUpdateIngredientRequest(c *gin.Context) (*UpdateIngredientRequest, bool) {
	var req UpdateIngredientRequest
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
