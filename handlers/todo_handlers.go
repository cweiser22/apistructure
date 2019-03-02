package handlers

import (
	"apistructure/store"
	"apistructure/store/todo"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
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
	respondwithJSON(w, http.StatusOK, todos)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
