package todo

import (
	"apistructure/models"
	"apistructure/store"
	"database/sql"
	"fmt"
)

type SqliteTodoStore struct {
	Conn *sql.DB
}

func NewSQLTodoStore(db *sql.DB) store.TodoStore {
	return &SqliteTodoStore{
		Conn: db,
	}
}

func (s *SqliteTodoStore) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	rows, err := s.Conn.Query("SELECT * FROM todos")
	fmt.Println("made query")
	if err != nil {
		fmt.Println("there qwas a problem")
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
	fmt.Println("rows closed")
	return todos, err

}

/*
func (p *postgresTodoStore) GetByID(id int) (models.Todo, error){

}
*/
