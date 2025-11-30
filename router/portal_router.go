package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/namduong/project-layout/internal/handlers"
	"github.com/namduong/project-layout/internal/middlewares"
	"github.com/namduong/project-layout/internal/services"
)

func PortalRouter(router *gin.RouterGroup, authPortalService services.AuthPortalServiceInterface) {
	portalHandler := handlers.PortalHandler{AuthPortalService: authPortalService}
	authRoutes := router.Group("/auth/portal")

	authRoutes.POST("/login", portalHandler.Login)
	authRoutes.POST("/refresh-token", portalHandler.RefreshToken)
	protected := authRoutes.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		authRoutes.POST("/logout", portalHandler.Logout)

	}
}
