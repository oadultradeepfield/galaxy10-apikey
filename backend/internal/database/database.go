package database

import (
	"log"
	"os"

	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		log.Fatalf("Database URL is not set in the environment variables")
	}

	db, err := gorm.Open(postgres.Open(database_url), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.APIKey{},
	); err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}

	log.Println("Database is connected and migrated successfully")
	return db
}
