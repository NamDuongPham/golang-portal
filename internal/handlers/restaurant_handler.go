package http

import (
	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/services"
)

type RestaurantHandler struct {
	AuthService services.AuthServiceInterface
}

func (h *RestaurantHandler) Create(c *gin.Context) {

}
