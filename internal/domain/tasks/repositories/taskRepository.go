package repositories

import "github.com/hudayberdipolat/go-ToDoList/internal/models"

type TaskRepository interface {
	GetAll(listID int) ([]models.Task, error)
	GetOne(listID, taskID int) (*models.Task, error)
	Create(task models.Task) error
	Update(listID, taskID int, task models.Task) error
	Delete(listID, taskID int) error
}
