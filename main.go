package main

import (
	"fmt"
	"net/http"

	"example.com/todo-app/db"
	"example.com/todo-app/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	_ = db.InitDB()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	routes.RegisterRoutes(router)

	fmt.Printf("Started server on port: %v \n", 3000)
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		panic("Error starting server")
	}

}
