package services

import "github.com/hudayberdipolat/go-ToDoList/internal/domain/tasks/dto"

type TaskService interface {
	GetAll(listID int) ([]dto.TaskResponse, error)
	GetOne(listID, taskID int) (*dto.TaskResponse, error)
	Create(listID int, createRequest dto.CreateTaskRequest) error
	Update(listID, taskID int, updateRequest dto.UpdateTaskRequest) error
	Delete(listID, taskID int) error
}
