package dto

import "github.com/hudayberdipolat/go-ToDoList/internal/models"

type TaskResponse struct {
	ID         uint
	TaskName   string
	TaskStatus string
	CreatedAt  string
	UpdatedAt  string
}

func NewTaskResponse(task models.Task) TaskResponse {
	return TaskResponse{
		ID:         task.ID,
		TaskName:   task.TaskName,
		TaskStatus: task.TaskStatus,
		CreatedAt:  task.CreatedAt.Format("01-02-2006"),
		UpdatedAt:  task.UpdatedAt.Format("01-02-2006"),
	}
}
