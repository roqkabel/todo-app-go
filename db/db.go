package db

import (
	"fmt"

	"example.com/todo-app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/app.db"), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to database")
	}

	fmt.Println("Successfully connected to database ...")

	migrations(db)

	DB = db
	return db
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Todo{})

	if err != nil {
		panic("Cannot perform migrations")
	}

	fmt.Println("Successfully migrated database ...")

}
