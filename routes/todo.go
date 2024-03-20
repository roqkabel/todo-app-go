package routes

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/todo-app/db"
	"example.com/todo-app/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type NewTodoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func HandleGetTodo(w http.ResponseWriter, r *http.Request) {

	var resultedTodos []models.Todo

	db.DB.Model(&models.Todo{}).Order("created_at desc").Find(&resultedTodos)

	render.JSON(w, r, resultedTodos)
}

func HandleCreateTodo(w http.ResponseWriter, r *http.Request) {

	todo := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(&todo)

	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		render.JSON(w, r, err)
		return
	}

	newTodo := db.DB.Model(&models.Todo{}).Create(&todo)

	if newTodo.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, newTodo.Error)
		return
	}

	fmt.Println(newTodo.Error)

	w.WriteHeader(201)
	render.JSON(w, r, "Created todo successfull")
}

func HandleEditTodo(w http.ResponseWriter, r *http.Request) {

	todo := models.Todo{}

	todoId := chi.URLParam(r, "todoId")

	// find if todo already exist

	foundTodo := db.DB.First(&todo, todoId)

	if foundTodo.Error != nil {
		render.Status(r, 404)
		render.JSON(w, r, map[string]any{
			"message": "Record not found",
		})
		return
	}

	err := json.NewDecoder(r.Body).Decode(&todo)

	defer r.Body.Close()

	if err != nil {
		render.Status(r, 500)
		render.JSON(w, r, map[string]any{
			"message": "Failed parsing json fields",
		})
		return
	}

	if tx := db.DB.Save(&todo); tx.Error != nil {
		render.Status(r, 500)
		render.JSON(w, r, map[string]any{
			"message": "An Error occured while updating record.",
		})
		return
	}

	render.JSON(w, r, map[string]any{
		"message": "Todo has been updated successfully",
		"result":  todo,
	})
}

func HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {

	todoId := chi.URLParam(r, "todoId")

	todo := models.Todo{}

	if foundTodo := db.DB.First(&todo, todoId); foundTodo.Error != nil {
		render.Status(r, 404)
		render.JSON(w, r, map[string]string{
			"message": "Record not found",
		})
		return
	}

	if tx := db.DB.Delete(&todo); tx.Error != nil {
		render.Status(r, 500)
		render.JSON(w, r, map[string]string{
			"message": "An error occured trying to delete todo",
		})
		return
	}

	render.JSON(w, r, map[string]any{
		"message": "Todo has been deleted successfully",
		"result":  todo.ID,
	})
}

func HandleToggleComplete(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")
	todo := models.Todo{}
	var toggleMessage string = "Todo has been marked as completed"

	if tx := db.DB.First(&todo, todoId); tx.Error != nil {
		w.WriteHeader(404)
		render.JSON(w, r, map[string]string{
			"message": "Todo was not found",
		})
		return
	}

	// toggle todo

	todo.Completed = !todo.Completed

	if tx := db.DB.Save(&todo); tx.Error != nil {
		w.WriteHeader(500)
		render.JSON(w, r, map[string]string{
			"message": "An error occured whiles marking todo as complete",
		})
		return
	}

	if !todo.Completed {
		toggleMessage = "Todo has been marked as incomplete"
	}

	render.JSON(w, r, map[string]string{
		"message": toggleMessage,
	})

}
