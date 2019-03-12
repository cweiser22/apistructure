package todo

import (
	"apistructure/models"
	"apistructure/store"
	"database/sql"
)

//impliments TodoStore interface
type SqliteTodoStore struct {
	Conn *sql.DB
}

//creates a SqliteTodoStore
func NewSQLTodoStore(db *sql.DB) store.TodoStore {
	return &SqliteTodoStore{
		Conn: db,
	}
}

//gets all todos from database
func (s *SqliteTodoStore) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	rows, err := s.Conn.Query("SELECT todo_id, task FROM todos")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.Task)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	rows.Close()
	return todos, err

}

func (s *SqliteTodoStore) GetById(id string) (models.Todo, error) {
	var todo models.Todo
	err := s.Conn.QueryRow("SELECT todo_id, task FROM todos WHERE todo_id = ?", id).Scan(&todo.ID, &todo.Task)
	if err != nil {
		return todo, err
	}
	return todo, err
}

/*
func (s *SqliteTodoStore) Create(t *models.Todo) error {

}
*/
