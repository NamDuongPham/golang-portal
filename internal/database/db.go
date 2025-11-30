package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	config "github.com/namduong/project-layout/configs"
	"github.com/namduong/project-layout/internal/logger"
	"github.com/namduong/project-layout/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *pgxpool.Pool
var db *gorm.DB

func Connect() {
	dsn := config.Cfg.Databases.Source.GetPostgresDSN()
	logger.GetLogger().Info("Connecting to PostgreSQL", zap.String("dsn", dsn))

	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		logger.GetLogger().Fatal("Failed to connect to DB", zap.Error(err))
	}

	logger.GetLogger().Info("Connected to PostgreSQL successfully")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.GetLogger().Fatal("Failed to connect to DB with GORM", zap.Error(err))
	}
	db.AutoMigrate(&models.Admin{}, &models.RefreshToken{}, &models.Restaurant{}, &models.Ingredient{}, &models.User{})
	logger.GetLogger().Info("Database migration completed")
}

func GetGormDB() *gorm.DB {
	return db
}
