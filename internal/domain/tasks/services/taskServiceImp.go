package services

import (
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/tasks/dto"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/tasks/repositories"
	"github.com/hudayberdipolat/go-ToDoList/internal/models"
	"time"
)

type taskServiceImp struct {
	taskRepo repositories.TaskRepository
}

func NewTaskRepository(repo repositories.TaskRepository) TaskService {
	return taskServiceImp{
		taskRepo: repo,
	}
}

func (t taskServiceImp) GetAll(listID int) ([]dto.TaskResponse, error) {
	tasks, err := t.taskRepo.GetAll(listID)
	if err != nil {
		return nil, err
	}
	var taskResponses []dto.TaskResponse

	for _, task := range tasks {
		taskResponse := dto.NewTaskResponse(task)
		taskResponses = append(taskResponses, taskResponse)
	}
	return taskResponses, err
}

func (t taskServiceImp) GetOne(listID, taskID int) (*dto.TaskResponse, error) {
	task, err := t.taskRepo.GetOne(listID, taskID)
	if err != nil {
		return nil, err
	}
	taskResponse := dto.NewTaskResponse(*task)

	return &taskResponse, nil
}

func (t taskServiceImp) Create(listID int, createRequest dto.CreateTaskRequest) error {

	createTask := models.Task{
		TaskName:   createRequest.TaskName,
		TaskStatus: "1",
		ListID:     uint(listID),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := t.taskRepo.Create(createTask); err != nil {
		return err
	}
	return nil
}

func (t taskServiceImp) Update(listID, taskID int, updateRequest dto.UpdateTaskRequest) error {
	updateTask := models.Task{
		TaskName:   updateRequest.TaskName,
		TaskStatus: updateRequest.TaskStatus,
		UpdatedAt:  time.Now(),
	}
	if err := t.taskRepo.Update(listID, taskID, updateTask); err != nil {
		return err
	}
	return nil
}

func (t taskServiceImp) Delete(listID, taskID int) error {
	if err := t.taskRepo.Delete(listID, taskID); err != nil {
		return err
	}
	return nil
}
