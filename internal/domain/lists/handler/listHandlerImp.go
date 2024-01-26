package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/lists/dto"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/lists/services"
	"github.com/hudayberdipolat/go-ToDoList/internal/utils/response"
	"github.com/hudayberdipolat/go-ToDoList/internal/utils/validate"
	"net/http"
	"strconv"
)

type listHandlerImp struct {
	listService services.ListService
}

func NewListHandler(service services.ListService) ListHandler {
	return listHandlerImp{
		listService: service,
	}
}

func (l listHandlerImp) GetAll(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	lists, err := l.listService.GetListsOfUser(userID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "user lists not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all lists of user", lists)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (l listHandlerImp) GetOne(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	listID, _ := strconv.Atoi(ctx.Params("listID"))
	list, err := l.listService.GetListOfUser(userID, listID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "user list not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get list of user", list)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (l listHandlerImp) Create(ctx *fiber.Ctx) error {
	var createListRequest dto.CreateListRequest
	userID := ctx.Locals("user_id").(int)
	if err := ctx.BodyParser(&createListRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(createListRequest); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// create list
	if err := l.listService.CreateList(userID, createListRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "list can't created", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "list created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (l listHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateListRequest dto.UpdateListRequest
	userID := ctx.Locals("user_id").(int)
	listID, _ := strconv.Atoi(ctx.Params("listID"))
	if err := ctx.BodyParser(&updateListRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(updateListRequest); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	if err := l.listService.UpdateList(userID, listID, updateListRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "list can't created", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "list updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (l listHandlerImp) Delete(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	listID, _ := strconv.Atoi(ctx.Params("listID"))
	if err := l.listService.DeleteList(userID, listID); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "list can't deleted", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "list deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
