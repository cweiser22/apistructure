package store

import "apistructure/models"

type TodoStore interface {
	GetAll() ([]models.Todo, error)
	//GetById(id int) (models.Todo, error)
	//Create(t *models.Todo) error
	//Update(t *models.Todo) error
}
