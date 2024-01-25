package constructor

import (
	"github.com/hudayberdipolat/go-ToDoList/internal/app"
	authConstructor "github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/constructor"
)

func Build(dependencies *app.Dependencies) {
	authConstructor.AuthRequirementCreator(dependencies.DB)

}
