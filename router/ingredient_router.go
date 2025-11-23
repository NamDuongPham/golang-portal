package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/namduong/project-layout/internal/handlers"
	"github.com/namduong/project-layout/internal/services"
)

func IngredientRouter(router *gin.RouterGroup, ingredientService services.IngredientServiceInterface) {

	ingredientHandler := handlers.IngredientHandler{IngredientService: ingredientService}

	{
		router.POST("/create", ingredientHandler.Create)
		router.GET("/", ingredientHandler.GetAll)
		router.GET("/search", ingredientHandler.Search)
		router.GET("/:id", ingredientHandler.GetByID)
		router.PUT("/:id", ingredientHandler.Update)
		router.DELETE("/:id", ingredientHandler.Delete)
	}
}
