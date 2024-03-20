package db

import (
	"fmt"

	"example.com/todo-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db_dns := "host=localhost user=postgres password=password dbname=todoappdb port=5555 sslmode=disable TimeZone=Africa/Accra"

	db, err := gorm.Open(postgres.Open(db_dns), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to database")
	}

	fmt.Println("Successfully connected to database ...")

	migrations(db)

	DB = db
	return db
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Todo{}, &models.User{})
	// err = db.AutoMigrate(&models.User{})

	if err != nil {
		panic("Cannot perform migrations")
	}

	fmt.Println("Successfully migrated database ...")

}
