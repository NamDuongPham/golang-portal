package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/middlewares"
	"github.com/namduong/project-layout/internal/services"
)

func InitRouter(authService services.AuthServiceInterface, restaurantService services.RestaurantServiceInterface, ingredientService services.IngredientServiceInterface) *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		AdminRouter(apiV1, authService)
		protected := apiV1.Group("")
		protected.Use(middlewares.AuthMiddleware())

		RestaurantRouter(protected.Group("/restaurants"), restaurantService)
		IngredientRouter(protected.Group("/ingredients"), ingredientService)

	}

	return router
}
