package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/http"
	"github.com/namduong/project-layout/internal/services"
)

func AdminRouter(router *gin.RouterGroup, authService services.AuthServiceInterface) {
	adminHandler := http.AdminHandler{
		AuthService: authService,
	}

	authRoutes := router.Group("/auth/admin")
	{
		authRoutes.POST("/login", adminHandler.Login)
		authRoutes.POST("/logout", adminHandler.Logout)
		authRoutes.POST("/refresh-token", adminHandler.RefreshToken)
	}
}
