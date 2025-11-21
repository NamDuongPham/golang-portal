package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/namduong/project-layout/internal/handlers"
	"github.com/namduong/project-layout/internal/services"
)

func RestaurantRouter(router *gin.RouterGroup, authService services.AuthServiceInterface) {
	restaurantHandler := handlers.RestaurantHandler{AuthService: authService}

	authRoutes := router.Group("/restaurant")

	{
		authRoutes.POST("/create", restaurantHandler.Create)
	}
}
