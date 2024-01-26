package consturtor

import (
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/tasks/handler"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/tasks/repositories"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/tasks/services"
	"gorm.io/gorm"
)

var taskRepo repositories.TaskRepository
var taskService services.TaskService
var TaskHandler handler.TaskHandler

func TaskRequirementCreator(db *gorm.DB) {
	taskRepo = repositories.NewTaskRepository(db)
	taskService = services.NewTaskRepository(taskRepo)
	TaskHandler = handler.NewTaskHandler(taskService)
}
