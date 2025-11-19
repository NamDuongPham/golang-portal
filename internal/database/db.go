package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	config "github.com/namduong/project-layout/configs"
	"github.com/namduong/project-layout/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *pgxpool.Pool
var db *gorm.DB

func Connect() {
	dsn := config.Cfg.Databases.Source.GetPostgresDSN()
	fmt.Println("Connecting to PostgreSQL with:", dsn)

	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	fmt.Println("Connected to PostgreSQL!")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	db.AutoMigrate(&models.Admin{}, &models.RefreshToken{})
}
