package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/namduong/project-layout/internal/handlers"
	"github.com/namduong/project-layout/internal/services"
)

func RestaurantRouter(router *gin.RouterGroup, restaurantService services.RestaurantServiceInterface) {

	restaurantHandler := handlers.RestaurantHandler{RestaurantService: restaurantService}

	{
		router.POST("/create", restaurantHandler.Create)
	}
}
