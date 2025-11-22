package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/services"
	"github.com/namduong/project-layout/internal/validators"
	"github.com/namduong/project-layout/utils"
)

type RestaurantHandler struct {
	RestaurantService services.RestaurantServiceInterface
}

func (h *RestaurantHandler) Create(c *gin.Context) {
	body, ok := validators.ValidateCreateRestaurantRequest(c)
	if !ok {
		return
	}
	res := h.RestaurantService.CreateRestaurant(body.Name, body.Address, body.Code)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *RestaurantHandler) GetAll(c *gin.Context) {
	page, pageSize := utils.GetPaginationParams(c)
	res := h.RestaurantService.GetAllRestaurants(page, pageSize)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *RestaurantHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "restaurant id is required",
		})
		return
	}

	res := h.RestaurantService.GetRestaurantByID(id)
	if !res.Status {
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *RestaurantHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "restaurant id is required",
		})
		return
	}

	body, ok := validators.ValidateUpdateRestaurantRequest(c)
	if !ok {
		return
	}

	updates := make(map[string]interface{})
	if body.Name != nil {
		updates["name"] = *body.Name
	}
	if body.Address != nil {
		updates["address"] = *body.Address
	}
	if body.Code != nil {
		updates["code"] = *body.Code
	}

	res := h.RestaurantService.UpdateRestaurant(id, updates)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *RestaurantHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "restaurant id is required",
		})
		return
	}

	res := h.RestaurantService.DeleteRestaurant(id)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *RestaurantHandler) Search(c *gin.Context) {
	query := utils.GetSearchQuery(c)
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "search query is required",
		})
		return
	}

	page, pageSize := utils.GetPaginationParams(c)
	res := h.RestaurantService.SearchRestaurants(query, page, pageSize)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
