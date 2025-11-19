package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/http"
	"github.com/namduong/project-layout/internal/services"
)

func InitRouter(authService services.AuthService) *gin.Engine {
	router := gin.Default()

	adminHandler := http.AdminHandler{
		AuthService: authService,
	}

	authRoutes := router.Group("api/v1/auth")
	{
		authRoutes.POST("/login", adminHandler.Login)
	}

	return router
}
