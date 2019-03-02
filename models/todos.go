package models

type Todo struct {
	ID   int    `json:"id" db:"todo_id"`
	Task string `json:"task" db:"task"`
}

/*
func GetAllTodos() []Todo {
	var todos []Todo
	rows, err := sql.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
}
*/
