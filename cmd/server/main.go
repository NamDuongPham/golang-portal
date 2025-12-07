package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/namduong/project-layout/internal/database"
	"github.com/namduong/project-layout/internal/logger"
	"github.com/namduong/project-layout/internal/repositories"
	"github.com/namduong/project-layout/internal/services"
	router "github.com/namduong/project-layout/router"
	"go.uber.org/zap"
)

func main() {
	if err := logger.InitLogger(); err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	log := logger.GetLogger()
	defer func() {
		_ = log.Sync()
	}()

	log.Info("Starting application")

	// Connect DB
	database.Connect()
	gormDB := database.GetGormDB()
	if gormDB == nil {
		log.Fatal("GORM DB is nil; ensure database.Connect() succeeded")
	}

	// Init repos
	adminRepo := repositories.NewAdminRepository(gormDB)
	refreshTokenRepo := repositories.NewRefreshTokenRepository(gormDB)
	restaurantRepo := repositories.NewRestaurantRepository(gormDB)
	ingredientRepo := repositories.NewIngredientRepository(gormDB)
	userRepo := repositories.NewUserRepository(gormDB)
	portalRepo := repositories.NewPortalRepository(gormDB)

	// Init services
	authPortalService := services.NewAuthPortalService(portalRepo, refreshTokenRepo)
	authService := services.NewAuthService(adminRepo, refreshTokenRepo)
	restaurantService := services.NewRestaurantService(restaurantRepo)
	ingredientService := services.NewIngredientService(ingredientRepo)
	userService := services.NewUserService(userRepo, restaurantRepo)

	r := router.InitRouter(authService, restaurantService, ingredientService, userService, authPortalService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Info("Server running", zap.String("address", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server startup error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Warn("Shutdown signal receive, shutting down server...")
	shutdownTimeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	go func() {
		log.Warn("Graceful shutdown initiated...")

		seconds := int(shutdownTimeout.Seconds())
		for i := seconds; i > 0; i-- {
			log.Warn("Server will force shutdown in", zap.Int("seconds_left", i))
			time.Sleep(1 * time.Second)
		}
	}()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", zap.Error(err))
	} else {
		log.Info("Server shutdown gracefully")
	}

	log.Info("Application stopped")
}
