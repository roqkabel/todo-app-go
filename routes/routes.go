package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func RegisterRoutes(r chi.Router, tokenAuth *jwtauth.JWTAuth) chi.Router {
	r.Group(func(r chi.Router) {

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/todos", HandleGetTodo)
		r.Post("/todos", HandleCreateTodo)
		r.Put("/todos/{todoId}", HandleEditTodo)
		r.Delete("/todos/{todoId}", HandleDeleteTodo)
		r.Patch("/todos/{todoId}/toggle-completed", HandleToggleComplete)

		// users

		r.Get("/users/todos", HandleUserGetTodos)
	})

	r.Post("/users/register", HandleUserRegistration)
	r.Post("/users/login", HandleUserLogin)

	return r
}
