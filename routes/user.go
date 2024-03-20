package routes

import (
	"encoding/json"
	"net/http"

	"example.com/todo-app/db"
	"example.com/todo-app/helpers"
	"example.com/todo-app/models"
)

func HandleUserRegistration(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		helpers.Response(w, r, helpers.ResponseParams{
			StatusCode: 500,
			Message:    "Failed trying to parse json",
			Result:     nil,
		})
		return
	}

	// verify email already exist

	if user.EmailAlreadyExist(db.DB) {
		helpers.Response(w, r, helpers.ResponseParams{
			StatusCode: 400,
			Message:    "Email already exist",
			Result:     nil,
		})
		return
	}

	// save data

	if tx := db.DB.Model(&models.User{}).Create(&user); tx.Error != nil {

		helpers.Response(w, r, helpers.ResponseParams{
			StatusCode: 500,
			Message:    "An error occured whiles registering user",
			Result:     tx.Error,
		})
		return
	}

	helpers.Response(w, r, helpers.ResponseParams{
		StatusCode: 201,
		Message:    "User created successfull",
	})
}

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		helpers.Response(w, r, helpers.ResponseParams{
			StatusCode: 500,
			Message:    "Failed to parse request body",
			Result:     nil,
		})
		return
	}

	// verify is user has an account

	var foundUser models.User

	if tx := db.DB.Model(models.User{Email: user.Email}).First(&foundUser); tx.Error != nil {
		helpers.Response(w, r, helpers.ResponseParams{
			StatusCode: 400,
			Message:    "Account with the email does not exist",
		})
		return
	}

	//  compare user password.

}
