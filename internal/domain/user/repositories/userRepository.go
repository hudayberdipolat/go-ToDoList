package repositories

import "github.com/hudayberdipolat/go-ToDoList/internal/models"

type UserRepository interface {
	UpdateData(userID int, user models.User) error
	UpdatePassword(userID int, password string) error
	GetUser(userID int) (*models.User, error)
}
