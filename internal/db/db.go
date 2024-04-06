package db

import (
	"fmt"

	"example.com/todo-app/config"
	"example.com/todo-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config config.Configuration) (*gorm.DB, error) {

	dbDSN := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT, config.DB_SSLMODE, config.DB_TIMEZONE)

	db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to database")
	}

	fmt.Println("Successfully connected to database ...")

	if err := migrations(db); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	DB = db
	return db, nil
}

func migrations(db *gorm.DB) error {

	if err := db.AutoMigrate(&models.Todo{}, &models.User{}); err != nil {
		return fmt.Errorf("failed to perform migrations: %v", err)
	}
	fmt.Println("Successfully migrated database ...")
	return nil

}
