package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/services"
	"github.com/namduong/project-layout/internal/validators"
	"github.com/namduong/project-layout/utils"
)

type UserHandler struct {
	UserService services.UserServiceInterface
}

func (h *UserHandler) Create(c *gin.Context) {
	body, ok := validators.ValidateCreateUserRequest(c)
	if !ok {
		return
	}
	res := h.UserService.CreateUser(body.Username, body.Password, body.ConfirmPassword, body.RestaurantID)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	page, pageSize := utils.GetPaginationParams(c)
	res := h.UserService.GetAllUsers(page, pageSize)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetByRestaurantID(c *gin.Context) {
	restaurantID := c.Param("id")
	if restaurantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "restaurant id is required",
		})
		return
	}

	page, pageSize := utils.GetPaginationParams(c)
	res := h.UserService.GetUsersByRestaurantID(restaurantID, page, pageSize)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "user id is required",
		})
		return
	}

	res := h.UserService.GetUserByID(id)
	if !res.Status {
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "user id is required",
		})
		return
	}

	body, ok := validators.ValidateUpdateUserRequest(c)
	if !ok {
		return
	}

	updates := make(map[string]interface{})
	if body.Username != nil {
		updates["username"] = *body.Username
	}
	if body.Password != nil {
		updates["password"] = *body.Password
	}
	if body.RestaurantID != nil {
		updates["restaurant_id"] = *body.RestaurantID
	}

	res := h.UserService.UpdateUser(id, updates)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "user id is required",
		})
		return
	}

	res := h.UserService.DeleteUser(id)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Search(c *gin.Context) {
	query := utils.GetSearchQuery(c)
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "search query is required",
		})
		return
	}

	page, pageSize := utils.GetPaginationParams(c)
	res := h.UserService.SearchUsers(query, page, pageSize)
	if !res.Status {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
