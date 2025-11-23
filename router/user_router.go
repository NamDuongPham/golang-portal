package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/namduong/project-layout/internal/handlers"
	"github.com/namduong/project-layout/internal/services"
)

func UserRouter(router *gin.RouterGroup, userService services.UserServiceInterface) {

	userHandler := handlers.UserHandler{UserService: userService}

	{
		router.POST("/create", userHandler.Create)
		router.GET("/", userHandler.GetAll)
		router.GET("/restaurant/:id", userHandler.GetByRestaurantID)
		router.GET("/search", userHandler.Search)
		router.GET("/:id", userHandler.GetByID)
		router.PUT("/:id", userHandler.Update)
		router.DELETE("/:id", userHandler.Delete)
	}
}
