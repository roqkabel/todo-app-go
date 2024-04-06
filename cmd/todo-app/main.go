package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/todo-app/config"
	"example.com/todo-app/internal/db"
	"example.com/todo-app/internal/helpers"
	"example.com/todo-app/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

func main() {

	// load environment vars
	config := config.Configuration{}
	config.LoadEnvironmentVars()

	_, err := db.InitDB(config)
	if err != nil {
		log.Fatal("An error occured trying to initialize database")
	}

	TokenAuth = jwtauth.New("HS256", []byte("6c9e885a17a4077243c039017a7de8e73aa4fb57e7f39d447a1a97a03fca877071742c077acbd7e2961a2c556c73d4a0"), nil) // replace with secret key
	helpers.AppTokenAuth = *TokenAuth

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	routes.RegisterRoutes(router, TokenAuth)

	fmt.Printf("Started server on port: %v \n", 5050)
	err = http.ListenAndServe(":5050", router)

	if err != nil {
		panic("Error starting server")
	}

}
