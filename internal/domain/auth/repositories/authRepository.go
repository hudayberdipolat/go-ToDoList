package repositories

import "github.com/hudayberdipolat/go-ToDoList/internal/models"

type AuthRepository interface {
	GetUserWithUsername(username string) (*models.User, error)
	CreateUser(user models.User) error
}
