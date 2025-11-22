package main

import (
	"github.com/namduong/project-layout/internal/database"
	"github.com/namduong/project-layout/internal/repositories"
	"github.com/namduong/project-layout/internal/services"
	router "github.com/namduong/project-layout/router"
)

func main() {
	database.Connect()
	gormDB := database.GetGormDB()
	if gormDB == nil {
		panic("gorm DB is nil; ensure database.Connect() succeeded")
	}
	adminRepo := repositories.NewAdminRepository(gormDB)
	refreshTokenRepo := repositories.NewRefreshTokenRepository(gormDB)
	restaurantRepo := repositories.NewRestaurantRepository(gormDB)

	service := services.NewAuthService(adminRepo, refreshTokenRepo)
	restaurantService := services.NewRestaurantService(restaurantRepo)

	r := router.InitRouter(service, restaurantService)

	r.Run()
}
