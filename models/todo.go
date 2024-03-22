package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      int64  `json:"user_id"`
	User        User   `gorm:"foreignKey:user_id" json:"user"`
}
