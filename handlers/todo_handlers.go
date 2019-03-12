package handlers

import (
	"apistructure/store"
	"apistructure/store/todo"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type TodoHandler struct {
	store store.TodoStore
}

func NewTodoHandler(db *sql.DB) *TodoHandler {
	return &TodoHandler{
		store: todo.NewSQLTodoStore(db),
	}
}

func (h *TodoHandler) ListAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.store.GetAll()
	fmt.Println(err)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong.")
		return
	}
	respondWithJSON(w, 200, todos)
}

func (h *TodoHandler) FetchTodoByID(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoId")

	todo, err := h.store.GetById(todoID)

	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, 404, "The todo does not exist.")
			return
		} else {
			respondWithError(w, 500, "Something went wrong.")
			return
		}
	}
	respondWithJSON(w, 200, todo)

}

/*

func (h *TodoHandler) CreateNewTodo(w http.ResponseWriter, r *http.Request) {

}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {

}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {

}
*/
