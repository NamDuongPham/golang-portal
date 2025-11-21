package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/services"
)

func InitRouter(authService services.AuthServiceInterface) *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		AdminRouter(apiV1, authService)
	}

	return router
}
