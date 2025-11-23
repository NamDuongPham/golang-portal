package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/services"
	"github.com/namduong/project-layout/internal/validators"
	"github.com/namduong/project-layout/utils"
)

type IngredientHandler struct {
	IngredientService services.IngredientServiceInterface
}

func (h *IngredientHandler) Create(c *gin.Context) {
	body, ok := validators.ValidateCreateIngredientRequest(c)
	if !ok {
		return
	}
	res := h.IngredientService.CreateIngredient(body.Name, body.Unit, body.Code)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *IngredientHandler) GetAll(c *gin.Context) {
	page, pageSize := utils.GetPaginationParams(c)
	res := h.IngredientService.GetAllIngredient(page, pageSize)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *IngredientHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "restaurant id is required",
		})
		return
	}

	res := h.IngredientService.GetIngredientByID(id)
	if !res.Status {
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *IngredientHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "restaurant id is required",
		})
		return
	}

	body, ok := validators.ValidateUpdateIngredientRequest(c)
	if !ok {
		return
	}

	updates := make(map[string]interface{})
	if body.Name != nil {
		updates["name"] = *body.Name
	}
	if body.Unit != nil {
		updates["unit"] = *body.Unit
	}
	if body.Code != nil {
		updates["code"] = *body.Code
	}

	res := h.IngredientService.UpdateIngredient(id, updates)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *IngredientHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "restaurant id is required",
		})
		return
	}

	res := h.IngredientService.DeleteIngredient(id)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *IngredientHandler) Search(c *gin.Context) {
	query := utils.GetSearchQuery(c)
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "search query is required",
		})
		return
	}

	page, pageSize := utils.GetPaginationParams(c)
	res := h.IngredientService.SearchIngredient(query, page, pageSize)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
