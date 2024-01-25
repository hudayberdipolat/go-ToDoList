package constructor

import (
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/user/handler"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/user/repositories"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/user/services"
	"gorm.io/gorm"
)

var userRepo repositories.UserRepository
var userService services.UserService
var UserHandler handler.UserHandler

func UserRequirementCreator(db *gorm.DB) {
	userRepo = repositories.NewUserRepository(db)
	userService = services.NewUserService(userRepo)
	UserHandler = handler.NewUserHandler(userService)
}
