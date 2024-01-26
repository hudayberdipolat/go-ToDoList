package constructor

import (
	"github.com/hudayberdipolat/go-ToDoList/internal/app"
	authConstructor "github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/constructor"
	listConstructor "github.com/hudayberdipolat/go-ToDoList/internal/domain/lists/constructor"
	taskConstructor "github.com/hudayberdipolat/go-ToDoList/internal/domain/tasks/consturtor"
	userConstructor "github.com/hudayberdipolat/go-ToDoList/internal/domain/user/constructor"
)

func Build(dependencies *app.Dependencies) {
	authConstructor.AuthRequirementCreator(dependencies.DB)
	userConstructor.UserRequirementCreator(dependencies.DB)
	listConstructor.ListRequirementCreator(dependencies.DB)
	taskConstructor.TaskRequirementCreator(dependencies.DB)
}
