package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/namduong/project-layout/internal/handlers"
	"github.com/namduong/project-layout/internal/middlewares"
	"github.com/namduong/project-layout/internal/services"
)

func AdminRouter(router *gin.RouterGroup, authService services.AuthServiceInterface) {
	adminHandler := handlers.AdminHandler{AuthService: authService}

	authRoutes := router.Group("/auth/admin")

	authRoutes.POST("/login", adminHandler.Login)
	authRoutes.POST("/refresh-token", adminHandler.RefreshToken)

	authRoutes.Use(middlewares.AuthMiddleware())
	{
		authRoutes.POST("/logout", adminHandler.Logout)
	}
}
