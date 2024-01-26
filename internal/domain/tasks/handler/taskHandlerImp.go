package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/tasks/dto"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/tasks/services"
	"github.com/hudayberdipolat/go-ToDoList/internal/utils/response"
	"github.com/hudayberdipolat/go-ToDoList/internal/utils/validate"
	"net/http"
	"strconv"
)

type taskHandlerImp struct {
	taskService services.TaskService
}

func NewTaskHandler(service services.TaskService) TaskHandler {
	return taskHandlerImp{
		taskService: service,
	}
}

func (t taskHandlerImp) GetAll(ctx *fiber.Ctx) error {
	listID, _ := strconv.Atoi(ctx.Params("listID"))
	tasks, err := t.taskService.GetAll(listID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "Tasks not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "All tasks of list", tasks)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (t taskHandlerImp) GetOne(ctx *fiber.Ctx) error {
	listID, _ := strconv.Atoi(ctx.Params("listID"))
	taskID, _ := strconv.Atoi(ctx.Params("taskID"))
	tasks, err := t.taskService.GetOne(listID, taskID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "Task not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "task response", tasks)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (t taskHandlerImp) Create(ctx *fiber.Ctx) error {
	listID, _ := strconv.Atoi(ctx.Params("listID"))
	var createRequest dto.CreateTaskRequest

	if err := ctx.BodyParser(&createRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(createRequest); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := t.taskService.Create(listID, createRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "task can't created", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "task created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (t taskHandlerImp) Update(ctx *fiber.Ctx) error {
	listID, _ := strconv.Atoi(ctx.Params("listID"))
	taskID, _ := strconv.Atoi(ctx.Params("taskID"))
	var updateRequest dto.UpdateTaskRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(updateRequest); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := t.taskService.Update(listID, taskID, updateRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "task can't updated", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "task updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)

}

func (t taskHandlerImp) Delete(ctx *fiber.Ctx) error {
	listID, _ := strconv.Atoi(ctx.Params("listID"))
	taskID, _ := strconv.Atoi(ctx.Params("taskID"))
	if err := t.taskService.Delete(listID, taskID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "task can't deleted", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "task deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
