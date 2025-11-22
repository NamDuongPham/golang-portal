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
		router.GET("/", restaurantHandler.GetAll)
		router.GET("/search", restaurantHandler.Search)
		router.GET("/:id", restaurantHandler.GetByID)
		router.PUT("/:id", restaurantHandler.Update)
		router.DELETE("/:id", restaurantHandler.Delete)
	}
}
