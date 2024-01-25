package constructor

import (
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/handler"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/repositories"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/services"
	"gorm.io/gorm"
)

var authRepo repositories.AuthRepository
var authService services.AuthService
var AuthHandler handler.AuthHandler

func AuthRequirementCreator(db *gorm.DB) {
	authRepo = repositories.NewAuthRepository(db)
	authService = services.NewAuthService(authRepo)
	AuthHandler = handler.NewAuthHandler(authService)
}
