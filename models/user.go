package models

import (
	"errors"

	// "example.com/todo-app/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Todos    []Todo
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	if u.EmailAlreadyExist(tx) {
		return errors.New("User already exist")
	}

	return
}

func (u *User) EmailAlreadyExist(db *gorm.DB) bool {

	var foundUser User = User{}

	if tx := db.Where("email = ?", u.Email).First(&foundUser); tx.Error != nil {
		// if there's an error fetching record
		return !errors.Is(tx.Error, gorm.ErrRecordNotFound)
	}

	if foundUser.Email == u.Email {
		return true
	}

	return false
}
