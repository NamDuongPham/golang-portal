package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/services"
	"github.com/namduong/project-layout/internal/validators"
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
