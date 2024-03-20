package routes

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) chi.Router {

	r.Get("/todos", HandleGetTodo)
	r.Post("/todos", HandleCreateTodo)
	r.Put("/todos/{todoId}", HandleEditTodo)
	r.Delete("/todos/{todoId}", HandleDeleteTodo)
	r.Patch("/todos/{todoId}/toggle-completed", HandleToggleComplete)

	return r
}
