package routes

import (
	"encoding/json"
	"net/http"

	"example.com/todo-app/db"
	"example.com/todo-app/helpers"
	"example.com/todo-app/models"
	"github.com/go-chi/render"
)

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

	// hash password

	hashedPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		helpers.Response(w, r, helpers.ResponseParams{
			StatusCode: 500,
			Message:    "An system error occured",
			Result:     nil,
		})
		return
	}

	user.Password = hashedPassword

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

	loginRequest := UserLoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		helpers.Response(w, r, helpers.ResponseParams{
			StatusCode: 500,
			Message:    "Failed to parse request body",
			Result:     nil,
		})
		return
	}

	defer r.Body.Close()

	var foundUser models.User

	if tx := db.DB.Model(models.User{Email: loginRequest.Email}).First(&foundUser); tx.Error != nil {
		helpers.Response(w, r, helpers.ResponseParams{
			StatusCode: 400,
			Message:    "Account with the email does not exist",
		})
		return
	}

	//  compare user password.

	if !helpers.MatchPasswords(loginRequest.Password, foundUser.Password) {
		helpers.Response(w, r, helpers.ResponseParams{
			StatusCode: 400,
			Message:    "Password does not match",
			Result:     nil,
		})
		return
	}

	// generate user tokens

	_, tokenString, err := helpers.AppTokenAuth.Encode(map[string]interface{}{
		"user_id": foundUser.ID,
	})

	if err != nil {
		helpers.Response(w, r, helpers.ResponseParams{
			StatusCode: 500,
			Message:    "An error occured trying to geneate jwt",
			Result:     nil,
		})
	}

	render.JSON(w, r, map[string]any{
		"token": tokenString,
		"name":  foundUser.Name,
	})

}
