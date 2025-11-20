package main

import (
	"github.com/namduong/project-layout/internal/database"
	"github.com/namduong/project-layout/internal/repositories"
	"github.com/namduong/project-layout/internal/services"
	"github.com/namduong/project-layout/routes"
)

func main() {
	database.Connect()
	gormDB := database.GetGormDB()
	if gormDB == nil {
		panic("gorm DB is nil; ensure database.Connect() succeeded")
	}
	adminRepo := repositories.NewAdminRepository(gormDB)
	refreshTokenRepo := repositories.NewRefreshTokenRepository(gormDB)

	service := services.NewAuthService(adminRepo, refreshTokenRepo)

	r := routes.InitRouter(service)

	r.Run()
}
