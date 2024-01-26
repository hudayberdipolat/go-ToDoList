package repositories

import "github.com/hudayberdipolat/go-ToDoList/internal/models"

type ListRepository interface {
	FindAll(userID int) ([]models.List, error)
	FindOne(userID, listID int) (*models.List, error)
	Create(list models.List) error
	Update(userID, listID int, list models.List) error
	Delete(userID, listID int) error
}
