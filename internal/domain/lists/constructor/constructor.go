package constructor

import (
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/lists/handler"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/lists/repositories"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/lists/services"
	"gorm.io/gorm"
)

var listRepo repositories.ListRepository
var listService services.ListService
var ListHandler handler.ListHandler

func ListRequirementCreator(db *gorm.DB) {
	listRepo = repositories.NewListRepository(db)
	listService = services.NewListService(listRepo)
	ListHandler = handler.NewListHandler(listService)
}
